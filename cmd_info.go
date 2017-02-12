package main

import (
	"fmt"
	"strings"
)

func cmd_Info(common VideoDefaults) {

	shellout_live("ffmpeg", []string{
		"-i", TOKEN_PREFIX + common.Source,
	}, show_info_lines)

	newsize := newSizeForFile(common.Source)
	fmt.Printf("New size: %s\n", newsize)
	fmt.Printf("New video bandwidth: %s\n", common.VideoBandwith)

}

func show_info_lines(stdin bool, line string) {

	//fmt.Printf("? %s\n", line)
	if strings.Contains(line, "Duration") || strings.Contains(line, "Stream #") {
		fmt.Printf("> %s\n", line)
	}

}

// Duration:
// Stream: #
