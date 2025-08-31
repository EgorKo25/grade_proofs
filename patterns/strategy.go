package patterns

import (
	"errors"
	"fmt"
)

// AuthStrategy интерфейс для всех стратегий авторизации
type AuthStrategy interface {
	Authenticate(credentials string) (string, error)
}

// JWTAuth стратегия для JWT авторизации
type JWTAuth struct{}

func (j *JWTAuth) Authenticate(token string) (string, error) {
	if token == "valid-jwt-token" {
		return "user123", nil
	}
	return "", errors.New("invalid JWT token")
}

// BasicAuth стратегия для Basic авторизации
type BasicAuth struct{}

func (b *BasicAuth) Authenticate(credentials string) (string, error) {
	if credentials == "user:password" {
		return "user123", nil
	}
	return "", errors.New("invalid basic credentials")
}

// OAuth2Auth стратегия для OAuth2 авторизации
type OAuth2Auth struct{}

func (o *OAuth2Auth) Authenticate(token string) (string, error) {
	if token == "valid-oauth-token" {
		return "user123", nil
	}
	return "", errors.New("invalid OAuth2 token")
}

type AuthContext struct {
	strategy AuthStrategy
}

// SetStrategy позволяет установить стратегию авторизации
func (a *AuthContext) SetStrategy(strategy AuthStrategy) {
	a.strategy = strategy
}

// Authenticate выполняет аутентификацию с использованием установленной стратегии
func (a *AuthContext) Authenticate(credentials string) (string, error) {
	return a.strategy.Authenticate(credentials)
}

func main() {
	context := &AuthContext{}

	context.SetStrategy(&JWTAuth{})
	userID, err := context.Authenticate("valid-jwt-token")
	if err != nil {
		fmt.Println("Ошибка аутентификации:", err)
	} else {
		fmt.Println("Аутентифицирован пользователь:", userID)
	}

	context.SetStrategy(&BasicAuth{})
	userID, err = context.Authenticate("user:password")
	if err != nil {
		fmt.Println("Ошибка аутентификации:", err)
	} else {
		fmt.Println("Аутентифицирован пользователь:", userID)
	}

	context.SetStrategy(&OAuth2Auth{})
	userID, err = context.Authenticate("valid-oauth-token")
	if err != nil {
		fmt.Println("Ошибка аутентификации:", err)
	} else {
		fmt.Println("Аутентифицирован пользователь:", userID)
	}
}
