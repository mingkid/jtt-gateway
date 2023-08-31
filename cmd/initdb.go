package main

import (
	"github.com/mingkid/jtt-gateway/db"
	"github.com/mingkid/jtt-gateway/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/mapper",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db.DB)
	g.ApplyBasic(
		model.Term{},
		model.Platform{},
	)
	g.Execute()
}
