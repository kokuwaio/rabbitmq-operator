// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqBinding) DeepCopyInto(out *RabbitmqBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqBinding.
func (in *RabbitmqBinding) DeepCopy() *RabbitmqBinding {
	if in == nil {
		return nil
	}
	out := new(RabbitmqBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqBindingList) DeepCopyInto(out *RabbitmqBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqBindingList.
func (in *RabbitmqBindingList) DeepCopy() *RabbitmqBindingList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqBindingSpec) DeepCopyInto(out *RabbitmqBindingSpec) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ClusterRef = in.ClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqBindingSpec.
func (in *RabbitmqBindingSpec) DeepCopy() *RabbitmqBindingSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqBindingStatus) DeepCopyInto(out *RabbitmqBindingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqBindingStatus.
func (in *RabbitmqBindingStatus) DeepCopy() *RabbitmqBindingStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqCluster) DeepCopyInto(out *RabbitmqCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqCluster.
func (in *RabbitmqCluster) DeepCopy() *RabbitmqCluster {
	if in == nil {
		return nil
	}
	out := new(RabbitmqCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterList) DeepCopyInto(out *RabbitmqClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterList.
func (in *RabbitmqClusterList) DeepCopy() *RabbitmqClusterList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterRef) DeepCopyInto(out *RabbitmqClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterRef.
func (in *RabbitmqClusterRef) DeepCopy() *RabbitmqClusterRef {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterSpec) DeepCopyInto(out *RabbitmqClusterSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterSpec.
func (in *RabbitmqClusterSpec) DeepCopy() *RabbitmqClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterStatus) DeepCopyInto(out *RabbitmqClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterStatus.
func (in *RabbitmqClusterStatus) DeepCopy() *RabbitmqClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqExchange) DeepCopyInto(out *RabbitmqExchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqExchange.
func (in *RabbitmqExchange) DeepCopy() *RabbitmqExchange {
	if in == nil {
		return nil
	}
	out := new(RabbitmqExchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqExchange) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqExchangeList) DeepCopyInto(out *RabbitmqExchangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqExchange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqExchangeList.
func (in *RabbitmqExchangeList) DeepCopy() *RabbitmqExchangeList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqExchangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqExchangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqExchangeSetting) DeepCopyInto(out *RabbitmqExchangeSetting) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqExchangeSetting.
func (in *RabbitmqExchangeSetting) DeepCopy() *RabbitmqExchangeSetting {
	if in == nil {
		return nil
	}
	out := new(RabbitmqExchangeSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqExchangeSpec) DeepCopyInto(out *RabbitmqExchangeSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	in.Settings.DeepCopyInto(&out.Settings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqExchangeSpec.
func (in *RabbitmqExchangeSpec) DeepCopy() *RabbitmqExchangeSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqExchangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqExchangeStatus) DeepCopyInto(out *RabbitmqExchangeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqExchangeStatus.
func (in *RabbitmqExchangeStatus) DeepCopy() *RabbitmqExchangeStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqExchangeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqPermisson) DeepCopyInto(out *RabbitmqPermisson) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqPermisson.
func (in *RabbitmqPermisson) DeepCopy() *RabbitmqPermisson {
	if in == nil {
		return nil
	}
	out := new(RabbitmqPermisson)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqPermisson) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqPermissonList) DeepCopyInto(out *RabbitmqPermissonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqPermisson, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqPermissonList.
func (in *RabbitmqPermissonList) DeepCopy() *RabbitmqPermissonList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqPermissonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqPermissonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqPermissonSpec) DeepCopyInto(out *RabbitmqPermissonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqPermissonSpec.
func (in *RabbitmqPermissonSpec) DeepCopy() *RabbitmqPermissonSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqPermissonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqPermissonStatus) DeepCopyInto(out *RabbitmqPermissonStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqPermissonStatus.
func (in *RabbitmqPermissonStatus) DeepCopy() *RabbitmqPermissonStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqPermissonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqQueue) DeepCopyInto(out *RabbitmqQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqQueue.
func (in *RabbitmqQueue) DeepCopy() *RabbitmqQueue {
	if in == nil {
		return nil
	}
	out := new(RabbitmqQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqQueueList) DeepCopyInto(out *RabbitmqQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqQueueList.
func (in *RabbitmqQueueList) DeepCopy() *RabbitmqQueueList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqQueueSetting) DeepCopyInto(out *RabbitmqQueueSetting) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqQueueSetting.
func (in *RabbitmqQueueSetting) DeepCopy() *RabbitmqQueueSetting {
	if in == nil {
		return nil
	}
	out := new(RabbitmqQueueSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqQueueSpec) DeepCopyInto(out *RabbitmqQueueSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	in.Settings.DeepCopyInto(&out.Settings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqQueueSpec.
func (in *RabbitmqQueueSpec) DeepCopy() *RabbitmqQueueSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqQueueStatus) DeepCopyInto(out *RabbitmqQueueStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqQueueStatus.
func (in *RabbitmqQueueStatus) DeepCopy() *RabbitmqQueueStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqUser) DeepCopyInto(out *RabbitmqUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqUser.
func (in *RabbitmqUser) DeepCopy() *RabbitmqUser {
	if in == nil {
		return nil
	}
	out := new(RabbitmqUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqUserList) DeepCopyInto(out *RabbitmqUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqUserList.
func (in *RabbitmqUserList) DeepCopy() *RabbitmqUserList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqUserSpec) DeepCopyInto(out *RabbitmqUserSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqUserSpec.
func (in *RabbitmqUserSpec) DeepCopy() *RabbitmqUserSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqUserStatus) DeepCopyInto(out *RabbitmqUserStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqUserStatus.
func (in *RabbitmqUserStatus) DeepCopy() *RabbitmqUserStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}
