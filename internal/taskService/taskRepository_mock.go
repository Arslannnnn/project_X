package taskService

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository - поддельный репозиторий для тестирования
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks() ([]Task, error) {
	args := m.Called()
	return args.Get(0).([]Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id string) (Task, error) {
	args := m.Called(id)
	return args.Get(0).(Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
