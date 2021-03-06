package main

import "bufio"
import "fmt"
import "net"
import "os"
import "math/rand"

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func loadQuotes() []string {
  f, err := os.Open("quotes")
  checkError(err)
  defer f.Close()

  var quotes []string
  quotes = make([]string, 0, 5)
  s := bufio.NewScanner(bufio.NewReader(f))

  for s.Scan() {
    quotes = append(quotes, s.Text())
  }
  return quotes
}

func handleConnection(c net.Conn, quotes []string) {
  // select a random quote; a quote of the moment, instead of the day
  q := quotes[rand.Intn(len(quotes))]
  // write the quote to the connection
  c.Write([]byte(q))
  c.Close()
}

func main() {
  l, err := net.Listen("tcp", ":17")
  checkError(err)
  defer l.Close()
  fmt.Println("port connected")

  quotes := loadQuotes()
  fmt.Println("quotes loaded; count =", len(quotes))

  fmt.Println("server started")

  for {
    c, err := l.Accept()
    checkError(err)

    go handleConnection(c, quotes)
  }
}
