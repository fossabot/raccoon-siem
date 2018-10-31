package main

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/jessevdk/go-flags"
	"log"
	"net"
	"os"
	"time"
)

var settings *playerSettings
var Connection net.Conn

type playerSettings struct {
	File        string        `long:"file" required:"y" description:"source file path"`
	Destination string        `long:"dst" default:"127.0.0.1:1514" description:"destination url"`
	Interval    time.Duration `long:"interval" default:"1ms" description:"send interval"`
	UseUDP      bool          `long:"udp" description:"use udp"`
	UsePCAP     bool          `long:"pcap" description:"use pcap parser"`
}

func main() {
	// Parse cmd flags
	settings = new(playerSettings)
	if _, err := flags.Parse(settings); err != nil {
		log.Fatal(err)
	}

	proto := "tcp"
	if settings.UseUDP {
		proto = "udp"
	}

	c, err := net.Dial(proto, settings.Destination)

	if err != nil {
		printErrorAndExit(err)
	}

	Connection = c

	fmt.Printf(">> Dialed to %s (%s)\n", settings.Destination, proto)

	if !settings.UsePCAP {
		replayTextFile(settings.File)
	} else {
		replayPcapFile(settings.File)
	}
}

func replayPcapFile(path string) {
	handle, err := pcap.OpenOffline(path)

	if err != nil {
		printErrorAndExit(err)
	}

	defer handle.Close()

	data := make([][]byte, 0)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		pld := packet.TransportLayer().LayerPayload()
		data = append(data, pld)
	}

	if settings.Interval > 0 {
		sendLogsInLoop(data)
	} else {
		sendLogs(data)
	}
}

func replayTextFile(path string) {
	fd, err := os.Open(path)

	if err != nil {
		printErrorAndExit(err)
	}

	defer fd.Close()

	lines := getLinesFromFile(fd)

	fmt.Printf(">> Loaded lines from %s\n", path)

	if settings.Interval > 0 {
		sendLogsInLoop(lines)
	} else {
		sendLogs(lines)
	}
}

func sendLogsInLoop(lines [][]byte) {
	fmt.Printf(">> Sending logs in loop with %s interval ...\n", settings.Interval)
	for {
		for i := range lines {
			Connection.Write(lines[i])
			time.Sleep(settings.Interval)
		}
	}
}

func sendLogs(lines [][]byte) {
	fmt.Printf(">> Sending logs once with %s interval ...\n", settings.Interval)
	for i := range lines {
		Connection.Write(lines[i])
		time.Sleep(settings.Interval)
	}
}

func getLinesFromFile(fd *os.File) (lines [][]byte) {
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) != 0 {
			line := append(bytes, '\n')
			lines = append(lines, line)
		}
	}
	return
}

func printErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
