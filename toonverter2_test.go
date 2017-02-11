package main

import (
	"testing"
)

const TST_BLOCK = 8
const TST_MAXWIDTH = 640

/**
Input #0, avi, from './file.avi':
  Metadata:
    encoder         : Some
  Duration: 01:43:45.76, start: 0.000000, bitrate: 1747 kb/s
    Stream #0:0: Video: mpeg4 (Advanced Simple Profile) (XVID / 0x44495658), yuv420p, 544x384 [SAR 1:1 DAR 17:12], 965 kb/s, 25 fps, 25 tbr, 25 tbn, 25 tbc
    Stream #0:1: Audio: ac3 ([0] [0][0] / 0x2000), 48000 Hz, 5.1(side), fltp, 384 kb/s
    Stream #0:2: Audio: ac3 ([0] [0][0] / 0x2000), 48000 Hz, 5.1(side), fltp, 384 kb/s
*/

func TestResizerText(t *testing.T) {

	in := `
Input #0, avi, from './file.avi':
  Metadata:
    encoder         : Some
  Duration: 01:43:45.76, start: 0.000000, bitrate: 1747 kb/s
    Stream #0:0: Video: mpeg4 (Advanced Simple Profile) (XVID / 0x44495658), yuv420p, 544x384 [SAR 1:1 DAR 17:12], 965 kb/s, 25 fps, 25 tbr, 25 tbn, 25 tbc
    Stream #0:1: Audio: ac3 ([0] [0][0] / 0x2000), 48000 Hz, 5.1(side), fltp, 384 kb/s
    Stream #0:2: Audio: ac3 ([0] [0][0] / 0x2000), 48000 Hz, 5.1(side), fltp, 384 kb/s`

	val := resizer(in)

	if val != "544x384" {
		t.Errorf("Wrong size %s", val)
	}
}

func TestCurrStreamSize(t *testing.T) {
	w, h := findCurrentStreamFrameSize("aa 10x20 bb")
	assertSize(t, "Stream size from string", 10, 20, w, h)
}

func TestResize_noResize(t *testing.T) {
	w, h := normalizeStreamSize(640, 320, TST_MAXWIDTH, TST_BLOCK)
	assertSize(t, "No resizing", 640, 320, w, h)
}

func TestResize_halve(t *testing.T) {
	w, h := normalizeStreamSize(1280, 640, TST_MAXWIDTH, TST_BLOCK)
	assertSize(t, "Halved", 640, 320, w, h)
}

func TestResize_not_8_smaller(t *testing.T) {
	w, h := normalizeStreamSize(321, 224, TST_MAXWIDTH, TST_BLOCK)
	assertSize(t, "Not 8 s ", 320, 224, w, h)
}

func TestResize_not_8_bigger(t *testing.T) {
	w, h := normalizeStreamSize(327, 327, TST_MAXWIDTH, TST_BLOCK)
	assertSize(t, "Not 8 b", 328, 328, w, h)
}

func assertSize(t *testing.T, spec string, expW int, expH int, myW int, myH int) {
	if myW != expW || myH != expH {
		t.Errorf("Spec: %s - Expected (%d x %d) - got (%d x %d)", spec, expW, expH, myW, myH)
	}
}
