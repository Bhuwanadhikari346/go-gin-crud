package main

import (
	"github.com/Bhuwanadhikari346/go-gin-crud/routes"
	"github.com/gin-gonic/gin"
)
func main() {
   router := gin.New()
 routes.UserRoute(router)
   router.Run(":8080")
}