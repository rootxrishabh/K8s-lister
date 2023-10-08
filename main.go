package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)
func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/rootxrishabh/.kube/config", "Kube config file for cluster authentication")
	kubeconf, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := kubernetes.NewForConfig(kubeconf)
	
	pod, _ := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})

	for _, pods := range pod.Items{
		fmt.Printf("%s", pods.Name)
	}
	// Getting deployments via the clientset

	deployment, _ := clientset.AppsV1().Deployments("default").List(context.Background(), v1.ListOptions{})

	for _, dep := range deployment.Items{
		fmt.Printf("%s", dep.Name)
	}
}
