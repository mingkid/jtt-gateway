package service

import (
	"fmt"

	"github.com/mingkid/jtt808-gateway/dal/mapper"
	"github.com/mingkid/jtt808-gateway/model"
	"github.com/mingkid/jtt808-gateway/pkg/errcode"

	"gorm.io/gorm"
)

type Terminal struct{}

// Locate 定位汇报
func (t Terminal) Locate(sn string, lng, lat float64) error {
	term, err := t.GetBySN(sn)
	if err != nil {
		return err
	}
	term.Lng = lng
	term.Lat = lat
	_, err = mapper.Q.Term.Where(mapper.Q.Term.SN.Like("%" + sn)).Updates(&term)
	return nil
}

func (t Terminal) All() ([]*model.Term, error) {
	return mapper.Q.Term.Find()
}

func (t Terminal) GetBySN(sn string) (term *model.Term, err error) {
	term, err = mapper.Q.Term.Where(mapper.Q.Term.SN.Like(fmt.Sprintf("%%%s", sn))).First()
	return
}

func (t Terminal) Save(sn, sim string) (err error) {
	// 只提供更新SIM
	if sn == "" {
		return errcode.TermSNNotNull
	}
	// 判断终端是否存在数据库
	term, err := t.GetBySN(sn)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	}
	if term != nil {
		// 更新终端
		term.SIM = sim
		_, err = mapper.Q.Term.Where(mapper.Q.Term.SN.Eq(sn)).Updates(&term)
	} else {
		err = mapper.Q.Term.Create(&model.Term{
			SN:  sn,
			SIM: sim,
		})
	}
	return
}

// Delete 删除终端
func (t Terminal) Delete(sn string) (err error) {
	var term *model.Term
	term, err = t.GetBySN(sn)
	if err != nil {
		return err
	}
	_, err = mapper.Q.Term.Delete(term)
	return err
}

func NewTerminal() *Terminal {
	return &Terminal{}
}
