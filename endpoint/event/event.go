package event

import (
	"context"
	"feature-distributor/common/subscribe"
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var ctx, cancelFunc = context.WithCancel(context.Background())

func Init() {
	subscriberName, err := os.Hostname()
	if err != nil {
		subscriberName = "feature-distributor-endpoint"
	}
	client := grpc.GetEventClient()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				createEventGrpcSubscriber(client, subscriberName)
			}
		}
	}()
}

func createEventGrpcSubscriber(client pb.EventServiceClient, subscriberName string) {
	logrus.Infof("create event subscriber: %v", subscriberName)
	stream, err := client.SubscribeEvents(ctx, &pb.Subscriber{
		Name: subscriberName,
	})
	if err != nil {
		logrus.Errorf("subscribe event failed, %v", err)
		time.Sleep(5 * time.Second)
		return
	}
FOR:
	for {
		select {
		case <-ctx.Done():
			break FOR
		default:
			event, err := stream.Recv()
			if err == io.EOF {
				break FOR
			}
			if err != nil {
				logrus.Warnf("receive event failed: %v", err)
				break FOR
			}
			data := event.GetData()
			if data == nil {
				logrus.Infof("receive event type: %v", event.Type)
				continue
			}
			logrus.Debugf("receive event: %v", data)
			subscribe.Pub(subscribe.ChannelEvent{
				ChangeType: event.GetType(),
				ProjectId:  data.GetProjectId(),
				ProjectKey: data.GetProjectKey(),
				ToggleId:   data.ToggleId,
				ToggleKey:  data.ToggleKey,
			})
		}
	}
	err = stream.CloseSend()
	if err != nil {
		logrus.Warnf("close send failed: %v", err)
	}
}

func Close() {
	cancelFunc()
}
