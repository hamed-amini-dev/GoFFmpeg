package main

import (
	"converter/ffmpeg"
	"converter/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var programName = path.Base(os.Args[0])

var rootCmd = &cobra.Command{
	Use: programName,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		from_flag, _ := cmd.Flags().GetString("from")
		to_flag, _ := cmd.Flags().GetString("to")
		fTimeDur, _ := cmd.Flags().GetString("t")
		thumbnailTimes = strings.Split(fTimeDur, ",")
		if from_flag == "mkv" || from_flag == "3gp" || from_flag == "rm" {
			if to_flag == "mp4" {
				FromFormat = from_flag
				ToFormat = to_flag
				execKnownFormatConvert = true
				execUnKnownFormatConvert = false
			}
		} else {
			execUnKnownFormatConvert = true
			execKnownFormatConvert = false

		}

		if from_flag == "*" {
			ToFormat = to_flag
			execUnKnownFormatConvert = true
			execKnownFormatConvert = false
		}

	},
}

func CalculateTimeDurations(durations string) ([]string, error) {
	var result []string
	for _, v := range thumbnailTimes {

		if len(v) != 0 {
			v1, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			v2, err := strconv.ParseFloat(durations, 64)
			if err != nil {
				return nil, err
			}
			a := int64(v1 * v2)
			//
			hours := a / 3600
			remainder := a - hours*3600
			mins := remainder / 60
			remainder = remainder - mins*60
			secs := remainder
			//
			result = append(result, fmt.Sprintf("%d", hours)+":"+fmt.Sprintf("%d", mins)+":"+fmt.Sprintf("%d", secs))

		}
	}
	return result, nil
}

/* var OutPutformatCmd = &cobra.Command{
	Use:   "format",
	Short: "Set OutPut Parameter",
	Long:  `This command set output convert file`,
	Run: func(cmd *cobra.Command, args []string) {
		from_flag, _ := cmd.Flags().GetString("from")
		to_flag, _ := cmd.Flags().GetString("to")

		log.Println(from_flag)
		log.Println(to_flag)
		ToFormat = args[0]
		execConvert = true
	},
}
*/
var ToFormat = "mp4"
var FromFormat = "3gp"
var thumbnailTimes []string
var overwrite = true
var execKnownFormatConvert = false
var execUnKnownFormatConvert = false

func init() {
	rootCmd.PersistentFlags().String("from", "", "input from format for convert.")
	rootCmd.PersistentFlags().String("to", "", "input to format for convert.")
	rootCmd.PersistentFlags().String("t", "", "input to format for convert.")
	//rootCmd.AddCommand(OutPutformatCmd)
	//rootCmd.Execute()
	rootCmd.Execute()
}

