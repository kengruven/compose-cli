/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package compose

import (
	"context"
	"fmt"

	"github.com/docker/compose-cli/api/compose"

	"github.com/spf13/cobra"

	"github.com/docker/compose-cli/api/client"
)

type convertOptions struct {
	*projectOptions
	Format string
}

func convertCommand(p *projectOptions) *cobra.Command {
	opts := convertOptions{
		projectOptions: p,
	}
	convertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Converts the compose file to a cloud format (default: cloudformation)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConvert(cmd.Context(), opts)
		},
	}
	flags := convertCmd.Flags()
	flags.StringVar(&opts.Format, "format", "yaml", "Format the output. Values: [yaml | json]")

	return convertCmd
}

func runConvert(ctx context.Context, opts convertOptions) error {
	var json []byte
	c, err := client.NewWithDefaultLocalBackend(ctx)
	if err != nil {
		return err
	}

	project, err := opts.toProject()
	if err != nil {
		return err
	}

	json, err = c.ComposeService().Convert(ctx, project, compose.ConvertOptions{
		Format: opts.Format,
	})
	if err != nil {
		return err
	}

	fmt.Println(string(json))
	return nil
}
