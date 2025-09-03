package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    fmt.Print("Enter message to send: ")
    input := bufio.NewReader(os.Stdin)
    msg, _ := input.ReadString('\n')

    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    conn.Write([]byte(msg))

    reply := bufio.NewReader(conn)
    response, _ := reply.ReadString('\n')
    fmt.Printf("Server reply: %s", response)
}