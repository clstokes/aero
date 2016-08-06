package structs

type MetadataItem struct {
	Url       string
	ParseFunc MetadataItemParseFunc
}

type MetadataItemParseFunc func(interface{}) (string, error)
