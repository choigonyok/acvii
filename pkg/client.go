package pkg

import (
	"context"
	"fmt"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func InClusterConfig() {

	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, pod := range pods.Items {
		println("Pod "+strconv.Itoa(i)+" Name : ", pod.Name)
		println("Pod "+strconv.Itoa(i)+" Name : ", pod.Namespace)
		for j, container := range pod.Spec.Containers {
			println("	Container " + strconv.Itoa(j) + " Name : " + container.Name)
			println("	Container " + strconv.Itoa(j) + " Image : " + container.Image)

			println("	Container " + strconv.Itoa(j) + " Memory Limit : " + container.Resources.Limits.Memory().String())
			println("	Container " + strconv.Itoa(j) + " CPU Limit : " + container.Resources.Limits.Cpu().String())

			println("	Container " + strconv.Itoa(j) + " Requests Limit : " + container.Resources.Requests.Memory().String())
			println("	Container " + strconv.Itoa(j) + " Request Limit : " + container.Resources.Requests.Cpu().String())

			println("	Container "+strconv.Itoa(j)+" Resources Args : ", container.ResizePolicy)

			for ar, arg := range container.Args {
				println("		Argument "+strconv.Itoa(ar)+" : ", arg)
			}
			for z, zi := range container.Ports {
				println("		Port "+strconv.Itoa(z)+" : ", zi.Name, " : ", zi.HostPort, " / ", zi.ContainerPort)
			}
			for l, li := range container.Resources.Claims {
				println("		Claim "+strconv.Itoa(l)+" : ", li.Name)
			}
		}
		println("")
		println("")
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
