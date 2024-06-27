package grpc

import (
	"feature-distributor/common/subscribe"
	"feature-distributor/core/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type EventServer struct {
	pb.UnimplementedEventServiceServer
}

func (EventServer) SubscribeEvents(sub *pb.Subscriber, srv pb.EventService_SubscribeEventsServer) error {
	logrus.Infof("subscribe event for [%s]", sub.GetName())
	_ = srv.Send(&pb.Event{
		Type: "init",
	})
	ctx, cancelFunc := context.WithCancel(context.Background())
	subscribe.Sub(sub.GetName(), func(event subscribe.ChannelEvent) {
		err := srv.Send(&pb.Event{
			Type: event.ChangeType,
			Data: &pb.UpdateData{
				ProjectId:  event.ProjectId,
				ProjectKey: event.ProjectKey,
				ToggleId:   event.ToggleId,
				ToggleKey:  event.ToggleKey,
			},
		})
		if err != nil {
			//判定客户端断连，此时主动关闭服务端的连接
			logrus.Warnf("send update event to [%s] failed: %v", sub.GetName(), err)
			cancelFunc()
		}
	})
	select {
	case <-ctx.Done():
		subscribe.Unsub(sub.GetName())
	case <-srv.Context().Done():
		subscribe.Unsub(sub.GetName())
	}
	return nil
}
