package dbv1

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pedafy/pedafy-tasks/src/database"
)

// GetAllTasks return all the assignments found in the database
func (d *DBv1) GetAllTasks() ([]database.Tasks, error) {
	sql := "SELECT * FROM `tasks`"

	resp, err := d.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	tasks := make([]database.Tasks, 0)

	for resp.Next() {
		var curr database.Tasks
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.StatusID, &curr.Title, &curr.Description, &curr.CreatedAt, &curr.LastEdit)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, curr)
	}
	return tasks, nil
}

// GetAllTasksByOrder get all assignments by order
func (d *DBv1) GetAllTasksByOrder(order string) ([]database.Tasks, error) {
	var sql string

	switch order {
	case "status":
		sql = "SELECT * FROM `tasks` ORDER BY `status_id` DESC"
	case "edit":
		sql = "SELECT * FROM `tasks` ORDER BY `last_edit` DESC"
	case "new":
		sql = "SELECT * FROM `tasks` ORDER BY `id` DESC"
	default:
		return d.GetAllTasks()
	}

	resp, err := d.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	tasks := make([]database.Tasks, 0)

	for resp.Next() {
		var curr database.Tasks
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.StatusID, &curr.Title, &curr.Description, &curr.CreatedAt, &curr.LastEdit)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, curr)
	}
	return tasks, nil
}

// GetAllTasksByFilter allows to get every task with a filter,
// when the filter is "status_id" and the value is "2" the function
// will query and return only tasks with a status_id equal to 2
func (d *DBv1) GetAllTasksByFilter(filter, value string) ([]database.Tasks, error) {
	sql := fmt.Sprintf("SELECT * FROM `tasks` WHERE `%s` = %s", filter, value)
	resp, err := d.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	tasks := make([]database.Tasks, 0)

	for resp.Next() {
		var curr database.Tasks
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.StatusID, &curr.Title, &curr.Description, &curr.CreatedAt, &curr.LastEdit)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, curr)
	}
	return tasks, nil
}

// NewTask will add the given task to the database and return it
func (d *DBv1) NewTask(task database.Tasks) (database.Tasks, error) {
	query, err := d.db.Prepare("INSERT INTO `tasks` (creator_id,status_id,title,description,last_edit) VALUES(?,?,?,?,?)")
	if err != nil {
		return database.Tasks{}, err
	}
	defer query.Close()

	result, err := query.Exec(task.CreatorID, task.StatusID, task.Title, task.Description, time.Now())
	if err != nil {
		return database.Tasks{}, err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return database.Tasks{}, err
	}

	newTask, err := d.GetAllTasksByFilter("id", strconv.Itoa(int(newID)))
	if err != nil || len(newTask) == 0 {
		return database.Tasks{}, err
	}
	return newTask[0], nil
}

// ModifyTask the whole task with the given ID and data
func (d *DBv1) ModifyTask(task database.Tasks, ID int) (database.Tasks, error) {
	if d.isInDatabase(ID) == false {
		return database.Tasks{}, nil
	}

	sql := "UPDATE `tasks` SET creator_id=?, status_id=?, title=?, description=?, last_edit=? WHERE id=?"
	query, err := d.db.Prepare(sql)
	if err != nil {
		return database.Tasks{}, err
	}
	defer query.Close()

	_, err = query.Exec(task.CreatorID, task.StatusID, task.Title, task.Description, time.Now(), ID)
	if err != nil {
		return database.Tasks{}, err
	}
	tasks, err := d.GetAllTasksByFilter("id", strconv.Itoa(ID))
	if err != nil || len(tasks) == 0 {
		return database.Tasks{}, err
	}
	return tasks[0], nil
}

// ArchiveTask will directly archive the given task
func (d *DBv1) ArchiveTask(ID int) (database.Tasks, error) {
	if d.isInDatabase(ID) == false {
		return database.Tasks{}, nil
	}

	archiveID, err := d.GetStatusByName("archived")
	if err != nil {
		return database.Tasks{}, err
	}
	sql := "UPDATE `tasks` SET status_id=?, last_edit=? WHERE id=?"
	query, err := d.db.Prepare(sql)
	if err != nil {
		return database.Tasks{}, err
	}
	defer query.Close()

	_, err = query.Exec(strconv.Itoa(archiveID.ID), time.Now(), strconv.Itoa(ID))
	if err != nil {
		return database.Tasks{}, err
	}
	tasks, err := d.GetAllTasksByFilter("id", strconv.Itoa(ID))
	if err != nil {
		return database.Tasks{}, err
	}
	return tasks[0], nil
}

func (d *DBv1) isInDatabase(ID int) bool {
	tasks, _ := d.GetAllTasksByFilter("id", strconv.Itoa(ID))
	if len(tasks) == 0 || tasks[0].ID == 0 {
		return false
	}
	return true
}
