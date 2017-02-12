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

// This form lets us assert a lot of cases.
var resizingTests = []struct {
	name string
	w    int // in
	h    int
	ew   int // expected w
	eh   int
}{
	{"no resize", 640, 320, 640, 320},
	{"halved", 1280, 640, 640, 320},
	{"8 boundary - small", 321, 224, 320, 224},
	{"8 boundary - large", 327, 327, 328, 328},
}

func TestResize_byCase(t *testing.T) {
	for _, tst := range resizingTests {
		w, h := normalizeStreamSize(tst.w, tst.h, TST_MAXWIDTH, TST_BLOCK)
		assertSize(t, tst.name, tst.ew, tst.eh, w, h)
	}
}

func assertSize(t *testing.T, spec string, expW int, expH int, myW int, myH int) {
	if myW != expW || myH != expH {
		t.Errorf("Spec: %s - Expected (%d x %d) - got (%d x %d)", spec, expW, expH, myW, myH)
	}
}
