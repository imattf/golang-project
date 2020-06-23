//build the client

//built ready to review

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/imattf/golang-project/04-grpc/04-chat/01-chat/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started...")

	//validate command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Must have a url to connect to as the first argument, and a username as the second argument")
		return
	}

	//createa a context -??
	ctx := context.Background()

	// establish a connection to grpc server and "let client know that is expected"
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//create the new chat client over the connection
	// initiate that stream for both directions
	c := chat.NewChatClient(conn)
	stream, err := c.Chat(ctx)
	if err != nil {
		panic(err)
	}

	//create a go channel
	waitc := make(chan struct{})
	go func() {
		for {
			//map9: messge comes out here from server
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			} else if err != nil {
				panic(err)
			}
			//map10: messge gets printed out to client here
			fmt.Println(msg.User + ": " + msg.Message)
		}

	}()

	fmt.Println("connection established, type \"quit\" or use ctrl+c to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//map1: message gets created here
		//...when someone types it comes in on this function
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}

		//map2: message gets sent out here to the server
		err := stream.Send(&chat.ChatMessage{
			User:    os.Args[2],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}

	<-waitc

}
