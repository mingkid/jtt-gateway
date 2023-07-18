package service

import (
	"github.com/mingkid/jtt808-gateway/dal"
	"github.com/mingkid/jtt808-gateway/model"
)

type Platform struct{}

func (p Platform) All() ([]*model.Platform, error) {
	return dal.Q.Platform.Find()
}

func NewPlatform() *Platform {
	return &Platform{}
}
