//build the server

//built ready to review

package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/imattf/golang-project/04-grpc/04-chat/01-chat/chat"
	"google.golang.org/grpc"
)

//Connection is a ...
type Connection struct {
	conn chat.Chat_ChatServer
	send chan *chat.ChatMessage
	quit chan struct{}
}

//NewConnection is a ...
func NewConnection(conn chat.Chat_ChatServer) *Connection {
	c := &Connection{
		conn: conn,
		send: make(chan *chat.ChatMessage),
		quit: make(chan struct{}),
	}
	go c.start()
	return c
}

//Close is a ...
func (c *Connection) Close() error {
	close(c.quit)
	close(c.send)
	return nil
}

//Send is a ...
//map 6: comes in here
func (c *Connection) Send(msg *chat.ChatMessage) {
	defer func() {
		//ignore any errors sending on a closed channel
		recover()
	}()
	//map7: gets send thru this send channel
	c.send <- msg
}

func (c *Connection) start() {
	running := true
	for running {
		select {
		//map8: this is the other end of the send channel
		//...msg comes out here and sent back to client here
		case msg := <-c.send:
			c.conn.Send(msg)
		case <-c.quit:
			running = false
		}
	}
}

//GetMessages is a ...
//map3: message comes in from client here
//...via the Recieve message on a connection
func (c *Connection) GetMessages(broadcast chan<- *chat.ChatMessage) error {
	for {
		msg, err := c.conn.Recv()
		if err == io.EOF {
			c.Close()
			return nil
		} else if err != nil {
			c.Close()
			return err
		}
		go func(msg *chat.ChatMessage) {
			select {
			//map4: messages gets sent out on this broadcast channel
			case broadcast <- msg:
			case <-c.quit:
			}
		}(msg)
	}
}

//ChatServer is a ...
type ChatServer struct {
	broadcast   chan *chat.ChatMessage
	quit        chan struct{}
	connections []*Connection
	connLock    sync.Mutex
}

//NewChatServer is a ...
func NewChatServer() *ChatServer {
	srv := &ChatServer{
		broadcast: make(chan *chat.ChatMessage),
		quit:      make(chan struct{}),
	}
	go srv.start()
	return srv
}

//Close is a ...
func (c *ChatServer) Close() error {
	close(c.quit)
	return nil
}

func (c *ChatServer) start() {
	running := true
	for running {
		select {
		//map 5: this is the other end of the broadcast channel
		//...msg comes out this end
		//...the server goes thru all the connections
		//...and broadcasts the message out to those connections
		//...via Send(msg)
		case msg := <-c.broadcast:
			c.connLock.Lock()
			for _, v := range c.connections {
				go v.Send(msg)
			}
			c.connLock.Unlock()
		case <-c.quit:
			running = false
		}
	}
}

//Chat is a ...
func (c *ChatServer) Chat(stream chat.Chat_ChatServer) error {
	conn := NewConnection(stream)

	c.connLock.Lock()
	c.connections = append(c.connections, conn)
	c.connLock.Unlock()

	err := conn.GetMessages(c.broadcast)

	c.connLock.Lock()
	for i, v := range c.connections {
		if v == conn {
			c.connections = append(c.connections[:], c.connections[i+1:]...)
		}
	}
	c.connLock.Unlock()

	return err
}

func main() {

	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	srv := NewChatServer()
	chat.RegisterChatServer(s, srv)

	fmt.Println("Server started at port 8080...")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}

}
