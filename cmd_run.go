package main

import (
	"fmt"
	"log"
	"regexp"
	
)

/**

rm -f $LOGFILE
$FFMPEG  -i "$FILEIN" $ONEMINUTE $VIDEO $MAPS $NOAUDIO -pass 1 -passlogfile $LOGFILE -y /dev/null
$FFMPEG  -i "$FILEIN" $ONEMINUTE $VIDEO $MAPS $AUDIO   -pass 2 -passlogfile $LOGFILE "$FILEOUT"

*/

func cmd_Run(common VideoDefaults) {
	log.Printf("Now running full conversion")

	logfile := "multipasslog_" + random_string()
	outfile := destfile(common.Source, ".avi")
	newsize := newSizeForFile(common.Source)

	shellout_live("ffmpeg", []string{
		"-progress /dev/stdout",
		"-i", TOKEN_PREFIX + common.Source,
		//"-t 0:30",
		common.VideoFormat,
		"-s " + newsize + " -r 25 -b:v " + common.VideoBandwith,
		common.AudioStreams,
		"-an",
		"-pass 1",
		"-passlogfile " + logfile,
		"-y /dev/null",
	}, mkTrackerFn("P1"))

	shellout_live("ffmpeg", []string{
		"-progress /dev/stdout",
		"-i", TOKEN_PREFIX + common.Source,
		//"-t 0:30",
		common.VideoFormat,
		"-s " + newsize + " -r 25 -b:v " + common.VideoBandwith,
		common.AudioStreams,
		common.AudioFormat,
		"-pass 2",
		"-passlogfile " + logfile,
		TOKEN_PREFIX + outfile,
	}, mkTrackerFn("P2"))

}

/**
Builds a tracking function. It prints the current time on every minute that changes.

frame=41266
I: fps=1070.8
I: stream_0_0_q=2.9
I: total_size=268728098
I: out_time_ms=1650680000
I: out_time=00:27:30.680000
I: dup_frames=0
I: drop_frames=0
I: progress=continue
*/

func mkTrackerFn(name string) func(bool, string) {

	r := regexp.MustCompile("^out_time=(..:..)")
	last_point := "";

	return func(stdin bool, line string) {
		result_slice := r.FindAllStringSubmatch(line, -1)
		if len(result_slice) == 1 {
			now := result_slice[0][1]
			if (now != last_point ) {
				fmt.Printf("%s %s", name, now)		
				last_point = now		
			} else {
				fmt.Printf(".")
			}
		}
	}
}
