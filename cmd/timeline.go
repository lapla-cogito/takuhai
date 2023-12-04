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

func (c *Cmd) tlExec(cmd *cobra.Command, args []string) error {
	if err := c.list.Load(); err != nil {
		return err
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
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resSagawa
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			for _, tl := range resSagawa {
				fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
			}
			return errors.New("failed to save latest status")
		}

		for _, tl := range resSagawa {
			fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
		}
		return nil
	}

	if resYamato, _ := track.Yamatotrack(gettracknum); resYamato != nil {
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resYamato
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			for _, tl := range resYamato {
				fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
			}
			return errors.New("failed to save latest status")
		}

		for _, tl := range resYamato {
			fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
		}
		return nil
	}

	if resJPost, _ := track.JPosttrack(gettracknum); resJPost != nil {
		for i, pkg := range c.list.Pkgs {
			if pkg.TN == gettracknum {
				c.list.Pkgs[i].Timeline = resJPost
				break
			}
		}
		if err := c.list.Save(""); err != nil {
			for _, tl := range resJPost {
				fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
			}
			return errors.New("failed to save latest status")
		}

		for _, tl := range resJPost {
			fmt.Printf("The package status: %s %s %s\n", tl.Date, tl.Status, tl.Office)
		}
		return nil
	}

	return nil
}

func (c *Cmd) newtlCmd() *cobra.Command {
	var tlCmd = &cobra.Command{
		Use:   "timeline",
		Short: "Shows the timeline of a specific package",
		Long: `Shows the timeline of a specific package.
You can input tracking number or alias name you registered as tracking key as follows:

$ takuhai timeline -t 1234567890
$ takuhai timeline -n hoge

Here, the -t option means to use tracking number as tracking key, alias name in the -n option.
`,
		RunE: c.tlExec,
	}
	tlCmd.Flags().StringVarP(&gettracknum, "tracknum", "t", "", "tracking number")
	tlCmd.Flags().StringVarP(&getalias, "name", "n", "", "name")

	return tlCmd
}
