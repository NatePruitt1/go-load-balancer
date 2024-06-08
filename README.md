# Go Balancing Server
This is a very simple layer 4 load balancer, which only accepts TCP connections.
It performs a simple round-robin algorithm on the connections received, and 
performs pass-through load balancing, meaning all requests go through this
server.

## Usage
After the load balancer has ben downloaded and compiled, it can be run. If it is
not given a configuration file in the format:\
```
{
    "port": <LISTENING PORT>,
    "endpoints": [
        <ENDPOINT IP ONE>,
        <ENDPOINT IP TWO>,
        ...
    ]
}
```
\This json file has to be given to the program as follows:\
```./loadbalancer -f <JSON FILE NAME>```
