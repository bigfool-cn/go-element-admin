package apis

import (
  "container/list"
  "encoding/json"
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
  "go-element-admin-api/utils"
  "log"
  "net"
  "net/http"
  "strconv"
  "strings"
  "sync"
  "time"
)

const (
	// 消息大小限制
	maxMessageSize = 512
	// 发送时间间隔
	sendLimitTime = 2

	// 单个ip链接数限制
	ipLimit = 50
)

type userAction interface {
	getName()(string,*list.Element)
	userJoin(name string,wsConn *wsConnection)
	userLeave(name string,wsConn *wsConnection)
}

type user struct {}

type userMessage struct {
	TypeId int `json:"type_id"`
	Names  []string `json:"names"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	Msg    string `json:"msg"`
}

var (
	// 最大的连接ID 每次连接都加1处理
	maxConnId int64

	// 用户
	us *user

	// 所有连接
	wsConnAll map[int64]*wsConnection

	// ip计数器
	connIpAll map[string]int64
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *userMessage // 读队列
	outChan  chan *userMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道,加锁处理
	isClosed  bool
	closeChan chan byte // 关闭通知
	ip        string // ip
	tokenInfo *utils.Claims // token信息
	sendTime  time.Time // 发送时间
	name      string // 分配的别名
	nameEl    *list.Element
	id        int64
}

// 获取客户端IP
func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}


func Chat(c *gin.Context)  {
	token := c.Query("token")
	tokenInfo, er := utils.Jwt.ParseToken(token)
	if er != nil {
    log.Println("token", token)
		log.Println("token验证失败", er.Error())
		http.Error(c.Writer,"token error",401)
		return
	}
	ip := clientIP(c.Request)
	if (ip == "" || connIpAll[ip] > ipLimit) {
		http.Error(c.Writer,"ip limit",403)
		return
	}

	// 应答客户端告知升级连接为ws
	wsScoket,err := upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Println("升级为ws失败", err.Error())
		return
	}

	connIpAll[ip] += 1
	maxConnId++
	log.Println("当前IP链接数",connIpAll)
	us = new(user)

	var stringBuild strings.Builder
	stringBuild.WriteString(tokenInfo.UserName)
  stringBuild.WriteString("_")
  stringBuild.WriteString(strconv.FormatInt(maxConnId,10))
	name := stringBuild.String()

	wsConn := &wsConnection{
		wsSocket:  wsScoket,
		inChan:    make(chan *userMessage, 1000),
		outChan:   make(chan *userMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		ip:        ip,
		tokenInfo: tokenInfo,
		sendTime:  time.Now(),
		name:      name,
		id:        maxConnId,
	}
	wsConnAll[maxConnId] = wsConn

	go us.userJoin(name,wsConn)

	// 读取消息
	go wsConn.wsReadLoop()

	// 发送消息
	go wsConn.wsWriteLoop()
}

// 用户加入处理
func (u *user) userJoin(name string, wsConn *wsConnection)  {
	names := make([]string,0)
	for _, wsc := range wsConnAll {
		names = append(names,wsc.name)
	}
	log.Println("在线用户",names)

	// 回复自己分配的名称
	meMsg := &userMessage{TypeId:1,Name:"system",Msg:name}
	err := wsConn.wsWrite(meMsg)
	if err != nil {
		log.Println("回复自己分配的名称错误",err.Error())
	}

	// 广播加入用户
	msg := &userMessage{TypeId:0,Names:names,Name:"system",Msg:name + " 加入群聊"}
	wsConn.wsWriteAll(msg)
}

// 用户离开处理
func (u *user) userLeave(name string, wsConn *wsConnection)  {
	connIpAll[wsConn.ip] -= 1
	if connIpAll[wsConn.ip] == 0 {
	  delete(connIpAll,wsConn.ip)
	}
	names := make([]string,0)
  for _, wsc := range wsConnAll {
    names = append(names,wsc.name)
  }
	// 广播加入用户
	msg := &userMessage{TypeId:0,Names:names,Name:"system",Msg:name + " 离开群聊"}
	wsConn.wsWriteAll(msg)
}

func (wsConn *wsConnection) wsReadLoop() {
	// 设置消息最大长度
	wsConn.wsSocket.SetReadLimit(maxMessageSize)
	for {
		// 读取一个message
		userMsg := &userMessage{}
		err := wsConn.wsSocket.ReadJSON(userMsg)
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Println("消息读取出现错误", err.Error())
			wsConn.close()
			return
		}

		diffTime := time.Now().Sub(wsConn.sendTime)
		// 限制发送频率
		if diffTime < sendLimitTime && userMsg.TypeId == 12 {
			userMsg = &userMessage{
				TypeId: -1,
				Name: "系统",
				Msg: "发送频率过快，请稍后再发！",
			}
			err := wsConn.wsWrite(userMsg)
			if err != nil {
				log.Println("发送消息失败",err)
			}
		} else {
			wsConn.sendTime = time.Now()
			wsConn.wsWriteAll(userMsg)
		}

		// 放入请求队列，消息入栈
		select {
		case wsConn.inChan <- userMsg:
		case <- wsConn.closeChan:
			return
		}
	}
}

func (wsConn *wsConnection) wsWriteLoop()  {
	tick := time.NewTicker(60 * time.Second)
	for {
		select {
		case userMsg := <- wsConn.outChan:
			msg ,er := json.Marshal(userMsg)
			if er != nil {
				log.Println("解析发送消息JSON失败", er.Error())
			}
			log.Println("发送消息",string(msg))
			if err := wsConn.wsSocket.WriteJSON(string(msg)); err != nil {
				log.Println("发送消息给客户端发生错误", err.Error())
				// 切断服务
				wsConn.close()
				return
			}
		case <- tick.C:
			if err := wsConn.wsSocket.WriteControl(websocket.PingMessage, []byte(""), time.Now().Add(30 * time.Second));err != nil {
				log.Println("ping error", err.Error())
			}

		case <-wsConn.closeChan:
			tick.Stop()
			// 获取到关闭通知
			return
		}
	}
}

// 写入消息到队列中
func (wsConn *wsConnection) wsWrite(userMsg *userMessage) error {
	select {
	case wsConn.outChan <- userMsg:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭：" + wsConn.name)
	}
	return nil
}

// 广播全部链接
func (wsConn *wsConnection) wsWriteAll(userMsg *userMessage)  {
	for _, wsCon := range wsConnAll {
		wsCon.outChan <- userMsg
	}
}

// 读取消息队列中的消息
func (wsConn *wsConnection) wsRead() (*userMessage, error) {
	select {
	case userMsg := <-wsConn.inChan:
		// 获取到消息队列中的消息
		return userMsg, nil
	case <-wsConn.closeChan:

	}
	return nil, errors.New("连接已经关闭")
}

// 关闭连接
func (wsConn *wsConnection) close() {
	log.Println("关闭连接被调用了",wsConn.name)
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()

	if wsConn.isClosed == false {
		wsConn.isClosed = true
		// 删除这个连接的变量
		delete(wsConnAll, wsConn.id)
		close(wsConn.closeChan)
    us.userLeave(wsConn.name,wsConn)
	}
}

func init()  {
  wsConnAll = make(map[int64]*wsConnection)
  connIpAll = make(map[string]int64)
}


