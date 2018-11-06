package v1

import (
	apimachineryconfigv1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiserverconfigv1 "k8s.io/apiserver/pkg/apis/config/v1alpha1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyAppConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ClientConnection configures the connection to Kubernetes
	ClientConnection apimachineryconfigv1.ClientConnectionConfiguration `json:"clientConnection"`
	// LeaderElection configures so the component can be deployed in HA mode on k8s
	LeaderElection apiserverconfigv1.LeaderElectionConfiguration `json:"leaderElection"`
	// Server holds configuration settings for the HTTPS server
	Server ServerConfiguration `json:"server"`
}

type ServerConfiguration struct {
	// Default: "0.0.0.0"
	// +optional
	Address string `json:"address"`
	// Default: 10250
	// +optional
	Port uint32 `json:"port"`
	// +optional
	TLSCertFile string `json:"tlsCertFile"`
	// +optional
	TLSPrivateKeyFile string `json:"tlsPrivateKeyFile"`
}
