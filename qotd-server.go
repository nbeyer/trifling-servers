package main

import "bufio"
import "fmt"
import "net"
import "os"

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  l, err := net.Listen("tcp", ":17")
  checkError(err)
  defer l.Close()
  fmt.Println("port connected")

  f, err := os.Open("quotes")
  checkError(err)
  defer f.Close()

  var quotes []string
  quotes = make([]string, 0, 5)
  s := bufio.NewScanner(bufio.NewReader(f))

  for s.Scan() {
    quotes = append(quotes, s.Text())
  }
  fmt.Println("quotes loaded; count =", len(quotes))

  fmt.Println("server started")

  for {
    c, err := l.Accept()
    if err != nil {
      panic(err)
    }
    fmt.Println("connection made ...")
    c.Write([]byte(quotes[0]))
    c.Close()
    fmt.Println("connection closed ...")
  }
}
