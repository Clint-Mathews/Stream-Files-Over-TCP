package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

type FileServer struct{}

func (fs *FileServer) Start() {
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		// n, err := io.Copy(buf, conn)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buf.Bytes())
		fmt.Printf("Received %d bytes over network\n", n)
	}
}
