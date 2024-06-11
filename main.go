package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ruraomsk/setpotop/command"
	"github.com/ruraomsk/setpotop/copyer"
	"golang.org/x/crypto/ssh"
)

var conn *ssh.Client
var err error
var Host *string
var PortSSH *int
var TempKey *bool
var User *string
var Password *string

func main() {
	fmt.Println("Программа настройки и обновления potop\nДля получения справки вызовите setpotop -help")
	Host = flag.String("host", "172.16.58.10", "IP адресс МФУ")
	PortSSH = flag.Int("port", 22, "Порт ssh")
	User = flag.String("user", "root", "Имя пользователя")
	Password = flag.String("pass", "12345678", "Пароль")
	flag.Parse()
	fmt.Printf("Производим настройку %s:%d user \"%s\" password \"%s\"\n", *Host, *PortSSH, *User, *Password)
	config := &ssh.ClientConfig{
		User: *User,
		Auth: []ssh.AuthMethod{
			ssh.Password(*Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", *Host, *PortSSH), config)
	if err != nil {
		fmt.Printf("Failed to dial: %s", err.Error())
	}
	defer conn.Close()
	command.Connection(conn)
	err = copyer.Connection(conn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	name := command.GetSystem()
	fmt.Println(name)
	switch name {
	case "KPDA":
		err = command.DeleteDir("/tmp/log")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.CreateDir("/tmp/log")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteFile("/tmp/test")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = copyer.WriteFile("/tmp/log/log.txt", []byte("Test log file"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = copyer.CopyFile("pdebugTest", "/tmp/test")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = copyer.Chmod("/tmp/test")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		command.KillProc("Photon")
	}
	fmt.Println("Конец работы")
}
