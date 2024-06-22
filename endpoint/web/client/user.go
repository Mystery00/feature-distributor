package client

import "feature-distributor/endpoint/pb"

type reqUser struct {
	RolloutKey string            `json:"rolloutKey" required:"true" binding:"required"`
	Attributes map[string]string `json:"attributes" required:"true" binding:"required"`
}

func (r reqUser) buildReqUser() *pb.ReqUser {
	return &pb.ReqUser{
		RolloutKey: r.RolloutKey,
		Attributes: r.Attributes,
	}
}
