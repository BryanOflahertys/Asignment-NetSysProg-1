package main

import (
	"fmt"
	"main/handler"
	"net"
	"time"
)

func main() {
    
    dial, err := net.DialTimeout("tcp", "localhost:6969", 5*time.Second)
    handler.ErrorHandler(err)
    defer dial.Close()

    dial.SetWriteDeadline(time.Now().Add(5 * time.Second))
    dial.SetReadDeadline(time.Now().Add(5 * time.Second))

    message := []byte("Hello from client")
    _, err = dial.Write(message)
    handler.ErrorHandler(err)

    
    buffer := make([]byte, 1024)
    _, err = dial.Read(buffer)
    handler.ErrorHandler(err)

    // Cetak balasan dari server
    fmt.Println("Response:", string(buffer))
}