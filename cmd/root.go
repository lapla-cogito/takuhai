/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"takuhai/track"
)

type Cmd struct {
	root *cobra.Command
	list *track.List
}

func (c *Cmd) Execute() {
	err := c.root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func New(li *track.List) *Cmd {
	var rootCmd = &Cmd{
		root: &cobra.Command{
			Use:   "takuhai",
			Short: "A CLI application to track packages",
			Long: `A CLI application to track packages you registered.
currently, this application can track:
- SAGAWA TRANSPORTATION CO., LTD.
- YAMATO TRANSPORT CO., LTD.
- Nippon Express Co., LTD.
	`,
		},
		list: li,
	}
	rootCmd.root.AddCommand(
		rootCmd.newregCmd(),
		rootCmd.newshowCmd(),
		rootCmd.newderegCmd(),
		rootCmd.newimportCmd(),
		rootCmd.newexportCmd(),
		rootCmd.newrenameCmd(),
	)

	return rootCmd

}
