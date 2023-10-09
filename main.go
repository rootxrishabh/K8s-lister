package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)
func main() {
	kubeconf, err := rest.InClusterConfig()
	if err != nil{
		fmt.Printf("Error %s retrieving cluster config through service account\n", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(kubeconf)
	if err != nil{
		fmt.Printf("Encountered an error while creating a clientset\n", err.Error())
	}
	
	pod, err := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	if err != nil{
		fmt.Printf("Encountered retrieveing pods\n", err.Error())
	}

	for _, pods := range pod.Items{
		fmt.Printf("%s", pods.Name)
	}
	// Getting deployments via the clientset

	deployment, err := clientset.AppsV1().Deployments("default").List(context.Background(), v1.ListOptions{})
	if err != nil{
		fmt.Printf("Encountered retrieveing deployments\n", err.Error())
	}

	for _, dep := range deployment.Items{
		fmt.Printf("%s", dep.Name)
	}
}
