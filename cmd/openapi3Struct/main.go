package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/urfave/cli/v3"
	"github.com/wexder/openapi3Struct"
)

func main() {
	cmd := &cli.Command{
		Name:  "generate",
		Usage: "generate schema from openapi",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "packages",
				Usage: "path for packages",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "output file",
				Value: "openapi.yaml",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			paths := cmd.StringSlice("packages")
			if len(paths) == 0 {
				return fmt.Errorf("packages flag is required")
			}
			fmt.Printf("Generating schema from packages: %v\n", paths)

			parser := openapi3Struct.NewParser(openapi3.T{}, openapi3Struct.WithPackagePaths(paths))

			err := parser.ParseSchemasFromStructs()
			if err != nil {
				return err
			}

			err = parser.SaveYamlToFile(cmd.String("output"))
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
