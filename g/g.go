package g

import (
	"sync"
)




type SafeToken struct {
	sync.RWMutex
	Token string
}

var (
	TokenSet = &SafeToken{}
)

func (this *SafeToken) Reinit(token string)  {
	this.Lock()
	defer this.Unlock()
	this.Token = token
}

func(this *SafeToken) Get() string {
	this.Lock()
	defer this.Unlock()
	return this.Token
}
