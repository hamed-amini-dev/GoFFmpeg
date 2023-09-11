package models

// Disposition ...
type IDisposition interface {
	GetDefault() int
	GetDub() int
	GetOriginal() int
	GetComment() int
	GetLyrics() int
	GetKaraoke() int
	GetForced() int
	GetHearingImpaired() int
	GetVisualImpaired() int
	GetCleanEffects() int
}

// Disposition ...
type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
}

//GetDefault ...
func (d Disposition) GetDefault() int {
	return d.Default
}

//GetDub ...
func (d Disposition) GetDub() int {
	return d.Dub
}

//GetOriginal ...
func (d Disposition) GetOriginal() int {
	return d.Original
}

//GetComment ...
func (d Disposition) GetComment() int {
	return d.Comment
}

//GetLyrics ...
func (d Disposition) GetLyrics() int {
	return d.Lyrics
}

//GetKaraoke ...
func (d Disposition) GetKaraoke() int {
	return d.Karaoke
}

//GetForced ...
func (d Disposition) GetForced() int {
	return d.Forced
}

//GetHearingImpaired ...
func (d Disposition) GetHearingImpaired() int {
	return d.HearingImpaired
}

//GetVisualImpaired ...
func (d Disposition) GetVisualImpaired() int {
	return d.VisualImpaired
}

//GetCleanEffects ...
func (d Disposition) GetCleanEffects() int {
	return d.CleanEffects
}
