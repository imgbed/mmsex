package gobox

import (
	"github.com/gohouse/golib/tgbot/adapter"
	"log"
)
type StoreObject struct {
	//Set map[interface{}]interface{}
	Set SetBox
	Hash HashBox
}
type Store2 map[string]interface{}

type GoBox struct {
	 *StoreObject
}

func NewGoBox() *GoBox {
	gb := &GoBox{}
	gb.initdata()
	//go gb.persist()
	return gb
}

func (gb *GoBox) Set(key, val interface{}) {
	gb.StoreObject.Set.Store(key, val)
	gb.persist()
}

func (gb *GoBox) Get(key interface{}) (interface{},bool) {
	return gb.StoreObject.Set.Load(key)
}

func (gb *GoBox) Hset(tab, key, val interface{}) {
	gb.StoreObject.Hash.Store(tab, key, val)
	gb.persist()
}

func (gb *GoBox) Hget(tab, key interface{}) (interface{},bool) {
	return gb.StoreObject.Hash.Load(tab, key)
}

func (gb *GoBox) persist()  {
	//for {
	//	time.Sleep(1*time.Second)

		err := adapter.NewGobPersist("gobox.gob").Store(gb.StoreObject)
		if err!=nil {
			log.Println(err.Error())
		}
	//}
}

func (gb *GoBox) initdata()  {
	var opt = StoreObject{
		Set: make(map[interface{}]interface{}),
		Hash: make(map[interface{}]map[interface{}]interface{}),
	}
	err := adapter.NewGobPersist("gobox.gob").Load(&opt)
	if err!=nil {
		log.Println(err.Error())
	}
	gb.StoreObject = &opt
}