package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()                    //１：Defaultルーター　Httpsリクエストを処理し適切なハンドラ関数にルーティング
	r.GET("/ping", func(c *gin.Context) { //２：エンドポイントのパス、リクエストの処理関数
		c.JSON(200, gin.H{ //ステータスコード、レスポンスボディ
			"message": "pong", //キーと値(map関数)
		})
	}) //エンドポイントの追加（GETやPOST
	r.Run("localhost:8080") //３：サーバーのアドレスやポート番号を指定する
}
