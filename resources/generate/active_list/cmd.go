package activeList

import (
	"encoding/csv"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"io"
	"os"
)

var Cmd = &cobra.Command{
	Use:   "al",
	Short: "generate active list",
	Args:  cobra.ExactArgs(0),
	RunE:  run,
}

var flags cmdFlags

func init() {
	// Source file path
	Cmd.Flags().StringVar(&flags.SourcePath, "from", "", "source file path")
	// Active list storage URL
	Cmd.Flags().StringVar(&flags.ALStorage, "storage", "localhost:6379", "active list storage URL")
	// Active list name
	Cmd.Flags().StringVar(&flags.ALName, "name", "", "active list name")
	// Delimiter
	Cmd.Flags().StringVar(&flags.Delimiter, "delimiter", ",", "delimiter")
	// Key field
	Cmd.Flags().StringVar(&flags.KeyField, "key", "", "key field")
	// Write
	Cmd.Flags().BoolVar(&flags.Write, "write", false, "do write to storage")
	// Required flags
	_ = Cmd.MarkFlagRequired("from")
	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("key")

}

func run(_ *cobra.Command, _ []string) error {
	if len(flags.Delimiter) != 1 {
		return fmt.Errorf("delimiter must be single byte ASCII character")
	}

	file, err := os.Open(flags.SourcePath)
	if err != nil {
		return err
	}

	defer file.Close()

	storageCli := redis.NewClient(&redis.Options{Addr: flags.ALStorage})
	tx := storageCli.TxPipeline()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.Comma = rune(flags.Delimiter[0])

	var columns []string
	firstLine := true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if firstLine {
			firstLine = false
			columns = record
			continue
		}

		var alKey string
		alValues := make(map[string]interface{})
		for i, value := range record {
			colName := columns[i]
			alValues[colName] = value
			if colName == flags.KeyField {
				alKey = value
			}
		}

		finalKey := activeLists.MakeFinalKey(flags.ALName, alKey)

		if flags.Write {
			tx.HMSet(finalKey, alValues)
		} else {
			fmt.Printf("%s: %v\n", finalKey, alValues)
		}
	}

	if flags.Write {
		if _, err := tx.Exec(); err != nil {
			return err
		}
	}

	fmt.Println("OK")
	return nil
}
