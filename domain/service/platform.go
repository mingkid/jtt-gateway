package service

import (
	"github.com/mingkid/jtt808-gateway/dal/mapper"
	"github.com/mingkid/jtt808-gateway/model"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"

	"gorm.io/gorm"
)

type Platform struct{}

func (p Platform) All() ([]*model.Platform, error) {
	return mapper.Q.Platform.Find()
}

func (p Platform) Save(opt PlatformSaveOpt) error {
	if opt.Identity == "" {
		return errcode.PlantformIdentityNotNull
	}

	platform, err := p.GetByIdentity(opt.Identity)
	if err == gorm.ErrRecordNotFound {
		// 新增
		err = mapper.Q.Platform.Create(&model.Platform{
			Identity:    opt.Identity,
			Host:        opt.Host,
			LocationAPI: opt.LocationAPI,
		})
	} else {
		// 更新
		platform.Host = opt.Host
		platform.LocationAPI = opt.LocationAPI
		_, err = mapper.Q.Platform.Where(mapper.Q.Platform.Identity.Eq(opt.Identity)).Updates(&platform)
	}
	return nil
}

func (p Platform) GetByIdentity(identity string) (platform *model.Platform, err error) {
	platform, err = mapper.Q.Platform.Where(mapper.Q.Platform.Identity.Eq(identity)).First()
	return
}

// PlatformSaveOpt 平台保存选项
type PlatformSaveOpt struct {
	Identity    string // 平台标识
	Host        string // 域名
	LocationAPI string // 定位 API
}
