package player

import (
	"github.com/spf13/cobra"
	"net"
	"time"
)

var (
	Cmd = &cobra.Command{
		Use:   "player",
		Short: "replay events",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	sourceFilePath, destinationURL string

	// Bool flags variables
	useUDP, sourceContainsPcap bool

	// Other flags variables
	interval time.Duration

	// Runtime variables
	connection net.Conn
)

func init() {
	// Replay file path
	Cmd.Flags().StringVarP(
		&sourceFilePath,
		"file",
		"f",
		"",
		"replay file path")

	// Destination URL
	Cmd.Flags().StringVarP(
		&destinationURL,
		"dst",
		"d",
		"127.0.0.1:1514",
		"destination URL")

	// Interval between message sends
	Cmd.Flags().DurationVarP(
		&interval,
		"interval",
		"i",
		100*time.Microsecond,
		"interval between message sends")

	// Replay file contains pcap dump
	Cmd.Flags().BoolVarP(
		&sourceContainsPcap,
		"pcap",
		"p",
		false,
		"replay file contains pcap dump")

	// Use UDP to connect to destination
	Cmd.Flags().BoolVarP(
		&useUDP,
		"udp",
		"u",
		false,
		"use UDP")

	Cmd.MarkFlagRequired("file")
}

func run(_ *cobra.Command, _ []string) error {
	proto := "tcp"
	if useUDP {
		proto = "udp"
	}

	c, err := net.Dial(proto, destinationURL)
	if err != nil {
		return err
	}

	connection = c

	if !sourceContainsPcap {
		replayTextFile(sourceFilePath)
	} else {
		replayPcapFile(sourceFilePath)
	}

	return nil
}
