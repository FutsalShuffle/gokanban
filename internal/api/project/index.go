package project

import (
	"github.com/gin-gonic/gin"
	"gokanban/internal/db"
	"net/http"
)

func IndexHandler(q *db.Queries) gin.HandlerFunc {
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