func main() {
	ExtraArg := make(map[string]interface{})
	ExtraArg["-frames:v 1"] = ""
	//ExtraArg["force_original_aspect_ratio"] = "decrease"
	inputPath := "./EncodeFile/input/"
	outputPath := "./EncodeFile/output/"

	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	opts := models.Options{
		//OutputFormat: &ToFormat,
		//Overwrite: &overwrite,
		//SeekTime:  &seekTime,
	}

	ffmpegConf := &models.Config{
		FFmpegBinPath:  "ffmpeg",
		FFprobeBinPath: "ffprobe",
		ShowProgress:   true,
	}

	// input folder use here
	// output folder use here
	if execKnownFormatConvert {
		// Get Thumbnail From Movie
		for _, file := range files {
			filename := strings.Split(file.Name(), ".")
			if filename[1] == FromFormat {
				progress := ffmpeg.New(ffmpegConf)
				progress.Input(inputPath + file.Name())
				mdata, err := progress.GetMetadata()
				if err != nil {
					log.Fatal(err)
					return
				}

				seeks, err := CalculateTimeDurations(mdata.GetFormat().GetDuration())
				if err != nil {
					log.Fatal(err)
					return
				}
				overwrite = true
				opts.Overwrite = &overwrite
				ffmpegConf.ShowProgress = true
				for i, v := range seeks {
					log.Printf("Start Take Thumbnail File %s Please Wait!", file.Name())
					opts.SeekTime = &v
					progress := ffmpeg.New(ffmpegConf)
					progress.Input(inputPath + file.Name())
					progress.Output(outputPath + filename[0] + strconv.FormatInt(int64(i), 10) + ".jpg")
					progress.WithOptions(opts)
					thumnailProgress, err := progress.Start(opts)
					if err != nil {
						log.Fatal(err)
						return
					}

					for msg := range thumnailProgress {
						log.Printf("Take a Thumbnail file %s %+v", file.Name(), msg)
					}

					log.Printf("Finish Take Thumbnail File %s ", file.Name())

				}

			}
		}
		// Convert Movie
		for _, file := range files {
			filename := strings.Split(file.Name(), ".")
			if filename[1] == FromFormat {
				log.Printf("Start Convert File %s Please Wait!", file.Name())
				overwrite = true
				opts.Overwrite = &overwrite
				opts.SeekTime = nil
				ffmpegConf.ShowProgress = true
				progress, err := ffmpeg.
					New(ffmpegConf).
					Input(inputPath + file.Name()).
					Output(outputPath + filename[0] + "." + ToFormat).
					WithOptions(opts).
					Start(opts)
				if err != nil {
					log.Fatal(err)
					return
				}

				for msg := range progress {
					log.Printf("Convert Video file %s %+v", file.Name(), msg)
				}

				log.Printf("Finish Convert File %s ", file.Name())

			}

		}
	} else if execUnKnownFormatConvert {
		// Get Thumbnail From Movie
		for _, file := range files {
			filename := strings.Split(file.Name(), ".")
			progress := ffmpeg.New(ffmpegConf)
			progress.Input(inputPath + file.Name())
			mdata, err := progress.GetMetadata()
			if err != nil {
				log.Fatal(err)
				return
			}

			seeks, err := CalculateTimeDurations(mdata.GetFormat().GetDuration())
			if err != nil {
				log.Fatal(err)
				return
			}
			overwrite = true
			opts.Overwrite = &overwrite
			ffmpegConf.ShowProgress = true
			for i, v := range seeks {
				log.Printf("Start Take Thumbnail File %s Please Wait!", file.Name())
				opts.SeekTime = &v
				progress := ffmpeg.New(ffmpegConf)
				progress.Input(inputPath + file.Name())
				progress.Output(outputPath + filename[0] + strconv.FormatInt(int64(i), 10) + ".jpg")
				progress.WithOptions(opts)
				thumnailProgress, err := progress.Start(opts)
				if err != nil {
					log.Fatal(err)
					return
				}

				for msg := range thumnailProgress {
					log.Printf("Take a Thumbnail file %s %+v", file.Name(), msg)
				}

				log.Printf("Finish Take Thumbnail File %s ", file.Name())

			}

		}
		// Convert Movie
		for _, file := range files {
			log.Printf("Start Convert File %s Please Wait!", file.Name())
			filename := strings.Split(file.Name(), ".")
			overwrite = true
			opts.Overwrite = &overwrite
			opts.SeekTime = nil
			ffmpegConf.ShowProgress = true
			progress, err := ffmpeg.
				New(ffmpegConf).
				Input(inputPath + file.Name()).
				Output(outputPath + filename[0] + "." + ToFormat).
				WithOptions(opts).
				Start(opts)

			if err != nil {
				log.Fatal(err)
			}

			for msg := range progress {
				log.Printf("Convert Video file %s %+v", file.Name(), msg)
			}

			log.Printf("Finish Convert File %s", file.Name())

		}
	} else {
		fmt.Println("Please input argument for example converter --from 3gp --to mp4 --t 0.1,0.8,0.9")
		fmt.Println("list of format from == 3gp and rm")
		fmt.Println("list of format to == mp4 and mp3")
		fmt.Println("for thumbnails use --t and insert multiple number between 0 to 1 , if you want one thumbnail use this sample -t 0.6")
	}

}
