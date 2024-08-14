package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gokanban/internal/api/project"
	q "gokanban/internal/db"
	"time"
)

type ProjectListResponse struct {
	Data []q.ListProjectResult `json:"data"`
}

func main() {
	//router := chi.NewMux()
	//api := humachi.New(router, huma.DefaultConfig("GoKanban", "0.0.1"))
	db, err := getDb()
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(time.Hour)
	if err != nil {
		panic("db connection error")
	}
	db.Ping()
	queries := q.New(db)
	//huma.Get(api, "/api/project", func(ctx context.Context, input *struct {
	//	All int
	//}) (*ProjectListResponse, error) {
	//	result, err := queries.GetListProjects(ctx, 56)
	//	response := &ProjectListResponse{}
	//	response.Body = result
	//
	//	return response, err
	//})
	//
	//http.ListenAndServe("0.0.0.0:8888", router)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//pprof.Register(r)
	r.GET("/api/project", project.IndexHandler(queries))
	r.Run("0.0.0.0:8888")
}

func getDb() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "dev",
		Passwd: "dev",
		Net:    "tcp",
		Addr:   "127.0.0.1:33060",
		DBName: "dev",
	}
	return sql.Open("mysql", cfg.FormatDSN())
}
