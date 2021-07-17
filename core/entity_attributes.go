package core

import (
	"net/url"
	"strconv"
	"sync"
)

type entityAttributes struct {
	data map[string]string
}

type EntityAttributes interface {
	Add(key, val string)
	Get(key string) string
	Del(key string)
	QueryString() string
}

func (e *entityAttributes) Add(key, val string) {
	e.data[key] = val
}

func (e *entityAttributes) Get(key string) string {
	return e.data[key]
}

func (e *entityAttributes) Del(key string) {
	delete(e.data, key)
}

func (e *entityAttributes) QueryString() string {
	return e.urlValues().Encode()
}

func (e *entityAttributes) urlValues() url.Values {
	ret := url.Values{}
	wg := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}
	index := 0

	for key, val := range e.data {
		wg.Add(1)
		index++
		go func(i int, k, v string) {
			defer wg.Done()
			rwMutex.Lock()
			ret.Add("attr-name"+strconv.Itoa(i), k)
			ret.Add("attr-value"+strconv.Itoa(i), v)
			rwMutex.Unlock()
		}(index, key, val)
	}

	wg.Wait()
	return ret
}

func NewEntityAttributes() EntityAttributes {
	return &entityAttributes{
		data: map[string]string{},
	}
}
