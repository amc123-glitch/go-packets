package main

import (
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
  "fmt"
  "log"
  "time"
)

var (
  device string = "en0"
  snaplen int32 = 65535
  promisc bool = false
  err error
  timeout time.Duration = -1 * time.Second
  handle *pcap.Handle
)

func main() {
  handle, err = pcap.OpenLive(device,snaplen,promisc,timeout)
  if err != nil {
    log.Fatal(err)
  }
  defer handle.Close()

  var filter string = "src host 127.0.0.1 and icmp"
  err = handle.SetBPFFilter(filter)
  if err != nil {
    log.Fatal(err)
  }

  packetSource := gopacket.NewPacketSource(handle,handle.LinkType())

  for packet := range packetSource.Packets() {
    fmt.Println("You have been pinged!!")
    fmt.Println("----------")
    fmt.Println(packet)
  }
}
