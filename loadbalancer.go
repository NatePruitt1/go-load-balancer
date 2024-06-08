package main

import ("log"; "os"; "net"; "encoding/json"; "flag"; "loadbalancer/blncsrvr")

type configuration struct {
  Port string 
  Endpoints []string 
}

func configFromFile(file []byte) *configuration {
  var c *configuration = new(configuration)
  err := json.Unmarshal(file, c)
  if err != nil {
    log.Fatal(err)
  }
  
  log.Println(c)

  return c
}

/*
Main runs a layer 4 load balancer as given by the provided configuration file
Uses configuration file that was opened, compiles it into server parameters,
and spins up a server to run it. 
*/
func main() {
  var file string
  flag.StringVar(&file, "f", "", "-f=<FILEPATH> the filepath to a json file containing the load balancers config")
  flag.Parse()

  //Get the configuration from the arguments

  config_file, err := os.ReadFile(file)
  if err != nil {
    log.Panic(err.Error())
  }

  var config *configuration = configFromFile(config_file)

  //Spin up the server listening loop. Currently will only start one listener
  server, err := net.Listen("tcp", "0.0.0.0:" + config.Port)

  if err != nil {
    log.Panic(err.Error())
  }

  //If the program exits, clean up.
  defer server.Close()
  
  var current_server int = 0
  
  for {
    connection, err := server.Accept()
    if err != nil {
      log.Panic(err.Error())
    }
    
    go blncsrvr.HandleConn(connection, config.Endpoints[current_server])

    current_server += 1
    if current_server >= len(config.Endpoints) {
      current_server = 0
    }
  }
}
