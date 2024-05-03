package gin

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() RequestHandler {
	engine := gin.Default()
	engine.LoadHTMLGlob("public/*")
	engine.Static("/public", "public")

	return RequestHandler{
		Gin: engine,
	}
}
