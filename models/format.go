package models

// Format ...
type IFormat interface {
	GetFilename() string
	GetNbStreams() int
	GetNbPrograms() int
	GetFormatName() string
	GetFormatLongName() string
	GetDuration() string
	GetSize() string
	GetBitRate() string
	GetProbeScore() int
	GetTags() Tags
}

// Format ...
type Format struct {
	Filename       string
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           Tags   `json:"tags"`
}

// GetFilename ...
func (f Format) GetFilename() string {
	return f.Filename
}

// GetNbStreams ...
func (f Format) GetNbStreams() int {
	return f.NbStreams
}

// GetNbPrograms ...
func (f Format) GetNbPrograms() int {
	return f.NbPrograms
}

// GetFormatName ...
func (f Format) GetFormatName() string {
	return f.FormatName
}

// GetFormatLongName ...
func (f Format) GetFormatLongName() string {
	return f.FormatLongName
}

// GetDuration ...
func (f Format) GetDuration() string {
	return f.Duration
}

// GetSize ...
func (f Format) GetSize() string {
	return f.Size
}

// GetBitRate ...
func (f Format) GetBitRate() string {
	return f.BitRate
}

// GetProbeScore ...
func (f Format) GetProbeScore() int {
	return f.ProbeScore
}

// GetTags ...
func (f Format) GetTags() Tags {
	return f.Tags
}
