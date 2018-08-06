package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"

// 	"github.com/lflxp/webMonitor/decoder/http"

// 	"github.com/google/gopacket"
// 	"github.com/google/gopacket/layers"
// 	"github.com/google/gopacket/pcap"
// )

// var (
// 	device       string = "enp3s0"
// 	snapshot_len int32  = 2048
// 	promiscuous  bool   = true
// 	err          error
// 	timeout      time.Duration = 30 * time.Second
// 	handle       *pcap.Handle
// )

// func main() {
// 	// Open device
// 	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer handle.Close()

// 	// Use the handle as a packet source to process all packets
// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// 	for p := range packetSource.Packets() {
// 		// Process packet here
// 		// fmt.Println(p)
// 		// eth := p.Layer(layers.LayerTypeEthernet)
// 		// if eth != nil {
// 		// 	fmt.Println("Ethernet layer detected.")
// 		// 	ethernetPacket, _ := eth.(*layers.Ethernet)
// 		// 	fmt.Println("Source MAC:", ethernetPacket.SrcMAC)
// 		// 	fmt.Println("Destionation MAC:", ethernetPacket.DstMAC)
// 		// 	// Ethernet type is typically IPv4 but could be ARP or other
// 		// 	fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
// 		// 	fmt.Println()
// 		// }

// 		// Dq := p.Layer(layers.LayerTypeDot1Q)
// 		// if Dq != nil {
// 		// 	fmt.Println("LayerTypeDot1Q layer detected.")
// 		// 	dd, _ := Dq.(*layers.Dot1Q)
// 		// 	fmt.Println("DropEligible:", dd.DropEligible)
// 		// 	fmt.Println("Priority:", dd.Priority)
// 		// 	// Ethernet type is typically IPv4 but could be ARP or other
// 		// 	fmt.Println("Dot1Q type: ", dd.Type)
// 		// 	fmt.Println("VLANIdentifier: ", dd.VLANIdentifier)
// 		// 	fmt.Println()
// 		// }

// 		// icmq := p.Layer(layers.LayerTypeICMPv4)
// 		// if icmq != nil {
// 		// 	fmt.Println("LayerTypeICMPv4 layer detected.")
// 		// 	ic, _ := icmq.(*layers.ICMPv4)
// 		// 	fmt.Println("Checksum:", ic.Checksum)
// 		// 	fmt.Println("Id:", ic.Id)
// 		// 	// Ethernet type is typically IPv4 but could be ARP or other
// 		// 	fmt.Println("Seq: ", ic.Seq)
// 		// 	fmt.Println("TypeCode: ", ic.TypeCode.String())
// 		// 	fmt.Println()
// 		// }

// 		// Let's see if the packet is IP ∂(even though the ether type told us)
// 		ipLayer := p.Layer(layers.LayerTypeIPv4)
// 		if ipLayer != nil {
// 			fmt.Println("IPv4 layer detected.")
// 			ip, _ := ipLayer.(*layers.IPv4)

// 			// IP layer variables:
// 			// Version (Either 4 or 6)
// 			// IHL (IP Header Length in 32-bit words)
// 			// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
// 			// Checksum, SrcIP, DstIP
// 			fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
// 			fmt.Println("Protocol: ", ip.Protocol)
// 			fmt.Println()
// 		}

// 		// Let's see if the packet is TCP
// 		udpLayer := p.Layer(layers.LayerTypeUDP)
// 		if udpLayer != nil {
// 			fmt.Println("UDP layer detected.")
// 			udp, _ := udpLayer.(*layers.UDP)

// 			// TCP layer variables:
// 			// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
// 			// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
// 			fmt.Printf("From port %d to %d\n", udp.SrcPort, udp.DstPort)
// 			fmt.Println("Checksum number: ", udp.Checksum)
// 			fmt.Println()

// 			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
// 			if pp.ErrorLayer() == nil {
// 				fmt.Println("UDP SFLOW detected")
// 			}
// 		}
// 		// sflowlayer := p.Layer(layers.LayerTypeSFlow)
// 		// if sflowlayer != nil {
// 		// 	fmt.Println("SFLOW layer detected")
// 		// }

// 		// When iterating through packet.Layers() above,
// 		// if it lists Payload layer then that is the same as
// 		// this applicationLayer. applicationLayer contains the payload
// 		applicationLayer := p.ApplicationLayer()
// 		if applicationLayer != nil {
// 			// fmt.Println("Application layer/Payload found.")
// 			// fmt.Printf("%d %s\n", len(applicationLayer.Payload()), applicationLayer.Payload())
// 			// Search for a string inside the payload
// 			if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
// 				fmt.Println("HTTP found!")
// 				// fmt.Println(applicationLayer.LayerType().String())
// 				// fmt.Println(p.Dump())
// 				decoder := http.Decoder{}
// 				decoder.SetFilter("")
// 				decoder.Buf = bufio.NewReader(bytes.NewReader(applicationLayer.Payload()))
// 				_data, err := decoder.DecodeHttp()
// 				if err != nil {
// 					fmt.Println(err)
// 				}

// 				// req := _data.(*http.HttpReq)
// 				// if req != nil {
// 				// 	fmt.Println("Request", req)
// 				// }

// 				// resp := _data.(*http.HttpResp)
// 				// if resp != nil {
// 				// 	fmt.Println("Response", resp)
// 				// }

// 				switch _data.(type) {
// 				case *http.HttpReq:
// 					fmt.Println("Request", _data.(*http.HttpReq))
// 				case *http.HttpResp:
// 					fmt.Println("Response", _data.(*http.HttpResp))
// 				}
// 			}
// 			// fmt.Println(applicationLayer.Payload())
// 			// fmt.Println()
// 		}

// 		// Check for errors
// 		// if err := p.ErrorLayer(); err != nil {
// 		// 	fmt.Println("Error decoding some part of the packet:", err)
// 		// }

// 		// for _, x := range p.Layers() {
// 		// 	fmt.Println(x.LayerType().String())
// 		// }

// 		// fmt.Println()
// 	}
// }
