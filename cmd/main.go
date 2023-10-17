package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/choigonyok/goopt/internal/test"
	example "github.com/choigonyok/goopt/pkg"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	eg := gin.Default()

	println(test.Test())
	example.Example()
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	for {
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

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		// _, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		// if errors.IsNotFound(err) {
		// 	fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		// 	fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		// } else if err != nil {
		// 	panic(err.Error())
		// } else {
		// 	fmt.Printf("Found example-xxxxx pod in default namespace\n")
		// }
		eg.GET("/", GoOptHandler)
		eg.Run(":8080")
	}
}

func GoOptHandler(c *gin.Context) {
	println("NOTHING")
}
