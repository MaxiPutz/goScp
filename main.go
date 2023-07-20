package main

import (
	"fmt"
	"strings"

	sshclient "github.maxiputz.com/ssh/sshClient/execute"
)

func main() {
	fmt.Println("hello wolrd")

	// scpcopy.CopyFileToServer(scpcopy.ClientCopy{
	// 	User:            "max",
	// 	IP:              "...",
	// 	Password:        "",
	// 	SourceFile:      "forTransfere.txt",
	// 	DestinationFile: "/Users/max/forTransfere.txt",
	// })

	tmp, _ := sshclient.Execute(sshclient.ClientExecute{
		User:     "max",
		IP:       "0.0.0.0",
		Password: "",
		Command:  "cat forTransfere.txt ",
	})

	arr := []string{}
	arr = append(arr, tmp)

	arr = Map[string, string](arr, func(ele string) string {
		return strings.ToUpper(ele)
	})

	fmt.Printf("arr: %v\n", arr)

	// scpcopy.CopyFileToHost(scpcopy.ClientCopy{
	// 	User:            "max",
	// 	IP:              "...",
	// 	Password:        "",
	// 	DestinationFile: "/Users/max/go/ssh/forTransfere.txt",
	// 	SourceFile:      "/Users/max/tmpSSHtest.txt",
	// })
}

func Map[K any, V any](a []K, function func(K) V) []V {
	out := make([]V, len(a))

	for i, ele := range a {
		out[i] = function(ele)
	}

	return out
}
