package main

import (
	"container/list"
	"sync"
)

type LRU struct {
	mux   sync.RWMutex
	size  int
	list  *list.List
	cache map[interface{}]*list.Element
}

type Data struct {
	key   string
	value interface{}
}

func NewLRU(size int) *LRU {
	return &LRU{
		size:  size,
		list:  list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

func (l *LRU) Get(key string) (interface{}, bool) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if ele, hit := l.cache[key]; hit {
		l.list.MoveToFront(ele)
		return ele.Value.(*Data).value, true
	}
	return nil, false
}

func (l *LRU) Set(key string, value interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.cache == nil {
		l.cache = make(map[interface{}]*list.Element)
		l.list = list.New()
	}
	if ele, ok := l.cache[key]; ok {
		l.list.MoveToFront(ele)
		return
	}

	ele := l.list.PushFront(&Data{
		key:   key,
		value: value,
	})
	l.cache[key] = ele

	if l.list.Len() > l.size {
		l.RemoveLatestElement()
	}
}

func (l *LRU) Delete(key string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.cache == nil {
		return
	}
	if ele, hit := l.cache[key]; hit {
		delete(l.cache, key)
		l.list.Remove(ele)
	}
}

func (l *LRU) RemoveLatestElement() {
	latestElement := l.list.Back()
	if latestElement == nil {
		return
	}
	l.list.Remove(latestElement)
	key := latestElement.Value.(*Data).key
	delete(l.cache, key)
}
