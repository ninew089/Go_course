# Go-lang

Terminal 
go mod init course-go

Run
Go run main.go

File main.go
package main
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

Time Stamp
date +%s


func main(){
  r := gin.Default()

  r.GET("/",func(ctx*gin.Context){
    ctx.String(http.StatusOK,"HelloWord")
  })
  r.Run()
}


go get -u github.com/cosmtrek/air
Run live version

$(go env GOPATH)/bin/air


Form
https://github.com/go-playground/validator

Upload file
http://permissions-calculator.org/


GORM DB
https://gorm.io/
https://gorm.io/docs/connecting_to_the_database.html

package models

import "github.com/jinzhu/gorm"


type Article struct {
  gorm.Model
  Title   string `gorm:"unique;not null"`
  Excerpt string `gorm:"not null"`
  Body    string `gorm:"not null"`
  Image   string `gorm:"not null"`
}





Migration
https://github.com/go-gormigrate/gormigrate


Get id 
date +%s 


package migrations

import (
  "course-go/models"
  

  "github.com/jinzhu/gorm"
  
  
  "gopkg.in/gormigrate.v1"
  
  
)

func m1615217031CreateArticlesTable() *gormigrate.Migration {


  return &gormigrate.Migration{
  
  
    ID: "1615217031",
    
    
    Migrate: func(tx *gorm.DB) error {
    
    
      return tx.AutoMigrate(&models.Article{}).Error
      
      
    },
    
    
    Rollback: func(tx *gorm.DB) error {
    
    
      return tx.DropTable("articles").Error
      
      
    },
    
    
  }
  
  
}

Pagination



Authen

Hook
https://gorm.io/docs/hooks.html
BeforeSave() ไม่นิยมใช้เเล้ว

Authentication
https://github.com/appleboy/gin-jwt
Cabin

Policy.csv

p, Admin, /api/v1/users, (GET)|(POST)

p, Admin, /api/v1/users/*, (GET)|(PATCH)|(DELETE)

p, Admin, /api/v1/categories, (POST)


p, Admin, /api/v1/categories/*, (PATCH)|(DELETE)


p, (Admin)|(Editor), /api/v1/articles, (POST)


p, (Admin)|(Editor), /api/v1/articles/*, (PATCH)|(DELETE)

