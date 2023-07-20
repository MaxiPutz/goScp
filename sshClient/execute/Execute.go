package sshclient

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type ClientExecute struct {
	User     string
	IP       string
	Password string
	Command  string
}

func Execute(exeInfo ClientExecute) (string, error) {
	conf := &ssh.ClientConfig{
		User: exeInfo.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(exeInfo.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", exeInfo.IP+":22", conf)
	if err != nil {
		fmt.Print("Failed to dial SSH server: %s", err)
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("not possible to connect")
		return "", err
	}

	var stdoutBuf bytes.Buffer

	session.Stdout = &stdoutBuf

	session.Stdin = os.Stdin
	err = session.Run(exeInfo.Command)

	if err != nil {
		fmt.Println("command not run")
		return "", err
	}

	return stdoutBuf.String(), nil
}
