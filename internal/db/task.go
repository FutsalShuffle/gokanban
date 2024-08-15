package db

import (
	"context"
)

type TaskListResult struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	CreatedById uint   `json:"created_by_id"`
	PriorityId  uint   `json:"priority_id"`
	StageId     uint   `json:"stage_id"`
	ComponentId uint   `json:"component_id"`
	TaskTypeId  uint   `json:"task_type_id"`
	Rank        string `json:"rank"`
}

const taskListQuery = "SELECT `tasks`.`id`,`tasks`.`name`,`tasks`.`created_by_id`," +
	"`tasks`.`priority_id`,`tasks`.`stage_id`,`tasks`.`component_id`," +
	"`tasks`.`task_type_id`,`tasks`.`rank` FROM `tasks`" +
	" WHERE project_id = ? " +
	"AND deleted_at is null " +
	"AND stage_id != 100 " +
	"AND stage_id != 101 " +
	"AND stage_id != 8 " +
	"AND `tasks`.`deleted_at` IS NULL"

func (q *Queries) GetListTasks(ctx context.Context, projectId int) ([]TaskListResult, error) {
	rows, err := q.db.QueryContext(ctx, taskListQuery, projectId)
	if err != nil {
		return nil, err
	}

	var items []TaskListResult
	for rows.Next() {
		var item TaskListResult
		rows.Scan(&item.Id, &item.Name, &item.CreatedById, &item.PriorityId, &item.StageId, &item.ComponentId, &item.TaskTypeId, &item.Rank)
		items = append(items, item)
	}
	rows.Close()

	return items, err
}
