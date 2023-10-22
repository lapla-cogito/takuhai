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

var comp, tracknum, alias string
var sagawa, yamato, jp bool

func (c *Cmd) regexec(cmd *cobra.Command, args []string) error {
	if err := c.list.Load(); err != nil {
		return err
	}
	if !sagawa && !yamato && !jp {
		comp = selectcompany()
	} else {
		if (sagawa && yamato) || (sagawa && jp) || (yamato && jp) {
			return errors.New("too many companies")
		}
	}

	if sagawa {
		comp = "sagawa"
	}
	if yamato {
		comp = "yamato"
	}
	if jp {
		comp = "jpost"
	}

	if comp == "" {
		return errors.New("company is empty")
	}

	if tracknum == "" {
		tracknum = inputtn()
	}

	if tracknum == "" {
		return errors.New("tracking number is empty")
	}

	nimotu := &track.Pkg{
		Company:  comp,
		TN:       tracknum,
		Alias:    alias,
		Timeline: []track.Stat{},
	}

	if err := nimotu.Tracknimotu(); err != nil {
		return nil
	}

	for _, p := range c.list.Pkgs {
		if p.Alias == alias {
			return errors.New("There is already exist the same name")
		}
	}

	if err := c.list.Addtoli(nimotu); err != nil {
		return err
	}

	if err := c.list.Save(""); err != nil {
		return err
	}

	fmt.Println("registered")
	return nil
}

func (c *Cmd) newregCmd() *cobra.Command {
	var regCmd = &cobra.Command{
		Use:   "reg",
		Short: "Registers a package",
		Long:  `Registers a package information by specifiying a tracking number and a company.`,
		RunE:  c.regexec,
	}
	regCmd.Flags().BoolVarP(&sagawa, "sagawa", "s", false, "sagawa")
	regCmd.Flags().BoolVarP(&yamato, "yamato", "y", false, "yamato")
	regCmd.Flags().BoolVarP(&jp, "jp", "j", false, "Japan post")
	regCmd.Flags().StringVarP(&tracknum, "tracknum", "t", "", "tracking number")
	regCmd.Flags().StringVarP(&alias, "name", "n", "", "name")

	return regCmd
}
