package main

import (
	"fmt"
	"net/http"
	// "github.com/BlockEater22/ChatApplication/pkg/websocket"
)


func main() {
	fmt.Println("hello!!!")
	setupRoutes()
	http.ListenAndServe(":9000",nil)
	
}

func setupRoutes(){
	
}