package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	apimachineryconfigv1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
	apiserverconfigv1 "k8s.io/apiserver/pkg/apis/config/v1alpha1"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_MyAppConfiguration defaults the object
func SetDefaults_MyAppConfiguration(obj *MyAppConfiguration) {
	if len(obj.ServerAddress) == 0 {
		obj.ServerAddress = "0.0.0.0"
	}
	if obj.HTTPSPort == 0 {
		obj.HTTPSPort = 9090
	}

	apimachineryconfigv1.RecommendedDefaultClientConnectionConfiguration(&obj.ClientConnection)
	apiserverconfigv1.RecommendedDefaultLeaderElectionConfiguration(&obj.LeaderElection)
}
