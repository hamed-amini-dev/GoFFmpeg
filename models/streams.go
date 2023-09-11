package models

// Streams ...
type IStreams interface {
	GetIndex() int
	GetID() string
	GetCodecName() string
	GetCodecLongName() string
	GetProfile() string
	GetCodecType() string
	GetCodecTimeBase() string
	GetCodecTagString() string
	GetCodecTag() string
	GetWidth() int
	GetHeight() int
	GetCodedWidth() int
	GetCodedHeight() int
	GetHasBFrames() int
	GetSampleAspectRatio() string
	GetDisplayAspectRatio() string
	GetPixFmt() string
	GetLevel() int
	GetChromaLocation() string
	GetRefs() int
	GetQuarterSample() string
	GetDivxPacked() string
	GetRFrameRrate() string
	GetAvgFrameRate() string
	GetTimeBase() string
	GetDurationTs() int
	GetDuration() string
	GetDisposition() Disposition
	GetBitRate() string
}

// Streams ...
type Streams struct {
	Index              int
	ID                 string      `json:"id"`
	CodecName          string      `json:"codec_name"`
	CodecLongName      string      `json:"codec_long_name"`
	Profile            string      `json:"profile"`
	CodecType          string      `json:"codec_type"`
	CodecTimeBase      string      `json:"codec_time_base"`
	CodecTagString     string      `json:"codec_tag_string"`
	CodecTag           string      `json:"codec_tag"`
	Width              int         `json:"width"`
	Height             int         `json:"height"`
	CodedWidth         int         `json:"coded_width"`
	CodedHeight        int         `json:"coded_height"`
	HasBFrames         int         `json:"has_b_frames"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio"`
	DisplayAspectRatio string      `json:"display_aspect_ratio"`
	PixFmt             string      `json:"pix_fmt"`
	Level              int         `json:"level"`
	ChromaLocation     string      `json:"chroma_location"`
	Refs               int         `json:"refs"`
	QuarterSample      string      `json:"quarter_sample"`
	DivxPacked         string      `json:"divx_packed"`
	RFrameRrate        string      `json:"r_frame_rate"`
	AvgFrameRate       string      `json:"avg_frame_rate"`
	TimeBase           string      `json:"time_base"`
	DurationTs         int         `json:"duration_ts"`
	Duration           string      `json:"duration"`
	Disposition        Disposition `json:"disposition"`
	BitRate            string      `json:"bit_rate"`
}

//GetIndex ...
func (s Streams) GetIndex() int {
	return s.Index
}

//GetID ...
func (s Streams) GetID() string {
	return s.ID
}

//GetCodecName ...
func (s Streams) GetCodecName() string {
	return s.CodecName
}

//GetCodecLongName ...
func (s Streams) GetCodecLongName() string {
	return s.CodecLongName
}

//GetProfile ...
func (s Streams) GetProfile() string {
	return s.Profile
}

//GetCodecType ...
func (s Streams) GetCodecType() string {
	return s.CodecType
}

//GetCodecTimeBase ...
func (s Streams) GetCodecTimeBase() string {
	return s.CodecTimeBase
}

//GetCodecTagString ...
func (s Streams) GetCodecTagString() string {
	return s.CodecTagString
}

//GetCodecTag ...
func (s Streams) GetCodecTag() string {
	return s.CodecTag
}

//GetWidth ...
func (s Streams) GetWidth() int {
	return s.Width
}

//GetHeight ...
func (s Streams) GetHeight() int {
	return s.Height
}

//GetCodedWidth ...
func (s Streams) GetCodedWidth() int {
	return s.CodedWidth
}

//GetCodedHeight ...
func (s Streams) GetCodedHeight() int {
	return s.CodedHeight
}

//GetHasBFrames ...
func (s Streams) GetHasBFrames() int {
	return s.HasBFrames
}

//GetSampleAspectRatio ...
func (s Streams) GetSampleAspectRatio() string {
	return s.SampleAspectRatio
}

//GetDisplayAspectRatio ...
func (s Streams) GetDisplayAspectRatio() string {
	return s.DisplayAspectRatio
}

//GetPixFmt ...
func (s Streams) GetPixFmt() string {
	return s.PixFmt
}

//GetLevel ...
func (s Streams) GetLevel() int {
	return s.Level
}

//GetChromaLocation ...
func (s Streams) GetChromaLocation() string {
	return s.ChromaLocation
}

//GetRefs ...
func (s Streams) GetRefs() int {
	return s.Refs
}

//GetQuarterSample ...
func (s Streams) GetQuarterSample() string {
	return s.QuarterSample
}

//GetDivxPacked ...
func (s Streams) GetDivxPacked() string {
	return s.DivxPacked
}

//GetRFrameRrate ...
func (s Streams) GetRFrameRrate() string {
	return s.RFrameRrate
}

//GetAvgFrameRate ...
func (s Streams) GetAvgFrameRate() string {
	return s.AvgFrameRate
}

//GetTimeBase ...
func (s Streams) GetTimeBase() string {
	return s.TimeBase
}

//GetDurationTs ...
func (s Streams) GetDurationTs() int {
	return s.DurationTs
}

//GetDuration ...
func (s Streams) GetDuration() string {
	return s.Duration
}

//GetDisposition ...
func (s Streams) GetDisposition() Disposition {
	return s.Disposition
}

//GetBitRate ...
func (s Streams) GetBitRate() string {
	return s.BitRate
}
