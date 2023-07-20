package service

import (
	"database/sql"

	"github.com/mingkid/jtt808-gateway/dal"
	"github.com/mingkid/jtt808-gateway/model"
	"github.com/mingkid/jtt808-gateway/server/web/common/errcode"

	"gorm.io/gorm"
)

type Terminal struct{}

func (t Terminal) GetBySN(sn string) (term *model.Term, err error) {
	term, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(sn)).First()
	return
}

func (t Terminal) Save(sn, sim string) (err error) {
	// 只提供更新SIM
	if sn == "" {
		return errcode.SNCanNotNull
	}
	// 判断终端是否存在数据库
	term, err := t.GetBySN(sn)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	}
	if term != nil {
		// 更新终端
		term.SIM = sim
		_, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(sn)).Updates(&term)
	} else {
		err = dal.Q.Term.Create(&model.Term{
			SN:  sn,
			SIM: sim,
			Alive: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
		})
	}
	return
}

// Activate 设置终端状态
func (t Terminal) Activate(sn string) (err error) {
	var term *model.Term
	term, err = t.GetBySN(sn)
	if err != nil {
		return err
	}
	term.Alive = sql.NullBool{
		Bool:  true,
		Valid: true,
	}
	_, err = dal.Q.Term.Where(dal.Q.Term.SN.Eq(sn)).Updates(term)
	return err
}

// Delete 删除终端
func (t Terminal) Delete(sn string) (err error) {
	var term *model.Term
	term, err = t.GetBySN(sn)
	if err != nil {
		return err
	}
	_, err = dal.Q.Term.Delete(term)
	return err
}

func NewTerminal() *Terminal {
	return &Terminal{}
}
