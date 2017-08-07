package server

import "net"
import "fmt"
import "bufio"

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
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
}
