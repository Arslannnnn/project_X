package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TaskRequest struct {
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
