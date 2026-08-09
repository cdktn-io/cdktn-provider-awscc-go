// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDmsEndpointOracleSettingsOutputReference interface {
	cdktn.ComplexObject
	AccessAlternateDirectly() cdktn.IResolvable
	AdditionalArchivedLogDestId() *float64
	AddSupplementalLogging() cdktn.IResolvable
	AllowSelectNestedTables() cdktn.IResolvable
	ArchivedLogDestId() *float64
	ArchivedLogsOnly() cdktn.IResolvable
	AsmPassword() *string
	AsmServer() *string
	AsmUser() *string
	CharLengthSemantics() *string
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
	DirectPathNoLog() cdktn.IResolvable
	DirectPathParallelLoad() cdktn.IResolvable
	EnableHomogenousTablespace() cdktn.IResolvable
	ExtraArchivedLogDestIds() *[]*float64
	FailTasksOnLobTruncation() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDmsEndpointOracleSettings
	SetInternalValue(val *DataAwsccDmsEndpointOracleSettings)
	NumberDatatypeScale() *float64
	OraclePathPrefix() *string
	ParallelAsmReadThreads() *float64
	ReadAheadBlocks() *float64
	ReadTableSpaceName() cdktn.IResolvable
	ReplacePathPrefix() cdktn.IResolvable
	RetryInterval() *float64
	SecretsManagerAccessRoleArn() *string
	SecretsManagerOracleAsmAccessRoleArn() *string
	SecretsManagerOracleAsmSecretId() *string
	SecretsManagerSecretId() *string
	SecurityDbEncryption() *string
	SecurityDbEncryptionName() *string
	SpatialDataOptionToGeoJsonFunctionName() *string
	StandbyDelayTime() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseAlternateFolderForOnline() cdktn.IResolvable
	UseBFile() cdktn.IResolvable
	UseDirectPathFullLoad() cdktn.IResolvable
	UseLogminerReader() cdktn.IResolvable
	UsePathPrefix() *string
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

// The jsii proxy struct for DataAwsccDmsEndpointOracleSettingsOutputReference
type jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AccessAlternateDirectly() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"accessAlternateDirectly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AdditionalArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalArchivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AddSupplementalLogging() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"addSupplementalLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AllowSelectNestedTables() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"allowSelectNestedTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ArchivedLogsOnly() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"archivedLogsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AsmPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AsmServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) AsmUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) CharLengthSemantics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"charLengthSemantics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) DirectPathNoLog() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"directPathNoLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) DirectPathParallelLoad() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"directPathParallelLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) EnableHomogenousTablespace() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableHomogenousTablespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ExtraArchivedLogDestIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"extraArchivedLogDestIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) FailTasksOnLobTruncation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) InternalValue() *DataAwsccDmsEndpointOracleSettings {
	var returns *DataAwsccDmsEndpointOracleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) NumberDatatypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberDatatypeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) OraclePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oraclePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ParallelAsmReadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelAsmReadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ReadAheadBlocks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readAheadBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ReadTableSpaceName() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readTableSpaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ReplacePathPrefix() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"replacePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) RetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecretsManagerOracleAsmSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecurityDbEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SecurityDbEncryptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) SpatialDataOptionToGeoJsonFunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialDataOptionToGeoJsonFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) StandbyDelayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"standbyDelayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) UseAlternateFolderForOnline() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useAlternateFolderForOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) UseBFile() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useBFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) UseDirectPathFullLoad() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useDirectPathFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) UseLogminerReader() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useLogminerReader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) UsePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePathPrefix",
		&returns,
	)
	return returns
}


func NewDataAwsccDmsEndpointOracleSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDmsEndpointOracleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDmsEndpointOracleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsEndpoint.DataAwsccDmsEndpointOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDmsEndpointOracleSettingsOutputReference_Override(d DataAwsccDmsEndpointOracleSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDmsEndpoint.DataAwsccDmsEndpointOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference)SetInternalValue(val *DataAwsccDmsEndpointOracleSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDmsEndpointOracleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

