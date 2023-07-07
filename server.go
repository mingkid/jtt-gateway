package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	jtt808 "github.com/mingkid/g-jtt808"
	"github.com/mingkid/g-jtt808/msg"
	msgCom "github.com/mingkid/g-jtt808/msg/common"
	"github.com/mingkid/jtt808-gateway/log"
)

// DefaultWriter 默认Writer
var DefaultWriter io.Writer = os.Stdout

type Server struct {
	ipAddr string
	port   uint
}

// IPAddr IP 地址
func (svr *Server) IPAddr() string {
	return svr.ipAddr
}

// Port 端口
func (svr *Server) Port() uint {
	return svr.port
}

// Serve 开启服务
func (svr *Server) Serve() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svr.ipAddr, svr.port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("[JTT808] 监听开始: %s! \n", l.Addr())

	for {
		c, err := l.Accept()
		fmt.Printf("[JTT808] 终端 %s 已连接！ \n", c.RemoteAddr())
		if err != nil {
			panic(err)
		}
		go func() {
			err := handleConnect(c)
			if err != nil {
				panic(err)
			}
		}()
	}
}

func handleConnect(c net.Conn) (err error) {
	for {
		var n int
		b := make([]byte, 1024)
		n, err = c.Read(b[:])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("[JTT808] 终端 %s 已断开连接！ \n", c.RemoteAddr())
				return
			}
			fmt.Println(err)
		}
		// 忽略空消息
		if n == 0 {
			continue
		}
		b = b[:n]

		// 不同的消息不同方式处理
		var rb []byte
		switch jtt808.ExtractMsgID(b) {
		case msgCom.TermRegister:
			rb, err = termRegister(c, b)
		case msgCom.TermAuth:
			rb, err = termAuth(c, b)
		case msgCom.TermHeartbeat:
			rb, err = termHeartbeat(c, b)
		case msgCom.TermLocationRepose:
			rb, err = termPositionRepose(c, b)
		case msgCom.TermLocationBatch:
			rb, err = termLocationBatch(c, b)
		default:
			rb, err = unknown(c, b)
		}
		if err != nil {
			fmt.Println(err)
		}

		if _, err = c.Write(rb); err != nil {
			fmt.Println(err)
		}
	}
}

func NewServer(ipAddr string, port uint) *Server {
	return &Server{
		ipAddr: ipAddr,
		port:   port,
	}
}

func termRegister(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgResB  msg.M0100
		msgRespH msg.Head
		msgRespB msg.M8100
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, &msgResB)
	if err = m.Decode(b); err != nil {
		return
	}

	// 业务处理
	// TODO

	// 结果处理
	res := msg.M8100Success
	if err != nil {
		//if err == application.TermNotFound {
		//	res = msg.M8100TermNotInDB
		//} else {
		//	fmt.Println(err.Error())
		//	return
		//}
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(res),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 组装响应
	msgRespH.SetID(msgCom.TermRegResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetToken("123123").SetSerialNumber(msgResH.SerialNum()).SetResult(res)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	resp, err = mr.Encode()
	return
}

func termAuth(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgResB  msg.M0102
		msgRespH msg.Head
		msgRespB msg.M8001
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, &msgResB)
	if err = m.Decode(b); err != nil {
		return
	}

	// 业务处理
	// TODO

	// 结果处理
	res := msg.M8001Success
	//if isAllow == false {
	//	res = msg.M8001Fail
	//}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(res),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(res)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}

func termHeartbeat(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgRespH msg.Head
		msgRespB msg.M8001
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, nil)
	if err = m.Decode(b); err != nil {
		return
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(msg.M8001Success),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 业务处理
	// TODO

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(msg.M8001Success)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}

func termLocationBatch(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgResB  msg.M0704
		msgRespH msg.Head
		msgRespB msg.M8001
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, &msgResB)
	if err = m.Decode(b); err != nil {
		return
	}

	// 业务处理
	// TODO

	// 结果处理
	res := msg.M8001Success
	if err != nil {
		fmt.Println(err.Error())
		res = msg.M8001Fail
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(res),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(res)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}

func termPositionRepose(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgResB  msg.M0200
		msgRespH msg.Head
		msgRespB msg.M8001
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, &msgResB)
	if err = m.Decode(b); err != nil {
		return
	}

	// 业务处理
	// TODO

	// 结果处理
	res := msg.M8001Success
	if err != nil {
		fmt.Println(err.Error())
		res = msg.M8001Fail
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(res),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(res)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}

func unknown(c net.Conn, b []byte) (resp []byte, err error) {
	var (
		msgResH  msg.Head
		msgRespH msg.Head
		msgRespB msg.M8001
	)

	// 解码
	m := msg.NewTermMsg(&msgResH, nil)
	if err = m.Decode(b); err != nil {
		return
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     c.RemoteAddr(),
		Phone:  msgResH.Phone(),
		Result: uint8(msg.M8001Success),
		MsgID:  b[1:3],
		Data:   b,
	}))

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(msg.M8001Success)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}
