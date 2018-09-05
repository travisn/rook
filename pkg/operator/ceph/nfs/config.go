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

import cephv1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"

func getGaneshaConfig(spec cephv1beta1.NFSGaneshaSpec) string {
	return getExportConfig(spec) + getRadosKVConfig(spec)
}

func getExportConfig(spec cephv1beta1.NFSGaneshaSpec) string {
	var config string
	for _, export := range spec.Exports {
		config += `
EXPORT
{
	Export_ID=100;
	Protocols = 4;
	Transports = TCP;
	Attr_Expiration_Time = 0;
	Delegations = R;
	Path = ` + export.Path + `;
	Pseudo = ` + getPseudoPathWithDefault(export) + `;
	Squash = ` + getRootSquashWithDefault(export.Squash) + `;
	Access_Type = ` + convertAccessType(export.AccessType) + `;
	FSAL {
		Name = CEPH;
	}
` + getAllowedClientConfig(export.AllowedClients) + `
}
`
	}
	return config
}

func getAllowedClientConfig(allowedClients []cephv1beta1.NFSAllowedClient) string {
	var config string
	for _, client := range allowedClients {
		config += `
	CLIENT
	{
		Clients = ` + client.Clients + `;
		Squash = ` + getSquashWithDefault(client.Squash) + `;
		Access_Type = ` + convertAccessType(client.AccessType) + `;
	}
 `
	}

	return config
}

func getRadosKVConfig(spec cephv1beta1.NFSGaneshaSpec) string {
	return `
RADOS_KV
{
	pool = "` + spec.ClientRecovery.Pool + `";
	namespace = "` + spec.ClientRecovery.Namespace + `";
}`
}

func getPseudoPathWithDefault(export cephv1beta1.GaneshaExportSpec) string {
	if export.PseudoPath != "" {
		return export.PseudoPath
	}
	// default to use the same pseudopath as the path
	return export.Path
}

func getSquashWithDefault(squash string) string {
	if squash == "" {
		// set the default squash to "none"
		return "None"
	}
	return squash
}

func getRootSquashWithDefault(squash string) string {
	if squash == "" {
		// set the default squash for the root
		return "No_root_squash"
	}
	return squash
}

func convertAccessType(mode string) string {
	if mode == "ReadOnly" {
		return "RO"
	} else if mode == "ReadWrite" {
		return "RW"
	}
	return mode
}
