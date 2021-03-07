package main

import (
	"course-go/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.Static("/upload","./upload")
	uploadDirs :=[...]string{"articles","users"}
	for _,dir := range uploadDirs{
		os.MkdirAll("upload"+dir,0755)

	}
	routes.Serve(r)
	r.Run()
}