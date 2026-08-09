// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointNeptuneSettingsOutputReference interface {
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
	ErrorRetryDuration() *float64
	SetErrorRetryDuration(val *float64)
	ErrorRetryDurationInput() *float64
	// Experimental.
	Fqn() *string
	IamAuthEnabled() interface{}
	SetIamAuthEnabled(val interface{})
	IamAuthEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	MaxRetryCount() *float64
	SetMaxRetryCount(val *float64)
	MaxRetryCountInput() *float64
	S3BucketFolder() *string
	SetS3BucketFolder(val *string)
	S3BucketFolderInput() *string
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
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
	ResetErrorRetryDuration()
	ResetIamAuthEnabled()
	ResetMaxFileSize()
	ResetMaxRetryCount()
	ResetS3BucketFolder()
	ResetS3BucketName()
	ResetServiceAccessRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointNeptuneSettingsOutputReference
type jsiiProxy_DmsEndpointNeptuneSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ErrorRetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ErrorRetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) IamAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) IamAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) MaxRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) MaxRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) S3BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) S3BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointNeptuneSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointNeptuneSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointNeptuneSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointNeptuneSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointNeptuneSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointNeptuneSettingsOutputReference_Override(d DmsEndpointNeptuneSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointNeptuneSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetErrorRetryDuration(val *float64) {
	if err := j.validateSetErrorRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorRetryDuration",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetIamAuthEnabled(val interface{}) {
	if err := j.validateSetIamAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetMaxRetryCount(val *float64) {
	if err := j.validateSetMaxRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryCount",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetS3BucketFolder(val *string) {
	if err := j.validateSetS3BucketFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketFolder",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetErrorRetryDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetErrorRetryDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetIamAuthEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIamAuthEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetMaxRetryCount() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRetryCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetS3BucketFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetS3BucketFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointNeptuneSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

