package routes

import (
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET/api/v1/article
// GET/api/v1/article:id
// POST/api/v1/article
// PATCH/api/v1/article/:id
// DELETE/api/v1/article/:id

type article struct {
	ID uint `json:"id" `
	Title string `json:"title"`
	Body string	`json:"body"`
	Image string	`json:"image"`

}
//binding:"required" คือ การกำหนดให้ใส่เข้ามาทุกครั้ง
type createArticle struct {
	Title string `form:"title" binding:"required"`
	Body string `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
	}
func Serve(r *gin.Engine){
	articles := []article{
		{ID:1, Title:"Title1",Body:"Body1"},
		{ID:2,Title:"Title2",Body:"Body2"},
		{ID:3,Title:"Title3",Body:"Body3"},
		{ID:4,Title:"Title4",Body:"Body4"},
		{ID:5,Title:"Title5",Body:"Body5"},
	}
	articleGroup :=r.Group("/api/v1/article")
	articleGroup.GET("",func(ctx*gin.Context){
		result := articles
		if limit := ctx.Query("limit"); limit != "" {
			n, _ := strconv.Atoi(limit)

			result = result[:n] // return "3"
		}
		ctx.JSON(http.StatusOK,gin.H{"articles":result})
	})
	articleGroup.GET("/:id",func(ctx*gin.Context){
		id,_ :=strconv.Atoi(ctx.Param("id"))
		for _,item := range articles{
			if item.ID == uint(id){
				ctx.JSON(http.StatusOK,gin.H{"articles":item})
				return
			}
		}
		ctx.JSON(http.StatusNotFound,gin.H{	"error":" Article Not found"})
	})
	//func(ctx *gin.Context)  รับข้อมูล
	articleGroup.POST("",func(ctx *gin.Context){
	var form createArticle
	if err:=ctx.ShouldBind(&form); err!= nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}

	a := article{
		ID:    uint(len(articles) + 1),
		Title: form.Title,
		Body:  form.Body,
	}
	//Get file
	file,_ := ctx.FormFile("image")

	//Create Path
	//ID=> 8,upload/articles/8/image.png
	path := "upload/articles/"+strconv.Itoa(int(a.ID))
	os.MkdirAll(path,0755)

	//Upload file
	filename := path +"/" + file.Filename // เอาชื่อของไฟล์
	if err :=ctx.SaveUploadedFile(file,filename);err != nil{
	//...//
	}
	a.Image = os.Getenv("HOST")+"/"+filename

	//Attach File to Article
	articles = append(articles, a)
	ctx.JSON(http.StatusCreated, gin.H{"article": a})
})
}
