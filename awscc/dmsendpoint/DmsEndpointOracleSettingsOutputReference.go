// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointOracleSettingsOutputReference interface {
	cdktn.ComplexObject
	AccessAlternateDirectly() interface{}
	SetAccessAlternateDirectly(val interface{})
	AccessAlternateDirectlyInput() interface{}
	AdditionalArchivedLogDestId() *float64
	SetAdditionalArchivedLogDestId(val *float64)
	AdditionalArchivedLogDestIdInput() *float64
	AddSupplementalLogging() interface{}
	SetAddSupplementalLogging(val interface{})
	AddSupplementalLoggingInput() interface{}
	AllowSelectNestedTables() interface{}
	SetAllowSelectNestedTables(val interface{})
	AllowSelectNestedTablesInput() interface{}
	ArchivedLogDestId() *float64
	SetArchivedLogDestId(val *float64)
	ArchivedLogDestIdInput() *float64
	ArchivedLogsOnly() interface{}
	SetArchivedLogsOnly(val interface{})
	ArchivedLogsOnlyInput() interface{}
	AsmPassword() *string
	SetAsmPassword(val *string)
	AsmPasswordInput() *string
	AsmServer() *string
	SetAsmServer(val *string)
	AsmServerInput() *string
	AsmUser() *string
	SetAsmUser(val *string)
	AsmUserInput() *string
	CharLengthSemantics() *string
	SetCharLengthSemantics(val *string)
	CharLengthSemanticsInput() *string
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
	DirectPathNoLog() interface{}
	SetDirectPathNoLog(val interface{})
	DirectPathNoLogInput() interface{}
	DirectPathParallelLoad() interface{}
	SetDirectPathParallelLoad(val interface{})
	DirectPathParallelLoadInput() interface{}
	EnableHomogenousTablespace() interface{}
	SetEnableHomogenousTablespace(val interface{})
	EnableHomogenousTablespaceInput() interface{}
	ExtraArchivedLogDestIds() *[]*float64
	SetExtraArchivedLogDestIds(val *[]*float64)
	ExtraArchivedLogDestIdsInput() *[]*float64
	FailTasksOnLobTruncation() interface{}
	SetFailTasksOnLobTruncation(val interface{})
	FailTasksOnLobTruncationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NumberDatatypeScale() *float64
	SetNumberDatatypeScale(val *float64)
	NumberDatatypeScaleInput() *float64
	OraclePathPrefix() *string
	SetOraclePathPrefix(val *string)
	OraclePathPrefixInput() *string
	ParallelAsmReadThreads() *float64
	SetParallelAsmReadThreads(val *float64)
	ParallelAsmReadThreadsInput() *float64
	ReadAheadBlocks() *float64
	SetReadAheadBlocks(val *float64)
	ReadAheadBlocksInput() *float64
	ReadTableSpaceName() interface{}
	SetReadTableSpaceName(val interface{})
	ReadTableSpaceNameInput() interface{}
	ReplacePathPrefix() interface{}
	SetReplacePathPrefix(val interface{})
	ReplacePathPrefixInput() interface{}
	RetryInterval() *float64
	SetRetryInterval(val *float64)
	RetryIntervalInput() *float64
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerOracleAsmAccessRoleArn() *string
	SetSecretsManagerOracleAsmAccessRoleArn(val *string)
	SecretsManagerOracleAsmAccessRoleArnInput() *string
	SecretsManagerOracleAsmSecretId() *string
	SetSecretsManagerOracleAsmSecretId(val *string)
	SecretsManagerOracleAsmSecretIdInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	SecurityDbEncryption() *string
	SetSecurityDbEncryption(val *string)
	SecurityDbEncryptionInput() *string
	SecurityDbEncryptionName() *string
	SetSecurityDbEncryptionName(val *string)
	SecurityDbEncryptionNameInput() *string
	SpatialDataOptionToGeoJsonFunctionName() *string
	SetSpatialDataOptionToGeoJsonFunctionName(val *string)
	SpatialDataOptionToGeoJsonFunctionNameInput() *string
	StandbyDelayTime() *float64
	SetStandbyDelayTime(val *float64)
	StandbyDelayTimeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseAlternateFolderForOnline() interface{}
	SetUseAlternateFolderForOnline(val interface{})
	UseAlternateFolderForOnlineInput() interface{}
	UseBFile() interface{}
	SetUseBFile(val interface{})
	UseBFileInput() interface{}
	UseDirectPathFullLoad() interface{}
	SetUseDirectPathFullLoad(val interface{})
	UseDirectPathFullLoadInput() interface{}
	UseLogminerReader() interface{}
	SetUseLogminerReader(val interface{})
	UseLogminerReaderInput() interface{}
	UsePathPrefix() *string
	SetUsePathPrefix(val *string)
	UsePathPrefixInput() *string
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
	ResetAccessAlternateDirectly()
	ResetAdditionalArchivedLogDestId()
	ResetAddSupplementalLogging()
	ResetAllowSelectNestedTables()
	ResetArchivedLogDestId()
	ResetArchivedLogsOnly()
	ResetAsmPassword()
	ResetAsmServer()
	ResetAsmUser()
	ResetCharLengthSemantics()
	ResetDirectPathNoLog()
	ResetDirectPathParallelLoad()
	ResetEnableHomogenousTablespace()
	ResetExtraArchivedLogDestIds()
	ResetFailTasksOnLobTruncation()
	ResetNumberDatatypeScale()
	ResetOraclePathPrefix()
	ResetParallelAsmReadThreads()
	ResetReadAheadBlocks()
	ResetReadTableSpaceName()
	ResetReplacePathPrefix()
	ResetRetryInterval()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerOracleAsmAccessRoleArn()
	ResetSecretsManagerOracleAsmSecretId()
	ResetSecretsManagerSecretId()
	ResetSecurityDbEncryption()
	ResetSecurityDbEncryptionName()
	ResetSpatialDataOptionToGeoJsonFunctionName()
	ResetStandbyDelayTime()
	ResetUseAlternateFolderForOnline()
	ResetUseBFile()
	ResetUseDirectPathFullLoad()
	ResetUseLogminerReader()
	ResetUsePathPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointOracleSettingsOutputReference
