# scanMachinePort
扫描指定机器端口是否占用
# Usage
``` go 

package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/zhangjunfang/scanMachinePort"
)

func main() {
	scanMachinePort.ScanPort(1024, 65535)
	scanMachinePort.ScanPreferredMachinePort("0.0.0.0", 1024, 65535)
}

```
