// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package policy

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_Eviction, InType: reflect.TypeOf(&Eviction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudget, InType: reflect.TypeOf(&PodDisruptionBudget{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetList, InType: reflect.TypeOf(&PodDisruptionBudgetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetSpec, InType: reflect.TypeOf(&PodDisruptionBudgetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_policy_PodDisruptionBudgetStatus, InType: reflect.TypeOf(&PodDisruptionBudgetStatus{})},
	)
}

// DeepCopy_policy_Eviction is an autogenerated deepcopy function.
func DeepCopy_policy_Eviction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Eviction)
		out := out.(*Eviction)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.DeleteOptions != nil {
			in, out := &in.DeleteOptions, &out.DeleteOptions
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.DeleteOptions)
			}
		}
		return nil
	}
}

// DeepCopy_policy_PodDisruptionBudget is an autogenerated deepcopy function.
func DeepCopy_policy_PodDisruptionBudget(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudget)
		out := out.(*PodDisruptionBudget)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_policy_PodDisruptionBudgetList is an autogenerated deepcopy function.
func DeepCopy_policy_PodDisruptionBudgetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetList)
		out := out.(*PodDisruptionBudgetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodDisruptionBudget, len(*in))
			for i := range *in {
				if err := DeepCopy_policy_PodDisruptionBudget(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_policy_PodDisruptionBudgetSpec is an autogenerated deepcopy function.
func DeepCopy_policy_PodDisruptionBudgetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetSpec)
		out := out.(*PodDisruptionBudgetSpec)
		*out = *in
		if in.MinAvailable != nil {
			in, out := &in.MinAvailable, &out.MinAvailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_policy_PodDisruptionBudgetStatus is an autogenerated deepcopy function.
func DeepCopy_policy_PodDisruptionBudgetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodDisruptionBudgetStatus)
		out := out.(*PodDisruptionBudgetStatus)
		*out = *in
		if in.DisruptedPods != nil {
			in, out := &in.DisruptedPods, &out.DisruptedPods
			*out = make(map[string]v1.Time, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
		return nil
	}
}
