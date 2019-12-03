// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	v1alpha2 "github.com/rook/rook/pkg/apis/rook.io/v1alpha2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditdConf) DeepCopyInto(out *AuditdConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditdConf.
func (in *AuditdConf) DeepCopy() *AuditdConf {
	if in == nil {
		return nil
	}
	out := new(AuditdConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowConf) DeepCopyInto(out *CcowConf) {
	*out = *in
	out.Trlog = in.Trlog
	out.Tenant = in.Tenant
	out.Network = in.Network
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowConf.
func (in *CcowConf) DeepCopy() *CcowConf {
	if in == nil {
		return nil
	}
	out := new(CcowConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowNetwork) DeepCopyInto(out *CcowNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowNetwork.
func (in *CcowNetwork) DeepCopy() *CcowNetwork {
	if in == nil {
		return nil
	}
	out := new(CcowNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowTenant) DeepCopyInto(out *CcowTenant) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowTenant.
func (in *CcowTenant) DeepCopy() *CcowTenant {
	if in == nil {
		return nil
	}
	out := new(CcowTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowTrlog) DeepCopyInto(out *CcowTrlog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowTrlog.
func (in *CcowTrlog) DeepCopy() *CcowTrlog {
	if in == nil {
		return nil
	}
	out := new(CcowTrlog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowdBgConfig) DeepCopyInto(out *CcowdBgConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowdBgConfig.
func (in *CcowdBgConfig) DeepCopy() *CcowdBgConfig {
	if in == nil {
		return nil
	}
	out := new(CcowdBgConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowdConf) DeepCopyInto(out *CcowdConf) {
	*out = *in
	out.BgConfig = in.BgConfig
	out.Network = in.Network
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowdConf.
func (in *CcowdConf) DeepCopy() *CcowdConf {
	if in == nil {
		return nil
	}
	out := new(CcowdConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcowdNetwork) DeepCopyInto(out *CcowdNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcowdNetwork.
func (in *CcowdNetwork) DeepCopy() *CcowdNetwork {
	if in == nil {
		return nil
	}
	out := new(CcowdNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeploymentConfig) DeepCopyInto(out *ClusterDeploymentConfig) {
	*out = *in
	if in.Directories != nil {
		in, out := &in.Directories, &out.Directories
		*out = make([]RtlfsDevice, len(*in))
		copy(*out, *in)
	}
	if in.DevConfig != nil {
		in, out := &in.DevConfig, &out.DevConfig
		*out = make(map[string]DevicesConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeploymentConfig.
func (in *ClusterDeploymentConfig) DeepCopy() *ClusterDeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterDeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.AnnotationsSpec, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(v1alpha2.Annotations, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = make(v1alpha2.PlacementSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Network.DeepCopyInto(&out.Network)
	out.Dashboard = in.Dashboard
	in.Resources.DeepCopyInto(&out.Resources)
	out.DataVolumeSize = in.DataVolumeSize.DeepCopy()
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	out.MaxContainerCapacity = in.MaxContainerCapacity.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpec) DeepCopyInto(out *DashboardSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpec.
func (in *DashboardSpec) DeepCopy() *DashboardSpec {
	if in == nil {
		return nil
	}
	out := new(DashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicesConfig) DeepCopyInto(out *DevicesConfig) {
	*out = *in
	in.Rtrd.DeepCopyInto(&out.Rtrd)
	if in.RtrdSlaves != nil {
		in, out := &in.RtrdSlaves, &out.RtrdSlaves
		*out = make([]RTDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Rtlfs.DeepCopyInto(&out.Rtlfs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicesConfig.
func (in *DevicesConfig) DeepCopy() *DevicesConfig {
	if in == nil {
		return nil
	}
	out := new(DevicesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicesResurrectOptions) DeepCopyInto(out *DevicesResurrectOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicesResurrectOptions.
func (in *DevicesResurrectOptions) DeepCopy() *DevicesResurrectOptions {
	if in == nil {
		return nil
	}
	out := new(DevicesResurrectOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISCSI) DeepCopyInto(out *ISCSI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISCSI.
func (in *ISCSI) DeepCopy() *ISCSI {
	if in == nil {
		return nil
	}
	out := new(ISCSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISCSI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISCSIList) DeepCopyInto(out *ISCSIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ISCSI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISCSIList.
func (in *ISCSIList) DeepCopy() *ISCSIList {
	if in == nil {
		return nil
	}
	out := new(ISCSIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISCSIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISCSISpec) DeepCopyInto(out *ISCSISpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.Annotations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.TargetParams = in.TargetParams
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISCSISpec.
func (in *ISCSISpec) DeepCopy() *ISCSISpec {
	if in == nil {
		return nil
	}
	out := new(ISCSISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISGW) DeepCopyInto(out *ISGW) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISGW.
func (in *ISGW) DeepCopy() *ISGW {
	if in == nil {
		return nil
	}
	out := new(ISGW)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISGW) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISGWList) DeepCopyInto(out *ISGWList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ISGW, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISGWList.
func (in *ISGWList) DeepCopy() *ISGWList {
	if in == nil {
		return nil
	}
	out := new(ISGWList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISGWList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISGWSpec) DeepCopyInto(out *ISGWSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.Annotations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISGWSpec.
func (in *ISGWSpec) DeepCopy() *ISGWSpec {
	if in == nil {
		return nil
	}
	out := new(ISGWSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFS) DeepCopyInto(out *NFS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFS.
func (in *NFS) DeepCopy() *NFS {
	if in == nil {
		return nil
	}
	out := new(NFS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSList) DeepCopyInto(out *NFSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSList.
func (in *NFSList) DeepCopy() *NFSList {
	if in == nil {
		return nil
	}
	out := new(NFSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSSpec) DeepCopyInto(out *NFSSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.Annotations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSSpec.
func (in *NFSSpec) DeepCopy() *NFSSpec {
	if in == nil {
		return nil
	}
	out := new(NFSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RTDevice) DeepCopyInto(out *RTDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RTDevice.
func (in *RTDevice) DeepCopy() *RTDevice {
	if in == nil {
		return nil
	}
	out := new(RTDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RTDevices) DeepCopyInto(out *RTDevices) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RTDevice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RTDevices.
func (in *RTDevices) DeepCopy() *RTDevices {
	if in == nil {
		return nil
	}
	out := new(RTDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RtlfsDevice) DeepCopyInto(out *RtlfsDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RtlfsDevice.
func (in *RtlfsDevice) DeepCopy() *RtlfsDevice {
	if in == nil {
		return nil
	}
	out := new(RtlfsDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RtlfsDevices) DeepCopyInto(out *RtlfsDevices) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RtlfsDevice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RtlfsDevices.
func (in *RtlfsDevices) DeepCopy() *RtlfsDevices {
	if in == nil {
		return nil
	}
	out := new(RtlfsDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3List) DeepCopyInto(out *S3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3List.
func (in *S3List) DeepCopy() *S3List {
	if in == nil {
		return nil
	}
	out := new(S3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.Annotations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3X) DeepCopyInto(out *S3X) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3X.
func (in *S3X) DeepCopy() *S3X {
	if in == nil {
		return nil
	}
	out := new(S3X)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3X) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3XList) DeepCopyInto(out *S3XList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3X, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3XList.
func (in *S3XList) DeepCopy() *S3XList {
	if in == nil {
		return nil
	}
	out := new(S3XList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3XList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3XSpec) DeepCopyInto(out *S3XSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(v1alpha2.Annotations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3XSpec.
func (in *S3XSpec) DeepCopy() *S3XSpec {
	if in == nil {
		return nil
	}
	out := new(S3XSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SWIFT) DeepCopyInto(out *SWIFT) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SWIFT.
func (in *SWIFT) DeepCopy() *SWIFT {
	if in == nil {
		return nil
	}
	out := new(SWIFT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SWIFT) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SWIFTList) DeepCopyInto(out *SWIFTList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SWIFT, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SWIFTList.
func (in *SWIFTList) DeepCopy() *SWIFTList {
	if in == nil {
		return nil
	}
	out := new(SWIFTList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SWIFTList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SWIFTSpec) DeepCopyInto(out *SWIFTSpec) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ChunkCacheSize = in.ChunkCacheSize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SWIFTSpec.
func (in *SWIFTSpec) DeepCopy() *SWIFTSpec {
	if in == nil {
		return nil
	}
	out := new(SWIFTSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetupNode) DeepCopyInto(out *SetupNode) {
	*out = *in
	out.Ccow = in.Ccow
	in.Ccowd.DeepCopyInto(&out.Ccowd)
	out.Auditd = in.Auditd
	if in.ClusterNodes != nil {
		in, out := &in.ClusterNodes, &out.ClusterNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Rtrd.DeepCopyInto(&out.Rtrd)
	if in.RtrdSlaves != nil {
		in, out := &in.RtrdSlaves, &out.RtrdSlaves
		*out = make([]RTDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Rtlfs.DeepCopyInto(&out.Rtlfs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetupNode.
func (in *SetupNode) DeepCopy() *SetupNode {
	if in == nil {
		return nil
	}
	out := new(SetupNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParametersSpec) DeepCopyInto(out *TargetParametersSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParametersSpec.
func (in *TargetParametersSpec) DeepCopy() *TargetParametersSpec {
	if in == nil {
		return nil
	}
	out := new(TargetParametersSpec)
	in.DeepCopyInto(out)
	return out
}
