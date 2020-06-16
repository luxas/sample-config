package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/component-base/config"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyAppConfiguration struct {
	metav1.TypeMeta

	// ClientConnection configures the connection to Kubernetes
	ClientConnection config.ClientConnectionConfiguration
	// LeaderElection configures so the component can be deployed in HA mode on k8s
	LeaderElection config.LeaderElectionConfiguration
	// Server holds configuration settings for the HTTPS server
	Server ServerConfiguration
}

type ServerConfiguration struct {
	// Default: "0.0.0.0"
	// +optional
	Address string
	// Default: 10250
	// +optional
	Port uint32
	// +optional
	TLSCertFile string
	// +optional
	TLSPrivateKeyFile string
}
