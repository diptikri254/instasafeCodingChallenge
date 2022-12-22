package router

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var routerVar *gin.Engine

func NewRouterVar(m *sync.Mutex) *gin.Engine {
	if routerVar == nil {
		m.Lock()
		if routerVar == nil {
			routerVar = gin.New()
		}
		m.Unlock()
	}
	return routerVar
}
