package service

import (
	"github.com/mingkid/jtt808-gateway/dal/mapper"
	"github.com/mingkid/jtt808-gateway/model"
)

type Platform struct{}

func (p Platform) All() ([]*model.Platform, error) {
	return mapper.Q.Platform.Find()
}

func NewPlatform() *Platform {
	return &Platform{}
}
