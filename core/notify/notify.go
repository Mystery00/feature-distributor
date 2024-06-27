package notify

import (
	"feature-distributor/common/subscribe"
	"feature-distributor/core/db/model"
)

func ProjectChange(changeType string, project model.Project) {
	subscribe.Pub(subscribe.ChannelEvent{
		ChangeType: changeType,
		ProjectId:  project.ID,
		ProjectKey: project.Key,
	})
}

func ToggleChange(changeType string, projectKey string, toggle model.Toggle) {
	subscribe.Pub(subscribe.ChannelEvent{
		ChangeType: changeType,
		ProjectId:  toggle.ProjectID,
		ProjectKey: projectKey,
		ToggleId:   &toggle.ID,
		ToggleKey:  &toggle.Key,
	})
}
