// package hexletpathsize
package main

import (
	path_size "code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	// fmt.Println(("hello"))

	(&cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			file_path := c.Args().Get(0)
			size, err := path_size.GetSize(file_path)
			if err != nil {
				return err
			}

			str := fmt.Sprintf("%dB\t%s", size, file_path)
			fmt.Println(str)

			return nil
		},
	}).Run(context.Background(), os.Args)
}
