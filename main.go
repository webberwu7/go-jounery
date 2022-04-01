package main

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"go-journey/hello"
	"go-journey/user"
	"go-journey/database"
	"github.com/jmoiron/sqlx"
)

func InitDb() *sqlx.DB {
	db := database.SqlConfig()
	return db
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		hello.PrintHello()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/count", func(c *gin.Context) {
		var a int64 = 0
		var i int64
		for i = 1; i <= 1000000000; i++ {
			a += i
		}
		c.JSON(200, gin.H{"answer": a})
	})
	router.GET("/user/:id", func(c *gin.Context){
		db := InitDb()
		userRepo := user.RepoInterface(db)
		userId, err := strconv.Atoi(c.Param("id"))
		user, err := userRepo.GetUser(userId)

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{"answer": user})
	})
	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
