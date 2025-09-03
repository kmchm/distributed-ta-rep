package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    fmt.Println("Server listening on port 8080...")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    msg, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading message:", err)
        return
    }
    fmt.Printf("Received: %s", msg)
    response := "server confirms message: " + msg
    conn.Write([]byte(response))
}