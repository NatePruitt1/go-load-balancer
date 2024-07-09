# Go Balancing Server
This is a very simple layer 4 load balancer, which only accepts TCP connections.
It performs a simple round-robin algorithm on the connections received, and 
performs pass-through load balancing, meaning all packets go through this
server.

## Usage
The command to use the load balancer looks like this:  
```./loadbalancer -f <CONFIG FILE NAME>```  
The config file must have the following attributes in JSON format:  
```
{
    "port": <LISTENING PORT>,
    "endpoints": [
        <ENDPOINT IP:PORT>,
        <ENDPOINT IP:PORT>,
        ...
    ]
}
```
