package scpcopy

import (
	"fmt"
	"os"

	"github.com/lkbhargav/go-scp"
	scp2 "github.com/povsister/scp"
	"golang.org/x/crypto/ssh"
)

type ClientCopy struct {
	User            string
	IP              string
	Password        string
	SourceFile      string
	DestinationFile string
}

func CopyFileToServer(clientInfo ClientCopy) {
	conf := &ssh.ClientConfig{
		User: clientInfo.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(clientInfo.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn := scp.NewClient(clientInfo.IP+":22", conf)

	conn.Connect()

	file, err := os.Open(clientInfo.SourceFile)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer file.Close()

	err = conn.CopyFromFile(*file, clientInfo.DestinationFile, "0665")

	if err != nil {
		fmt.Println("not send ", err)
	}

}

func CopyFileToHost(clientInfo ClientCopy) {
	conf := &ssh.ClientConfig{
		User: clientInfo.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(clientInfo.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, _ := ssh.Dial("tcp", clientInfo.IP+":22", conf)

	defer conn.Close()

	session, _ := conn.NewSession()
	defer session.Close()

	conn2, _ := scp2.NewClient(clientInfo.IP+":22", conf, &scp2.ClientOption{})

	err := conn2.CopyFileFromRemote(clientInfo.SourceFile, clientInfo.DestinationFile, &scp2.FileTransferOption{})

	if err != nil {
		panic(err)
	}
}
