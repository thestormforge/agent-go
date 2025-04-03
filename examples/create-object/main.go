package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	optimizev1 "github.com/thestormforge/agent-go/pkg/api/optimize/v1"
	optimizescheme "github.com/thestormforge/agent-go/pkg/stormforge/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/client-go/util/homedir"
)

func main() {
	// Flags to input pod name and namespace
	var (
		workloadName string
		namespace    string
		resourceName string
	)
	flag.StringVar(&workloadName, "workload", "", "Name of the workload optimize object to be created, same as the workload to be pointed to")
	flag.StringVar(&namespace, "namespace", "default", "Namespace for the workload")
	flag.StringVar(&resourceName, "resource", "", "Resource of the workload")

	flag.Parse()

	if workloadName == "" {
		fmt.Println("Workload name must be provided using --workload")
		os.Exit(1)
	}

	if resourceName == "" {
		fmt.Println("Resource name must be provided using --resource")
		os.Exit(1)
	}

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// adding CRD to the scheme
	if err := optimizescheme.AddToScheme(scheme.Scheme); err != nil {
		log.Fatal(err)
	}

	// Create a new client using controller-runtime's client package
	k8sClient, err := client.New(config, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		log.Fatal(err)
	}

	// Workload Optimizer object definition
	wo := &optimizev1.WorkloadOptimizer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      workloadName,
			Namespace: namespace,
		},
		Spec: optimizev1.WorkloadOptimizerSpec{
			TargetRef: &optimizev1.TargetRef{
				Kind: resourceName,
				Name: workloadName,
			},
		},
	}

	// Create the WO object
	if err := k8sClient.Create(context.Background(), wo); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Workload Optimizer '%s' created successfully in namespace '%s'\n", wo.Name, wo.Namespace)
}
