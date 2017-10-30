package netx

import (
	"errors"
	"net"
)

type IfacePred func(net.Interface) bool

func LocalIPs(ifacePred IfacePred) ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0)
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		if ifacePred != nil && !ifacePred(iface) {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip)
		}
	}
	return ips, nil
}

func LocalIP(ifacePred IfacePred) (net.IP, error) {
	ips, err := LocalIPs(ifacePred)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, errors.New("Not found IP")
	}
	return ips[0], nil
}

func LocalIPStr(ifacePred IfacePred, def string) string {
	ip, err := LocalIP(ifacePred)
	if err != nil {
		return def
	}
	return ip.String()
}

func OnlyEth0(iface net.Interface) bool {
	return iface.Name == "eth0"
}

func OnlyEth1(iface net.Interface) bool {
	return iface.Name == "eth1"
}
