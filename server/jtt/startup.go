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

var Svr *jtt.Engine

func Serve(port uint) {
	Svr = jtt.Default()
	Svr.PhoneToTermID = func(phone string) (string, error) {
		term, err := service.Terminal{}.GetBySN(phone[3:])
		if err != nil {
			return "", err
		}
		return term.SN, nil
	}

	// 注册控制器函数
	Svr.RegisterHandler(msg.MsgIDTermRegister, termRegister)
	Svr.RegisterHandler(msg.MsgIDTermAuth, termAuth)
	Svr.RegisterHandler(msg.MsgIDTermLocationReport, locationReport)
	Svr.RegisterHandler(msg.MsgIDTermLocationBatch, locationBatchReport)
	Svr.RegisterHandler(msg.MsgIDTermHeartbeat, termHearBeat)

	Svr.Serve("", port)
}

func termRegister(ctx *jtt.Context) {
	var (
		m       msg.MsgWith[msg.M0100]
		decoder codec.Decoder
		err     error

		res = msg.M8100ResultSuccess
	)

	_ = decoder.Decode(&m, ctx.Data())

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
		m       msg.MsgWith[msg.M0102]
		decoder codec.Decoder

		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(&m, ctx.Data())

	// 业务处理：token 校验
	token := m.Body.Token
	if token != "123123" {
		res = msg.M8001ResultFail
	}

	_ = ctx.Generic(res)
}

func locationReport(ctx *jtt.Context) {
	var (
		m       msg.MsgWith[msg.M0200]
		decoder codec.Decoder
		err     error

		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(&m, ctx.Data())

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
		m       msg.MsgWith[msg.M0704]
		decoder codec.Decoder
		err     error

		res = msg.M8001ResultSuccess
	)

	_ = decoder.Decode(&m, ctx.Data())

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
