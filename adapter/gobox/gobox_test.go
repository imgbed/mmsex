package gobox

import "testing"

var gb = NewGoBox()
func TestGoBox_Set(t *testing.T) {
	gb.Set("key",1)
	gb.Set("key2",2)
	gb.Hset("tab","tabkey",222)
}

func TestGoBox_Get(t *testing.T) {
	//TestGoBox_Set(t)
	gb.initdata()
	t.Log(gb.Get("key2"))
	t.Log(gb.Hget("tab","tabkey"))
}
