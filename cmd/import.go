/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"takuhai/track"
)

var importpath string
var ymlData interface{}

func (c *Cmd) importExec(cmd *cobra.Command, args []string) error {
	if track.IsExist(importpath) {
		return errors.New("No such file: " + importpath)
	}

	if err := c.list.Load(); err != nil {
		return err
	}

	buf, err := os.ReadFile(importpath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &ymlData)
	if err != nil {
		return err
	}

	err = c.list.Merge(ymlData)
	if err != nil {
		return err
	}

	if err := c.list.Save(""); err != nil {
		return err
	}

	return nil
}

func (c *Cmd) newimportCmd() *cobra.Command {
	var regCmd = &cobra.Command{
		Use:   "import",
		Short: "Import package informations from exported yaml file",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
	
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: c.importExec,
	}
	regCmd.Flags().StringVarP(&importpath, "path", "p", "", "file path")

	return regCmd
}
