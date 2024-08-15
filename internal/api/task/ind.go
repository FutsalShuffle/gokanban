package task

import (
	"github.com/gin-gonic/gin"
	"gokanban/internal/db"
	"gokanban/internal/domain/models"
	"gokanban/internal/domain/models/model"
	"gokanban/internal/domain/services"
	"gorm.io/gorm"
	"net/http"
)

type TaskListResult struct {
	Id                     uint    `json:"id"`
	Name                   string  `json:"name"`
	CreatedById            uint    `json:"created_by_id"`
	PriorityId             uint    `json:"priority_id"`
	StageId                uint    `json:"stage_id"`
	ComponentId            uint    `json:"component_id"`
	TaskTypeId             uint    `json:"task_type_id"`
	Rank                   string  `json:"rank"`
	EpicName               string  `json:"epic_name"`
	PossibleTaskNextStages []int32 `json:"possibleTaskNextStages"`
}

func Index(orm *gorm.DB, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project *models.Project
		orm.Model(model.Project{}).WithContext(c).Where("slug = ?", c.Param("slug")).Find(&project)
		rows, err := orm.Model(model.Task{}).Where("tasks.project_id = ?", project.Id).
			Where("tasks.deleted_at is null").
			Where("tasks.stage_id != ?", models.END).
			Where("tasks.stage_id != ?", models.BACKLOG).
			Where("tasks.stage_id != ?", models.RELEASE).
			Select("tasks.id", "tasks.name", "tasks.created_by_id", "tasks.priority_id", "tasks.stage_id", "tasks.component_id", "tasks.task_type_id", "tasks.rank", "e.name").
			Joins("LEFT JOIN tasks e ON e.id = tasks.epic_id").
			Rows()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		items := make(map[uint]TaskListResult)
		projectFlow := services.GetFlowBySlug(project.Workflow)
		user := &models.User{IsAdmin: true, Id: 56}

		for rows.Next() {
			var item TaskListResult
			rows.Scan(&item.Id, &item.Name, &item.CreatedById, &item.PriorityId, &item.StageId, &item.ComponentId, &item.TaskTypeId, &item.Rank, &item.EpicName)
			item.PossibleTaskNextStages = projectFlow.PossibleTaskNextStages(&models.Task{
				StageId:     int32(item.StageId),
				TaskTypeId:  int32(item.TaskTypeId),
				ComponentId: int32(item.ComponentId),
			}, user, project,
			)
			items[item.Id] = item
		}
		rows.Close()

		c.JSON(http.StatusOK, gin.H{
			"data": items,
		})
	}
}
