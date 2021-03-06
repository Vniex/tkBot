package websocket

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
	"net/http"
	"fmt"
	"errors"
	log "github.com/sirupsen/logrus"
)


// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接
type WsConnection struct {
	clientName string
	wsSocket *websocket.Conn // 底层websocket
	inChan chan *RobotHubMsg	// 读队列
	outChan chan *RobotHubMsg // 写队列

	mutex sync.Mutex	// 避免重复关闭管道
	isClosed bool
	closeChan chan byte  // 关闭通知
}

func NewWsConnection (name string,conn *websocket.Conn) *WsConnection{
	return &WsConnection{
		clientName:name,
		wsSocket: conn,
		inChan: make(chan *RobotHubMsg, 1000),
		outChan: make(chan *RobotHubMsg, 1000),
		closeChan: make(chan byte),
		isClosed: false,
	}
}

func (wsConn *WsConnection)WsReadLoop() {
	for {
		// 读一个message
		_, data, err := wsConn.wsSocket.ReadMessage()

		if err != nil {
			goto ERR
		}
		req := ParseRobotMsg(data)
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <- wsConn.closeChan:
			goto closed
		}
	}
ERR:
	wsConn.WsClose()
closed:
	clientClose()
}

func (wsConn *WsConnection)WsWriteLoop() {
	var(
		jsonMsg []byte
		err error
	)
	for {
		select {
		// 取一个应答
		case msg := <- wsConn.outChan:
			// 写给websocket
			//log.Println(msg)

			if jsonMsg,err=msg.ToBytes();err!=nil{
				goto ERR
			}
			if err := wsConn.wsSocket.WriteMessage(websocket.TextMessage,jsonMsg); err != nil {
				goto ERR
			}

		case <- wsConn.closeChan:
			goto closed
		}
	}
ERR:
	wsConn.WsClose()
closed:
	clientClose()
}


func (wsConn *WsConnection)Heartbeat(interval int,msg *RobotHubMsg) {

	for {
		time.Sleep(time.Duration(interval )* time.Second)
		if err := wsConn.WsWrite(msg); err != nil {
			fmt.Println("heartbeat fail")
			wsConn.WsClose()
			break
		}
	}
}

func (wsConn *WsConnection)ProcLoop(procHandler func(msg *RobotHubMsg) ) {


	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.WsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		procHandler(msg)
		//err = wsConn.WsWrite(msg)
		//if err != nil {
		//	fmt.Println("write fail")
		//	break
		//}
	}
}


func (wsConn *WsConnection)WsWrite(msg *RobotHubMsg) error {
	select {
	case wsConn.outChan <- msg:
	case <- wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *WsConnection)WsRead() (*RobotHubMsg, error) {
	select {
	case msg := <- wsConn.inChan:
		return msg, nil
	case <- wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *WsConnection)WsClose() {
	log.Error("close conn")
	wsConn.wsSocket.Close()
	delete(HubWsConn, wsConn.clientName)
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

func clientClose(){
	fmt.Println("already close...")
}