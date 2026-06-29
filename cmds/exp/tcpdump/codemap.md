# cmds/exp/tcpdump/

## Responsibility
Network packet capture and analysis tool with BPF filtering, protocol decoding, and timestamp options.

## Integration Points
- github.com/gopacket/gopacket: used for packet decoding
- github.com/gopacket/gopacket/layers: used for protocol layer parsing (Ethernet, IP, TCP, UDP, DNS, ICMP)
- github.com/packetcap/go-pcap: used for live packet capture from interfaces
- pkg/uroot/unixflag: used for flag parsing
