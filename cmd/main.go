package main

import (
	"fmt"
	"github.com/AllenKaplan/hackpack-go/api"
	"github.com/AllenKaplan/hackpack-go/repo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//config := repo.Config{
	//	Host:     os.Getenv("POSTGRES_HOST"),
	//	User:     os.Getenv("POSTGRES_USER"),
	//	Password: os.Getenv("POSTGRES_PASSWORD"),
	//	DBName:   os.Getenv("POSTGRES_NAME"),
	//	Port:     os.Getenv("POSTGRES_PORT"),
	//}

	//repo, err := repo.NewDB(config)
	repo := repo.NewFakeDB()
	//if err != nil {
	//	panic(err)
	//}

	s, err := api.NewServer(repo)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", api.PingHandler)
	r.GET("/", api.IndexHandler)
	r.GET("/app/:id", s.GetTodoHandler)
	r.GET("/app", s.GetTodosHandler)
	r.POST("/app", s.PostTodoHandler)
	r.PUT("/app", s.PutTodoHandler)
	r.DELETE("/app", s.DeleteTodoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
