/*
Copyright 2016 The Rook Authors. All rights reserved.

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
package ceph

import (
	"github.com/rook/rook/cmd/rook/rook"
	"github.com/rook/rook/pkg/daemon/ceph/ganesha"
	"github.com/rook/rook/pkg/daemon/ceph/mon"
	"github.com/rook/rook/pkg/util/flags"
	"github.com/spf13/cobra"
)

var (
	ganeshaName             string
	ganeshaCopyBinariesPath string
)

var ganeshaCmd = &cobra.Command{
	Use:    "ganesha",
	Short:  "Configures and runs the nfs ganesha daemon",
	Hidden: true,
}
var ganeshaConfigCmd = &cobra.Command{
	Use:    "init",
	Short:  "Updates ceph.conf for ganesha",
	Hidden: true,
}
var ganeshaRunCmd = &cobra.Command{
	Use:    "run",
	Short:  "Runs the nfs ganesha daemon",
	Hidden: true,
}

func init() {
	ganeshaRunCmd.Flags().StringVar(&ganeshaName, "ganesha-name", "", "name of the ganesha server")
	ganeshaConfigCmd.Flags().StringVar(&ganeshaCopyBinariesPath, "copy-binaries-path", "", "If specified, copy the rook binaries to this path for use by the daemon container")
	addCephFlags(ganeshaConfigCmd)

	ganeshaCmd.AddCommand(ganeshaConfigCmd)
	ganeshaCmd.AddCommand(ganeshaRunCmd)

	flags.SetFlagsFromEnv(ganeshaCmd.Flags(), rook.RookEnvVarPrefix)
	flags.SetFlagsFromEnv(ganeshaConfigCmd.Flags(), rook.RookEnvVarPrefix)
	flags.SetFlagsFromEnv(ganeshaRunCmd.Flags(), rook.RookEnvVarPrefix)

	ganeshaConfigCmd.RunE = configGanesha
	ganeshaRunCmd.RunE = startGanesha
}

func configGanesha(cmd *cobra.Command, args []string) error {
	required := []string{"copy-binaries-path", "mon-endpoints", "cluster-name", "admin-secret", "public-ip", "private-ip"}
	if err := flags.VerifyRequiredFlags(ganeshaConfigCmd, required); err != nil {
		return err
	}

	rook.SetLogLevel()
	rook.LogStartupInfo(ganeshaConfigCmd.Flags())

	clusterInfo.Monitors = mon.ParseMonEndpoints(cfg.monEndpoints)

	// generate the ceph config
	err := ganesha.Initialize(createContext(), &clusterInfo)
	if err != nil {
		rook.TerminateFatal(err)
	}

	// copy the rook and tini binaries for use by the daemon container
	copyBinaries(ganeshaCopyBinariesPath)

	return nil
}

func startGanesha(cmd *cobra.Command, args []string) error {
	required := []string{"ganesha-name"}
	if err := flags.VerifyRequiredFlags(ganeshaRunCmd, required); err != nil {
		return err
	}

	rook.SetLogLevel()
	rook.LogStartupInfo(ganeshaCmd.Flags())

	err := ganesha.Run(createContext(), ganeshaName)
	if err != nil {
		rook.TerminateFatal(err)
	}

	return nil
}
