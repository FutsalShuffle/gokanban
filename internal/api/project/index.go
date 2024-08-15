package project

import (
	"github.com/gin-gonic/gin"
	"gokanban/internal/db"
	"gokanban/internal/domain/models/model"
	"gorm.io/gorm"
	"net/http"
)

func IndexHandler(q *db.Queries, orm *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := q.GetListProjects(c, 56)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func Show(orm *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project model.Project
		err := orm.Model(&model.Project{}).
			Where("slug = ?", c.Param("slug")).
			Preload("Users").
			Find(&project).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": project,
		})
	}
}
