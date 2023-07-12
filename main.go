package main

import (
	"log"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	dClient := discovery.NewDiscoveryClientForConfigOrDie(config)

	serverInfo, err := dClient.ServerVersion()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("kube-api server version: %s.%s\n", serverInfo.Major, serverInfo.Minor)
	log.Printf("kube-api server gitVersion: %s and gitCommit: %s\n", serverInfo.GitVersion, serverInfo.GitCommit)
}
