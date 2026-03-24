package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

type Config struct {
	human     bool
	hidden    bool
	recursive bool
}

func main() {
	cfg := Config{human: false, hidden: false, recursive: false}

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "human-readable sizes (auto-select unit)",
				Destination: &cfg.human,
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "include hidden files and directories",
				Destination: &cfg.hidden,
			},
			&cli.BoolFlag{
				Name:        "resursive",
				Aliases:     []string{"r"},
				Usage:       "recursive size of directories",
				Destination: &cfg.recursive,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			file_path := c.Args().Get(0)
			size, err := code.GetPathSize(file_path, cfg.hidden, cfg.recursive)
			if err != nil {
				return err
			}
			sizeFmt := code.FormatSize(size, cfg.human)

			str := fmt.Sprintf("%s\t%s", sizeFmt, file_path)
			fmt.Println(str)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
