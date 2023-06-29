package main

import (
	"database/sql"
	"flag"
	"fmt"

	h "github.com/geeky-robot/golang-gin-crud/handler"
	r "github.com/geeky-robot/golang-gin-crud/repository"
	s "github.com/geeky-robot/golang-gin-crud/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	dbName = flag.String("dbname", "postgres", "Database name")
	dbHost = flag.String("dbhost", "127.0.0.1", "Database host")
	dbPort = flag.String("dbport", "3000", "Port for the database")
	dbUser = flag.String("dbuser", "postgres", "Database user")
	dbPass = flag.String("dbpass", "admin", "Database password")
)

func main() {
	var db *sql.DB = InitDb(*dbHost, *dbPort, *dbUser, *dbPass, *dbName)
	repo := r.UserDbRepo{Db: db}
	service := &s.UserService{Repo: &repo}
	handler := h.NewUserHandler(service)
	server = gin.Default()
	server.POST("/user", handler.CreateUser)
	server.POST("/users", handler.CreateUsers)
	server.GET("/user/:id", handler.GetUser)
	server.GET("/user", handler.GetUsers)
	server.PUT("/user", handler.UpdateUser)
	server.DELETE("/user/:id", handler.DeleteUser)
	server.Run(":5000")
}

func InitDb(dbHost, dbPort, dbUser, dbPass, dbName string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database connection error", err.Error())
	}
	return db
}
