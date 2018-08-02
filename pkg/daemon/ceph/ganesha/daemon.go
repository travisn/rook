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
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/coreos/pkg/capnslog"
	"github.com/rook/rook/pkg/clusterd"
	cephconfig "github.com/rook/rook/pkg/daemon/ceph/config"
)

var logger = capnslog.NewPackageLogger("github.com/rook/rook", "ganesha")

const (
	cephConfigPath    = "/etc/ceph/ceph.conf"
	ganeshaConfigPath = "/etc/ganesha/ganesha.conf"
	ganeshaConfig     = `NFS_CORE_PARAM {
	Enable_NLM = false;
	Enable_RQUOTA = false;
	Protocols = 4;
}

CACHEINODE {
	Dir_Max = 1;
	Dir_Chunk = 0;
	Cache_FDs = false;
	NParts = 1;
	Cache_Size = 1;
}

EXPORT_DEFAULTS {
	Attr_Expiration_Time = 0;
}

RADOS_URLS {
ceph_conf = '$(CEPH_CONFIG_PATH)';
	userid = '$(USER_ID)';
}
# TODO: Replace the rados url with the reference path to /etc/ganesha/ganesha.conf
#%url rados://mypool/object

NFSv4 {
RecoveryBackend = 'rados_kv';
Minor_Versions = 1, 2;
}
RADOS_KV {
ceph_conf = '$(CEPH_CONFIG_PATH)';
userid = '$(USER_ID)';
#pool = 'mypool';
}
`
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

	// save the keyring to the given path
	if err := os.MkdirAll(filepath.Dir(ganeshaConfigPath), 0744); err != nil {
		return fmt.Errorf("failed to create ganesha config directory %s: %+v", ganeshaConfigPath, err)
	}
	if err := ioutil.WriteFile(ganeshaConfigPath, []byte(generateGaneshaConfig(config)), 0644); err != nil {
		return fmt.Errorf("failed to ganesha config to %s: %+v", ganeshaConfigPath, err)
	}

	return nil
}

func generateGaneshaConfig(config *Config) string {
	configPath := "/etc/ganesha/config/export.conf"
	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Warningf("failed to read file %s. %+v", configPath, err)
	} else {
		logger.Infof("config file %s contents:\n%s", configPath, string(contents))
	}

	// Replace the placeholder settings in the ganesh config
	r := strings.NewReplacer(
		"$(USER_ID)", "admin",
		"$(CEPH_CONFIG_PATH)", cephConfigPath)
	return r.Replace(ganeshaConfig)
}

func startGanesha(context *clusterd.Context, config *Config) error {
	logger.Infof("starting rpcbind for ganesha")
	if err := context.Executor.ExecuteCommand(false, "", "rpcbind"); err != nil {
		return fmt.Errorf("failed to start mds. %+v", err)
	}

	logger.Infof("starting ganesha server %s", config.Name)
	// For debug logging, add the params: "-N", "NIV_DEBUG"
	if err := context.Executor.ExecuteCommand(false, "", "ganesha.nfsd", "-F", "-L", "STDOUT"); err != nil {
		return fmt.Errorf("failed to start ganesha. %+v", err)
	}

	logger.Infof("started ganesha")
	return nil
}
