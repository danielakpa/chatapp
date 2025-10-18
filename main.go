package main

// client is a single chatting user in a room
type client struct {

	// a web socket for this user
	socket *websocket.com

	// receive is a channel to receive messages from other client
	receive chan []byte

	room *room
}
