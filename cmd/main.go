package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jangidRkt08/pizza-tracker-go/internal/models"
)

func main(){
	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DbPath)
	if err != nil{
		slog.Error("failed to initialize db","error",err)
		os.Exit(1)
	}

	slog.Info("Database initialized successfully")
	RegistorCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()
	if err := loadTemplates(router); err != nil{
		slog.Error("failed to load templates","error",err)
		os.Exit(1)
	}
	setupRoutes(router,h)
	slog.Info("Server started successfully", "url", "http://localhost:"+cfg.Port)
	router.Run(":"+cfg.Port)
}