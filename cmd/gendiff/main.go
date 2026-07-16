package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"code"
	"code/internal/formatter"
)

func main() {
	cmd := &cli.Command{
		Name:                  "gendiff",
		Usage:                 "Compares two configuration files and shows a difference.",
		UsageText:             "gendiff [global options]",
		ArgsUsage:             "first_file second_file",
		HideHelpCommand:       true,
		EnableShellCompletion: false,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   formatter.OutputFormatStylish,
				Usage:   "output format",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 2 {
				return cli.ShowAppHelp(cmd)
			}

			firstFile := cmd.Args().Get(0)
			secondFile := cmd.Args().Get(1)
			format := cmd.String("format")

			diff, err := code.GenDiff(firstFile, secondFile, format)
			if err != nil {
				return err
			}

			fmt.Println(diff)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
