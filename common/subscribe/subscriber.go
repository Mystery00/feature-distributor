package subscribe

type Subscriber struct {
	Name     string
	Notifier Notify
}

type Notify func(event ChannelEvent)
