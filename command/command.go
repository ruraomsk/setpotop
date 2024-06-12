package command

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

var conn *ssh.Client

func Connection(c *ssh.Client) {
	conn = c
}
func KillProc(name string) error {
	fmt.Printf("ps -e | grep %s\n", name)
	res := make([]string, 0)
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(fmt.Sprintf("ps -e | grep %s", name))
	scanner := bufio.NewScanner(bytes.NewReader(b.Bytes()))
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimLeft(s, " ")
		bs := strings.Split(s, " ")
		res = append(res, bs[0])
	}
	for _, v := range res {
		kill(v)
	}
	return nil
}
func kill(kp string) {
	fmt.Printf("kill %s\n", kp)
	session, err := conn.NewSession()
	if err != nil {
		return
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(fmt.Sprintf("kill -9 %s\n", kp))
}
func DeleteFile(path string) error {
	fmt.Printf("rm %s\n", path)
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(fmt.Sprintf("rm %s", path))
	return nil
}
func CreateDir(path string) error {
	fmt.Printf("mkdir %s\n", path)
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(fmt.Sprintf("mkdir %s", path)); err != nil {
		fmt.Printf("Failed to run: %s\n", err.Error())
		return err
	}
	return nil
}

func DeleteDir(path string) error {
	fmt.Printf("rm -r %s\n", path)
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(fmt.Sprintf("rm -r %s", path))
	return nil
}

func GetSystem() string {
	session, err := conn.NewSession()
	if err != nil {
		fmt.Printf("Failed new session %s\n", err.Error())
		return ""
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("uname -a"); err != nil {
		fmt.Printf("Failed to run: %s\n", err.Error())
		return ""
	}
	bs := b.String()
	bbs := strings.Split(bs, " ")
	return bbs[1]
}
