// HTMLの扱い
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
    // htmlのディレクトリを指定
    engine.LoadHTMLGlob("templates/*")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "message": "おはようせかい！！！！！！",
        })
    })
    engine.Run(":3000")
}