// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mgnnetworkmigrationdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference interface {
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
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3BucketOwner() *string
	SetS3BucketOwner(val *string)
	S3BucketOwnerInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference
type jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3BucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3BucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnNetworkMigrationDefinition.MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference_Override(m MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnNetworkMigrationDefinition.MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetS3BucketOwner(val *string) {
	if err := j.validateSetS3BucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketOwner",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3ConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

