// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference interface {
	cdktn.ComplexObject
	AccessPointId() *string
	SetAccessPointId(val *string)
	AccessPointIdInput() *string
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
	Iam() *string
	SetIam(val *string)
	IamInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetAccessPointId()
	ResetIam()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference
type jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) AccessPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) AccessPointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) Iam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) IamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference_Override(b BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetAccessPointId(val *string) {
	if err := j.validateSetAccessPointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPointId",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetIam(val *string) {
	if err := j.validateSetIamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iam",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ResetAccessPointId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccessPointId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ResetIam() {
	_jsii_.InvokeVoid(
		b,
		"resetIam",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationAuthorizationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

