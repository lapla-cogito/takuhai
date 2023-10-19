/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var deregtracknum, deregalias string

func (c *Cmd) deregExec(cmd *cobra.Command, args []string) error {
	if err := c.list.Load(); err != nil {
		return err
	}

	if deregtracknum != "" && deregalias != "" {
		return errors.New("too many options")
	}

	if deregalias != "" {
		deregtracknum = c.list.Gettn(deregalias)
		if deregtracknum == "" {
			return errors.New("no such alias")
		}
	}
	if deregtracknum == "" {
		return errors.New("no tracking number specified")
	}

	for i, pkg := range c.list.Pkgs {
		if pkg.TN == deregtracknum {
			c.list.Pkgs = append(c.list.Pkgs[:i], c.list.Pkgs[i+1:]...)
			break
		}
	}

	if err := c.list.Save(""); err != nil {
		return err
	}

	fmt.Println("deregistered")
	return nil
}

func (c *Cmd) newderegCmd() *cobra.Command {
	var deregCmd = &cobra.Command{
		Use:   "dereg",
		Short: "Deregister the specific package",
		Long: `Deregister the specific package.
You can input tracking number or package name you registered as tracking key as follows:

$ takuhai -t 1234567890
$ takuhai -n hoge

Here, the -t option means to use tracking number as tracking key, package name in the -n option.
`,
		RunE: c.deregExec,
	}
	deregCmd.Flags().StringVarP(&deregtracknum, "tracknum", "t", "", "tracking number")
	deregCmd.Flags().StringVarP(&deregalias, "name", "n", "", "name")

	return deregCmd
}
