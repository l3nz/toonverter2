package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/**
Given a file name, returns the new size to use.
*/

func newSizeForFile(file string) string {
	_, _, stderr := shellout("ffmpeg", []string{
		"-i", TOKEN_PREFIX + file,
	})
	return resizer(stderr)
}

/**
Given the results (stderr) of ffmpeg, find the
frame size and resize it.
*/

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

	realwidth := adjSize(width, blksize)
	if realwidth > maxWidth {
		realwidth = maxWidth
	}

	realheight := adjSize(int(float64(realwidth)*ratio), blksize)
	return realwidth, realheight
}

/**
Creates a new destination file.

It tries removing the old extension, unless it is the same
as the current one, in which case it is added.
*/

func destfile(inputfile string, suffix string) string {

	if strings.HasSuffix(inputfile, suffix) {
		return inputfile + suffix
	} else {
		if strings.HasSuffix(inputfile, ".mkv") {
			inputfile = strings.TrimSuffix(inputfile, ".mkv")
		} else if strings.HasSuffix(inputfile, ".mp4") {
			inputfile = strings.TrimSuffix(inputfile, ".mp4")
		} else if strings.HasSuffix(inputfile, ".mpeg") {
			inputfile = strings.TrimSuffix(inputfile, ".mpeg")
		}
	}

	return inputfile + suffix
}

/**
Assert that a file exists (or does not exist).
Panics on violations.
*/

func assertFile(filename string, shouldexist bool) {

}
