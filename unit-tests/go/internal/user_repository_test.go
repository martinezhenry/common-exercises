package internal

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockDB struct {
	// Mocked database connection
	mock.Mock
}

// Exec implements SQLDatabase.
func (m *MockDB) Exec(query string, args ...any) (sql.Result, error) {
	// Mock the Exec method
	a := m.Called(query, args)
	if a.Error(0) != nil {
		return nil, a.Error(0)
	}

	// Return a mock result
	return a.Get(0).(sql.Result), nil
}

// Get implements SQLDatabase.
func (m *MockDB) Get(dest any, query string, args ...any) error {
	// Mock the Get method
	a := m.Called(dest, query, args)
	if a.Error(0) != nil {
		return a.Error(0)
	}

	// Return the result
	return nil
}

// Select implements SQLDatabase.
func (m *MockDB) Select(dest any, query string, args ...any) error {
	// Mock the Select method
	a := m.Called(dest, query, args)
	if a.Error(0) != nil {
		return a.Error(0)
	}

	// Return the result
	return nil
}

func Test_GetUserByID(t *testing.T) {
	// Given
	// Mock db
	mockDB := &MockDB{}
	// Mock the Get method
	mockDB.On("Get", mock.Anything, "SELECT * FROM users WHERE id = ?", []any{"1"}).
		Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*User)
		user.ID = "1"
		user.Name = "Test User"
		user.Age = 30
	})

	// Create a new user repository
	userRepository := NewUserRepository(mockDB)

	// When
	// Call the GetUserByID method
	user, err := userRepository.GetUserByID("1")

	// Then
	require.NoError(t, err)
	// Check the response
	require.Equal(t, "1", user.ID)
	require.Equal(t, "Test User", user.Name)
	require.Equal(t, 30, user.Age)
	// Check that the Get method was called with the correct arguments
	mockDB.AssertExpectations(t)
}
