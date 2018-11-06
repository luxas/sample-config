package main

import (
	"io/ioutil"
	"os"
	"fmt"

	"github.com/luxas/sample-config/pkg/apis/config/scheme"
	"github.com/luxas/sample-config/pkg/apis/config"
	"github.com/luxas/sample-config/pkg/apis/config/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func main() {
	if err := readConfigOrWriteDefault(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func readConfigOrWriteDefault() error {
	// The internal config object to be populated and written to STDOUT
	cfg := &config.MyAppConfiguration{}
	// If an argument was specified, try to deserialize the file
	var err error
	if len(os.Args) > 1 && len(os.Args[1]) != 0 {
		err = decodeFileInto(os.Args[1], cfg)
	} else {
		err = populateV1Defaults(cfg)
	}
	if err != nil {
		return err
	}

	// Finally, always print the resulting config in the v1 version
	cfgbytes, err := marshalYAML(cfg, v1.SchemeGroupVersion)
	if err != nil {
		return err
	}
	fmt.Printf("%s", cfgbytes)
	return nil
}

// decodeFileInto reads a file and decodes the it into an internal type
func decodeFileInto(filePath string, obj runtime.Object) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	// Regardless of if the bytes are of the v1 or v1beta1 version,
	// it will be read successfully and converted into the internal version
	return runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), content, obj)
}
// populateV1Defaults populates cfg based on v1 defaults
func populateV1Defaults(cfg *config.MyAppConfiguration) error {
	// Create a new config of some external version,
	// default it, convert it into the internal version
	v1cfg := &v1.MyAppConfiguration{}
	scheme.Scheme.Default(v1cfg)
	return scheme.Scheme.Convert(v1cfg, cfg, nil)
}
// marshalYAML marshals any ComponentConfig object registered in the scheme for the specific version
func marshalYAML(obj runtime.Object, groupVersion schema.GroupVersion) ([]byte, error) {
	// yamlEncoder is a generic-purpose encoder to YAML for this scheme
	yamlEncoder := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	// versionSpecificEncoder writes out YAML bytes for exactly this v1beta1 version
	versionSpecificEncoder := scheme.Codecs.EncoderForVersion(yamlEncoder, groupVersion)
	// Encode the object to YAML for the given version
	return runtime.Encode(versionSpecificEncoder, obj)
}
