package main

import (
	"fmt"
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

	// Manage resource
	err = kubeclient.Delete("besu-validators-configmap", &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigMap Deleted successfully!")
}
