package rpc

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
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
	log.Printf(args.Who)
	reply.Message = "Hello, " + args.Who + "!"
	log.Printf(reply.Message)
	return nil
}

func main() {
	r := mux.NewRouter()
	jsonRPC := rpc.NewServer()
	jsonCodec := json.NewCodec()
	jsonRPC.RegisterCodec(jsonCodec, "application/json")
	jsonRPC.RegisterCodec(jsonCodec, "application/json; charset=UTF-8") // For firefox 11 and other browsers which append the charset=UTF-8
	jsonRPC.RegisterService(new(HelloService), "")
	r.Handle("/api/", jsonRPC)
	http.ListenAndServe(":1111", r)
}
