package main

import (
	"net"
	"context"
	"idefinitive"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
type server struct{}
func main(){
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	math.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	} 
}
func (S *server) AddService(ctx context.Context, request *math.Request ) (*math.Response,error){
	fnum,snum := request.GetFirstnumber(),request.GetSecondnumber()
	result := fnum + snum
	return &math.Response{Result:result}, nil
}


