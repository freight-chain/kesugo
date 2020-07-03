package besu-operators

import (
	"github.com/AlexsJones/kubeops/lib/subscription"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type BesuPodOperator struct{}

func (BesuPodOperator) WithElectedResource() interface{} {

	return &v1.Pod{}
}

func (BesuPodOperator) WithEventType() []watch.EventType {

	return []watch.EventType {watch.Added, watch.Deleted, watch.Modified}
}

func (BesuPodOperator) OnEvent(msg subscription.Message) {

	pod := msg.Event.Object.(*v1.Pod)
	log.Debugf("Besu pod event from %s Incoming",pod.Name)
	if pod.Labels["app.kubernetes.io/name"] == "kubeops" {

		log.Debugf("%v",pod)
		existingLabels := pod.Labels

		if _, ok := existingLabels["besu-pod"]; !ok {
			//Let's add a label...
			existingLabels["besu-pod-capsule"] = "besu-capsule"
			pod.SetLabels(existingLabels)
			// Invoke a new pod client interface
			pi := msg.Client.CoreV1().Pods(pod.Namespace)
			if _, err := pi.Update(pod); err != nil {
				log.Error(err)
			}else {
				log.Debug("Added a new label...")
			}
		}
	}
}
