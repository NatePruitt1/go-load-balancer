package blncsrvr

import ("log"; "net")

func readFromConn(connection net.Conn, read_chan chan []byte, error_chan chan int) {
  defer log.Println("Read connection closed")

  for {
    var buffer []byte = make([]byte, 1024)
    n, err := connection.Read(buffer)
    if err != nil {
      error_chan <- -1
      log.Println("Reading error recieved")
      return
    }

    read_chan <- buffer[:n]
  }
}

func writeChannelToConn(connection net.Conn, write_chan chan []byte, error_chan chan int) {
  defer log.Println("Write connection closed")

  for {
    var buffer []byte = <- write_chan
    
    var written int = 0
    for written < len(buffer) {
      n, err := connection.Write(buffer)
      if err != nil {
        error_chan <- -1
        log.Println("Writing error recieved")
        return
      }
      written += n
    }
  }
}

// handlConn handles a connection from a client to the load balancer
// Takes a connection that has been made, and an endpoint that should be passed
// to.
// Establishes a connection with endpoint, takes packets from connection, 
// passes them to endpoint, recieves from endpoint, and passes to connection. 
func HandleConn(connection net.Conn, endpoint string) {
  log.Printf("Connection opened with %s\n", connection.RemoteAddr().String())

  //Connect to the endpoint
  server, err := net.Dial("tcp", endpoint)
  if err != nil {
    log.Panic(err)
  }

  read_chan := make(chan []byte, 4)
  write_chan := make(chan []byte, 4)
  read_error_code := make(chan int)
  write_error_code := make(chan int)
  
  reading_closed := false
  writing_closed := false

  go readFromConn(connection, read_chan, read_error_code)
  go writeChannelToConn(server, write_chan, write_error_code)

  for !(reading_closed || writing_closed) {
    select {
    case recieved := <- read_chan:
      log.Println(string(recieved))
      write_chan <- recieved
    case <- read_error_code:
      reading_closed = true
      connection.Close()
      server.Close()
    case <- write_error_code:
      writing_closed = true
      connection.Close()
      server.Close()
    default:
      continue;
    }
  }
}
