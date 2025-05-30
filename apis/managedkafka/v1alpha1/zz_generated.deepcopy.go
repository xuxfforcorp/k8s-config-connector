//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConfig) DeepCopyInto(out *AccessConfig) {
	*out = *in
	if in.NetworkConfigs != nil {
		in, out := &in.NetworkConfigs, &out.NetworkConfigs
		*out = make([]NetworkConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConfig.
func (in *AccessConfig) DeepCopy() *AccessConfig {
	if in == nil {
		return nil
	}
	out := new(AccessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityConfig) DeepCopyInto(out *CapacityConfig) {
	*out = *in
	if in.VcpuCount != nil {
		in, out := &in.VcpuCount, &out.VcpuCount
		*out = new(int64)
		**out = **in
	}
	if in.MemoryBytes != nil {
		in, out := &in.MemoryBytes, &out.MemoryBytes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityConfig.
func (in *CapacityConfig) DeepCopy() *CapacityConfig {
	if in == nil {
		return nil
	}
	out := new(CapacityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIdentity) DeepCopyInto(out *ClusterIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(ClusterParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIdentity.
func (in *ClusterIdentity) DeepCopy() *ClusterIdentity {
	if in == nil {
		return nil
	}
	out := new(ClusterIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParent) DeepCopyInto(out *ClusterParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParent.
func (in *ClusterParent) DeepCopy() *ClusterParent {
	if in == nil {
		return nil
	}
	out := new(ClusterParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRef) DeepCopyInto(out *ClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRef.
func (in *ClusterRef) DeepCopy() *ClusterRef {
	if in == nil {
		return nil
	}
	out := new(ClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupIdentity) DeepCopyInto(out *ConsumerGroupIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(ConsumerGroupParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupIdentity.
func (in *ConsumerGroupIdentity) DeepCopy() *ConsumerGroupIdentity {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupParent) DeepCopyInto(out *ConsumerGroupParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupParent.
func (in *ConsumerGroupParent) DeepCopy() *ConsumerGroupParent {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupRef) DeepCopyInto(out *ConsumerGroupRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupRef.
func (in *ConsumerGroupRef) DeepCopy() *ConsumerGroupRef {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerPartitionMetadata) DeepCopyInto(out *ConsumerPartitionMetadata) {
	*out = *in
	if in.Offset != nil {
		in, out := &in.Offset, &out.Offset
		*out = new(int64)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerPartitionMetadata.
func (in *ConsumerPartitionMetadata) DeepCopy() *ConsumerPartitionMetadata {
	if in == nil {
		return nil
	}
	out := new(ConsumerPartitionMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerTopicMetadata) DeepCopyInto(out *ConsumerTopicMetadata) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]*ConsumerPartitionMetadata, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConsumerPartitionMetadata)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerTopicMetadata.
func (in *ConsumerTopicMetadata) DeepCopy() *ConsumerTopicMetadata {
	if in == nil {
		return nil
	}
	out := new(ConsumerTopicMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpConfig) DeepCopyInto(out *GcpConfig) {
	*out = *in
	if in.AccessConfig != nil {
		in, out := &in.AccessConfig, &out.AccessConfig
		*out = new(AccessConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KmsKeyRef != nil {
		in, out := &in.KmsKeyRef, &out.KmsKeyRef
		*out = new(v1beta1.KMSCryptoKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpConfig.
func (in *GcpConfig) DeepCopy() *GcpConfig {
	if in == nil {
		return nil
	}
	out := new(GcpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaCluster) DeepCopyInto(out *ManagedKafkaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaCluster.
func (in *ManagedKafkaCluster) DeepCopy() *ManagedKafkaCluster {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaClusterList) DeepCopyInto(out *ManagedKafkaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedKafkaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaClusterList.
func (in *ManagedKafkaClusterList) DeepCopy() *ManagedKafkaClusterList {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaClusterObservedState) DeepCopyInto(out *ManagedKafkaClusterObservedState) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaClusterObservedState.
func (in *ManagedKafkaClusterObservedState) DeepCopy() *ManagedKafkaClusterObservedState {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaClusterObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaClusterSpec) DeepCopyInto(out *ManagedKafkaClusterSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.GcpConfig != nil {
		in, out := &in.GcpConfig, &out.GcpConfig
		*out = new(GcpConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CapacityConfig != nil {
		in, out := &in.CapacityConfig, &out.CapacityConfig
		*out = new(CapacityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RebalanceConfig != nil {
		in, out := &in.RebalanceConfig, &out.RebalanceConfig
		*out = new(RebalanceConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaClusterSpec.
func (in *ManagedKafkaClusterSpec) DeepCopy() *ManagedKafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaClusterStatus) DeepCopyInto(out *ManagedKafkaClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(ManagedKafkaClusterObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaClusterStatus.
func (in *ManagedKafkaClusterStatus) DeepCopy() *ManagedKafkaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConsumerGroup) DeepCopyInto(out *ManagedKafkaConsumerGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConsumerGroup.
func (in *ManagedKafkaConsumerGroup) DeepCopy() *ManagedKafkaConsumerGroup {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConsumerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaConsumerGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConsumerGroupList) DeepCopyInto(out *ManagedKafkaConsumerGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedKafkaConsumerGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConsumerGroupList.
func (in *ManagedKafkaConsumerGroupList) DeepCopy() *ManagedKafkaConsumerGroupList {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConsumerGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaConsumerGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConsumerGroupObservedState) DeepCopyInto(out *ManagedKafkaConsumerGroupObservedState) {
	*out = *in
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make(map[string]*ConsumerTopicMetadata, len(*in))
		for key, val := range *in {
			var outVal *ConsumerTopicMetadata
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(ConsumerTopicMetadata)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConsumerGroupObservedState.
func (in *ManagedKafkaConsumerGroupObservedState) DeepCopy() *ManagedKafkaConsumerGroupObservedState {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConsumerGroupObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConsumerGroupSpec) DeepCopyInto(out *ManagedKafkaConsumerGroupSpec) {
	*out = *in
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(Parent)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConsumerGroupSpec.
func (in *ManagedKafkaConsumerGroupSpec) DeepCopy() *ManagedKafkaConsumerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConsumerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConsumerGroupStatus) DeepCopyInto(out *ManagedKafkaConsumerGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(ManagedKafkaConsumerGroupObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConsumerGroupStatus.
func (in *ManagedKafkaConsumerGroupStatus) DeepCopy() *ManagedKafkaConsumerGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConsumerGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaTopic) DeepCopyInto(out *ManagedKafkaTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaTopic.
func (in *ManagedKafkaTopic) DeepCopy() *ManagedKafkaTopic {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaTopicList) DeepCopyInto(out *ManagedKafkaTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedKafkaTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaTopicList.
func (in *ManagedKafkaTopicList) DeepCopy() *ManagedKafkaTopicList {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaTopicSpec) DeepCopyInto(out *ManagedKafkaTopicSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(ClusterRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(int32)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int32)
		**out = **in
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaTopicSpec.
func (in *ManagedKafkaTopicSpec) DeepCopy() *ManagedKafkaTopicSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaTopicStatus) DeepCopyInto(out *ManagedKafkaTopicStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaTopicStatus.
func (in *ManagedKafkaTopicStatus) DeepCopy() *ManagedKafkaTopicStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	if in.SubnetworkRef != nil {
		in, out := &in.SubnetworkRef, &out.SubnetworkRef
		*out = new(v1beta1.ComputeSubnetworkRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(ClusterRef)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebalanceConfig) DeepCopyInto(out *RebalanceConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebalanceConfig.
func (in *RebalanceConfig) DeepCopy() *RebalanceConfig {
	if in == nil {
		return nil
	}
	out := new(RebalanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicIdentity) DeepCopyInto(out *TopicIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(TopicParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicIdentity.
func (in *TopicIdentity) DeepCopy() *TopicIdentity {
	if in == nil {
		return nil
	}
	out := new(TopicIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicParent) DeepCopyInto(out *TopicParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicParent.
func (in *TopicParent) DeepCopy() *TopicParent {
	if in == nil {
		return nil
	}
	out := new(TopicParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicRef) DeepCopyInto(out *TopicRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicRef.
func (in *TopicRef) DeepCopy() *TopicRef {
	if in == nil {
		return nil
	}
	out := new(TopicRef)
	in.DeepCopyInto(out)
	return out
}
