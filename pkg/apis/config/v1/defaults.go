package v1

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
	if len(obj.Server.Address) == 0 {
		obj.Server.Address = "0.0.0.0"
	}
	if obj.Server.Port == 0 {
		obj.Server.Port = 9090
	}

	apimachineryconfigv1.RecommendedDefaultClientConnectionConfiguration(&obj.ClientConnection)
	apiserverconfigv1.RecommendedDefaultLeaderElectionConfiguration(&obj.LeaderElection)
}
