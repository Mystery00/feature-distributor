package subscribe

type ChannelEvent struct {
	ProjectId  int64
	ProjectKey string
	ToggleId   *int64
	ToggleKey  *string
}
