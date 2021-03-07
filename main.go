package main
import (
	"github.com/gin-gonic/gin"
	"course-go/routes"
)

func main(){
	r := gin.Default()
	routes.Serve(r)
	r.Run()
}