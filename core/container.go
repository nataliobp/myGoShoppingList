package core

import (
	"database/sql"
	"github.com/nataliobp/myGoShoppingList/identityAccess"
	"log"
)

type Container struct {
	services map[string]interface{}
}

func NewContainer() *Container {
	services := make(map[string]interface{})
	return &Container{services}
}

func (c *Container) Set(name string, service interface{}) {
	c.services[name] = service
}

func (c *Container) Get(name string, ) interface{} {
	return c.services[name]
}

func (c *Container) Init() *Container {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	c.Set("inMemoryUserRepository", identityAccess.NewInMemoryUserRepository())
	c.Set("userRepository", identityAccess.NewSqliteUserRepository(db))

	c.Set("signUpService", identityAccess.NewSignUpService(
		c.Get("userRepository").(identityAccess.UserRepository)))

	c.Set("findUserService", identityAccess.NewFindUserService(
		c.Get("userRepository").(identityAccess.UserRepository)))

	return c
}
