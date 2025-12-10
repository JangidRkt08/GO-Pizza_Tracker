package main

import (
	"encoding/json"
	"html/template"
	// "net/http"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "github.com/google/cel-go/common/functions"
)

type Config struct{
	Port string
	DbPath string
	SessionSecretKey string

}

func loadConfig() Config {
	return Config{
		Port : GetEnv("PORT","8080"),
		DbPath : GetEnv("DB_PATH","./db/orders.db"),
		SessionSecretKey : GetEnv("SESSION_SECRET_KEY","secret_key"),
	}
}

func GetEnv(key, fallback string) string {
	if value:= os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func loadTemplates(router *gin.Engine) error{
	functions := template.FuncMap{
		"add": func(a, b int) int{return a+b},
		"json": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}

	tmpl, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}	
	router.SetHTMLTemplate(tmpl)
	return nil
}

func setUpsessionStore(db *gorm.DB, secretKey []byte) sessions.Store{
	store := gormsessions.NewStore(db, true, secretKey)
	store.Options(sessions.Options{
		Path: "/",
		MaxAge: 86400,
		HttpOnly: true,
		SameSite: 3,
	})
	return store

}

func SetSessionValue(c *gin.Context, key string, value interface{}) error{
	session := sessions.Default(c)
	session.Set(key, value)
	return session.Save()
}

func GetSessionString(c *gin.Context, key string) string{
	session := sessions.Default(c)
	val :=session.Get(key)
	if val == nil{
		return ""
	}
	str, _ := val.(string)
	return str 

}

func ClearSession(c *gin.Context) error{
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}