type jsiiProxy_DmsEndpointOracleSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AccessAlternateDirectly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAlternateDirectly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AccessAlternateDirectlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAlternateDirectlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AdditionalArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalArchivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AdditionalArchivedLogDestIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalArchivedLogDestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AddSupplementalLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addSupplementalLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AddSupplementalLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addSupplementalLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AllowSelectNestedTables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelectNestedTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AllowSelectNestedTablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelectNestedTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ArchivedLogDestIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archivedLogDestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ArchivedLogsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedLogsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ArchivedLogsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedLogsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) AsmUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) CharLengthSemantics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"charLengthSemantics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) CharLengthSemanticsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"charLengthSemanticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) DirectPathNoLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathNoLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) DirectPathNoLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathNoLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) DirectPathParallelLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathParallelLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) DirectPathParallelLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathParallelLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) EnableHomogenousTablespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHomogenousTablespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) EnableHomogenousTablespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHomogenousTablespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ExtraArchivedLogDestIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"extraArchivedLogDestIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ExtraArchivedLogDestIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"extraArchivedLogDestIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) FailTasksOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) FailTasksOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) NumberDatatypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberDatatypeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) NumberDatatypeScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberDatatypeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) OraclePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oraclePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) OraclePathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oraclePathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ParallelAsmReadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelAsmReadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ParallelAsmReadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelAsmReadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReadAheadBlocks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readAheadBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReadAheadBlocksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readAheadBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReadTableSpaceName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readTableSpaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReadTableSpaceNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readTableSpaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReplacePathPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replacePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ReplacePathPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replacePathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) RetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) RetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecurityDbEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecurityDbEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecurityDbEncryptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SecurityDbEncryptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SpatialDataOptionToGeoJsonFunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialDataOptionToGeoJsonFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) SpatialDataOptionToGeoJsonFunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialDataOptionToGeoJsonFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) StandbyDelayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"standbyDelayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) StandbyDelayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"standbyDelayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseAlternateFolderForOnline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAlternateFolderForOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseAlternateFolderForOnlineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAlternateFolderForOnlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseBFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseBFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseDirectPathFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDirectPathFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseDirectPathFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDirectPathFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseLogminerReader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogminerReader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UseLogminerReaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogminerReaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UsePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference) UsePathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePathPrefixInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointOracleSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointOracleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointOracleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointOracleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointOracleSettingsOutputReference_Override(d DmsEndpointOracleSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAccessAlternateDirectly(val interface{}) {
	if err := j.validateSetAccessAlternateDirectlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessAlternateDirectly",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAdditionalArchivedLogDestId(val *float64) {
	if err := j.validateSetAdditionalArchivedLogDestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalArchivedLogDestId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAddSupplementalLogging(val interface{}) {
	if err := j.validateSetAddSupplementalLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addSupplementalLogging",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAllowSelectNestedTables(val interface{}) {
	if err := j.validateSetAllowSelectNestedTablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSelectNestedTables",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetArchivedLogDestId(val *float64) {
	if err := j.validateSetArchivedLogDestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivedLogDestId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetArchivedLogsOnly(val interface{}) {
	if err := j.validateSetArchivedLogsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivedLogsOnly",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAsmPassword(val *string) {
	if err := j.validateSetAsmPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAsmServer(val *string) {
	if err := j.validateSetAsmServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmServer",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetAsmUser(val *string) {
	if err := j.validateSetAsmUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmUser",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetCharLengthSemantics(val *string) {
	if err := j.validateSetCharLengthSemanticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"charLengthSemantics",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetDirectPathNoLog(val interface{}) {
	if err := j.validateSetDirectPathNoLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPathNoLog",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetDirectPathParallelLoad(val interface{}) {
	if err := j.validateSetDirectPathParallelLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPathParallelLoad",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetEnableHomogenousTablespace(val interface{}) {
	if err := j.validateSetEnableHomogenousTablespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHomogenousTablespace",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetExtraArchivedLogDestIds(val *[]*float64) {
	if err := j.validateSetExtraArchivedLogDestIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraArchivedLogDestIds",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetFailTasksOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTasksOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTasksOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetNumberDatatypeScale(val *float64) {
	if err := j.validateSetNumberDatatypeScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberDatatypeScale",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetOraclePathPrefix(val *string) {
	if err := j.validateSetOraclePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oraclePathPrefix",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetParallelAsmReadThreads(val *float64) {
	if err := j.validateSetParallelAsmReadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelAsmReadThreads",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetReadAheadBlocks(val *float64) {
	if err := j.validateSetReadAheadBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAheadBlocks",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetReadTableSpaceName(val interface{}) {
	if err := j.validateSetReadTableSpaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readTableSpaceName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetReplacePathPrefix(val interface{}) {
	if err := j.validateSetReplacePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacePathPrefix",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetRetryInterval(val *float64) {
	if err := j.validateSetRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryInterval",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecretsManagerOracleAsmAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecretsManagerOracleAsmSecretId(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecurityDbEncryption(val *string) {
	if err := j.validateSetSecurityDbEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDbEncryption",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSecurityDbEncryptionName(val *string) {
	if err := j.validateSetSecurityDbEncryptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDbEncryptionName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetSpatialDataOptionToGeoJsonFunctionName(val *string) {
	if err := j.validateSetSpatialDataOptionToGeoJsonFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spatialDataOptionToGeoJsonFunctionName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetStandbyDelayTime(val *float64) {
	if err := j.validateSetStandbyDelayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyDelayTime",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetUseAlternateFolderForOnline(val interface{}) {
	if err := j.validateSetUseAlternateFolderForOnlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAlternateFolderForOnline",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetUseBFile(val interface{}) {
	if err := j.validateSetUseBFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBFile",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetUseDirectPathFullLoad(val interface{}) {
	if err := j.validateSetUseDirectPathFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDirectPathFullLoad",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetUseLogminerReader(val interface{}) {
	if err := j.validateSetUseLogminerReaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLogminerReader",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointOracleSettingsOutputReference)SetUsePathPrefix(val *string) {
	if err := j.validateSetUsePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePathPrefix",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAccessAlternateDirectly() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessAlternateDirectly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAdditionalArchivedLogDestId() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalArchivedLogDestId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAddSupplementalLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetAddSupplementalLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAllowSelectNestedTables() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowSelectNestedTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetArchivedLogDestId() {
	_jsii_.InvokeVoid(
		d,
		"resetArchivedLogDestId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetArchivedLogsOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetArchivedLogsOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAsmPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetAsmPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAsmServer() {
	_jsii_.InvokeVoid(
		d,
		"resetAsmServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetAsmUser() {
	_jsii_.InvokeVoid(
		d,
		"resetAsmUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetCharLengthSemantics() {
	_jsii_.InvokeVoid(
		d,
		"resetCharLengthSemantics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetDirectPathNoLog() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectPathNoLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetDirectPathParallelLoad() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectPathParallelLoad",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetEnableHomogenousTablespace() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableHomogenousTablespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetExtraArchivedLogDestIds() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraArchivedLogDestIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetFailTasksOnLobTruncation() {
	_jsii_.InvokeVoid(
		d,
		"resetFailTasksOnLobTruncation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetNumberDatatypeScale() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberDatatypeScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetOraclePathPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetOraclePathPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetParallelAsmReadThreads() {
	_jsii_.InvokeVoid(
		d,
		"resetParallelAsmReadThreads",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetReadAheadBlocks() {
	_jsii_.InvokeVoid(
		d,
		"resetReadAheadBlocks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetReadTableSpaceName() {
	_jsii_.InvokeVoid(
		d,
		"resetReadTableSpaceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetReplacePathPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetReplacePathPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetRetryInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecretsManagerOracleAsmAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerOracleAsmAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecretsManagerOracleAsmSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerOracleAsmSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecurityDbEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityDbEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSecurityDbEncryptionName() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityDbEncryptionName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetSpatialDataOptionToGeoJsonFunctionName() {
	_jsii_.InvokeVoid(
		d,
		"resetSpatialDataOptionToGeoJsonFunctionName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetStandbyDelayTime() {
	_jsii_.InvokeVoid(
		d,
		"resetStandbyDelayTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetUseAlternateFolderForOnline() {
	_jsii_.InvokeVoid(
		d,
		"resetUseAlternateFolderForOnline",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetUseBFile() {
	_jsii_.InvokeVoid(
		d,
		"resetUseBFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetUseDirectPathFullLoad() {
	_jsii_.InvokeVoid(
		d,
		"resetUseDirectPathFullLoad",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetUseLogminerReader() {
	_jsii_.InvokeVoid(
		d,
		"resetUseLogminerReader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ResetUsePathPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetUsePathPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointOracleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

