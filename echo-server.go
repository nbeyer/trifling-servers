package main

import "fmt"
import "net"

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  l, err := net.Listen("tcp", ":7")
  checkError(err)
  defer l.Close()
  fmt.Println("port connected")

  fmt.Println("server started")

  for {
    c, err := l.Accept()
    checkError(err)

    c.Write([]byte("stub"))
    c.Close()
  }
}
