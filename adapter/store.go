package adapter

type Storer interface {
	Set(key,val interface{})
	Get(key interface{}) (val interface{})
	Hset(tab string,key,val interface{})
	Hget(tab string,key interface{}) (val interface{})
}
