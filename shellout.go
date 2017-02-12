package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/**
Runs a command with its options - notifies a tracker function on every line written to stout/stderr.
Makes a copy of stdout and stderr and returns it.
*/

func shellout_live(command string, options []string, fnTracker func(bool, string)) (error, string, string) {

	tokens := expand_tokens(options)
	log.Printf("%s %1v", command, tokens)
	cmd := exec.Command(command, tokens...)

	var sout bytes.Buffer
	var serr bytes.Buffer

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err, "", ""
	}

	cmdReaderE, errE := cmd.StderrPipe()
	if errE != nil {
		fmt.Fprintln(os.Stderr, "Error creating StderrPipe for Cmd", errE)
		return err, "", ""
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			sout.Write([]byte(text))
			if fnTracker != nil {
				fnTracker(true, text)
			}

		}
	}()

	scannerE := bufio.NewScanner(cmdReaderE)
	go func() {
		for scannerE.Scan() {
			text := scannerE.Text()
			serr.Write([]byte(text))

			if fnTracker != nil {
				fnTracker(false, text)
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		//fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err, "", ""
	}

	err = cmd.Wait()

	txtOut := sout.String()
	txtErr := serr.String()

	if err != nil {
		//fmt.Fprintf(os.Stderr, "Error waiting for Cmd %+v", err)
		return err, txtOut, txtErr
	}

	return nil, txtOut, txtErr
}

/**
Like shellout_live but uses no tracking function.
*/

func shellout(command string, options []string) (error, string, string) {
	return shellout_live(command, options, nil)
}

const TOKEN_PREFIX = "@@"

/**
 * Expands a list of tokens.
 */

func expand_tokens(options []string) []string {

	var out []string
	for _, v := range options {

		if len(v) > 0 {

			if strings.HasPrefix(v, TOKEN_PREFIX) {
				v1 := strings.TrimPrefix(v, TOKEN_PREFIX)
				out = append(out, v1)

			} else {
				vals := strings.Split(v, " ")
				for _, vv := range vals {
					out = append(out, vv)
				}
			}
		}
	}

	return out

}

func random_string() string {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}
