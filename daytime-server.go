package main

import "fmt"
import "net"
import "time"

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func generateDaytime() string {
  t := time.Now().UTC()
  return t.Format(time.RFC3339)
}

func main() {
  l, err := net.Listen("tcp", ":13")
  checkError(err)
  defer l.Close()
  fmt.Println("port connected")

  fmt.Println("server started")

  for {
    c, err := l.Accept()
    checkError(err)

    d := generateDaytime()
    // write the daytime to the connection
    c.Write([]byte(d))
    c.Close()
  }
}
