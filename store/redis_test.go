package store

import (
	"github.com/go-redis/redis/v7"
	"strings"
	"testing"
)

func TestRedis(t *testing.T) {
	rds := BootRedis()

	res,err := rds.HGet("aaa","a").Result()
	if err!=nil {
		t.Fatal(err.Error(), err==redis.Nil)
	}
	t.Log(res)
}

func TestImg(t *testing.T) {
	// 为 mm和gif,添加索引
	rds := BootRedis()
	var cursor uint64
	var total int64
	var count int64 = 100
	for total <= 50000 {
		var result []string
		var err error
		var curr uint64
		var scan = rds.HScan("img", cursor, "", count)
		result, curr, err = scan.Result()
		if err!=nil {
			t.Fatal(err.Error())
		}
		//t.Log(cursor)
		for _,v := range result{
			if !strings.HasPrefix(v,"http") {
				//t.Log(v)
				rds.SAdd("set:img",v)
			}
		}
		cursor = curr
		total += count
	}

	t.Log(cursor,total)
}

func TestGif(t *testing.T) {
	// 为 mm和gif,添加索引
	rds := BootRedis()
	var cursor uint64
	var total int64
	var count int64 = 1
	for total < 121 {
		var result []string
		var err error
		var curr uint64
		var scan = rds.HScan("gif", cursor, "", count)
		result, curr, err = scan.Result()
		if err!=nil {
			t.Fatal(err.Error())
		}
		//t.Log(cursor)
		for _,v := range result{
			if !strings.HasPrefix(v,"http") {
				//t.Log(v)
				rds.SAdd("set:gif",v)
			}
		}
		cursor = curr
		total += count
	}

	t.Log(cursor,total)
}