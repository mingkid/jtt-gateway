package publish

import (
	"bytes"
	"encoding/json"

	"github.com/mingkid/g-jtt808/msg"
)

type LocationOpt struct {
	Phone     string `json:"phone"`
	Warning   uint32 `json:"warning"`
	Status    uint32 `json:"status"`
	Lat       uint32 `json:"lat"`
	Lnt       uint32 `json:"lnt"`
	Altitude  uint16 `json:"altitude"`
	Speed     uint16 `json:"speed"`
	Direction uint16 `json:"direction"`
	Time      uint32 `json:"time"`
}

func (l LocationOpt) Buffer() (*bytes.Buffer, error) {
	locationJson, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(locationJson), nil
}

func NewLocation(h msg.Head, b msg.M0200) *LocationOpt {
	return &LocationOpt{
		Phone:    h.Phone(),
		Warning:  uint32(b.Warn()),
		Status:   uint32(b.Status()),
		Lat:      b.Latitude(),
		Lnt:      b.Longitude(),
		Altitude: b.Altitude(),
	}
}
