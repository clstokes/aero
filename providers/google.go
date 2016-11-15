package providers

import (
	"fmt"
	"github.com/clstokes/aero/structs"
	"strings"
)

type Google struct {
	Mapping structs.ProviderMapping
}

func (g Google) Name() string {
	return structs.NAME_GOOGLE
}

func (g Google) IsCurrentProvider() bool {
	_, err := g.Read(structs.KEY_ADDRESS_PRIVATE)
	if err != nil {
		return false
	}
	return true
}

func (g Google) Read(s string) (string, error) {
	metadataItem, exists := g.Mapping.MetadataItems[s]
	if !exists {
		return "", fmt.Errorf("No lookup support for key [%v]", s)
	}

	url := g.Mapping.MetadataAddress + metadataItem.Url
	value, err := GetMetadata(url, map[string]string{"Metadata-Flavor": "Google"})

	if err != nil {
		return "", err
	}

	if metadataItem.ParseFunc != nil {
		parsedValue, err := metadataItem.ParseFunc(value)
		return parsedValue, err
	}

	return value, nil
}

// https://cloud.google.com/compute/docs/metadata
// Metadata-Flavor: Google
func InitGoogle(defaults structs.ProviderMapping) structs.Provider {
	mapping := structs.ProviderMapping{
		MetadataAddress: "http://169.254.169.254",
		MetadataItems: map[string]structs.MetadataItem{
			structs.KEY_ADDRESS_PRIVATE: structs.MetadataItem{
				Url: "/computeMetadata/v1/instance/network-interfaces/0/ip",
			},
			structs.KEY_ADDRESS_PUBLIC: structs.MetadataItem{
				Url: "/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip",
			},
			structs.KEY_INSTANCE_NAME: structs.MetadataItem{
				Url: "/computeMetadata/v1/instance/hostname",
				ParseFunc: func(v interface{}) (string, error) {
					// value is in form "<name>.c.<project-name>.internal"
					value := v.(string)
					valueSplit := strings.Split(value, ".")

					if len(valueSplit) != 4 {
						return "", fmt.Errorf("Unparseable value [%s]", value)
					}

					return valueSplit[0], nil
				},
			},
			structs.KEY_PROVIDER: structs.MetadataItem{
				// kind of a hack to not introduce another special field/function
				Url: "/computeMetadata/v1/instance/hostname",
				ParseFunc: func(v interface{}) (string, error) {
					return structs.NAME_GOOGLE, nil
				},
			},
			structs.KEY_ZONE: structs.MetadataItem{
				Url: "/computeMetadata/v1/instance/zone",
				ParseFunc: func(v interface{}) (string, error) {
					// value is in form "projects/<project-id>/zones/<zone>"
					value := v.(string)
					valueSplit := strings.Split(value, "/")

					if len(valueSplit) != 4 {
						return "", fmt.Errorf("Unparseable value [%s]", value)
					}

					return valueSplit[3], nil
				},
			},
			structs.KEY_REGION: structs.MetadataItem{
				Url: "/computeMetadata/v1/instance/zone",
				ParseFunc: func(v interface{}) (string, error) {
					// value is in form "projects/<project-id>/zones/<zone>"
					value := v.(string)
					valueSplit := strings.Split(value, "/")

					// value is in form "<region>-<location><number>-<zone-letter>"
					zoneValue := valueSplit[3]
					return fmt.Sprintf("%s", zoneValue[0:len(zoneValue)-2]), nil
				},
			},
		},
	}
	ChangeProviderMappingDefaults(&mapping, &defaults)
	google := Google{Mapping: mapping}
	return google
}
