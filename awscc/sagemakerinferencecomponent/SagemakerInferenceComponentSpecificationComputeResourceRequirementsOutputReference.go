// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerinferencecomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxMemoryRequiredInMb() *float64
	SetMaxMemoryRequiredInMb(val *float64)
	MaxMemoryRequiredInMbInput() *float64
	MinMemoryRequiredInMb() *float64
	SetMinMemoryRequiredInMb(val *float64)
	MinMemoryRequiredInMbInput() *float64
	NumberOfAcceleratorDevicesRequired() *float64
	SetNumberOfAcceleratorDevicesRequired(val *float64)
	NumberOfAcceleratorDevicesRequiredInput() *float64
	NumberOfCpuCoresRequired() *float64
	SetNumberOfCpuCoresRequired(val *float64)
	NumberOfCpuCoresRequiredInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetMaxMemoryRequiredInMb()
	ResetMinMemoryRequiredInMb()
	ResetNumberOfAcceleratorDevicesRequired()
	ResetNumberOfCpuCoresRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference
type jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) MaxMemoryRequiredInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryRequiredInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) MaxMemoryRequiredInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryRequiredInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) MinMemoryRequiredInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryRequiredInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) MinMemoryRequiredInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryRequiredInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) NumberOfAcceleratorDevicesRequired() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfAcceleratorDevicesRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) NumberOfAcceleratorDevicesRequiredInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfAcceleratorDevicesRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) NumberOfCpuCoresRequired() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCpuCoresRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) NumberOfCpuCoresRequiredInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCpuCoresRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference_Override(s SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetMaxMemoryRequiredInMb(val *float64) {
	if err := j.validateSetMaxMemoryRequiredInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemoryRequiredInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetMinMemoryRequiredInMb(val *float64) {
	if err := j.validateSetMinMemoryRequiredInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemoryRequiredInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetNumberOfAcceleratorDevicesRequired(val *float64) {
	if err := j.validateSetNumberOfAcceleratorDevicesRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfAcceleratorDevicesRequired",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetNumberOfCpuCoresRequired(val *float64) {
	if err := j.validateSetNumberOfCpuCoresRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfCpuCoresRequired",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ResetMaxMemoryRequiredInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxMemoryRequiredInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ResetMinMemoryRequiredInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMinMemoryRequiredInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ResetNumberOfAcceleratorDevicesRequired() {
	_jsii_.InvokeVoid(
		s,
		"resetNumberOfAcceleratorDevicesRequired",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ResetNumberOfCpuCoresRequired() {
	_jsii_.InvokeVoid(
		s,
		"resetNumberOfCpuCoresRequired",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

