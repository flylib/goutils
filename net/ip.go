package net

import (
	"fmt"
	"net"
)

// Get preferred outbound ip of this machine(局域网ip）
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

func GetLocalIP() (net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			// 检查地址类型是否为 *net.IPNet，并且不是 loopback 地址
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil { // IPv4 地址
					return ipNet.IP, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("本机没有找到有效的局域网 IP 地址")
}
