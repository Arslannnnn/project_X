package taskService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		isDone    bool
		mockID    string
		mockSetup func(m *MockTaskRepository, task Task)
		wantTask  Task
		wantErr   bool
	}{
		{
			name:   "успешное создание задачи",
			text:   "Test",
			isDone: false,
			mockID: "102ufa",
			mockSetup: func(m *MockTaskRepository, task Task) {
				m.On("CreateTask", task).Return(nil)
			},
			wantTask: Task{
				ID:     "102ufa",
				Text:   "Test",
				IsDone: false,
			},
			wantErr: false,
		},
		{
			name:   "ошибка при создании",
			text:   "Bad task",
			isDone: true,
			mockID: "77mosckow",
			mockSetup: func(m *MockTaskRepository, task Task) {
				m.On("CreateTask", task).Return(errors.New("ошибка базы данных"))
			},
			wantTask: Task{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)

			task := Task{
				ID:     tt.mockID,
				Text:   tt.text,
				IsDone: tt.isDone,
			}

			oldID := newID
			newID = func() string { return tt.mockID }
			defer func() { newID = oldID }()

			tt.mockSetup(mockRepo, task)

			service := NewTaskService(mockRepo)
			result, err := service.CreateTask(tt.text, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantTask, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTask, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		ID        string
		Task      Task
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное получение задачи",
			ID:   "555",
			mockSetup: func(m *MockTaskRepository) {
				task := Task{ID: "555", Text: "Test task", IsDone: false}
				m.On("GetTaskByID", "555").Return(task, nil)
			},
			Task:    Task{ID: "555", Text: "Test task", IsDone: false},
			wantErr: false,
		},
		{
			name: "задача не найдена",
			ID:   "111",
			Task: Task{},
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskByID", "111").Return(Task{}, errors.New("task not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			result, err := service.GetTaskByID(tt.ID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Task, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name      string
		ID        string
		text      string
		isDone    bool
		Task      Task
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name:   "успешное обновление задачи",
			ID:     "555",
			text:   "Updated task",
			isDone: true,
			mockSetup: func(m *MockTaskRepository) {
				existingTask := Task{ID: "555", Text: "Original task", IsDone: false}
				updatedTask := Task{ID: "555", Text: "Updated task", IsDone: true}
				m.On("GetTaskByID", "555").Return(existingTask, nil)
				m.On("UpdateTask", updatedTask).Return(nil)
			},
			Task:    Task{ID: "555", Text: "Updated task", IsDone: true},
			wantErr: false,
		},
		{
			name:   "ошибка при получении задачи",
			ID:     "111",
			text:   "Updated task",
			isDone: true,
			Task:   Task{},
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskByID", "111").Return(Task{}, errors.New("task not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			result, err := service.UpdateTask(tt.ID, tt.text, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Task, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		ID        string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное удаление задачи",
			ID:   "555",
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", "555").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении задачи",
			ID:   "111",
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", "111").Return(errors.New("failed to delete"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			err := service.DeleteTask(tt.ID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
