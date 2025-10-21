package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(text string, isDone bool, userID string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id, text string, isDone bool) (Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type tasskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &tasskService{repo: r}
}

func (s *tasskService) CreateTask(text string, isDone bool, userID string) (Task, error) {
	task := Task{
		ID:     newID(),
		Text:   text,
		IsDone: isDone,
		UserID: &userID,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *tasskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *tasskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *tasskService) UpdateTask(id, text string, isDone bool) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Text = text
	task.IsDone = isDone

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *tasskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

var newID = func() string {
	return uuid.NewString()
}

func (s *tasskService) GetTasksByUserID(userID string) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
