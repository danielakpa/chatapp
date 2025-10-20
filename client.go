package main

import 
// client is a single chatting user in a room
type client struct {

	// a web socket for this user
	socket *websocket.com

	// receive is a channel to receive messages from other client
	receive chan []byte

	room *room
}

// send messages function
func ( c *client) read(){

	// close the connetion when we are done
	defer c.socket.Closer()
	
	// as long as there is an input, forward it
	for{
		_, msg, err := c.socket.Readmessage()

		if err != nil {
			return
		}

		c.room.forward <- msg

	}
}

// use to recieve messages
func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.recieve {
		err := c.socket.WriteMessage(websocket.Textmessage, msg)
		if err != nil {
			return
		}
	}

}