package main

import (
	"errors"
	"fmt"
)

// // S — Single Responsibility
// User отвечает только за хранение данных пользователя
type User struct {
	ID   int
	Name string
}

// // O — Open/Closed
// Интерфейс UserRepository открыт для расширения (можно добавить новую реализацию),
// но закрыт для модификации (контракт остаётся тем же)
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

type InMemoryUserRepo struct {
	store map[int]*User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{store: make(map[int]*User)}
}

func (r *InMemoryUserRepo) FindByID(id int) (*User, error) {
	if user, ok := r.store[id]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}

func (r *InMemoryUserRepo) Save(user *User) error {
	r.store[user.ID] = user
	return nil
}

// // L — Liskov Substitution
// Добавляем PostgresUserRepo — можно подставить вместо InMemoryUserRepo
// и UserService продолжит работать без изменений
type PostgresUserRepo struct{}

func (r *PostgresUserRepo) FindByID(id int) (*User, error) {
	return &User{ID: 1, Name: "From Postgres"}, nil
}

func (r *PostgresUserRepo) Save(user *User) error {
	fmt.Println("User saved in Postgres:", user)
	return nil
}

// // I — Interface Segregation
// Выделяем отдельный интерфейс для уведомлений, чтобы не смешивать с репозиторием
type UserNotifier interface {
	Notify(user *User, msg string) error
}

type EmailNotifier struct{}

func (n *EmailNotifier) Notify(user *User, msg string) error {
	fmt.Printf("Email to %s: %s\n", user.Name, msg)
	return nil
}

// // D — Dependency Inversion
// UserService зависит от абстракций (UserRepository, UserNotifier), а не от конкретных реализаций
type UserService struct {
	repo     UserRepository
	notifier UserNotifier
}

func NewUserService(repo UserRepository, notifier UserNotifier) *UserService {
	return &UserService{repo: repo, notifier: notifier}
}

func (s *UserService) RegisterUser(id int, name string) (*User, error) {
	user := &User{ID: id, Name: name}
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	_ = s.notifier.Notify(user, "Welcome to our system!")
	return user, nil
}

// // Main — демонстрация работы
func main() {
	repo := NewInMemoryUserRepo()
	notifier := &EmailNotifier{}
	service := NewUserService(repo, notifier)

	user, _ := service.RegisterUser(1, "Egor")
	fmt.Println("Created:", user)

	// Подставим другую реализацию репозитория (Postgres)
	postgresRepo := &PostgresUserRepo{}
	service2 := NewUserService(postgresRepo, notifier)
	_, _ = service2.RegisterUser(2, "Ivan")
}
