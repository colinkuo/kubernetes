/*
Copyright 2014 Google Inc. All rights reserved.

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

package cmd

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd/util"
)

func (f *Factory) NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information.",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunVersion(f, out, cmd)
			util.CheckErr(err)
		},
	}
	cmd.Flags().BoolP("client", "c", false, "Client version only (no server required).")
	return cmd
}

func RunVersion(f *Factory, out io.Writer, cmd *cobra.Command) error {
	if util.GetFlagBool(cmd, "client") {
		kubectl.GetClientVersion(out)
		return nil
	}

	client, err := f.Client(cmd)
	if err != nil {
		return err
	}

	kubectl.GetVersion(out, client)
	return nil
}
