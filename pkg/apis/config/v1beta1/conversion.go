package v1beta1

import (
	"github.com/luxas/sample-config/pkg/apis/config"

	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1beta1_MyAppConfiguration_To_config_MyAppConfiguration(in *MyAppConfiguration, out *config.MyAppConfiguration, s conversion.Scope) error {
	if err := autoConvert_v1beta1_MyAppConfiguration_To_config_MyAppConfiguration(in, out, s); err != nil {
		return err
	}
	out.Server.Address = in.ServerAddress
	out.Server.Port = in.HTTPSPort
	out.Server.TLSCertFile = in.TLSConfig.TLSCertFile
	out.Server.TLSPrivateKeyFile = in.TLSConfig.TLSPrivateKeyFile
	return nil
}

func Convert_config_MyAppConfiguration_To_v1beta1_MyAppConfiguration(in *config.MyAppConfiguration, out *MyAppConfiguration, s conversion.Scope) error {
	if err := autoConvert_config_MyAppConfiguration_To_v1beta1_MyAppConfiguration(in, out, s); err != nil {
		return err
	}
	out.ServerAddress = in.Server.Address
	out.HTTPSPort = in.Server.Port
	out.TLSConfig.TLSCertFile = in.Server.TLSCertFile
	out.TLSConfig.TLSPrivateKeyFile = in.Server.TLSPrivateKeyFile
	return nil
}
