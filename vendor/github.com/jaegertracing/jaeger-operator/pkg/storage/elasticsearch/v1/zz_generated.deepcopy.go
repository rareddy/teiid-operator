// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// TODO remove this file, it's temporary copied from es-operator due to old SDK dependency
//   https://github.com/jaegertracing/jaeger-operator/issues/206

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elasticsearch) DeepCopyInto(out *Elasticsearch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elasticsearch.
func (in *Elasticsearch) DeepCopy() *Elasticsearch {
	if in == nil {
		return nil
	}
	out := new(Elasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elasticsearch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchList) DeepCopyInto(out *ElasticsearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elasticsearch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchList.
func (in *ElasticsearchList) DeepCopy() *ElasticsearchList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchNode) DeepCopyInto(out *ElasticsearchNode) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ElasticsearchNodeRole, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Storage.DeepCopyInto(&out.Storage)
	if in.GenUUID != nil {
		in, out := &in.GenUUID, &out.GenUUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchNode.
func (in *ElasticsearchNode) DeepCopy() *ElasticsearchNode {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchNodeSpec) DeepCopyInto(out *ElasticsearchNodeSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchNodeSpec.
func (in *ElasticsearchNodeSpec) DeepCopy() *ElasticsearchNodeSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchNodeStatus) DeepCopyInto(out *ElasticsearchNodeStatus) {
	*out = *in
	out.UpgradeStatus = in.UpgradeStatus
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ElasticsearchNodeRole, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchNodeStatus.
func (in *ElasticsearchNodeStatus) DeepCopy() *ElasticsearchNodeStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchNodeUpgradeStatus) DeepCopyInto(out *ElasticsearchNodeUpgradeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchNodeUpgradeStatus.
func (in *ElasticsearchNodeUpgradeStatus) DeepCopy() *ElasticsearchNodeUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchNodeUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ElasticsearchNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchStatus) DeepCopyInto(out *ElasticsearchStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ElasticsearchNodeStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[ElasticsearchNodeRole]PodStateMap, len(*in))
		for key, val := range *in {
			var outVal map[PodStateType][]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PodStateMap, len(*in))
				for key, val := range *in {
					var outVal []string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = make([]string, len(*in))
						copy(*out, *in)
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchStatus.
func (in *ElasticsearchStatus) DeepCopy() *ElasticsearchStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchStorageSpec) DeepCopyInto(out *ElasticsearchStorageSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchStorageSpec.
func (in *ElasticsearchStorageSpec) DeepCopy() *ElasticsearchStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PodStateMap) DeepCopyInto(out *PodStateMap) {
	{
		in := &in
		*out = make(PodStateMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStateMap.
func (in PodStateMap) DeepCopy() PodStateMap {
	if in == nil {
		return nil
	}
	out := new(PodStateMap)
	in.DeepCopyInto(out)
	return *out
}
