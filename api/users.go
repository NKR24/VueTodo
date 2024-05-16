package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID      uuid.UUID `json:"id"`
	NOTE    string    `json:"note"`
	COMPETE bool      `json:"compelete"`
}

func createTodo(c echo.Context) error {
	newTodo := new(Todo)
	newTodo.ID = uuid.New()
	c.Bind(newTodo)
	key := fmt.Sprintf("Todos:%s", newTodo.ID)
	rh.JSONSet(key, ".", newTodo)
	fmt.Println(newTodo)
	return c.JSON(http.StatusCreated, newTodo.ID)
}

func getTodo(c echo.Context) error {
	id := c.Param("id")
	key := fmt.Sprintf("Todos:%s", id)
	todoJSON, err := redis.Bytes(rh.JSONGet(key, "."))
	todo := new(Todo)
	json.Unmarshal(todoJSON, &todo)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, &todo)
}

func deleteTodo(c echo.Context) error {
	id := c.Param("id")
	key := fmt.Sprintf("Todos:%s", id)
	_, err := rh.JSONDel(key, ".")
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.NoContent(http.StatusOK)
}

func putTodo(c echo.Context) error {
	id := c.Param("id")
	key := fmt.Sprintf("Todos:%s", id)
	newTodo := new(Todo)
	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	rh.JSONSet(key, ".", newTodo)
	return c.JSON(http.StatusOK, newTodo.NOTE)
}
