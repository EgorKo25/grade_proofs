package patterns

import "fmt"

// ==== Абстракции ====

// User модель
type User struct {
	ID   int
	Name string
}

// UserRepository общий интерфейс
type UserRepository interface {
	Save(user *User) error
	FindByID(id int) (*User, error)
}

// ==== Реализации ====

// PostgresRepo реализация
type PostgresRepo struct{}

func (r *PostgresRepo) Save(user *User) error {
	fmt.Println("Saving user in Postgres:", user.Name)
	return nil
}
func (r *PostgresRepo) FindByID(id int) (*User, error) {
	return &User{ID: id, Name: "From Postgres"}, nil
}

// MySQLRepo реализация
type MySQLRepo struct{}

func (r *MySQLRepo) Save(user *User) error {
	fmt.Println("Saving user in MySQL:", user.Name)
	return nil
}
func (r *MySQLRepo) FindByID(id int) (*User, error) {
	return &User{ID: id, Name: "From MySQL"}, nil
}

// ==== Фабрика ====

// Типы репозиториев
const (
	Postgres = "postgres"
	MySQL    = "mysql"
)

// RepositoryFactory фабрика, создающая репозитории
type RepositoryFactory struct{}

func (f *RepositoryFactory) Create(repoType string) (UserRepository, error) {
	switch repoType {
	case Postgres:
		return &PostgresRepo{}, nil
	case MySQL:
		return &MySQLRepo{}, nil
	default:
		return nil, fmt.Errorf("unknown repository type: %s", repoType)
	}
}

// ==== Использование ====
func main() {
	factory := &RepositoryFactory{}

	postgresRepo, _ := factory.Create(Postgres)
	mysqlRepo, _ := factory.Create(MySQL)

	postgresRepo.Save(&User{ID: 1, Name: "Egor"})
	mysqlRepo.Save(&User{ID: 2, Name: "Ivan"})
}
