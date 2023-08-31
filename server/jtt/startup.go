package jtt

import (
	"fmt"

	"github.com/mingkid/jtt-gateway/domain/service"
	"github.com/mingkid/jtt-gateway/model"
	"github.com/mingkid/jtt-gateway/server/jtt/publish"

	jtt "github.com/mingkid/g-jtt"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

func Serve(port uint) {
	jttSvr := jtt.Default()
	jttSvr.PhoneToTermID = func(ctx *jtt.Context) (string, error) {
		term, err := service.Terminal{}.GetBySN(ctx.Head().Phone[3:])
		if err != nil {
			if ctx.Head().MsgID == msg.MsgIDTermRegister {
				_ = ctx.Register(msg.M8100ResultCarNotInDB, "")
			} else {
				_ = ctx.Generic(msg.M8001ResultFail)
			}
			return "", err
		}
		return term.SN, nil
	}

	// 注册控制器函数
	jttSvr.RegisterHandler(msg.MsgIDTermRegister, termRegister)
	jttSvr.RegisterHandler(msg.MsgIDTermAuth, termAuth)
	jttSvr.RegisterHandler(msg.MsgIDTermLocationReport, locationReport)
	jttSvr.RegisterHandler(msg.MsgIDTermLocationBatch, locationBatchReport)
	jttSvr.RegisterHandler(msg.MsgIDTermHeartbeat, termHearBeat)

	jttSvr.Serve("", port)
}

func termRegister(ctx *jtt.Context) {
	var (
		decoder codec.Decoder
		err     error

		m   = msg.New[msg.M0100]()
		res = msg.M8100ResultSuccess
	)

	_ = decoder.Decode(m, ctx.Data())

	// 业务处理
	termService := service.NewTerminal()
	if _, err = termService.GetBySN(m.Body.TermID); err != nil {
		res = msg.M8100TermNotInDB
		ctx.Register(res, "")
		return
	}

	ctx.Register(res, "123123")
}

func termAuth(ctx *jtt.Context) {
	var (
		decoder codec.Decoder

		m   = msg.New[msg.M0102]()
		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(m, ctx.Data())

	// 业务处理：token 校验
	token := m.Body.Token
	if token != "123123" {
		res = msg.M8001ResultFail
	}

	_ = ctx.Generic(res)
}

func locationReport(ctx *jtt.Context) {
	var (
		decoder codec.Decoder
		err     error

		m   = msg.New[msg.M0200]()
		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(m, ctx.Data())

	// 业务处理：终端定位更新
	termService := service.Terminal{}
	lng := float64(m.Body.Longitude) / 1000000.0
	lat := float64(m.Body.Latitude) / 1000000.0
	if err = termService.Locate(m.Head.Phone[3:], lng, lat); err != nil {
		res = msg.M8001ResultFail
		ctx.Generic(res)
		return
	}

	// 业务处理：终端定位推送到业务平台
	var platforms []*model.Platform
	platformService := service.Platform{}
	if platforms, err = platformService.All(); err != nil {
		res = msg.M8001ResultFail
		ctx.Generic(res)
		return
	}

	for _, platform := range platforms {
		pusher := publish.New(platform.Host, platform.LocationAPI)
		if err = pusher.Locate(publish.NewLocationOpt(m.Phone, &m.Body, false)); err != nil {
			res = msg.M8001ResultFail
			ctx.Generic(res)
			return
		}
	}

	ctx.Generic(res)
}

func locationBatchReport(ctx *jtt.Context) {
	var (
		decoder codec.Decoder
		err     error

		m   = msg.New[msg.M0704]()
		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(m, ctx.Data())

	// 业务处理
	platformService := service.Platform{}
	platforms, err := platformService.All()
	if err != nil {
		fmt.Println(err.Error())
		res = msg.M8001ResultFail
		ctx.Generic(res)
		return
	}

	for _, platform := range platforms {
		pusher := publish.New(platform.Host, platform.LocationAPI)
		locations, _ := m.Body.Items()
		for _, b := range locations {
			m0200b := new(msg.M0200)
			decoder.Decode(m0200b, b)
			_ = pusher.Locate(publish.NewLocationOpt(m.Head.Phone, m0200b, true))
		}
	}

	ctx.Generic(res)
}

func termHearBeat(ctx *jtt.Context) {
	var (
		decoder codec.Decoder
		m       msg.Head

		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(&m, ctx.Data())

	// 业务处理
	_, err := service.NewTerminal().GetBySN(ctx.Head().Phone[3:])
	if err != nil {
		res = msg.M8001ResultFail
		ctx.Generic(res)
		return
	}
	//svr.updateSession(term.SN)

	ctx.Generic(res)
}
