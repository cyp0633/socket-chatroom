package internal

import (
	"net"
	"regexp"
	"time"

	"github.com/seancfoley/ipaddress-go/ipaddr"
	"go.uber.org/zap"
)

func Probe() (ip string) {
	ifaces, err := net.Interfaces()
	var ipChan = make(chan string)
	go listenProbePkt(&ipChan)
	if err != nil {
		Logger.Error("Reading interfaces failed", zap.Error(err))
		return
	}
	for _, i := range ifaces {
		// get broadcast addresses
		Logger.Info("Interface", zap.String("name", i.Name))
		addrs, err := i.Addrs()
		if err != nil {
			Logger.Error("Reading addresses failed", zap.Error(err))
			return
		}
		for _, addr := range addrs {
			var ipaddr1 *ipaddr.IPAddress
			switch v := addr.(type) {
			case *net.IPNet:
				ipaddr1, _ = ipaddr.NewIPAddressFromNetIPNet(v)
			case *net.IPAddr:
				ipaddr1, _ = ipaddr.NewIPAddressFromNetIPAddr(v)
			}
			if ipaddr1 == nil || ipaddr1.IsIPv6() {
				continue
			}
			bcast, _ := ipaddr1.ToIPv4().ToBroadcastAddress()
			Logger.Info("Broadcast address", zap.String("address", bcast.String()))
			sendProbePkt(bcast.String())
		}
	}
	return <-ipChan
}

func sendProbePkt(addr string) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(addr),
		Port: 65432,
	})
	if err != nil {
		Logger.Error("DialUDP failed", zap.Error(err))
		return
	}
	defer socket.Close()
	_, err = socket.Write([]byte("PROBE"))
	// wait 0.1 second
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		Logger.Error("Write failed", zap.Error(err))
		return
	}
}

var probeReplyRegex = regexp.MustCompile(`^HERE`)

func listenProbePkt(ch *chan string) {
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 65435,
	})
	if err != nil {
		Logger.Error("ListenUDP failed", zap.Error(err))
		return
	}
	defer socket.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := socket.ReadFromUDP(buf)
		if err != nil {
			Logger.Error("ReadFromUDP failed", zap.Error(err))
			// continue
			return
		}
		if probeReplyRegex.Match(buf[:n]) {
			*ch <- addr.IP.String()
			Logger.Info("Probe reply", zap.String("address", addr.IP.String()))
			return
		}
	}
}
