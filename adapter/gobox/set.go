package gobox

import "sync"

type SetBox map[interface{}]interface{}

func (sb SetBox) Store(key, val interface{})  {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	sb[key] = val
}
func (sb SetBox) Load(key interface{}) (res interface{},ok bool) {
	var mu sync.RWMutex
	mu.RLock()
	defer mu.RUnlock()
	res,ok = sb[key]
	return
}
