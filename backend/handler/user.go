package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tatsuyaHello-template/model"
	"github.com/tatsuyaHello-template/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var login model.RequestUser

	if err := c.BindJSON(&login); err != nil {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}

	hash, err := repository.GetHashByEmail(login.Email)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "メールアドレスまたはパスワードが間違っています",
		})
		return
	}
	err = passwordVerify(hash, login.Password)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "メールアドレスまたはパスワードが間違っています",
		})
		return
	}

	user, ok := repository.GetUser(login.Email)
	if !ok {
		c.JSON(404, gin.H{
			"message": "メールアドレスまたはパスワードが間違っています",
		})
		return
	}

	// session
	// session := sessions.Default(c)
	// var count int
	// v := session.Get("count")
	// if v == nil {
	// 	count = 0
	// } else {
	// 	count = v.(int)
	// 	count++
	// }
	// session.Set("count", count)
	// session.Save()

	c.JSON(200, gin.H{
		"result":  user,
		"message": "認証に成功しました",
	})
	return
}

func Signup(c *gin.Context) {
	var err error
	var signup model.RequestUser

	if err = c.BindJSON(&signup); err != nil {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}

	_, ok := repository.GetUser(signup.Email)
	if ok {
		c.JSON(404, gin.H{
			"message": "すでにユーザが存在しています",
		})
		return
	}

	hash, err := passwordHash(signup.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "サーバエラーです。時間を置いてからやり直してください",
		})
		return
	}
	signup.Password = hash

	user, err := repository.CreateUser(&signup)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "サーバエラーです。時間を置いてからやり直してください",
		})
		return
	}
	c.JSON(201, gin.H{
		//NOTE
		// DBにハッシュして保存されているパスワードをreturnしているので、実際に使用するときはマスクとかした方が良いかも
		"result":  user,
		"message": "サインアップに成功しました。",
	})
	return
}

func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func passwordVerify(hash *string, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(*hash), []byte(pw))
}
