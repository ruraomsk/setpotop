package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/ruraomsk/setpotop/command"
	"github.com/ruraomsk/setpotop/scp"
	"golang.org/x/crypto/ssh"
)

var conn *ssh.Client
var err error
var Host *string
var PortSSH *int
var TempKey *bool
var User *string
var Password *string

//go:embed tolt5x tobanana
var resource embed.FS

func main() {
	fmt.Println("Программа настройки и обновления potop\nДля получения справки вызовите setpotop -help")
	Host = flag.String("host", "192.168.88.1", "IP адресс МФУ")
	PortSSH = flag.Int("port", 22, "Порт ssh")
	User = flag.String("user", "root", "Имя пользователя")
	Password = flag.String("pass", "root", "Пароль")
	flag.Parse()
	fmt.Printf("ip address :%s new value or enter", *Host)
	var vs string
	fmt.Scanf("%s\n", &vs)
	if vs != "" {
		*Host = vs
	}
	fmt.Printf("user :%s new value or enter", *User)
	var vsu string
	fmt.Scanf("%s\n", &vsu)
	if vsu != "" {
		*User = vsu
	}
	var vsp string
	fmt.Printf("password :%s new value or enter", *Password)
	fmt.Scanf("%s\n", &vsp)
	if vsp != "" {
		*Password = vsp
	}
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
		fmt.Printf("Failed to dial: %s\n", err.Error())
		os.Exit(-1)
	}
	defer conn.Close()
	command.Connection(conn)
	name := command.GetSystem()
	fmt.Printf("Устройство :%s\n", name)
	switch name {
	case "bananapim2ultra":
		command.KillProc("gobanana")
		command.KillProc("potop")
		err = scp.Connection(conn)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteDir("/home/rura")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.CreateDir("/home/rura")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		data, _ := resource.ReadFile("tobanana/gobanana.sh")
		err = scp.WriteFile("/home/rura/gobanana.sh", data, true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		data, _ = resource.ReadFile("tobanana/rc.local")
		err = scp.WriteFile("/etc/rc.local", data, true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = scp.CopyFile("potop", "/home/rura/potop", true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	case "LT5x":
		command.KillProc("potop")
		err = scp.Connection(conn)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteDir("/cache/log")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteDir("/cache/db")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteFile("/root/start")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = command.DeleteFile("/cache/config.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		data, _ := resource.ReadFile("tolt5x/gopotop.sh")
		err = scp.WriteFile("/root/gopotop.sh", data, true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		data, _ = resource.ReadFile("tolt5x/rc.local")
		err = scp.WriteFile("/etc/rc.local", data, true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		err = scp.CopyFile("potop", "/cache/potop", true)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}
	fmt.Println("Конец работы... нажмите Enter...")
	var none string
	fmt.Scanf("%s\n", &none)

}
