package client

import (
	"fmt"
	"os"
	"strconv"
)

func GetPort() string {
	if len(os.Args) == 2 {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil || len(os.Args[1]) != 4 {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return ""
		}
		return ":" + os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return ""
	}
	return ":8989"
}
