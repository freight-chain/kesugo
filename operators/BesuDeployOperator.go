package besu-operators

import (
	"github.com/AlexsJones/kubeops/lib/subscription"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type BesuDeployOperator struct{}

func (BesuDeployOperator) WithElectedResource() interface{} {

	return &v1.Deployment{}
}

func (BesuDeployOperator) WithEventType() []watch.EventType {

	return []watch.EventType {watch.Added, watch.Deleted, watch.Modified}
}

func (BesuDeployOperator) OnEvent(msg subscription.Message) {

	d := msg.Event.Object.(*v1.Deployment)

	log.Debugf("Deployment %s has %d Available replicas",d.Name,d.Status.AvailableReplicas)

}
