package main

import (
	"course-go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err:=godotenv.Load()
	if err!= nil{
		log.Fatal("Failed to load .env file")
	}
	r := gin.Default()
	r.Static("/upload","./upload")
	uploadDirs :=[...]string{"articles","users"}
	for _,dir := range uploadDirs{
		os.MkdirAll("upload"+dir,0755)
	}
	routes.Serve(r)
	r.Run(":"+os.Getenv("PORT"))
}