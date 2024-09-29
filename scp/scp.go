package scp

import (
	"bytes"
	"context"
	"fmt"
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

var client scp.Client
var err error

func Connection(conn *ssh.Client) error {
	client, err = scp.NewClientBySSH(conn)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(path string, body []byte, exec bool) error {
	fmt.Printf("write file %s %v\n", path, exec)
	bs := bytes.NewReader(body)
	permissons := "0664"
	if exec {
		permissons = "0700"
	}
	err := client.CopyFile(context.Background(), bs, path, permissons)
	return err
}
func CopyFile(src, dst string, exec bool) error {
	fmt.Printf("copy %s to %s\n", src, dst)
	permissons := "0600"
	if exec {
		permissons = "0700"
	}
	file, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	bs := bytes.NewReader(file)
	err = client.CopyFile(context.Background(), bs, dst, permissons)
	return err
}
