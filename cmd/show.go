/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"takuhai/track"
)

var gettracknum, getalias string
var all bool

func (c *Cmd) showExec(cmd *cobra.Command, args []string) error {
	if err := c.list.Load(); err != nil {
		return err
	}

	if all {
		if len(c.list.Pkgs) == 0 {
			return errors.New("no packages")
		}

		for _, pkg := range c.list.Pkgs {
			fmt.Printf("Alias: %s -> %s\n", pkg.Alias, pkg.Timeline[len(pkg.Timeline)-1].Status)
		}

		return nil
	}

	if gettracknum != "" && getalias != "" {
		return errors.New("too many options")
	}

	if getalias != "" {
		gettracknum = c.list.Gettn(getalias)
		if gettracknum == "" {
			return errors.New("no such alias")
		}
	}

	if gettracknum == "" {
		return errors.New("no tracking number specified")
	}

	if resSagawa, _ := track.Sagawatrack(gettracknum); resSagawa != nil {
		fmt.Printf("The package status: %s %s %s\n", resSagawa[len(resSagawa)-1].Date, resSagawa[len(resSagawa)-1].Status, resSagawa[len(resSagawa)-1].Office)
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resSagawa
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			return err
		}
		return nil
	}

	if resYamato, _ := track.Yamatotrack(gettracknum); resYamato != nil {
		fmt.Printf("The package status: %s %s %s\n", resYamato[len(resYamato)-1].Date, resYamato[len(resYamato)-1].Status, resYamato[len(resYamato)-1].Office)
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resYamato
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			return err
		}
		return nil
	}

	if resJPost, _ := track.JPosttrack(gettracknum); resJPost != nil {
		fmt.Printf("The package status: %s %s %s\n", resJPost[len(resJPost)-1].Date, resJPost[len(resJPost)-1].Status, resJPost[len(resJPost)-1].Office)
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resJPost
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (c *Cmd) newshowCmd() *cobra.Command {
	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "Shows the state of a specific package",
		Long: `Shows the state of a specific package.
You can input tracking number or alias name you registered as tracking key as follows:

$ takuhai -t 1234567890
$ takuhai -n hoge

Here, the -t option means to use tracking number as tracking key, alias name in the -n option.
`,
		RunE: c.showExec,
	}
	showCmd.Flags().StringVarP(&gettracknum, "tracknum", "t", "", "tracking number")
	showCmd.Flags().StringVarP(&getalias, "name", "n", "", "name")
	showCmd.Flags().BoolVarP(&all, "all", "a", false, "all packages")

	return showCmd
}
