/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var oldname, newname string

func (c *Cmd) renameexec(cmd *cobra.Command, args []string) error {

	if oldname == "" || newname == "" {
		return errors.New("please specify both oldname and newname")
	}

	if err := c.list.Load(); err != nil {
		return err
	}

	for i, p := range c.list.Pkgs {
		if p.Alias == oldname {
			fmt.Println("found a package named " + oldname + "\nrenaming...")
			c.list.Pkgs[i].Alias = newname
			if err := c.list.Save(""); err != nil {
				return err
			}
			fmt.Println("renamed from " + oldname + " to " + newname)
			return nil
		}
	}
	fmt.Println("no such package named " + oldname)
	return nil
}

func (c *Cmd) newrenameCmd() *cobra.Command {
	var renameCmd = &cobra.Command{
		Use:   "rename",
		Short: "Rename a package",
		Long:  `Registers a package information by specifiying a tracking number and a company.`,
		RunE:  c.renameexec,
	}

	renameCmd.Flags().StringVarP(&oldname, "old", "o", "", "oldname")
	renameCmd.Flags().StringVarP(&newname, "new", "n", "", "newname")

	return renameCmd
}
