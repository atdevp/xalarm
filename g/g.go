package g

import (
	"sync"
)



type SafeGlobalToken struct {
	sync.RWMutex
	Token string
}

type SafeLocalToken struct {
	sync.RWMutex
	Token string
}

var (
	GlobalTokenSet = &SafeGlobalToken{}
	LocalTokenSet = &SafeLocalToken{}
)

func (this *SafeGlobalToken) Reinit(token string)  {
	this.Lock()
	defer this.Unlock()
	this.Token = token
}

func(this *SafeGlobalToken) Get() string {
	this.Lock()
	defer this.Unlock()
	return this.Token
}

func (this *SafeLocalToken) Reinit(token string)  {
	this.Lock()
	defer this.Unlock()
	this.Token = token
}

func(this *SafeLocalToken) Get() string {
	this.Lock()
	defer this.Unlock()
	return this.Token
}