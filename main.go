package main

import (
	"context"
	"fmt"
	"instasafeCodingChallenge/router"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := &sync.Mutex{}
	routerV := router.NewRouterVar(m)
	router.APIMultiplexer(ctx)

	ListenAndServe("9000", routerV)
}

func ListenAndServe(port string, router *gin.Engine) {
	fmt.Println("Listening on:", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
