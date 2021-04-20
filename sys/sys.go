package sys

import (
	"net"
	"os/exec"
	"regexp"
)

/**
 * 获取电脑CPUId
 */
func GetCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "BFEBFBFF000906EA"
	}
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	if len(str) > 11 {
		return str[11:]
	}
	return "BFEBFBFF000906EA"
}

/*
 *获取本机的MAC地址
 */
func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "20:0d:b0:49:1c:1b"
	}
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	return mac
}
