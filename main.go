package main

import (
	"fmt"

	scpcopy "github.maxiputz.com/ssh/sshClient/SCPCopy"
	sshclient "github.maxiputz.com/ssh/sshClient/execute"
)

func main() {
	fmt.Println("hello wolrd")

	scpcopy.CopyFileToServer(scpcopy.ClientCopy{
		User:            "max",
		IP:              "...",
		Password:        "",
		SourceFile:      "forTransfere.txt",
		DestinationFile: "/Users/max/forTransfere.txt",
	})

	sshclient.Execute(sshclient.ClientExecute{
		User:     "max",
		IP:       "...",
		Password: "",
		Command:  "cat forTransfere.txt ",
	})

	scpcopy.CopyFileToHost(scpcopy.ClientCopy{
		User:            "max",
		IP:              "...",
		Password:        "",
		DestinationFile: "/Users/max/go/ssh/forTransfere.txt",
		SourceFile:      "/Users/max/tmpSSHtest.txt",
	})
}
