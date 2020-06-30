package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func main() {
	// Create client
	var kubeconfig string
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kubeclient := clientset.AppsV1().StatefulSets("besu")

	// Create resource object
	object := &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "node",
			Namespace: "besu",
			Labels: map[string]string{
				"app": "node",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: ptrint32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "node",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "node",
					},
					Annotations: map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "9545",
						"prometheus.io/scrape": "true",
					},
				},
				Spec: corev1.PodSpec{},
			},
			ServiceName:    "besu-node",
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{},
		},
	}

	// Manage resource
	_, err = kubeclient.Update(object)
	if err != nil {
		panic(err)
	}
	fmt.Println("StatefulSet Updated successfully!")
}

func ptrint32(p int32) *int32 {
	return &p
}
