package player

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
	"time"
)

func replayPcapFile(path string) error {
	handle, err := pcap.OpenOffline(path)
	if err != nil {
		return err
	}

	defer handle.Close()

	data := make([][]byte, 0)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		pld := packet.TransportLayer().LayerPayload()
		data = append(data, pld)
	}

	return sendEventsInLoop(data, interval)
}

func replayTextFile(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	return sendEventsInLoop(getLinesFromFile(fd), interval)
}

func sendEventsInLoop(lines [][]byte, interval time.Duration) error {
	fmt.Printf(">> Sending logs in loop with %s interval ...\n", interval)
	for {
		for i := range lines {
			idx := 0
			for {
				sent, err := connection.Write(lines[i][idx:])
				if err != nil {
					return err
				}

				if sent == len(lines[i]) {
					break
				}

				idx = sent + 1
			}
			time.Sleep(interval)
		}
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
