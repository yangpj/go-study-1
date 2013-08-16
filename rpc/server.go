package main

import (
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
)

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json\r\n")
	s.RegisterCodec(json2.NewCodec(), "application/json; charset=UTF-8")
	s.RegisterCodec(json2.NewCodec(), "application/x-www-form-urlencoded")
	s.RegisterService(new(HelloService), "home")
	http.Handle("/rpc", s)
	http.ListenAndServe(":8080", nil)
}
