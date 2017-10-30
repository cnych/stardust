package syncx

import (
	"sync"
)

func LockDo(l *sync.Mutex, action func()) {
	if l != nil {
		l.Lock()
		defer l.Unlock()
	}
	action()
}

func WLockDo(rwm *sync.RWMutex, action func()) {
	if rwm != nil {
		rwm.Lock()
		defer rwm.Unlock()
	}
	action()
}

func RLockDo(rwm *sync.RWMutex, action func()) {
	if rwm != nil {
		rwm.RLock()
		defer rwm.RUnlock()
	}
	action()
}
