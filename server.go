package main

import "fmt"
import "bufio"
import "os"
import "net"
//import "io/ioutil"
import "time"

func main(){
  port := ""

  if len(os.Args) > 1 {
    port = os.Args[1]
  } else {
    port = "5454"
  }

  socket, err := net.Listen("tcp", ":" + port)

  for err != nil {
    fmt.Println("Cannot listen at port. Retrying...")
    time.Sleep(500 * time.Millisecond)
    socket, err = net.Listen("tcp", ":" + port)
  }

  fmt.Println("Listening to port " + port + "!...")

  for{
        conn, err := socket.Accept()
        if err != nil {
            fmt.Println("Connection failed :(")
            continue
        } else {
          fmt.Println("Recieved connection!")
          go printRequest(conn)
          go respond(conn)
        }
    }
}


func printRequest(conn net.Conn){
  packet, err := bufio.NewReader(conn).ReadBytes('\n')
  if err != nil {
    fmt.Println("Error in reading request")
  } else {
    fmt.Println(packet)
    fmt.Println(string(packet)[:len(packet)-1])
  }
}

func respond(conn net.Conn){
  defer conn.Close()
  time.Sleep(200 * time.Millisecond)
  data := "200 OK Hello World!"
  conn.Write([]byte(string(data)))
}
