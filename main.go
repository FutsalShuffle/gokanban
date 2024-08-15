package main

import (
	"github.com/gin-gonic/gin"
	mysqld "github.com/go-sql-driver/mysql"
	"gokanban/internal/api/project"
	"gokanban/internal/api/task"
	q "gokanban/internal/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ProjectListResponse struct {
	Data []q.ListProjectResult `json:"data"`
}

func main() {
	//router := chi.NewMux()
	//api := humachi.New(router, huma.DefaultConfig("GoKanban", "0.0.1"))
	dbg, err := GetDb()
	db, err := dbg.DB()
	defer db.Close()
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
	r.GET("/api/project", project.IndexHandler(queries, dbg))
	r.GET("/api/project/:slug/task", task.Index(dbg, queries))
	r.GET("/api/project/:slug", project.Show(dbg))
	r.Run("0.0.0.0:8888")
}

func GetDbDSN() string {
	config := mysqld.Config{
		User:   "dev",
		Passwd: "dev",
		Net:    "tcp",
		Addr:   "127.0.0.1:33060",
		DBName: "dev",
	}
	return config.FormatDSN()
}

func GetDb() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(GetDbDSN() + "&parseTime=True"))
}
