package main

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/google/cel-go/common/functions"
)

type Config struct{
	Port string
	DbPath string

}

func loadConfig() Config {
	return Config{
		Port : GetEnv("PORT","8080"),
		DbPath : GetEnv("DB_PATH","./db/orders.db"),
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
	}

	tmpl, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}	
	router.SetHTMLTemplate(tmpl)
	return nil
}
