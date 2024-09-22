package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//1.解析模板
	////函数的主要目的是将多个文件的内容解析成模板，并将这些模板关联在一起
	//如果所有文件都被成功解析，则返回模板指针t 和 nil 作为错误。
	//如果解析过程中发生错误，则返回 nil 和该错误。
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		fmt.Println("解析失败，错误是:", err)
		return
	}
	//2.渲染模板
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("渲染失败，失败原因：", err)
		return
	}

}

func main() {
	ginServer := gin.Default()
	//相应一个页面给前端
	ginServer.LoadHTMLGlob("templates/*")
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "hello.html", gin.H{
			"msg": "后台数据",
		})
	})

	//接收前端传来的数据
	ginServer.GET("user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		fmt.Println(userid)
		username := context.Query("username")
		fmt.Println(username)
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"msg":      "ok",
			"username": username,
		})
	})

	//http.HandleFunc("/", sayHello)
	ginServer.Run(":8080")
}
