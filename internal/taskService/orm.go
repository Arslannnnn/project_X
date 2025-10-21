package taskService

type Task struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Text   string  `json:"task"`
	IsDone bool    `json:"is_done"`
	UserID *string `gorm:"user_id" json:"user_id"`
}

type TaskRequest struct {
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
