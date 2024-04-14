package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Clint-Mathews/Stream-Files-Over-TCP/server"
)

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":4000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.LittleEndian, int64(size))
	// n, err := io.Copy(conn, bytes.NewReader(file))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))

	// n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("Written %d bytes over network\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(200000)
	}()
	fmt.Println("Stream files over TCP!")

	fs := &server.FileServer{}
	fs.Start()

}
