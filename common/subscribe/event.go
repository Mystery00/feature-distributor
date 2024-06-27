package subscribe

type ChannelEvent struct {
	ChangeType string
	ProjectId  int64
	ProjectKey string
	ToggleId   *int64
	ToggleKey  *string
}
