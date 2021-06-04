package designpattern

import "sync"

type singleton struct {
	Name string
}

var instance *singleton

// 不考虑并发
func GetInstance() *singleton {
	if instance == nil{
		instance = &singleton{"pj"}
	}
	return instance
}

// sync.Once 原子操作并加锁
var once sync.Once

func GetInstanceWithOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}