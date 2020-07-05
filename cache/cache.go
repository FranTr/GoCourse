package cache

import (
	"github.com/pabloos/http/greet"
)

type Cache struct {
	Messages map[string]greet.Greet
}

func (c *Cache) Set(t greet.Greet) {
	c.Messages[t.Name] = t
}

func (c *Cache) Get(k greet.Greet) (greet.Greet, bool) {
	message, found := c.Messages[k.Name]
	return message, found
}

func (c *Cache) GetMessages() map[string]greet.Greet {
	return c.Messages
}
