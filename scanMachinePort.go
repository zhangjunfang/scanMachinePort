package scanMachinePort

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

/**
  描述：
     默认扫描本机端口

  参数说明：
  startPort  扫描开始端口 包含这个端口
  endPort    扫描结束端口 不包含这个端口

*/
func ScanPreferredMachinePort(ip string, startPort, endPort int64) []int64 {
	usedPort := make([]int64, 0)
	for startPort < endPort {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, startPort))
		startPort = startPort + 1
		if err != nil {
			if value, ok := err.(*net.OpError); ok {
				port, err := strconv.Atoi(strings.Split(value.Error(), ":")[1])
				if err != nil {
					continue
				}
				usedPort = append(usedPort, int64(port))
				//fmt.Println(port)
			}
			continue
		}
		if ln == nil {
			continue
		}
		ln.Close()
	}
	return usedPort
}

/**
  描述：
     默认扫描本机端口

  参数说明：
  startPort  扫描开始端口 包含这个端口
  endPort    扫描结束端口 不包含这个端口

*/
func ScanPort(startPort, endPort int64) []int64 {
	usedPort := make([]int64, 0)
	for startPort < endPort {
		ln, err := net.Listen("tcp", "0.0.0.0:"+strconv.FormatInt(int64(startPort), 10))
		startPort = startPort + 1
		if err != nil {
			if value, ok := err.(*net.OpError); ok {
				port, err := strconv.Atoi(strings.Split(value.Error(), ":")[1])
				if err != nil {
					continue
				}
				usedPort = append(usedPort, int64(port))
				//fmt.Println(port)
			}
			continue
		}
		if ln == nil {
			continue
		}
		ln.Close()
	}
	return usedPort
}
