package strategies

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/redmeros/htrade/internal/dirty"
	"github.com/redmeros/htrade/models"
	h "github.com/redmeros/htrade/web/helpers"

	"github.com/gorilla/websocket"
)

func findStrategy(codeName string) (models.IStrategy, error) {
	switch codeName {
	case "buyHold":
		return dirty.NewBuyAndHoldStrategy(), nil
	default:
		return nil, fmt.Errorf("Cannot find strategy with CodeName: %s", codeName)
	}
}

// StrategyInfo to endpoint dla info o strategii
func StrategyInfo(c *gin.Context) {
	codeName := c.Param("codeName")
	if len(codeName) == 0 {
		h.Bad(c, "Cannot get codeName", 400)
		return
	}

	s, err := findStrategy(codeName)
	if err != nil {
		h.Badf(c, "Error: %s", 400, err.Error())
	}

	c.JSON(200, &s)
}

var myhub = newHub()
var wsh = wsHandler{h: myhub}

func HandleWs(c *gin.Context) {
	wsh.ServeHTTP(c.Writer, c.Request)
}

//https://github.com/utiq/go-in-5-minutes/blob/master/episode4/hub.go
type hub struct {
	connectionsMx sync.RWMutex
	connections   map[*connection]struct{}
	broadcast     chan []byte

	logMx sync.RWMutex
	log   [][]byte
}

func newHub() *hub {
	h := &hub{
		connectionsMx: sync.RWMutex{},
		broadcast:     make(chan []byte),
		connections:   make(map[*connection]struct{}),
	}

	go func() {
		for {
			msg := <-h.broadcast
			h.connectionsMx.RLock()
			for c := range h.connections {
				select {
				case c.send <- msg:
				case <-time.After(1 * time.Second):
					log.Printf("shutting down connection %s", c)
					h.removeConnection(c)
				}
			}
			h.connectionsMx.RUnlock()
		}
	}()
	return h
}

func (h *hub) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{} //?? dlaczego??
}

func (h *hub) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send)
	}
}

type connection struct {
	send         chan []byte
	cancelToken  chan bool
	otherChannel chan bool

	h *hub
}

func startTesting(out chan []byte, cancel chan bool, other chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		i++
		select {
		case <-ticker.C:
			if i == 10 {
				ticker.Stop()
			}
			msg := models.NewNotify(fmt.Sprintf("Still testing: %d", i))
			out <- msg.Json()
		case <-other:
			msg := models.NewNotify("Inny kanal !")
			out <- msg.Json()
		case <-cancel:
			msg := models.NewError("Canceled by user???", nil)
			out <- msg.Json()
			return
		}
	}
}

func (c *connection) processMessage(bmessage []byte) {
	log.Print("Processing message")

	var inMsg models.WsMessage
	json.Unmarshal(bmessage, &inMsg)

	if inMsg.Action == "start_testing" {
		rmsg := models.NewNotify("testing started")
		go startTesting(c.send, c.cancelToken, c.otherChannel)
		c.send <- rmsg.Json()
	} else if inMsg.Action == "stop_testing" {
		rmsg := models.NewNotify("testing stopped")
		c.cancelToken <- true
		c.send <- rmsg.Json()
	} else if inMsg.Action == "other_action" {
		c.otherChannel <- true
	} else {
		rmsg := models.NewError("cannot recognize msg", inMsg)
		c.send <- rmsg.Json()
	}
}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	c.cancelToken = make(chan bool)
	c.otherChannel = make(chan bool)
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}
		log.Println("reader: " + string(message))
		c.processMessage(message)
		// c.h.broadcast <- message
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for message := range c.send {
		err := wsConn.WriteMessage(websocket.TextMessage, message)
		log.Println("writer: " + string(message))

		if err != nil {
			break
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsHandler struct {
	h *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}
	c := &connection{send: make(chan []byte, 256), h: wsh.h}
	c.h.addConnection(c)
	defer c.h.removeConnection(c)
	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}
