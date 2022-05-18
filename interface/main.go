package main

import (
	"fmt"
	"goweb/Cuiyu"
	"goweb/routes"

	//	"math/big"
	"net/http"
	"goweb/blockchains"
	"github.com/gin-gonic/gin"
	"goweb/runSwap"
)

func main() {

	m := Cuiyu.Testht()
	fmt.Print(m)
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./form.html", "./data/index.html")
	// 获取form表单提交数据
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	})
	//r.StaticFile("/cy", "./data/index.html")
	//r.GET("/show", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)

	//})
	r.Static("/static", "./data")
	routes.Issue(r)
	// var username string
	// var password string
	//获取数据
	// r.POST("/login", func(c *gin.Context) {
	// 	// 方式一
	// 	username = c.PostForm("username")
	// 	password = c.PostForm("password")
	// 	//c.HTML(http.StatusOK, "index.html", nil)
	// 	//r.Static("", "./data")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// 	fmt.Println("username:", username, "password", password)

	// })

	// err1 := blockchains.Charge(10)
	// if err1 != nil {
	// 	return
	// }
	//fmt.Print("fffffffffffffff")
	for i := 0; i<2;i++{
		err := blockchains.Charge(12)
		if err != nil{
			return
		}
	}



	if runSwap.RunSwap() != nil {
		fmt.Println("error")
		return
	}




	r.Run(":9994")
	
	//	r.Run(":9997")
}
