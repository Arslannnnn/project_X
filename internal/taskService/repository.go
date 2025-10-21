package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type tasskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tasskRepository{db: db}
}

func (r *tasskRepository) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *tasskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tasskRepository) GetTaskByID(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *tasskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *tasskRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}

func (r *tasskRepository) GetTasksByUserID(userID string) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
