// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference interface {
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
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleConfigArn() *string
	SetLifecycleConfigArn(val *string)
	LifecycleConfigArnInput() *string
	SageMakerImageArn() *string
	SetSageMakerImageArn(val *string)
	SageMakerImageArnInput() *string
	SageMakerImageVersionArn() *string
	SetSageMakerImageVersionArn(val *string)
	SageMakerImageVersionArnInput() *string
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
	ResetInstanceType()
	ResetLifecycleConfigArn()
	ResetSageMakerImageArn()
	ResetSageMakerImageVersionArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference
type jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) LifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) LifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) SageMakerImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sageMakerImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) SageMakerImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sageMakerImageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) SageMakerImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sageMakerImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) SageMakerImageVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sageMakerImageVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference_Override(s SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetLifecycleConfigArn(val *string) {
	if err := j.validateSetLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetSageMakerImageArn(val *string) {
	if err := j.validateSetSageMakerImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sageMakerImageArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetSageMakerImageVersionArn(val *string) {
	if err := j.validateSetSageMakerImageVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sageMakerImageVersionArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ResetLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleConfigArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ResetSageMakerImageArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSageMakerImageArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ResetSageMakerImageVersionArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSageMakerImageVersionArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

