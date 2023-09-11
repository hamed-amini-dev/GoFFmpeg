package models

// Metadata ...
type IMetadata interface {
	GetFormat() Format
	GetStreams() []Streams
}

// Metadata ...
type Metadata struct {
	Format  Format    `json:"format"`
	Streams []Streams `json:"streams"`
}

// GetFormat ...
func (m Metadata) GetFormat() Format {
	return m.Format
}

// GetStreams ...
func (m Metadata) GetStreams() (streams []Streams) {
	for _, element := range m.Streams {
		streams = append(streams, element)
	}
	return streams
}
