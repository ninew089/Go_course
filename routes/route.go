package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"

)

// GET/api/v1/article
// GET/api/v1/article:id
// POST/api/v1/article
// PATCH/api/v1/article/:id
// DELETE/api/v1/article/:id

type article struct {
	ID uint
	Title string
	Body string
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
}