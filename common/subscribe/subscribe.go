package subscribe

import (
	"github.com/sirupsen/logrus"
	"sync"
)

func Pub(event ChannelEvent) {
	mu.RLock()
	defer mu.RUnlock()
	for _, subscriber := range subscriberMap {
		subscriber.Notifier(event)
	}
}

var mu sync.RWMutex
var subscriberMap = make(map[string]Subscriber)

func Sub(name string, notify Notify) {
	subscriber := Subscriber{
		Name:     name,
		Notifier: notify,
	}
	mu.Lock()
	defer mu.Unlock()
	logrus.Debugf("sub [%s]", name)
	subscriberMap[name] = subscriber
}

func Unsub(name string) {
	mu.Lock()
	defer mu.Unlock()
	logrus.Debugf("unsub [%s]", name)
	delete(subscriberMap, name)
}
