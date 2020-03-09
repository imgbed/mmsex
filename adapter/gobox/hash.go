package gobox

import "sync"

type HashBox map[interface{}]map[interface{}]interface{}

func (box HashBox) Store(tab, key, val interface{}) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	if _, ok := box.Load(tab, key); !ok {
		var tmpmap = map[interface{}]interface{}{
			key: val,
		}
		box[tab] = tmpmap
	} else {
		box[tab][key] = val
	}
}
func (box HashBox) Load(tab, key interface{}) (res interface{}, ok bool) {
	var mu sync.RWMutex
	mu.RLock()
	defer mu.RUnlock()
	if res1, ok1 := box[tab]; ok1 {
		res, ok = res1[key]
	}
	return
}
