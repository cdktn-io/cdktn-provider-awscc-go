// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointRedshiftSettingsOutputReference interface {
	cdktn.ComplexObject
	AcceptAnyDate() interface{}
	SetAcceptAnyDate(val interface{})
	AcceptAnyDateInput() interface{}
	AfterConnectScript() *string
	SetAfterConnectScript(val *string)
	AfterConnectScriptInput() *string
	BucketFolder() *string
	SetBucketFolder(val *string)
	BucketFolderInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	CaseSensitiveNames() interface{}
	SetCaseSensitiveNames(val interface{})
	CaseSensitiveNamesInput() interface{}
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
	CompUpdate() interface{}
	SetCompUpdate(val interface{})
	CompUpdateInput() interface{}
	ConnectionTimeout() *float64
	SetConnectionTimeout(val *float64)
	ConnectionTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DateFormat() *string
	SetDateFormat(val *string)
	DateFormatInput() *string
	EmptyAsNull() interface{}
	SetEmptyAsNull(val interface{})
	EmptyAsNullInput() interface{}
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	EncryptionModeInput() *string
	ExplicitIds() interface{}
	SetExplicitIds(val interface{})
	ExplicitIdsInput() interface{}
	FileTransferUploadStreams() *float64
	SetFileTransferUploadStreams(val *float64)
	FileTransferUploadStreamsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadTimeout() *float64
	SetLoadTimeout(val *float64)
	LoadTimeoutInput() *float64
	MapBooleanAsBoolean() interface{}
	SetMapBooleanAsBoolean(val interface{})
	MapBooleanAsBooleanInput() interface{}
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	RemoveQuotes() interface{}
	SetRemoveQuotes(val interface{})
	RemoveQuotesInput() interface{}
	ReplaceChars() *string
	SetReplaceChars(val *string)
	ReplaceCharsInput() *string
	ReplaceInvalidChars() *string
	SetReplaceInvalidChars(val *string)
	ReplaceInvalidCharsInput() *string
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	ServerSideEncryptionKmsKeyId() *string
	SetServerSideEncryptionKmsKeyId(val *string)
	ServerSideEncryptionKmsKeyIdInput() *string
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
	TimeFormat() *string
	SetTimeFormat(val *string)
	TimeFormatInput() *string
	TrimBlanks() interface{}
	SetTrimBlanks(val interface{})
	TrimBlanksInput() interface{}
	TruncateColumns() interface{}
	SetTruncateColumns(val interface{})
	TruncateColumnsInput() interface{}
	WriteBufferSize() *float64
	SetWriteBufferSize(val *float64)
	WriteBufferSizeInput() *float64
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
	ResetAcceptAnyDate()
	ResetAfterConnectScript()
	ResetBucketFolder()
	ResetBucketName()
	ResetCaseSensitiveNames()
	ResetCompUpdate()
	ResetConnectionTimeout()
	ResetDateFormat()
	ResetEmptyAsNull()
	ResetEncryptionMode()
	ResetExplicitIds()
	ResetFileTransferUploadStreams()
	ResetLoadTimeout()
	ResetMapBooleanAsBoolean()
	ResetMaxFileSize()
	ResetRemoveQuotes()
	ResetReplaceChars()
	ResetReplaceInvalidChars()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	ResetServerSideEncryptionKmsKeyId()
	ResetServiceAccessRoleArn()
	ResetTimeFormat()
	ResetTrimBlanks()
	ResetTruncateColumns()
	ResetWriteBufferSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointRedshiftSettingsOutputReference
type jsiiProxy_DmsEndpointRedshiftSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) AcceptAnyDate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAnyDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) AcceptAnyDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAnyDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) CaseSensitiveNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) CaseSensitiveNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) CompUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) CompUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) EmptyAsNull() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) EmptyAsNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyAsNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ExplicitIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explicitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ExplicitIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explicitIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) FileTransferUploadStreams() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileTransferUploadStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) FileTransferUploadStreamsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileTransferUploadStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) LoadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) LoadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) MapBooleanAsBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) MapBooleanAsBooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBooleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) RemoveQuotes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeQuotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) RemoveQuotesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeQuotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ReplaceChars() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ReplaceCharsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ReplaceInvalidChars() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ReplaceInvalidCharsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ServerSideEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TrimBlanks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimBlanks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TrimBlanksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimBlanksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TruncateColumns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truncateColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) TruncateColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truncateColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) WriteBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) WriteBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBufferSizeInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointRedshiftSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointRedshiftSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointRedshiftSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointRedshiftSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointRedshiftSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointRedshiftSettingsOutputReference_Override(d DmsEndpointRedshiftSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointRedshiftSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetAcceptAnyDate(val interface{}) {
	if err := j.validateSetAcceptAnyDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptAnyDate",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetBucketFolder(val *string) {
	if err := j.validateSetBucketFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketFolder",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetCaseSensitiveNames(val interface{}) {
	if err := j.validateSetCaseSensitiveNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitiveNames",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetCompUpdate(val interface{}) {
	if err := j.validateSetCompUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compUpdate",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetEmptyAsNull(val interface{}) {
	if err := j.validateSetEmptyAsNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyAsNull",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetEncryptionMode(val *string) {
	if err := j.validateSetEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetExplicitIds(val interface{}) {
	if err := j.validateSetExplicitIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitIds",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetFileTransferUploadStreams(val *float64) {
	if err := j.validateSetFileTransferUploadStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileTransferUploadStreams",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetLoadTimeout(val *float64) {
	if err := j.validateSetLoadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetMapBooleanAsBoolean(val interface{}) {
	if err := j.validateSetMapBooleanAsBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBooleanAsBoolean",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetRemoveQuotes(val interface{}) {
	if err := j.validateSetRemoveQuotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeQuotes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetReplaceChars(val *string) {
	if err := j.validateSetReplaceCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceChars",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetReplaceInvalidChars(val *string) {
	if err := j.validateSetReplaceInvalidCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidChars",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetServerSideEncryptionKmsKeyId(val *string) {
	if err := j.validateSetServerSideEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetTimeFormat(val *string) {
	if err := j.validateSetTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetTrimBlanks(val interface{}) {
	if err := j.validateSetTrimBlanksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimBlanks",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetTruncateColumns(val interface{}) {
	if err := j.validateSetTruncateColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truncateColumns",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference)SetWriteBufferSize(val *float64) {
	if err := j.validateSetWriteBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeBufferSize",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetAcceptAnyDate() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptAnyDate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		d,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetBucketFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetCaseSensitiveNames() {
	_jsii_.InvokeVoid(
		d,
		"resetCaseSensitiveNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetCompUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetCompUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetEmptyAsNull() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyAsNull",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetExplicitIds() {
	_jsii_.InvokeVoid(
		d,
		"resetExplicitIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetFileTransferUploadStreams() {
	_jsii_.InvokeVoid(
		d,
		"resetFileTransferUploadStreams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetLoadTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetMapBooleanAsBoolean() {
	_jsii_.InvokeVoid(
		d,
		"resetMapBooleanAsBoolean",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetRemoveQuotes() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoveQuotes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetReplaceChars() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceChars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetReplaceInvalidChars() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceInvalidChars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetServerSideEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetServerSideEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetTimeFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetTrimBlanks() {
	_jsii_.InvokeVoid(
		d,
		"resetTrimBlanks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetTruncateColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetTruncateColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ResetWriteBufferSize() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteBufferSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointRedshiftSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

