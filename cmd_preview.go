package main

import (
	"fmt"
	"log"
)

func cmd_Preview(common VideoDefaults) {

	log.Printf("Now running preview")

	outfile := destfile(common.Source, ".preview.avi")
	newsize := newSizeForFile(common.Source)

	/*
		ffmpeg([]string{
			"-i", "\"" + common.Source + "\"",
			"-t", "5:00",
			common.VideoFormat,
			"-s 640x320 -r 25 -b:v " + common.VideoBandwith,
			"", // audio mapiing
			common.AudioFormat,
			"\"" + outfile + "\"",

		})
	*/

	/*
		ffmpeg([]string{
			"-i \"" + common.Source + "\"",
			"-t 0:30",
			"xxxx.avi",
		})
	*/
	shellout_live("ffmpeg", []string{
		"-i", TOKEN_PREFIX + common.Source,
		"-t 0:30",
		common.VideoFormat,
		"-s " + newsize + " -r 25 -b:v " + common.VideoBandwith,
		common.AudioStreams,
		common.AudioFormat,
		TOKEN_PREFIX + outfile,
	}, shelltracker)

}

func shelltracker(stdin bool, to string) {

	if stdin {
		fmt.Printf("I: %s\n", to)
	} else {
		fmt.Printf("E: %s\n", to)
	}

}
