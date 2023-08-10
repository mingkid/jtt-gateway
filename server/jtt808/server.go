package jtt808

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/mingkid/jtt808-gateway/domain"
	"github.com/mingkid/jtt808-gateway/domain/service"
	"github.com/mingkid/jtt808-gateway/server/jtt808/publish"

	jtt808 "github.com/mingkid/g-jtt808"
	"github.com/mingkid/g-jtt808/msg"
	msgCom "github.com/mingkid/g-jtt808/msg/common"
	"github.com/mingkid/jtt808-gateway/log"
	"gorm.io/gorm"
)

// DefaultWriter 默认Writer
var DefaultWriter io.Writer = os.Stdout

type Server struct {
	ipAddr   string
	port     uint
	sessions domain.SessionCache
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
			err := svr.handleConnect(c)
			if err != nil {
				panic(err)
			}
		}()
	}
}

func (svr *Server) handleConnect(c net.Conn) (err error) {
	for {
		var n int
		b := make([]byte, 1024)
		n, err = c.Read(b[:])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("[JTT808] 终端 %s 已断开连接！ \n", c.RemoteAddr())
				return nil
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
			rb, err = svr.termRegister(c, b)
		case msgCom.TermAuth:
			rb, err = svr.termAuth(c, b)
		case msgCom.TermHeartbeat:
			rb, err = svr.termHeartbeat(c, b)
		case msgCom.TermLocationRepose:
			rb, err = svr.termPositionRepose(c.RemoteAddr(), b)
		case msgCom.TermLocationBatch:
			rb, err = svr.termLocationBatch(c, b)
		default:
			rb, err = svr.unknown(c, b)
		}
		if err != nil {
			fmt.Println(err)
		}

		if _, err = c.Write(rb); err != nil {
			fmt.Println(err)
		}
	}
}

func NewServer(ipAddr string, port uint, cache domain.SessionCache) *Server {
	return &Server{
		ipAddr:   ipAddr,
		port:     port,
		sessions: cache,
	}
}

// termRegister 终端注册
func (svr *Server) termRegister(c net.Conn, b []byte) (resp []byte, err error) {
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
	res := msg.M8100Success
	termService := service.NewTerminal()
	var termID string
	termID, err = msgResB.TermID()
	if err != nil {
		return nil, err
	}
	_, err = termService.GetBySN(termID)
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}

	// 结果处理
	if err == gorm.ErrRecordNotFound {
		// 终端不存在
		res = msg.M8100TermRegistered
	} else {
		// 终端已注册
		res = msg.M8100TermRegistered
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

// termAuth 终端鉴权
func (svr *Server) termAuth(c net.Conn, b []byte) (resp []byte, err error) {
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

	// 结果处理
	res := msg.M8001Success
	//if isAllow == false {
	//	res = msg.M8001Fail
	//}

	// 业务处理
	token, _ := msgResB.Token()
	if token != "123123" {
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

// termHeartbeat 心跳
func (svr *Server) termHeartbeat(c net.Conn, b []byte) (resp []byte, err error) {
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

	result := msg.M8001Success

	// 业务处理
	term, err := service.NewTerminal().GetBySN(msgResH.Phone()[3:])
	if err != nil {
		return nil, err
	}
	svr.updateSession(term.SN)

	// 组装响应
	msgRespH.SetID(msgCom.PlatformCommResp)
	msgRespH.SetPhone(msgResH.Phone())
	msgRespB.SetMsgID(msgResH.MsgID()).SetSerialNumber(msgResH.SerialNum()).SetResult(result)
	mr := msg.NewPlantFormMsg(&msgRespH, msgRespB)
	return mr.Encode()
}

func (svr *Server) termLocationBatch(c net.Conn, b []byte) (resp []byte, err error) {
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

func (svr *Server) termPositionRepose(addr net.Addr, b []byte) (resp []byte, err error) {
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

	// 结果处理
	res := msg.M8001Success
	if err != nil {
		fmt.Println(err.Error())
		res = msg.M8001Fail
	}

	// 业务处理
	platformService := service.Platform{}
	platforms, err := platformService.All()
	if err != nil {
		fmt.Println(err.Error())
		res = msg.M8001Fail
	}

	for _, platform := range platforms {
		pusher := publish.New(platform.Host, platform.LocationAPI)
		_ = pusher.Locate(publish.NewLocationOpt(msgResH.Phone(), msgResB))
	}

	// 打印日志
	fmt.Fprint(DefaultWriter, log.DefaultInfoFormatter(log.InfoFormatterParams{
		Time:   time.Now(),
		IP:     addr,
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

func (svr *Server) unknown(c net.Conn, b []byte) (resp []byte, err error) {
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

// updateSession 更新会话信息
func (svr *Server) updateSession(sn string) {
	session := svr.sessions.Get(sn)
	if session == nil {
		session = new(domain.Session)
		svr.sessions.Set(sn, session)
	}
	session.Expire = int64(time.Minute * 30)
	svr.sessions.Set(sn, session)
}
