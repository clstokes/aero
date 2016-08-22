package structs

const (
	NAME_AMAZON = "amazon"
	NAME_GOOGLE = "google"

	KEY_ADDRESS_PRIVATE = "address-private"
	KEY_ADDRESS_PUBLIC  = "address-public"
	KEY_INSTANCE_NAME   = "instance"
	KEY_PROVIDER        = "provider"
	KEY_REGION          = "region"
	KEY_ZONE            = "zone"
)

var AllKeys = []string{
	KEY_ADDRESS_PRIVATE,
	KEY_ADDRESS_PUBLIC,
	KEY_INSTANCE_NAME,
	KEY_PROVIDER,
	KEY_REGION,
	KEY_ZONE,
}

type Provider interface {
	Name() string
	IsCurrentProvider() bool
	Read(key string) (string, error)
}

type ProviderMapping struct {
	MetadataAddress string
	MetadataItems   map[string]MetadataItem
}
