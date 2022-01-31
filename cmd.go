package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdbox.New("config", "file", "print", "dir", "dump", "saved", "updated")
	x.Version = `v1.0.1`
	x.Summary = `manage local configuration settings`
	x.Usage = `(file|dir|dump|print|saved|updated|<name> [<value>])`
	x.Description = `
		The *config* subcommand provides access to the underlying
		configuration data stored as JSON files at an expected location and
		ensures that any configuration changes by another program are not
		overriden by the current program using the same configuration. This
		helps avoid accidental contention between two running programs.

		The JSON data is always saved as a one-dimensional, flat map of
		strings. Depth may be achieved by using dotted or dashed notation
		for the keys but ultimately everything is saved in a single map of
		key/value pairs where the keys and values are both quoted UTF-8
		strings. This improves efficiency and avoids type inference issues. `

	x.Method = func(args []string) error {
		config, err := conf.New()
		if err != nil {
			return err
		}
		err = config.Load()
		if err != nil {
			return err
		}
		switch len(args) {
		case 0:
			config.PrettyPrint()
		case 1:
			switch args[0] {
			case "file":
				fmt.Println(config.Path())
			case "dir":
				fmt.Println(config.Dir)
			case "dump":
				fmt.Println(config)

			case "saved":
				fmt.Println(config.Saved)
			case "updated":
				fmt.Println(config.Updated)
			default:
				fmt.Println(config.Get(args[0]))
			}
		case 2:
			return config.SetSave(args[0], args[1])
		default:
			return x.UsageError()
		}
		return nil
	}
}
