package main

import (
	"flag"
	"fmt"
)

var IpMFU *string
var PortSSH *int
var TempKey *bool

func main() {
	fmt.Println("Программа настройки и обновления potop\nДля получения справки вызовите setpotop -help")
	IpMFU = flag.String("ip", "192.168.88.1", "Ip МФУ")
	PortSSH = flag.Int("port", 22, "Порт ssh")
	TempKey = flag.Bool("key", true, "Временный ключ ssh")
	flag.Parse()
}
