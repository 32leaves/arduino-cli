// This file is part of arduino-cli.
//
// Copyright 2019 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to modify or
// otherwise use the software for commercial activities involving the Arduino
// software without disclosing the source code of your own applications. To purchase
// a commercial license, send an email to license@arduino.cc.

package feedback

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Feedback wraps an io.Writer and provides an uniform API the CLI can use to
// provide feedback to the users.
type Feedback struct {
	out io.Writer
	err io.Writer
}

// New creates a Feedback instance
func New(out, err io.Writer) *Feedback {
	return &Feedback{
		out: out,
		err: err,
	}
}

// DefaultFeedback provides a basic feedback object to be used as default.
func DefaultFeedback() *Feedback {
	return New(os.Stdout, os.Stderr)
}

// OutputWriter returns the underlying io.Writer to be used when the Print*
// api is not enough.
func (fb *Feedback) OutputWriter() io.Writer {
	return fb.out
}

// ErrorWriter is the same as OutputWriter but exposes the underlying error
// writer.
func (fb *Feedback) ErrorWriter() io.Writer {
	return fb.out
}

// Printf behaves like fmt.Printf but writes on the out writer and adds a newline.
func (fb *Feedback) Printf(format string, v ...interface{}) {
	fb.Print(fmt.Sprintf(format, v...))
}

// Print behaves like fmt.Print but writes on the out writer and adds a newline.
func (fb *Feedback) Print(v ...interface{}) {
	fmt.Fprintln(fb.out, v...)
}

// Errorf behaves like fmt.Printf but writes on the error writer and adds a
// newline. It also logs the error.
func (fb *Feedback) Errorf(format string, v ...interface{}) {
	fb.Error(fmt.Sprintf(format, v...))
}

// Error behaves like fmt.Print but writes on the error writer and adds a
// newline. It also logs the error.
func (fb *Feedback) Error(v ...interface{}) {
	fmt.Fprintln(fb.err, v...)
	logrus.Error(fmt.Sprint(v...))
}

// PrintJSON is a convenient wrapper to provide feedback by printing the
// desired output in a pretty JSON format. It adds a newline to the output.
func (fb *Feedback) PrintJSON(v interface{}) {
	if d, err := json.MarshalIndent(v, "", "  "); err != nil {
		fb.Errorf("Error during JSON encoding of the output: %v", err)
	} else {
		fb.Print(string(d))
	}
}
