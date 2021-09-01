// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/pingcap/ticdc/pkg/cmd/redo"
	"github.com/pingcap/ticdc/pkg/cmd/server"
	"github.com/pingcap/ticdc/pkg/cmd/version"
	"github.com/spf13/cobra"
)

// NewCmd creates the root command.
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cdc",
		Short: "CDC",
		Long:  `Change Data Capture`,
	}
}

// Run runs the root command.
func Run() {
	cmd := NewCmd()

	// Outputs cmd.Print to stdout.
	cmd.SetOut(os.Stdout)

	cmd.AddCommand(server.NewCmdServer())
	cmd.AddCommand(version.NewCmdVersion())
	cmd.AddCommand(redo.NewCmdRedo())

	if err := cmd.Execute(); err != nil {
		cmd.Println(err)
		os.Exit(1)
	}
}