package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

type PostgresUserRepo struct{}

func (r *PostgresUserRepo) FindByID(id int) (*User, error) {
	return &User{ID: id, Name: "User from Postgres"}, nil
}
func (r *PostgresUserRepo) Save(user *User) error {
	fmt.Printf("Saved user %v in Postgres\n", user)
	return nil
}

type InMemoryUserRepo struct {
	users map[int]*User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make(map[int]*User)}
}

func (r *InMemoryUserRepo) FindByID(id int) (*User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}
func (r *InMemoryUserRepo) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(id int, name string) (*User, error) {
	user := &User{ID: id, Name: name}
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func main() {
	postgresRepo := &PostgresUserRepo{}
	service := NewUserService(postgresRepo)
	u, _ := service.RegisterUser(1, "Egor")
	fmt.Println("Created:", u)

	mockRepo := NewInMemoryUserRepo()
	testService := NewUserService(mockRepo)
	_, _ = testService.RegisterUser(2, "TestUser")
	found, _ := testService.GetUser(2)
	fmt.Println("From test repo:", found)
}
