package ffmpeg

import (
	"bufio"
	"bytes"
	"context"
	"converter/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Encoder ...
type IEncoder interface {
	Start(opts models.Options) (<-chan models.Progress, error)
	Input(i string) *Encoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) *Encoder
	Output(o string) *Encoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) *Encoder
	WithOptions(opts models.Options) *Encoder
	WithAdditionalOptions(opts models.Options) *Encoder
	WithContext(ctx *context.Context) *Encoder
	GetMetadata() (*models.Metadata, error)
}

type Encoder struct {
	config           *models.Config
	input            string
	output           []string
	options          [][]string
	metadata         models.Metadata
	inputPipeReader  *io.ReadCloser
	outputPipeReader *io.ReadCloser
	inputPipeWriter  *io.WriteCloser
	outputPipeWriter *io.WriteCloser
	commandContext   *context.Context
}

// New ...
func New(cfg *models.Config) IEncoder {
	return &Encoder{config: cfg}
}

// Start ...
func (t *Encoder) Start(opts models.Options) (<-chan models.Progress, error) {

	var stderrIn io.ReadCloser

	out := make(chan models.Progress)

	defer t.closePipes()

	// Validates config
	if err := t.validate(); err != nil {
		log.Println("1")
		return nil, err
	}

	// Get file metadata
	_, err := t.GetMetadata()
	if err != nil {
		log.Println("2")

		return nil, err
	}

	// Append input file and standard options
	args := append([]string{"-i", t.input}, opts.GetStrArguments()...)
	outputLength := len(t.output)
	optionsLength := len(t.options)

	if outputLength == 1 && optionsLength == 0 {
		// Just append the 1 output file we've got
		args = append(args, t.output[0])
	} else {
		for index, out := range t.output {
			// Get executable flags
			// If we are at the last output file but still have several options, append them all at once
			if index == outputLength-1 && outputLength < optionsLength {
				for i := index; i < len(t.options); i++ {
					args = append(args, t.options[i]...)
				}
				// Otherwise just append the current options
			} else {
				args = append(args, t.options[index]...)
			}

			// Append output flag
			args = append(args, out)
		}
	}

	//log.Println("args", args)

	// Initialize command
	// If a context object was supplied to this Encoder before
	// starting, use this context when creating the command to allow
	// the command to be killed when the context expires
	var cmd *exec.Cmd
	if t.commandContext == nil {
		cmd = exec.Command(t.config.FFmpegBinPath, args...)
	} else {
		cmd = exec.CommandContext(*t.commandContext, t.config.FFprobeBinPath, args...)
	}

	// If progresss enabled, get stderr pipe and start progress process
	if t.config.ShowProgress && !t.config.Verbose {
		stderrIn, err = cmd.StderrPipe()
		if err != nil {
			log.Println("3")

			return nil, fmt.Errorf("Failed getting transcoding progress (%s) with args (%s) with error %s", t.config.FFmpegBinPath, args, err)
		}
	}

	if t.config.Verbose {
		cmd.Stderr = os.Stdout
	}

	// Start process
	err = cmd.Start()
	if err != nil {
		log.Println("4")

		return nil, fmt.Errorf("Failed starting transcoding (%s) with args (%s) with error %s", t.config.FFmpegBinPath, args, err)
	}

	if t.config.ShowProgress && !t.config.Verbose {
		go func() {
			t.progress(stderrIn, out)
		}()

		go func() {
			defer close(out)
			err = cmd.Wait()
		}()
	} else {
		err = cmd.Wait()
	}

	return out, nil
}

// Input ...
func (t *Encoder) Input(arg string) *Encoder {
	t.input = arg
	return t
}

// Output ...
func (t *Encoder) Output(arg string) *Encoder {
	t.output = append(t.output, arg)
	return t
}

// InputPipe ...
func (t *Encoder) InputPipe(w *io.WriteCloser, r *io.ReadCloser) *Encoder {
	if &t.input == nil {
		t.inputPipeWriter = w
		t.inputPipeReader = r
	}
	return t
}

// OutputPipe ...
func (t *Encoder) OutputPipe(w *io.WriteCloser, r *io.ReadCloser) *Encoder {
	if &t.output == nil {
		t.outputPipeWriter = w
		t.outputPipeReader = r
	}
	return t
}

