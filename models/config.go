package models

// Config ...
type IConfig interface{}

// Config ...
type Config struct {
	FFmpegBinPath  string
	FFprobeBinPath string
	ShowProgress   bool
	Verbose        bool
}
