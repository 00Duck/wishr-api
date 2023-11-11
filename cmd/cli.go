package cmd

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/00Duck/wishr-api/database"
)

var cmdSocketPath = "./sockets/cmd.sock"

type CLI struct {
	listener net.Listener
	db       *database.DB
}

func NewCLI(db *database.DB) *CLI {
	os.MkdirAll("./sockets", os.ModePerm)
	os.Remove(cmdSocketPath)
	sock, err := net.Listen("unix", cmdSocketPath)
	if err != nil {
		log.Fatal("Could not open command socket - dying :(.")
	}
	return &CLI{
		listener: sock,
		db:       db,
	}
}

func (c *CLI) Start() {
	for {
		conn, err := c.listener.Accept()
		if err != nil {
			log.Println("CLI Error: " + err.Error())
		}

		go func(conn net.Conn) {
			defer conn.Close()
			buf := make([]byte, 4096)

			//read incoming data into buffer
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("Error reading input: " + err.Error())
				return
			}

			//convert to a string by reading from the start to the length of the input. Trim excess spaces
			argStr := strings.TrimSpace(string(buf[:n]))
			msg, err := c.dispatch(strings.Split(argStr, " "))
			if err != nil {
				conn.Write([]byte(err.Error()))
				return
			}

			if msg != "" {
				_, err = conn.Write([]byte(msg + "\n"))
				if err != nil {
					log.Println("Error writing output: " + err.Error())
				}
			}

		}(conn)
	}
}

func (c *CLI) Exit() {
	os.Remove(cmdSocketPath)
}
