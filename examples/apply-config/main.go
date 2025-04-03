package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	optimizeapplyconfigv1 "github.com/thestormforge/agent-go/pkg/applyconfigurations/optimize/v1"
	optimizeclient "github.com/thestormforge/agent-go/pkg/stormforge"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/util/homedir"
)

func main() {
	// Flags to input pod name and namespace
	var (
		workloadName string
		namespace    string
		resourceName string
		schedule     string
	)
	flag.StringVar(&workloadName, "workload", "", "Name of the workload optimize object to be created, same as the workload to be pointed to")
	flag.StringVar(&namespace, "namespace", "default", "Namespace for the workload")
	flag.StringVar(&resourceName, "resource", "", "Resource of the workload")
	flag.StringVar(&schedule, "schedule", "P1D", "Schedule of the recommendation")

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

	clientset, err := optimizeclient.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	woApply := optimizeapplyconfigv1.WorkloadOptimizer(workloadName, namespace).
		WithSpec(
			optimizeapplyconfigv1.WorkloadOptimizerSpec().
				WithTargetRef(
					&optimizeapplyconfigv1.TargetRefApplyConfiguration{
						Kind: &resourceName,
						Name: &workloadName,
					},
				).
				WithSchedule(schedule),
		)

	// printing the object in JSON format
	data, err := json.MarshalIndent(woApply, "", "  ")
	if err != nil {
		fmt.Println("error marshalling:", err)
	} else {
		fmt.Println(string(data))
	}

	wo, err :=
		clientset.OptimizeV1().
			WorkloadOptimizers(namespace).
			Apply(context.Background(), woApply, v1.ApplyOptions{FieldManager: "stormforge"})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Workload Optimizer '%s' created successfully in namespace '%s'\n", wo.Name, wo.Namespace)
}
