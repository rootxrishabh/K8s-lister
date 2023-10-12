package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconf := flag.String("kubeconfing", "/Users/rootxrishabh/.kube/config", "location to kubernetes config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconf)
	if err != nil {
		fmt.Printf("Error %s fetching config from host", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s while fetching config from cluster", err.Error())
		}
	}
	if err != nil {
		fmt.Printf("Error %s retrieving cluster config through service account\n", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Encountered an error while creating a clientset\n", err.Error())
	}

	pod, err := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered retrieveing pods\n", err.Error())
	}

	for _, pods := range pod.Items {
		fmt.Printf("%s\n", pods.Name)
	}

	// Getting deployments via the clientset
	deployment, err := clientset.AppsV1().Deployments("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered retrieveing deployments\n", err.Error())
	}

	// prints the label on the node
	resource, err := clientset.CoreV1().Nodes().Get(context.Background(), "kind-control-plane", v1.GetOptions{})
	if err != nil {
		fmt.Printf("Error %s encountered while getting node", err.Error())
	}
	i := 1
	for nodes := range resource.GetLabels() {
		fmt.Printf("Labels no %v on the node is %v \n", i, nodes)
		i++
	}

	for _, dep := range deployment.Items {
		fmt.Printf("%s", dep.Name)
	}
	fmt.Print("\n")

	// Creating an informer factory which will help keep a realtime track of changes inside the cluster for various resources
	// informerFactory := informers.NewSharedInformerFactory(clientset, 10*time.Minute)
	// podInformer := informerFactory.Core().V1().Pods()
	// podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: func(new interface{}) {
	// 		fmt.Printf("Write your bussiness logic to handle AddFunc\n")
	// 	},
	// 	UpdateFunc: func(oldObj, newObj interface{}) {
	// 		fmt.Printf("Write your bussiness logic to handle UpdateFunc\n")
	// 	},
	// 	DeleteFunc: func(obj interface{}) {
	// 		fmt.Printf("Write your bussiness logic to handle DeleteFunc\n")
	// 	},
	// })

	// informerFactory.Start(wait.NeverStop)
	// informerFactory.WaitForCacheSync(wait.NeverStop)
	// if realTimePodInfo, err := podInformer.Lister().Pods("default").Get("nginx"); err != nil {
	// 	fmt.Printf("Error %s while retrieveing pods from he informer", err.Error())
	// } else {
	// 	fmt.Printf("Pods retrieved %T", realTimePodInfo)
	// }
}
