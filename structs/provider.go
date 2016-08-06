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
	KEY_ZONE            = "zone"
	KEY_ADDRESS_PRIVATE = "address-private"
	KEY_ADDRESS_PUBLIC  = "address-public"
)

var AllKeys = []string{
	KEY_ADDRESS_PRIVATE,
	KEY_ADDRESS_PUBLIC,
	KEY_ZONE,
}
