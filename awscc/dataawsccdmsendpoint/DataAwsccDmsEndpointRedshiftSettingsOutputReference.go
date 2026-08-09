// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDmsEndpointRedshiftSettingsOutputReference interface {
	cdktn.ComplexObject
	AcceptAnyDate() cdktn.IResolvable
	AfterConnectScript() *string
	BucketFolder() *string
	BucketName() *string
	CaseSensitiveNames() cdktn.IResolvable
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
	CompUpdate() cdktn.IResolvable
	ConnectionTimeout() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DateFormat() *string
	EmptyAsNull() cdktn.IResolvable
	EncryptionMode() *string
	ExplicitIds() cdktn.IResolvable
	FileTransferUploadStreams() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDmsEndpointRedshiftSettings
	SetInternalValue(val *DataAwsccDmsEndpointRedshiftSettings)
	LoadTimeout() *float64
	MapBooleanAsBoolean() cdktn.IResolvable
	MaxFileSize() *float64
	RemoveQuotes() cdktn.IResolvable
	ReplaceChars() *string
	ReplaceInvalidChars() *string
	SecretsManagerAccessRoleArn() *string
	SecretsManagerSecretId() *string
	ServerSideEncryptionKmsKeyId() *string
	ServiceAccessRoleArn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeFormat() *string
	TrimBlanks() cdktn.IResolvable
	TruncateColumns() cdktn.IResolvable
	WriteBufferSize() *float64
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

// The jsii proxy struct for DataAwsccDmsEndpointRedshiftSettingsOutputReference
type jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) AcceptAnyDate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"acceptAnyDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) CaseSensitiveNames() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"caseSensitiveNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) CompUpdate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"compUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) EmptyAsNull() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"emptyAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ExplicitIds() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"explicitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) FileTransferUploadStreams() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileTransferUploadStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) InternalValue() *DataAwsccDmsEndpointRedshiftSettings {
	var returns *DataAwsccDmsEndpointRedshiftSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) LoadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) MapBooleanAsBoolean() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) RemoveQuotes() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"removeQuotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ReplaceChars() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ReplaceInvalidChars() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) TrimBlanks() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"trimBlanks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) TruncateColumns() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"truncateColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) WriteBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBufferSize",
		&returns,
	)
	return returns
}


func NewDataAwsccDmsEndpointRedshiftSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDmsEndpointRedshiftSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDmsEndpointRedshiftSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsEndpoint.DataAwsccDmsEndpointRedshiftSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDmsEndpointRedshiftSettingsOutputReference_Override(d DataAwsccDmsEndpointRedshiftSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsEndpoint.DataAwsccDmsEndpointRedshiftSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference)SetInternalValue(val *DataAwsccDmsEndpointRedshiftSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDmsEndpointRedshiftSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

