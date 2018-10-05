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
package ganesha

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/pkg/capnslog"
	"github.com/rook/rook/pkg/clusterd"
	cephconfig "github.com/rook/rook/pkg/daemon/ceph/config"
)

var logger = capnslog.NewPackageLogger("github.com/rook/rook", "ganesha")

const (
	cephConfigPath = "/etc/ceph/ceph.conf"
)

type Config struct {
	Name        string
	ClusterInfo *cephconfig.ClusterInfo
}

func Run(context *clusterd.Context, config *Config) error {

	err := generateConfigFiles(context, config)
	if err != nil {
		return fmt.Errorf("failed to generate ganesha config files. %+v", err)
	}

	err = startGanesha(context, config)
	if err != nil {
		return fmt.Errorf("failed to run ganesha. %+v", err)
	}

	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan struct{})
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			logger.Infof("shutdown signal received, exiting...")
			close(stopChan)
			return nil
		}
	}
}

func generateConfigFiles(context *clusterd.Context, config *Config) error {
	// write the latest config to the config dir
	if err := cephconfig.GenerateAdminConnectionConfig(context, config.ClusterInfo); err != nil {
		return fmt.Errorf("failed to write connection config. %+v", err)
	}

	return nil
}

func startGanesha(context *clusterd.Context, config *Config) error {
	logger.Infof("starting ganesha server %s", config.Name)
	// For debug logging, add the params: "-N", "NIV_DEBUG"
	if err := context.Executor.ExecuteCommand(false, "", "ganesha.nfsd", "-F", "-L", "STDOUT"); err != nil {
		return fmt.Errorf("failed to start ganesha. %+v", err)
	}

	logger.Infof("started ganesha")
	return nil
}
