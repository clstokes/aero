package providers

import (
	"fmt"
	"github.com/clstokes/aero/structs"
)

type Amazon struct {
	Mapping structs.ProviderMapping
}

func (a Amazon) Name() string {
	return "amazon"
}

func (a Amazon) IsCurrentProvider() bool {
	_, err := a.Read(structs.KEY_ADDRESS_PRIVATE)
	if err != nil {
		return false
	}
	return true
}

func (a Amazon) Read(s string) (string, error) {
	metadataItem, exists := a.Mapping.MetadataItems[s]
	if !exists {
		return "", fmt.Errorf("No lookup support for key [%v]", s)
	}

	url := "http://" + a.Mapping.MetadataAddress + metadataItem.Url
	value, err := GetMetadata(url, nil)

	if err != nil {
		return "", err
	}

	if metadataItem.ParseFunc != nil {
		parsedValue, err := metadataItem.ParseFunc(value)
		return parsedValue, err
	}

	return value, nil
}

// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
func InitAmazon() structs.Provider {
	mapping := structs.ProviderMapping{
		MetadataAddress: "169.254.169.254",
		MetadataItems: map[string]structs.MetadataItem{
			structs.KEY_ADDRESS_PRIVATE: structs.MetadataItem{
				Url: "/latest/meta-data/local-ipv4",
			},
			structs.KEY_ADDRESS_PUBLIC: structs.MetadataItem{
				Url: "/latest/meta-data/public-ipv4",
			},
			structs.KEY_INSTANCE_NAME: structs.MetadataItem{
				Url: "/latest/meta-data/instance-id",
			},
			structs.KEY_REGION: structs.MetadataItem{
				Url: "/latest/meta-data/placement/availability-zone",
				ParseFunc: func(v interface{}) (string, error) {
					// value is in form "<region>-<location>-<number><zone-letter>"
					value := v.(string)
					return fmt.Sprintf("%s", value[0:len(value)-1]), nil
				},
			},
			structs.KEY_ZONE: structs.MetadataItem{
				Url: "/latest/meta-data/placement/availability-zone",
			},
		},
	}
	aws := Amazon{Mapping: mapping}
	return aws
}
