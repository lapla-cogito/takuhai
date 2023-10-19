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

var exportpath string
var aliases, tns []string

func (c *Cmd) exportExec(cmd *cobra.Command, args []string) error {
	fmt.Println("exporting...")
	var exportedData = track.New()
	var added bool

	if exportpath == "" {
		return errors.New("no file path specified")
	}

	if len(aliases) == 0 && len(tns) == 0 {
		return errors.New("no package specified")
	}

	if err := c.list.Load(); err != nil {
		return err
	}

	merged := make([][]string, 0)
	for _, alias := range aliases {
		merged = append(merged, []string{alias, "aliases"})
	}
	for _, tn := range tns {
		merged = append(merged, []string{tn, "tns"})
	}

	for _, el := range merged {
		for _, pkg := range c.list.Pkgs {
			if pkg.Alias == el[0] && el[1] == "aliases" {
				if err := exportedData.Addtoli(&pkg); err != nil {
					return err
				}
				added = true
				break
			} else if pkg.TN == el[0] && el[1] == "tns" {
				if err := exportedData.Addtoli(&pkg); err != nil {
					return err
				}
				added = true
				break
			}
		}
		if !added {
			fmt.Println("no such package: " + alias)
		}
	}

	if err := exportedData.Save(exportpath); err != nil {
		return err
	}

	fmt.Println("exported to " + exportpath)

	return nil
}

func (c *Cmd) newexportCmd() *cobra.Command {
	var regCmd = &cobra.Command{
		Use:   "export",
		Short: "Export package informations",
		Long:  `Export package informations to a file.\n`,
		RunE:  c.exportExec,
	}
	regCmd.Flags().StringVarP(&exportpath, "path", "p", "", "file path")
	regCmd.Flags().StringArrayVarP(&aliases, "name", "n", []string{}, "package name")
	regCmd.Flags().StringArrayVarP(&tns, "tracknum", "t", []string{}, "tracking numbers")

	return regCmd
}
