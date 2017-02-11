package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	//"syscall"
)

// Generic constants
const VERSION = "0.0.0"
const MAXWIDTH = 640

func main() {

	app := cli.NewApp()
	app.Name = "toonverter"
	app.Usage = "Converts some video to DIVX"
	app.Version = VERSION

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Convers some video as high-quality file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "source, s",
					Value: "",
					Usage: "The source file to be processed",
				},
			},

			Action: func(cliCtxt *cli.Context) {
				cmd_Run()
			},
		},
	}

	log.Printf("Starting")
	app.Run(os.Args)

	execCmd("ffmpeg", "-?")
}

func execCmd(executable string, parms string) {
	// see http://www.darrencoxall.com/golang/executing-commands-in-go/

	cmd := exec.Command(executable, parms)
	stdout, stderr := cmd.CombinedOutput()
	log.Printf("o: %s\n", stdout)
	log.Printf("e: %s\n", stderr)

	/**
		var waitStatus syscall.WaitStatus
		if err := cmd.Run(); err != nil {
			printError(err)
			// Did the command fail because of an unsuccessful exit code
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus = exitError.Sys().(syscall.WaitStatus)
				printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			}
		} else {
			// Command was successful
			waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	**/
}

func printError(err error) {
	if err != nil {
		log.Printf("==> Error: %s\n", err.Error())
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		log.Printf("==> Output: %s\n", string(outs))
	}
}

func resizer(result string) string {
	w, h := findCurrentStreamFrameSize(result)
	w1, h1 := normalizeStreamSize(w, h, MAXWIDTH, 8)
	return fmt.Sprintf("%dx%d", w1, h1)
}

/**
 * Get the size of the current stream.
 * The framse size is the fist item that looks like "90x50".
 */

func findCurrentStreamFrameSize(text string) (int, int) {

	r := regexp.MustCompile(" (\\d\\d+)x(\\d\\d+)")
	result_slice := r.FindAllStringSubmatch(text, -1)
	if len(result_slice) == 1 {
		sw := result_slice[0][1]
		sh := result_slice[0][2]
		iw, _ := strconv.Atoi(sw)
		ih, _ := strconv.Atoi(sh)
		return iw, ih
	} else {
		log.Printf("Not found!")
		return 0, 0
	}
}

/**
 * Make sure length aligned on block boundary.
 * Chooses whether to elarge ir shrink based on size.
 */

func adjSize(size int, block int) int {
	remainder := size % block
	if remainder != 0 {
		if remainder < (block / 2) {
			size = size - remainder
		} else {
			size = size + (block - remainder)
		}
	}
	return size
}

/**
 * Normalizes a stream size, keeping ratio into consideration.
 */

func normalizeStreamSize(width int, height int, maxWidth int, blksize int) (int, int) {

	var ratio float64 = float64(height) / float64(width)
	//log.Printf( "W %d H %d Ratio: %f", width, height, ratio)

	realwidth := adjSize(width, blksize)

	if realwidth > maxWidth {
		realwidth = maxWidth
	}

	realheight := adjSize(int(float64(realwidth)*ratio), blksize)

	return realwidth, realheight
}
