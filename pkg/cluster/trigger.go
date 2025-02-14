package cluster

import (
	"context"
	"time"

	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"

	"github.com/vanus-labs/vanus/pkg/cluster/raw_client"
)

type triggerService struct {
	client ctrlpb.TriggerControllerClient
}

func newTriggerService(cc *raw_client.Conn) TriggerService {
	return &triggerService{client: raw_client.NewTriggerClient(cc)}
}

func (es *triggerService) RawClient() ctrlpb.TriggerControllerClient {
	return es.client
}

func (es *triggerService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	return raw_client.RegisterHeartbeat(ctx, interval, es.client, reqFunc)
}
