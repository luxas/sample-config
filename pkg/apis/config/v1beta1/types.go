package v1beta1

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
	// Default: "0.0.0.0"
	ServerAddress string `json:"serverAddress"`
	// Default: 9090
	// +optional
	HTTPSPort uint32 `json:"httpsPort"`
	// TLSConfig holds settings for the TLS configuration
	TLSConfig TLSConfig `json:"tlsConfig"`
}

type TLSConfig struct {
	// +optional
	TLSCertFile string `json:"tlsCertFile"`
	// +optional
	TLSPrivateKeyFile string `json:"tlsPrivateKeyFile"`
}