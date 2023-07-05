package model

import "database/sql"

type Term struct {
	SN    string
	SIM   string
	Alive sql.NullBool
}
