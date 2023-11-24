package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	f "net-cat/client"
)

type Client struct {
	name    string
	message []byte
	conn    net.Conn
}

var History []byte

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Client
	conns      map[net.Conn]string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Client, 10),
		conns:      make(map[net.Conn]string),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	var mutex sync.Mutex
	mutex.Lock()
	go s.acceptLoop(&mutex)
	mutex.Unlock()
	<-s.quitch
	close(s.msgch)
	return nil
}

func (s *Server) getUsername(conn net.Conn) string {
	pengi, err := os.ReadFile("logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write(pengi)
	///////////////get username
	buff := make([]byte, 2048)
	for {
		// bufio.NewScanner(r)
		u, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if u == 1 || !usernameValidity(buff[:u-1]) {
			conn.Write([]byte("Please enter your name:"))
			continue
		}
		username := string(buff[:u])
		username = strings.Trim(username, "\r\n")
		if s.usernameExists(username) {
			conn.Write([]byte("This name already exists, please enter another name:"))
			continue
		}
		s.conns[conn] = username
		return username
	}
}

func usernameValidity(name []byte) bool {
	for _, v := range name {
		if v < 33 || v > 126 {
			return false
		}
	}
	return true
}

func (s *Server) usernameExists(username string) bool {
	for _, name := range s.conns {
		if name == username {
			return true
		}
	}
	return false
}

func (s *Server) acceptLoop(mutex *sync.Mutex) {
	// fmt.Println("Accept Loop")
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		// s.conns[conn] = true /////////add conn
		mutex.Lock()
		s.conns[conn] = ""
		mutex.Unlock()
		if len(s.conns) > 10 {
			conn.Write([]byte("Number of users reached its limit"))
			break
		}
		//////first message pengi
		username := s.getUsername(conn)
		conn.Write(History)
		// username := <-s.msgch
		joinMsg := fmt.Sprintf("%s has joined our chat...\n", username)

		mutex.Lock()
		History = append(History, []byte(joinMsg)...)
		mutex.Unlock()
		s.msgch <- Client{
			name:    username,
			message: []byte(joinMsg),
			conn:    conn,
		}

		go s.broadcast(mutex)

		go s.readLoop(conn, username, mutex)

	}
}

func (s *Server) readLoop(conn net.Conn, username string, mutex *sync.Mutex) {
	for {
		buf := make([]byte, 2048)

		identifier := []byte("[" + time.Now().Format("2006-01-02 15:04:05") + "][" + username + "]:")
		conn.Write(identifier)

		n, err := conn.Read(buf)
		if err != nil {

			leaveMsg := fmt.Sprintf("%s has left our chat...\n", username)

			History = append(History, []byte(leaveMsg)...)

			s.msgch <- Client{
				name:    username,
				message: []byte(leaveMsg),
				conn:    conn,
			}
			mutex.Lock()
			delete(s.conns, conn)
			go s.broadcast(mutex)

			mutex.Unlock()
			break

		}
		mes := strings.TrimSpace(string(buf[:n]))

		if len(mes) == 0 {
			conn.Write([]byte("Please don't send empty messages\n"))
			continue
		}
		msg := append(identifier, buf[:n]...)

		mutex.Lock()
		History = append(History, msg...)
		mutex.Unlock()

		if buf[:n][0] != 10 {
			go s.broadcast(mutex)
		}
		s.msgch <- Client{
			name:    username,
			message: msg,
			conn:    conn,
		}

	}
}

func (s *Server) broadcast(mutex *sync.Mutex) {
	msg := <-s.msgch
	mutex.Lock()
	for conn, username := range s.conns {
		if conn != msg.conn {
			// go func(conn net.Conn) {
			// 	if _, err := conn.Write(msg.message); err != nil {
			// 		fmt.Println("write error:", err)
			// 	}
			// }(conn)
			fmt.Fprint(conn, "\n", string(msg.message), fmt.Sprintf("[%v][%s]:", time.Now().Format("2006-01-02 15:04:05"), username))
		}
	}
	mutex.Unlock()
}

func main() {
	server := NewServer(f.GetPort())
	// go func() {
	// 	for msg := range server.msgch {
	// 		fmt.Printf("received message from connection(%s):%s\n", msg.name, string(msg.message))
	// 	}
	// }()

	log.Fatal(server.Start())
}
