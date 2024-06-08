# Go Balancing Server
This is a very simple layer 4 load balancer, which only accepts TCP connections.
It performs a simple round-robin algorithm on the connections received, and 
performs pass-through load balancing, meaning all requests go through this
server.

## Usage
After downloading, you can compile it with a simple
```go build -buildvcs=false```
and can be run with 
```./loadbalancer -f <CONFIG_FILE>```
