package notify

import (
	"feature-distributor/common/subscribe"
	"feature-distributor/core/db/model"
)

func ProjectChange(project model.Project) {
	subscribe.Pub(subscribe.ChannelEvent{
		ProjectId:  project.ID,
		ProjectKey: project.Key,
	})
}

func ToggleChange(projectKey string, toggle model.Toggle) {
	subscribe.Pub(subscribe.ChannelEvent{
		ProjectId:  toggle.ProjectID,
		ProjectKey: projectKey,
		ToggleId:   &toggle.ID,
		ToggleKey:  &toggle.Key,
	})
}
