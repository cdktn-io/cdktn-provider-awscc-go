// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/sagemakercomputequota/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference interface {
	cdktn.ComplexObject
	AcceleratorPartition() SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartitionOutputReference
	AcceleratorPartitionInput() interface{}
	Accelerators() *float64
	SetAccelerators(val *float64)
	AcceleratorsInput() *float64
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemoryInGiB() *float64
	SetMemoryInGiB(val *float64)
	MemoryInGiBInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VCpu() *float64
	SetVCpu(val *float64)
	VCpuInput() *float64
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
	PutAcceleratorPartition(value *SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartition)
	ResetAcceleratorPartition()
	ResetAccelerators()
	ResetCount()
	ResetInstanceType()
	ResetMemoryInGiB()
	ResetVCpu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference
type jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) AcceleratorPartition() SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartitionOutputReference {
	var returns SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartitionOutputReference
	_jsii_.Get(
		j,
		"acceleratorPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) AcceleratorPartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) Accelerators() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) AcceleratorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) MemoryInGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) MemoryInGiBInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInGiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) VCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) VCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vCpuInput",
		&returns,
	)
	return returns
}


func NewSagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerComputeQuota.SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference_Override(s SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerComputeQuota.SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetAccelerators(val *float64) {
	if err := j.validateSetAcceleratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerators",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetMemoryInGiB(val *float64) {
	if err := j.validateSetMemoryInGiBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryInGiB",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference)SetVCpu(val *float64) {
	if err := j.validateSetVCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vCpu",
		val,
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) PutAcceleratorPartition(value *SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartition) {
	if err := s.validatePutAcceleratorPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAcceleratorPartition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetAcceleratorPartition() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorPartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		s,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		s,
		"resetCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetMemoryInGiB() {
	_jsii_.InvokeVoid(
		s,
		"resetMemoryInGiB",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ResetVCpu() {
	_jsii_.InvokeVoid(
		s,
		"resetVCpu",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

