package config

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "config.luxaslabs.com"

var (
	SchemeBuilder = runtime.NewSchemeBuilder(
		addKnownTypes,
	)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme = localSchemeBuilder.AddToScheme
	// SchemeGroupVersion is'the group & version for this scheme
	SchemeGroupVersion = schema.GroupVersion{
		Group: GroupName,
		Version: runtime.APIVersionInternal,
	}
)

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &MyAppConfiguration{})
	return nil
}
