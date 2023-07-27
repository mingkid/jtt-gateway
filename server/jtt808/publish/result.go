package publish

import "encoding/json"

type LocationRes struct {
	Token string `json:"token"`
}

func LoadLocationRes(b []byte) (*LocationRes, error) {
	r := new(LocationRes)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
