// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecscluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference interface {
	cdktn.ComplexObject
	CloudwatchEncryptionEnabled() interface{}
	SetCloudwatchEncryptionEnabled(val interface{})
	CloudwatchEncryptionEnabledInput() interface{}
	CloudwatchLogGroupName() *string
	SetCloudwatchLogGroupName(val *string)
	CloudwatchLogGroupNameInput() *string
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
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3EncryptionEnabled() interface{}
	SetS3EncryptionEnabled(val interface{})
	S3EncryptionEnabledInput() interface{}
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KeyPrefixInput() *string
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
	ResetCloudwatchEncryptionEnabled()
	ResetCloudwatchLogGroupName()
	ResetS3BucketName()
	ResetS3EncryptionEnabled()
	ResetS3KeyPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference
type jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudwatchEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudwatchEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudwatchLogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudwatchLogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3EncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3EncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3EncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3EncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsCluster.EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference_Override(e EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsCluster.EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetCloudwatchEncryptionEnabled(val interface{}) {
	if err := j.validateSetCloudwatchEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetCloudwatchLogGroupName(val *string) {
	if err := j.validateSetCloudwatchLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogGroupName",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3EncryptionEnabled(val interface{}) {
	if err := j.validateSetS3EncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3EncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3KeyPrefix(val *string) {
	if err := j.validateSetS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetCloudwatchEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudwatchEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetCloudwatchLogGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudwatchLogGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		e,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3EncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetS3EncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

