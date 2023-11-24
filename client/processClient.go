package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func ProcessClient(connection net.Conn) {
	r := bufio.NewReader(os.Stdin)
	username, err := r.ReadString('\n')
	// _, err := io.Copy(connection, r)
	// _, err = io.Copy(connection, connection)
	if err != nil {
		log.Fatal(err)
	}
	username = strings.Trim(username, "\r\n")

	welcomeMsg := fmt.Sprintf("Welcome %s.\n", username)
	connection.Write([]byte(welcomeMsg))

	defer connection.Close()
}
