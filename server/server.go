package server

import "net"
import "fmt"
import "bufio"
import "strings"

func Start() {
	s, err := net.Listen("tcp", ":3002")
	if s == nil {
		panic("couldn't start listening: " + err.Error())
	}
	conns := clientConns(s)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	client.Write([]byte("WELCOME TO HDT, STATE YOUR VERSION\n"))
	lastLine := ""
	i := 0
	b := bufio.NewReader(client)
	for {
		if i == 1 {
			if lastLine == "1" {
				client.Write([]byte("I THINK THE LONGEST CHAIN IS BLAH\n"))
			} else {
				client.Write([]byte("PLEASE UPGRADE\n"))
				break
			}
		}
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		lastLine = strings.TrimSpace(string(line))
		i++
	}
}
