package main

// import "bufio"
import "fmt"
import "net"

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func generateDaytime() string {
  return "today"
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
