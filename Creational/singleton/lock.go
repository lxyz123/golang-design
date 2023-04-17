package main

import (
	"sync"
	"sync/atomic"
)

// lock
var instance1 *singleton

var lock sync.Mutex

func GetInstanceLock() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		return new(singleton)
	} else {
		return instance1
	}
}

// 使用atomic处理，因为上面加锁影响性能
var instance2 *singleton

var initialized uint32

func GetInstance2() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance1
	}
	lock.Lock()
	defer lock.Unlock()
	if initialized == 0 {
		instance2 = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance2
}

// once.Do

var once sync.Once

var instance3 *singleton

func GetInstance3() *singleton {
	once.Do(func() {
		instance3 = new(singleton)
	})
	return instance3
}
