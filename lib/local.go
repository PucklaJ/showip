package lib

import "net"

type LocalAddress struct {
	IP        net.IP
	Interface net.Interface
}

func GetLocalIPv4Address() ([]LocalAddress, error) {
	var ipAddrs []LocalAddress

	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range iFaces {
		addrs, err := i.Addrs()
		if err != nil {
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

			if ip.IsLoopback() {
				continue
			}

			ip = ip.To4()

			if ip == nil {
				continue
			}

			ipAddrs = append(ipAddrs, LocalAddress{
				IP:        ip,
				Interface: i,
			})
		}
	}

	return ipAddrs, nil
}
