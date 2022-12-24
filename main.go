package main

import (
	"context"
	"fmt"
	"instasafeCodingChallenge/location"
	"instasafeCodingChallenge/router"
	"instasafeCodingChallenge/transaction"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := &sync.Mutex{}
	routerV := router.NewRouterVar(m)

	transactionDbService := transaction.NewTransactionDbService()
	transactionService := transaction.NewTransactionService(transactionDbService)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	locationDbService := location.NewLocationDbService()
	locationService := location.NewLocationService(locationDbService)
	locationHandler := location.NewLocationHandler(locationService)

	apiSvc := router.NewApi(transactionHandler, locationHandler)
	router.APIMultiplexer(ctx, apiSvc)

	ListenAndServe("9000", routerV)
}

func ListenAndServe(port string, router *gin.Engine) {
	fmt.Println("Listening on:", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
