/***********************************************************************************
 *                         This file is part of showip
 *                    https://github.com/PucklaMotzer09/showip
 ***********************************************************************************
 * Copyright (c) 2022 PucklaMotzer09
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from the
 * use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not claim
 * that you wrote the original software. If you use this software in a product,
 * an acknowledgment in the product documentation would be appreciated but is
 * not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source distribution.
 ************************************************************************************/

package lib

import "net"

type LocalAddress struct {
	IP        net.IP
	Interface net.Interface
}

func GetLocalIPAddress() ([]LocalAddress, error) {
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

			if ip == nil || ip.IsLoopback() {
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

func IsIPv4(ip net.IP) bool {
	return (len(ip) == net.IPv4len) || (len(ip) == net.IPv6len && isZeros(ip[0:10]) && ip[10] == 0xff && ip[11] == 0xff)
}

func IsIPv6(ip net.IP) bool {
	return !IsIPv4(ip)
}

// Is p all zeros?
func isZeros(p net.IP) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			return false
		}
	}
	return true
}
