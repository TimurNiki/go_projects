package main

import (
	"github.com/TimurNiki/go_api_tutorial/v5-grpc/services/order/http"
	
)
func main(){
	httpServer:= NewHttpServer(":8000")
	go httpServer.Run()
}