// WithOptions Sets the options object
func (t *Encoder) WithOptions(opts models.Options) *Encoder {
	t.options = [][]string{opts.GetStrArguments()}
	return t
}

// WithAdditionalOptions Appends an additional options object
func (t *Encoder) WithAdditionalOptions(opts models.Options) *Encoder {
	t.options = append(t.options, opts.GetStrArguments())
	return t
}

// WithContext is to be used on a Encoder *before Starting* to
// pass in a context.Context object that can be used to kill
// a running models process. Usage of this method is optional
func (t *Encoder) WithContext(ctx *context.Context) *Encoder {
	t.commandContext = ctx
	return t
}

// validate ...
func (t *Encoder) validate() error {
	if t.config.FFmpegBinPath == "" {
		return errors.New("ffmpeg binary path not found")
	}

	if t.input == "" {
		return errors.New("missing input option")
	}

	outputLength := len(t.output)

	if outputLength == 0 {
		return errors.New("missing output option")
	}

	// length of output files being greater than length of options would produce an invalid ffmpeg command
	// unless there is only 1 output file, which obviously wouldn't be a problem
	if outputLength > len(t.options) && outputLength != 1 {
		return errors.New("number of options and output files does not match")
	}

	for index, output := range t.output {
		if output == "" {
			return fmt.Errorf("output at index %d is an empty string", index)
		}
	}

	return nil
}

// GetMetadata Returns metadata for the specified input file
func (t *Encoder) GetMetadata() (*models.Metadata, error) {

	if t.config.FFprobeBinPath != "" {
		var outb, errb bytes.Buffer

		input := t.input

		if t.inputPipeReader != nil {
			input = "pipe:"
		}

		args := []string{"-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}

		cmd := exec.Command(t.config.FFprobeBinPath, args...)
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FFprobeBinPath, args, err, outb.String(), errb.String())
		}

		var metadata models.Metadata

		if err = json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
			return nil, err
		}

		t.metadata = metadata

		return &metadata, nil
	}

	return nil, errors.New("ffprobe binary not found")
}

// progress sends through given channel the transcoding status
func (t *Encoder) progress(stream io.ReadCloser, out chan models.Progress) {

	defer stream.Close()

	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(stream)
	scanner.Split(split)

	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		Progress := new(models.Progress)
		line := scanner.Text()

		if strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
			var re = regexp.MustCompile(`=\s+`)
			st := re.ReplaceAllString(line, `=`)

			f := strings.Fields(st)

			var framesProcessed string
			var currentTime string
			var currentBitrate string
			var currentSpeed string

			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")

				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]

					if fieldname == "frame" {
						framesProcessed = fieldvalue
					}

					if fieldname == "time" {
						currentTime = fieldvalue
					}

					if fieldname == "bitrate" {
						currentBitrate = fieldvalue
					}
					if fieldname == "speed" {
						currentSpeed = fieldvalue
					}
				}
			}

			timesec := t.DurToSec(currentTime)
			dursec, _ := strconv.ParseFloat(t.metadata.GetFormat().GetDuration(), 64)

			progress := (timesec * 100) / dursec
			Progress.Progress = progress

			Progress.CurrentBitrate = currentBitrate
			Progress.FramesProcessed = framesProcessed
			Progress.CurrentTime = currentTime
			Progress.Speed = currentSpeed

			out <- *Progress
		}
	}
}

// closePipes Closes pipes if opened
func (t *Encoder) closePipes() {
	if t.inputPipeReader != nil {
		ipr := *t.inputPipeReader
		ipr.Close()
	}

	if t.outputPipeWriter != nil {
		opr := *t.outputPipeWriter
		opr.Close()
	}
}

func (t *Encoder) DurToSec(dur string) (sec float64) {
	durAry := strings.Split(dur, ":")
	var secs float64
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.ParseFloat(durAry[0], 64)
	secs = hr * (60 * 60)
	min, _ := strconv.ParseFloat(durAry[1], 64)
	secs += min * (60)
	second, _ := strconv.ParseFloat(durAry[2], 64)
	secs += second
	return secs
}
