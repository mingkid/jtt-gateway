package service

import (
	"github.com/mingkid/jtt-gateway/dal/mapper"
	"github.com/mingkid/jtt-gateway/model"
	"github.com/mingkid/jtt-gateway/pkg/errcode"

	"gorm.io/gorm"
)

type Platform struct{}

func (p Platform) All() ([]*model.Platform, error) {
	return mapper.Q.Platform.Find()
}

func (p Platform) Save(opt PlatformSaveOpt) error {
	if opt.Identity == "" {
		return errcode.PlatformIdentityNotNull
	}

	platform, err := p.GetByIdentity(opt.Identity)
	if err == gorm.ErrRecordNotFound {
		// 新增
		err = mapper.Q.Platform.Create(&model.Platform{
			Identity:    opt.Identity,
			Host:        opt.Host,
			LocationAPI: opt.LocationAPI,
			AccessToken: opt.AccessToken,
		})
	} else {
		// 更新
		platform.Host = opt.Host
		platform.LocationAPI = opt.LocationAPI
		platform.AccessToken = opt.AccessToken
		_, err = mapper.Q.Platform.Where(mapper.Q.Platform.Identity.Eq(opt.Identity)).Updates(&platform)
	}
	return nil
}

func (p Platform) GetByIdentity(ident string) (platform *model.Platform, err error) {
	platform, err = mapper.Q.Platform.Where(mapper.Q.Platform.Identity.Eq(ident)).First()
	return
}

func (p Platform) Delete(ident string) error {
	platform, err := p.GetByIdentity(ident)
	if err != nil {
		return err
	}
	_, err = mapper.Q.Platform.Delete(platform)
	return err
}

// PlatformSaveOpt 平台保存选项
type PlatformSaveOpt struct {
	Identity    string // 平台标识
	Host        string // 域名
	LocationAPI string // 定位 API
	AccessToken string // 访问令牌
}
