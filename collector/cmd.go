package collector

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers"
	"runtime"
)

var (
	Cmd = &cobra.Command{
		Use:   "collector",
		Short: "start raccoon collector",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}
	flags cmdFlags
	cfg   Config
)

func init() {
	// Config file path
	Cmd.Flags().StringVar(&flags.ConfigPath, "config", "", "configuration file")
	// Raccoon Collector ID
	Cmd.Flags().StringVar(&flags.ID, "id", "", "collector ID")
	// Raccoon core URL
	Cmd.Flags().StringVar(&flags.CoreURL, "core", "http://localhost:7220", "core URL")
	// Raccoon bus URL
	Cmd.Flags().StringVar(&flags.BusURL, "bus", "nats://localhost:4222", "bus URL")
	// Raccoon storage URL
	Cmd.Flags().StringVar(&flags.StorageURL, "storage", "http://localhost:9200", "storage URL")
	// Prometheus metrics port
	Cmd.Flags().StringVar(&flags.MetricsPort, "metrics", "7221", "metrics port")
	// Worker count
	Cmd.Flags().IntVar(&flags.Workers, "workers", runtime.NumCPU(), "worker count")
}

func run(_ *cobra.Command, _ []string) (err error) {
	//
	// Check command line flags
	//

	if flags.ID == "" && flags.CoreURL == "" && flags.ConfigPath == "" {
		return fmt.Errorf("either config file or core URL and collector ID must be specified")
	}

	//
	// Load configuration
	//

	if flags.ConfigPath != "" {
		if err := helpers.ReadConfigFromFile(flags.ConfigPath, &cfg); err != nil {
			return err
		}
	} else {
		if err := helpers.ReadConfigFromCore(flags.CoreURL, "collector", flags.ID, &cfg); err != nil {
			return err
		}
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	//
	// Prepare processor for initialization
	//

	proc := Processor{
		hostname:     helpers.GetHostName(),
		ipAddress:    helpers.GetIPAddress(),
		metrics:      newMetrics(flags.MetricsPort),
		inputChannel: make(connectors.OutputChannel),
		enrichment:   cfg.Enrichment,
		workers:      flags.Workers,
	}

	//
	// Initialize active lists
	//

	globals.ActiveLists, err = activeLists.NewContainer(cfg.ActiveLists, cfg.Name, flags.BusURL, flags.StorageURL)
	if err != nil {
		return err
	}

	//
	// Initialize dictionaries
	//

	globals.Dictionaries = dictionaries.NewStorage(cfg.Dictionaries)

	//
	// Initialize normalizer
	//

	proc.normalizer, err = normalizers.New(cfg.Normalizer)
	if err != nil {
		return err
	}

	//
	// Initialize filters
	//

	for _, cfg := range cfg.Filters {
		f, err := filters.NewFilter(cfg)
		if err != nil {
			return err
		}
		proc.filters = append(proc.filters, f)
	}

	//
	// Initialize aggregation rules
	//

	for _, cfg := range cfg.Rules {
		rule, err := aggregation.NewRule(cfg, proc.output)
		if err != nil {
			return err
		}
		rule.Start()
		proc.aggregationRules = append(proc.aggregationRules, rule)
	}

	//
	// Initialize destinations
	//

	for _, cfg := range cfg.Destinations {
		dst, err := destinations.New(cfg)
		if err != nil {
			return err
		}
		if err := dst.Start(); err != nil {
			return err
		}
		proc.destinations = append(proc.destinations, dst)
	}

	//
	// Initialize connector
	//

	connector, err := connectors.New(cfg.Connector, proc.inputChannel)
	if err != nil {
		return err
	}

	if err := connector.Start(); err != nil {
		return err
	}

	//
	// Begin processing
	//

	proc.Start()
	runtime.Goexit()
	return nil
}
