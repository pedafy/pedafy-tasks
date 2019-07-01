package database

import "time"

// Tasks refers to the tasks table in the database
type Tasks struct {
	ID          int        `db:"id" json:"id"`
	CreatorID   string     `db:"creator_id" json:"creator_id"`
	StatusID    int        `db:"status_id" json:"status_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	LastEdit    *time.Time `db:"last_edit" json:"last_edit"`
}

// Status represents a task's status
type Status struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
