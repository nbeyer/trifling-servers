package main

import "net"
import "fmt"

func main() {
  l, err := net.Listen("tcp", ":17")
  if err != nil {
    panic(err)
  }

  fmt.Println("server started")

  defer l.Close()

  for {
    c, err := l.Accept()
    if err != nil {
      panic(err)
    }
    fmt.Println("Connection made ...")
    c.Write([]byte("Insert witty quote here."))
    c.Close()
    fmt.Println("Connection closed ...")
  }
}
