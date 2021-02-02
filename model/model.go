package model

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/jinzhu/gorm"
	"api-auth-test/auth"
)

type Server struct {
	DB *gorm.DB
}

var (
	Model modelInterface = &Server{}
)

type modelInterface interface {
	Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error)

	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

	CreateTodo(*Todo) (*Todo, error)

	FetchAuth(*auth.Details) (*Auth, error)
	DeleteAuth(*auth.Details) error
	CreateAuth(uint64) (*Auth, error)
}

func (s *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmodedisable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		return nil, err
	}
	s.DB.Debug().AutoMigrate(
		&User{},
		&Auth{},
		&Todo{},
	)
	return s.DB, nil
}

