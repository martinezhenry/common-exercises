package internal

type UserRepository interface {
	GetUserByID(id string) (User, error)
	SetUser(user User) error
	DeleteUser(id string) error
	GetAllUsers() ([]User, error)
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int   `json:"age"`
}

type userRepository struct {
	storage Storage[string, User]
}

func NewUserRepository(storage Storage[string, User]) UserRepository {
	return &userRepository{storage: storage}
}

func (r *userRepository) GetUserByID(id string) (User, error) {
	user, err := r.storage.Get(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) SetUser(user User) error {
	return r.storage.Set(user.ID, user)
}

func (r *userRepository) DeleteUser(id string) error {
	return r.storage.Delete(id)
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	users, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}