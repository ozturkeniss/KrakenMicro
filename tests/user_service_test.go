package tests

import (
	"context"
	"testing"
	"time"

	"gomicro/internal/user/model"
	"gomicro/internal/user/service"
)

// MockUserRepository implements repository.UserRepository interface
type MockUserRepository struct {
	users map[uint]*model.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[uint]*model.User),
	}
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	user.ID = uint(len(m.users) + 1)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return nil, nil
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *model.User) error {
	if _, exists := m.users[user.ID]; exists {
		user.UpdatedAt = time.Now()
		m.users[user.ID] = user
		return nil
	}
	return nil
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint) error {
	if _, exists := m.users[id]; exists {
		delete(m.users, id)
		return nil
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		user     *model.User
		wantErr  bool
		checkFields bool
	}{
		{
			name: "valid user",
			user: &model.User{
				Email:     "test@example.com",
				Password:  "password123",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr:     false,
			checkFields: true,
		},
		{
			name: "invalid email",
			user: &model.User{
				Email:     "invalid-email",
				Password:  "password123",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr:     true,
			checkFields: false,
		},
		{
			name: "short password",
			user: &model.User{
				Email:     "test@example.com",
				Password:  "123",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr:     true,
			checkFields: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			repo := NewMockUserRepository()
			userService := service.NewUserService(repo)

			// Execute
			err := userService.CreateUser(context.Background(), tt.user)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateUser() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("CreateUser() unexpected error: %v", err)
				return
			}

			if tt.checkFields {
				// Verify user was created
				created, err := repo.GetByID(context.Background(), tt.user.ID)
				if err != nil {
					t.Errorf("Failed to get created user: %v", err)
					return
				}

				if created == nil {
					t.Error("CreateUser() user was not created")
					return
				}

				// Verify fields
				if created.Email != tt.user.Email {
					t.Errorf("CreateUser() email = %v, want %v", created.Email, tt.user.Email)
				}
				if created.FirstName != tt.user.FirstName {
					t.Errorf("CreateUser() firstName = %v, want %v", created.FirstName, tt.user.FirstName)
				}
				if created.LastName != tt.user.LastName {
					t.Errorf("CreateUser() lastName = %v, want %v", created.LastName, tt.user.LastName)
				}
				if created.Password == tt.user.Password {
					t.Error("CreateUser() password was not hashed")
				}
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	repo := NewMockUserRepository()
	userService := service.NewUserService(repo)

	// Create a test user
	testUser := &model.User{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
	}
	repo.Create(context.Background(), testUser)

	tests := []struct {
		name    string
		userID  uint
		wantErr bool
	}{
		{
			name:    "existing user",
			userID:  1,
			wantErr: false,
		},
		{
			name:    "non-existing user",
			userID:  999,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			user, err := userService.GetUser(context.Background(), tt.userID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetUser() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("GetUser() unexpected error: %v", err)
				return
			}

			if tt.userID == 1 {
				if user == nil {
					t.Error("GetUser() returned nil user for existing ID")
					return
				}
				if user.ID != tt.userID {
					t.Errorf("GetUser() user ID = %v, want %v", user.ID, tt.userID)
				}
			} else {
				if user != nil {
					t.Error("GetUser() returned user for non-existing ID")
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup
	repo := NewMockUserRepository()
	userService := service.NewUserService(repo)

	// Create a test user
	testUser := &model.User{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
	}
	repo.Create(context.Background(), testUser)

	tests := []struct {
		name    string
		user    *model.User
		wantErr bool
	}{
		{
			name: "valid update",
			user: &model.User{
				ID:        1,
				Email:     "updated@example.com",
				FirstName: "Jane",
				LastName:  "Smith",
			},
			wantErr: false,
		},
		{
			name: "non-existing user",
			user: &model.User{
				ID:        999,
				Email:     "nonexisting@example.com",
				FirstName: "None",
				LastName:  "Existing",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			user: &model.User{
				ID:        1,
				Email:     "invalid-email",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			err := userService.UpdateUser(context.Background(), tt.user)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("UpdateUser() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("UpdateUser() unexpected error: %v", err)
				return
			}

			if tt.user.ID == 1 {
				// Verify user was updated
				updated, err := repo.GetByID(context.Background(), tt.user.ID)
				if err != nil {
					t.Errorf("Failed to get updated user: %v", err)
					return
				}

				if updated == nil {
					t.Error("UpdateUser() user was not found")
					return
				}

				// Verify fields
				if updated.Email != tt.user.Email {
					t.Errorf("UpdateUser() email = %v, want %v", updated.Email, tt.user.Email)
				}
				if updated.FirstName != tt.user.FirstName {
					t.Errorf("UpdateUser() firstName = %v, want %v", updated.FirstName, tt.user.FirstName)
				}
				if updated.LastName != tt.user.LastName {
					t.Errorf("UpdateUser() lastName = %v, want %v", updated.LastName, tt.user.LastName)
				}
			}
		})
	}
} 