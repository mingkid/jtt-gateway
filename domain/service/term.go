package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mingkid/jtt-gateway/dal/mapper"
	"github.com/mingkid/jtt-gateway/model"
	"github.com/mingkid/jtt-gateway/pkg/errcode"

	"gorm.io/gorm"
)

type Terminal struct{}

// Locate 定位汇报
func (t Terminal) Locate(sim string, lng, lat float64, locateAt time.Time) error {
	term, err := t.GetBySIM(sim)
	if err != nil {
		return err
	}
	term.Lng = lng
	term.Lat = lat
	term.LocateAt = sql.NullInt64{
		Int64: locateAt.Unix(),
		Valid: true,
	}
	_, err = mapper.Q.Term.Where(mapper.Q.Term.SIM.Like("%" + sim)).Updates(&term)
	return nil
}

func (t Terminal) All() ([]*model.Term, error) {
	return mapper.Q.Term.Find()
}

func (t Terminal) GetBySIM(sim string) (term *model.Term, err error) {
	term, err = mapper.Q.Term.Where(mapper.Q.Term.SIM.Like(fmt.Sprintf("%%%s", sim))).First()
	return
}

func (t Terminal) Save(sim string) (err error) {
	// 只提供更新SIM
	if sim == "" {
		return errcode.TermSNNotNull
	}
	// 判断终端是否存在数据库
	term, err := t.GetBySIM(sim)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	}
	if term != nil {
		// 更新终端
		term.SIM = sim
		_, err = mapper.Q.Term.Where(mapper.Q.Term.SIM.Eq(sim)).Updates(&term)
	} else {
		err = mapper.Q.Term.Create(&model.Term{
			SIM: sim,
		})
	}
	return
}

// Delete 删除终端
func (t Terminal) Delete(sn string) (err error) {
	var term *model.Term
	term, err = t.GetBySIM(sn)
	if err != nil {
		return err
	}
	_, err = mapper.Q.Term.Delete(term)
	return err
}

func NewTerminal() *Terminal {
	return &Terminal{}
}
