package main

import (
	//"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

// Generic constants
const VERSION = "0.0.0"
const MAXWIDTH = 640

type VideoDefaults struct {
	Source        string
	VideoBandwith string
	MaxWidth      int
	VideoFormat   string
	AudioFormat   string
	AudioStreams  string
}

func main() {

	app := cli.NewApp()
	app.Name = "toonverter"
	app.Usage = "Converts some video to DIVX"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "source, s",
			Value: "",
			Usage: "The source file to be processed",
		},
		cli.StringFlag{
			Name:  "bw-video",
			Value: "1300k",
			Usage: "The video bandwidth",
		},
		cli.IntFlag{
			Name:  "max-width",
			Value: 640,
			Usage: "The maximum video width",
		},
		cli.StringFlag{
			Name:  "audio-stream",
			Value: "",
			Usage: "The audio stream to include",
		},
		cli.StringFlag{
			Name:  "video-format",
			Value: "-vtag DIVX -c:v mpeg4 -f avi",
			Usage: "",
		},
		cli.StringFlag{
			Name:  "audio-format",
			Value: "-acodec libmp3lame -ac 2 -vol 256 -b:a 128k",
			Usage: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Converts some video as high-quality file",

			Action: func(cliCtxt *cli.Context) {
				common := loadCommon(cliCtxt)
				cmd_Run(common)
			},
		},

		{
			Name:    "preview",
			Aliases: []string{"p"},
			Usage:   "Converts 5 minutes of video for preview",

			Action: func(cliCtxt *cli.Context) {
				common := loadCommon(cliCtxt)
				cmd_Preview(common)
			},
		},

		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Display file information",
			Action: func(cliCtxt *cli.Context) {
				common := loadCommon(cliCtxt)
				cmd_Info(common)
			},
		},
	}

	log.Printf("Starting")
	app.Run(os.Args)

	//execCmd("ffmpeg", "-?")
}

func loadCommon(cliCtxt *cli.Context) VideoDefaults {
	def := VideoDefaults{
		Source:        cliCtxt.Parent().String("source"),
		VideoBandwith: cliCtxt.Parent().String("bw-video"),
		MaxWidth:      cliCtxt.Parent().Int("max-width"),
		VideoFormat:   cliCtxt.Parent().String("video-format"),
		AudioFormat:   cliCtxt.Parent().String("audio-format"),
		AudioStreams:  "",
	}

	log.Printf("Common: %#v", def)

	return def
}
