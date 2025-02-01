package socket

type Room struct {
	clients map[string]*Client
	game    *Game
}

func newRoom() *Room {
	return &Room{
		clients: make(map[string]*Client),
		game:    newGame(),
	}
}

func (r *Room) addClient(c *Client) {
	r.clients[c.id] = c
}
