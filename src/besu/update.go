package main

import (
	"fmt"
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
	kubeclient := clientset.CoreV1().ConfigMaps("besu")

	// Create resource object
	object := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "besu-validators-configmap",
			Namespace: "besu",
			Labels: map[string]string{
				"app": "besu-validators-configmap",
			},
		},
		Data: map[string]string{
			"validator1PubKey": "5d812c3c25ff398ab416968fce9009c2be7ed70a87abc8ea30bd667ce17a9287a6341fbf6ce757bb8148436c39c71296639ea81afcc94cdf908b6e1344f26188",
			"validator2PubKey": "b2fba529681ea7f4619556753d40c8689b936fb1c621bc91f94d2938eb58c285d4911457ae4887b9c3bd593b2d608d319c6dc384d6acae2d043a4657029178d3",
			"validator3PubKey": "00b20ab6a385a2403d64637b3d93cb6d83215a08f29adb6feb4b8bf03387b734444e8b060f53150dea4b9b897823540d19918c13d6f57a5153d190b5fad7bf51",
			"validator4PubKey": "5fc1f8dc9f0c03087128e4bd724530e883d7de1a431269876dff9c95b8952f73c7e85ac7b49d85a2ad4950e967319482af435e07a0eab0a98d98449437787a00",
		},
	}

	// Manage resource
	_, err = kubeclient.Update(object)
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigMap Updated successfully!")
}
