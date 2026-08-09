// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3storagelens/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference interface {
	cdktn.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	Encryption() S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionOutputReference
	EncryptionInput() interface{}
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputSchemaVersion() *string
	SetOutputSchemaVersion(val *string)
	OutputSchemaVersionInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	PutEncryption(value *S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryption)
	ResetAccountId()
	ResetArn()
	ResetEncryption()
	ResetFormat()
	ResetOutputSchemaVersion()
	ResetPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference
type jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Encryption() S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionOutputReference {
	var returns S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) OutputSchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) OutputSchemaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference_Override(s S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetOutputSchemaVersion(val *string) {
	if err := j.validateSetOutputSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) PutEncryption(value *S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryption) {
	if err := s.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEncryption",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetArn() {
	_jsii_.InvokeVoid(
		s,
		"resetArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetOutputSchemaVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputSchemaVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

