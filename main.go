package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


// Function to call in case you want to create a pod. NOTE: this function is optional and it does not add to the purpose of this project. It was created just for learning purposes.
func k8sClient(){
	config := flag.String("Kubeconfig", "/Users/rootxrishabh/.kube/config", "Kubeconfig from hostOS")
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", *config)
	if err != nil{
		fmt.Printf("Error %s deriving config from host", err.Error())
		kubeconfig, err = rest.InClusterConfig()
		if err != nil{
			fmt.Printf("Error %s deriving config from cluster. Check roles!", err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil{
		fmt.Printf("Error creating client set from config")
	}

	pod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mysql-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "mysql",
					Image: "mysql",
					Env: []v1.EnvVar{
						{
							Name:  "MYSQL_ROOT_PASSWORD",
							Value: "your-root-password",
						},
					},
					Ports: []v1.ContainerPort{
						{
							ContainerPort: 3306,
						},
					},
				},
			},
		},
	}

	clientset.CoreV1().Pods("default").Create(context.Background(), &pod, metav1.CreateOptions{})
}

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

	pod, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered retrieveing pods\n", err.Error())
	}

	for _, pods := range pod.Items {
		fmt.Printf("%s\n", pods.Name)
	}

	// Getting deployments via the clientset
	deployment, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered retrieveing deployments\n", err.Error())
	}

	// prints the label on the node
	resource, err := clientset.CoreV1().Nodes().Get(context.Background(), "kind-control-plane", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error %s encountered while getting node", err.Error())
	}
	i := 1
	for nodes := range resource.GetLabels() {
		fmt.Printf("Labels no %v on the node is %v \n", i, nodes)
		i++``
	}

	for _, dep := range deployment.Items {
		fmt.Printf("%s", dep.Name)
	}
	fmt.Print("\n")

}
