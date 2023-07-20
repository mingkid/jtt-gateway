package service

import (
	"database/sql"

	"github.com/mingkid/jtt808-gateway/dal"
	"github.com/mingkid/jtt808-gateway/model"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"

	"gorm.io/gorm"
)

type Terminal struct{}

type TermActiveCommander struct {
	SN     string // 序列号
	Active bool   // 状态 true在线 false离线
}

func (t Terminal) GetBySN(sn string) (term *model.Term, err error) {
	term, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(sn)).First()
	return
}

func (t Terminal) Save(opt TermSaveOpt) (err error) {
	// 只提供更新SIM
	if opt.SN == "" {
		return errcode.SNCanNotNull
	}
	// 判断终端是否存在数据库
	term, err := t.GetBySN(opt.SN)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	}
	if term != nil {
		// 更新终端
		term.SIM = opt.SIM
		_, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(opt.SN)).Updates(&term)
	} else {
		err = dal.Q.Term.Create(&model.Term{
			SN:  opt.SN,
			SIM: opt.SIM,
			Alive: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
		})
	}
	return
}

// SetTermActive 设置终端状态
func (t Terminal) SetTermActive(cmd TermActiveCommander) (err error) {
	sn, err := t.GetBySN(cmd.SN)
	if err != nil {
		return err
	}
	sn.Alive = sql.NullBool{
		Bool:  cmd.Active,
		Valid: true,
	}
	_, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(cmd.SN)).Updates(sn)
	return err
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

// TermSaveOpt 终端保存选项
type TermSaveOpt struct {
	SN  string `json:"sn" binding:"required"` // 序列号
	SIM string `json:"sim"`                   // 流量卡号
}
