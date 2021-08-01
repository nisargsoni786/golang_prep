package main
import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"time"
	"rpcobjects"
)

// https://www.educative.io/courses/the-way-to-go/my4mprqBVQG

func main() {
	calc := new(rpcobjects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}