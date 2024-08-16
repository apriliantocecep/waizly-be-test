package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"waizly/internal/entity"
)

func ClearAll() {
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	var user *entity.User
	err := db.First(&user).Error
	assert.Nil(t, err)
	return user
}
