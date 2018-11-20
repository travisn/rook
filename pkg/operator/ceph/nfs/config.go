/*
Copyright 2018 The Rook Authors. All rights reserved.

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

// Package nfs for NFS ganesha
package nfs

import "fmt"
import cephv1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"

const (
	cephConfigPath = "/etc/ceph/ceph.conf"
	userID         = "admin"
)

func getGaneshaNodeID(n cephv1beta1.NFSGanesha, name string) string {
	return fmt.Sprintf("%s.%s", n.Name, name)
}

func getGaneshaConfig(n cephv1beta1.NFSGanesha, name string) string {
	nodeID := getGaneshaNodeID(n, name)
	return `
NFS_CORE_PARAM {
	Enable_NLM = false;
	Enable_RQUOTA = false;
	Protocols = 4;
}

CACHEINODE {
	Dir_Chunk = 0;
	NParts = 1;
	Cache_Size = 1;
}

EXPORT_DEFAULTS {
	Attr_Expiration_Time = 0;
}

NFSv4 {
	Delegations = false;
	RecoveryBackend = 'rados_cluster';
	Minor_Versions = 1, 2;
}

RADOS_KV {
	ceph_conf = '` + cephConfigPath + `';
	userid = ` + userID + `;
	nodeid = ` + nodeID + `;
	pool = "` + n.Spec.RADOS.Pool + `";
	namespace = "` + n.Spec.RADOS.Namespace + `";
}
`
}
