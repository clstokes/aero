package structs

type Provider interface {
	Name() string
	IsCurrentProvider() bool
	Read(key string) (string, error)
}

type ProviderMapping struct {
	MetadataAddress string
	MetadataItems   map[string]MetadataItem
}

const (
	KEY_ADDRESS_PRIVATE = "address-private"
	KEY_ADDRESS_PUBLIC  = "address-public"
	KEY_INSTANCE_NAME   = "instance-name"
	KEY_REGION          = "region"
	KEY_ZONE            = "zone"
)

var AllKeys = []string{
	KEY_ADDRESS_PRIVATE,
	KEY_ADDRESS_PUBLIC,
	KEY_INSTANCE_NAME,
	KEY_REGION,
	KEY_ZONE,
}
