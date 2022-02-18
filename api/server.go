package api

import (
	"github.com/AllenKaplan/hackpack-go/repo"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AllenKaplan/hackpack-go/models"
)

type Server struct {
	Repo repo.Repository
}

func NewServer(repo repo.Repository) (Server, error) {
	return Server{Repo: repo}, nil
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Todo Tracker",
	})
}

func (s Server) GetTodoHandler(c *gin.Context) {
	var Id uint
	c.Bind(&Id)
	res, err := s.Repo.GetTodo(c.Request.Context(), Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) GetTodosHandler(c *gin.Context) {
	res, err := s.Repo.GetTodos(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) PostTodoHandler(c *gin.Context) {
	var Todo models.Todo
	c.Bind(&Todo)
	res, err := s.Repo.CreateTodo(c.Request.Context(), Todo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) PutTodoHandler(c *gin.Context) {
	var Todo models.Todo
	c.Bind(&Todo)
	res, err := s.Repo.UpdateTodo(c.Request.Context(), Todo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) DeleteTodoHandler(c *gin.Context) {
	var Id uint
	c.Bind(&Id)
	res, err := s.Repo.DeleteTodo(c.Request.Context(), Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}
