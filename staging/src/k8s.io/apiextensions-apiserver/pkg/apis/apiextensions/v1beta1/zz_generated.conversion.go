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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition,
		Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition,
		Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition,
		Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition,
		Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList,
		Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList,
		Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames,
		Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames,
		Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec,
		Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec,
		Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus,
		Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus,
		Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation,
		Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation,
		Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation,
		Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation,
		Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps,
		Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps,
		Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray,
		Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray,
		Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool,
		Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool,
		Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray,
		Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray,
	)
}

func autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = apiextensions.CustomResourceDefinitionConditionType(in.Type)
	out.Status = apiextensions.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = CustomResourceDefinitionConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apiextensions.CustomResourceDefinition)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]CustomResourceDefinition, 0)
	} else {
		out.Items = *(*[]CustomResourceDefinition)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = apiextensions.ResourceScope(in.Scope)
	if err := Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(&in.Validation, &out.Validation, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = ResourceScope(in.Scope)
	if err := Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(&in.Validation, &out.Validation, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	out.Conditions = *(*[]apiextensions.CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	if in.Conditions == nil {
		out.Conditions = make([]CustomResourceDefinitionCondition, 0)
	} else {
		out.Conditions = *(*[]CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	}
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in, out, s)
}

func autoConvert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in *CustomResourceValidation, out *apiextensions.CustomResourceValidation, s conversion.Scope) error {
	out.OpenAPIV3Schema = (*apiextensions.JSONSchemaProps)(unsafe.Pointer(in.OpenAPIV3Schema))
	return nil
}

// Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in *CustomResourceValidation, out *apiextensions.CustomResourceValidation, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in, out, s)
}

func autoConvert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in *apiextensions.CustomResourceValidation, out *CustomResourceValidation, s conversion.Scope) error {
	out.OpenAPIV3Schema = (*JSONSchemaProps)(unsafe.Pointer(in.OpenAPIV3Schema))
	return nil
}

// Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in *apiextensions.CustomResourceValidation, out *CustomResourceValidation, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in, out, s)
}

func autoConvert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in *ExternalDocumentation, out *apiextensions.ExternalDocumentation, s conversion.Scope) error {
	out.Description = in.Description
	out.URL = in.URL
	return nil
}

// Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation is an autogenerated conversion function.
func Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in *ExternalDocumentation, out *apiextensions.ExternalDocumentation, s conversion.Scope) error {
	return autoConvert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in, out, s)
}

func autoConvert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in *apiextensions.ExternalDocumentation, out *ExternalDocumentation, s conversion.Scope) error {
	out.Description = in.Description
	out.URL = in.URL
	return nil
}

// Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation is an autogenerated conversion function.
func Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in *apiextensions.ExternalDocumentation, out *ExternalDocumentation, s conversion.Scope) error {
	return autoConvert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in *JSONSchemaProps, out *apiextensions.JSONSchemaProps, s conversion.Scope) error {
	out.ID = in.ID
	out.Schema = apiextensions.JSONSchemaURL(in.Schema)
	out.Ref = (*string)(unsafe.Pointer(in.Ref))
	out.Description = in.Description
	out.Type = *(*apiextensions.StringOrArray)(unsafe.Pointer(&in.Type))
	out.Format = in.Format
	out.Title = in.Title
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Default, &out.Default, 0); err != nil {
		return err
	}
	out.Maximum = (*float64)(unsafe.Pointer(in.Maximum))
	out.ExclusiveMaximum = in.ExclusiveMaximum
	out.Minimum = (*float64)(unsafe.Pointer(in.Minimum))
	out.ExclusiveMinimum = in.ExclusiveMinimum
	out.MaxLength = (*int64)(unsafe.Pointer(in.MaxLength))
	out.MinLength = (*int64)(unsafe.Pointer(in.MinLength))
	out.Pattern = in.Pattern
	out.MaxItems = (*int64)(unsafe.Pointer(in.MaxItems))
	out.MinItems = (*int64)(unsafe.Pointer(in.MinItems))
	out.UniqueItems = in.UniqueItems
	out.MultipleOf = (*float64)(unsafe.Pointer(in.MultipleOf))
	out.Enum = *(*[]interface{})(unsafe.Pointer(&in.Enum))
	out.MaxProperties = (*int64)(unsafe.Pointer(in.MaxProperties))
	out.MinProperties = (*int64)(unsafe.Pointer(in.MinProperties))
	out.Required = *(*[]string)(unsafe.Pointer(&in.Required))
	out.Items = (*apiextensions.JSONSchemaPropsOrArray)(unsafe.Pointer(in.Items))
	out.AllOf = *(*[]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.AllOf))
	out.OneOf = *(*[]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.OneOf))
	out.AnyOf = *(*[]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.AnyOf))
	out.Not = (*apiextensions.JSONSchemaProps)(unsafe.Pointer(in.Not))
	out.Properties = *(*map[string]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.Properties))
	out.AdditionalProperties = (*apiextensions.JSONSchemaPropsOrBool)(unsafe.Pointer(in.AdditionalProperties))
	out.PatternProperties = *(*map[string]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.PatternProperties))
	out.Dependencies = *(*apiextensions.JSONSchemaDependencies)(unsafe.Pointer(&in.Dependencies))
	out.AdditionalItems = (*apiextensions.JSONSchemaPropsOrBool)(unsafe.Pointer(in.AdditionalItems))
	out.Definitions = *(*apiextensions.JSONSchemaDefinitions)(unsafe.Pointer(&in.Definitions))
	out.ExternalDocs = (*apiextensions.ExternalDocumentation)(unsafe.Pointer(in.ExternalDocs))
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Example, &out.Example, 0); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in *JSONSchemaProps, out *apiextensions.JSONSchemaProps, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(in *apiextensions.JSONSchemaProps, out *JSONSchemaProps, s conversion.Scope) error {
	out.ID = in.ID
	out.Schema = JSONSchemaURL(in.Schema)
	out.Ref = (*string)(unsafe.Pointer(in.Ref))
	out.Description = in.Description
	out.Type = *(*StringOrArray)(unsafe.Pointer(&in.Type))
	out.Format = in.Format
	out.Title = in.Title
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Default, &out.Default, 0); err != nil {
		return err
	}
	out.Maximum = (*float64)(unsafe.Pointer(in.Maximum))
	out.ExclusiveMaximum = in.ExclusiveMaximum
	out.Minimum = (*float64)(unsafe.Pointer(in.Minimum))
	out.ExclusiveMinimum = in.ExclusiveMinimum
	out.MaxLength = (*int64)(unsafe.Pointer(in.MaxLength))
	out.MinLength = (*int64)(unsafe.Pointer(in.MinLength))
	out.Pattern = in.Pattern
	out.MaxItems = (*int64)(unsafe.Pointer(in.MaxItems))
	out.MinItems = (*int64)(unsafe.Pointer(in.MinItems))
	out.UniqueItems = in.UniqueItems
	out.MultipleOf = (*float64)(unsafe.Pointer(in.MultipleOf))
	out.Enum = *(*[]interface{})(unsafe.Pointer(&in.Enum))
	out.MaxProperties = (*int64)(unsafe.Pointer(in.MaxProperties))
	out.MinProperties = (*int64)(unsafe.Pointer(in.MinProperties))
	out.Required = *(*[]string)(unsafe.Pointer(&in.Required))
	out.Items = (*JSONSchemaPropsOrArray)(unsafe.Pointer(in.Items))
	out.AllOf = *(*[]JSONSchemaProps)(unsafe.Pointer(&in.AllOf))
	out.OneOf = *(*[]JSONSchemaProps)(unsafe.Pointer(&in.OneOf))
	out.AnyOf = *(*[]JSONSchemaProps)(unsafe.Pointer(&in.AnyOf))
	out.Not = (*JSONSchemaProps)(unsafe.Pointer(in.Not))
	out.Properties = *(*map[string]JSONSchemaProps)(unsafe.Pointer(&in.Properties))
	out.AdditionalProperties = (*JSONSchemaPropsOrBool)(unsafe.Pointer(in.AdditionalProperties))
	out.PatternProperties = *(*map[string]JSONSchemaProps)(unsafe.Pointer(&in.PatternProperties))
	out.Dependencies = *(*JSONSchemaDependencies)(unsafe.Pointer(&in.Dependencies))
	out.AdditionalItems = (*JSONSchemaPropsOrBool)(unsafe.Pointer(in.AdditionalItems))
	out.Definitions = *(*JSONSchemaDefinitions)(unsafe.Pointer(&in.Definitions))
	out.ExternalDocs = (*ExternalDocumentation)(unsafe.Pointer(in.ExternalDocs))
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Example, &out.Example, 0); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(in *apiextensions.JSONSchemaProps, out *JSONSchemaProps, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in *JSONSchemaPropsOrArray, out *apiextensions.JSONSchemaPropsOrArray, s conversion.Scope) error {
	out.Schema = (*apiextensions.JSONSchemaProps)(unsafe.Pointer(in.Schema))
	out.JSONSchemas = *(*[]apiextensions.JSONSchemaProps)(unsafe.Pointer(&in.JSONSchemas))
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in *JSONSchemaPropsOrArray, out *apiextensions.JSONSchemaPropsOrArray, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in *apiextensions.JSONSchemaPropsOrArray, out *JSONSchemaPropsOrArray, s conversion.Scope) error {
	out.Schema = (*JSONSchemaProps)(unsafe.Pointer(in.Schema))
	out.JSONSchemas = *(*[]JSONSchemaProps)(unsafe.Pointer(&in.JSONSchemas))
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in *apiextensions.JSONSchemaPropsOrArray, out *JSONSchemaPropsOrArray, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in *JSONSchemaPropsOrBool, out *apiextensions.JSONSchemaPropsOrBool, s conversion.Scope) error {
	out.Allows = in.Allows
	out.Schema = (*apiextensions.JSONSchemaProps)(unsafe.Pointer(in.Schema))
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in *JSONSchemaPropsOrBool, out *apiextensions.JSONSchemaPropsOrBool, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in *apiextensions.JSONSchemaPropsOrBool, out *JSONSchemaPropsOrBool, s conversion.Scope) error {
	out.Allows = in.Allows
	out.Schema = (*JSONSchemaProps)(unsafe.Pointer(in.Schema))
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in *apiextensions.JSONSchemaPropsOrBool, out *JSONSchemaPropsOrBool, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in *JSONSchemaPropsOrStringArray, out *apiextensions.JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	out.Schema = (*apiextensions.JSONSchemaProps)(unsafe.Pointer(in.Schema))
	out.Property = *(*[]string)(unsafe.Pointer(&in.Property))
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in *JSONSchemaPropsOrStringArray, out *apiextensions.JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in *apiextensions.JSONSchemaPropsOrStringArray, out *JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	out.Schema = (*JSONSchemaProps)(unsafe.Pointer(in.Schema))
	out.Property = *(*[]string)(unsafe.Pointer(&in.Property))
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in *apiextensions.JSONSchemaPropsOrStringArray, out *JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in, out, s)
}
