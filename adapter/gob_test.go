package adapter

import "testing"

type User struct {
	Name string
	Age int
}
var ps = NewGobPersist("xxx.gob")
var ps2 = NewGobPersist("xxx2.gob")
func TestStore(t *testing.T) {
	var u = User{
		Name: "kevin",
		Age:  18,
	}
	err := ps.Store(&u)
	if err!=nil {
		t.Error(err.Error())
		return
	}
	ps2.Store(map[string][]string{"imgs":{"xxxx"},"gifs":{"fffff"}})
	t.Log("store success: xxx.gob")
}

func TestLoad(t *testing.T) {
	var u User
	err := ps.Load(&u)
	if err!=nil {
		t.Error(err.Error())
		return
	}
	var img = map[string][]string{}
	ps2.Load(&img)
	t.Logf("load success: %+v", u)
	t.Logf("load success: %+v", img)
}
