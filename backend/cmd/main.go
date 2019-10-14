package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tatsuyaHello-template/handler"
)

func main() {

	// logファイルの設定
	logfile, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("cannot open logfile:%v", err)
	}
	defer logfile.Close()
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.Print("ok?")

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// ルーティングの登録を以下で行う
	router.GET("ping/json", handler.PingJson)
	router.GET("ping/string", handler.PingString)

	router.POST("/login", handler.Login)
	router.POST("/signup", handler.Signup)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	log.Fatal(router.Run(":" + os.Getenv("SERVER_PORT")))
}
