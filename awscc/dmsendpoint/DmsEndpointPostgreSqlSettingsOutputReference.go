// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointPostgreSqlSettingsOutputReference interface {
	cdktn.ComplexObject
	AfterConnectScript() *string
	SetAfterConnectScript(val *string)
	AfterConnectScriptInput() *string
	BabelfishDatabaseName() *string
	SetBabelfishDatabaseName(val *string)
	BabelfishDatabaseNameInput() *string
	CaptureDdls() interface{}
	SetCaptureDdls(val interface{})
	CaptureDdlsInput() interface{}
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
	DatabaseMode() *string
	SetDatabaseMode(val *string)
	DatabaseModeInput() *string
	DdlArtifactsSchema() *string
	SetDdlArtifactsSchema(val *string)
	DdlArtifactsSchemaInput() *string
	ExecuteTimeout() *float64
	SetExecuteTimeout(val *float64)
	ExecuteTimeoutInput() *float64
	FailTasksOnLobTruncation() interface{}
	SetFailTasksOnLobTruncation(val interface{})
	FailTasksOnLobTruncationInput() interface{}
	// Experimental.
	Fqn() *string
	HeartbeatEnable() interface{}
	SetHeartbeatEnable(val interface{})
	HeartbeatEnableInput() interface{}
	HeartbeatFrequency() *float64
	SetHeartbeatFrequency(val *float64)
	HeartbeatFrequencyInput() *float64
	HeartbeatSchema() *string
	SetHeartbeatSchema(val *string)
	HeartbeatSchemaInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MapBooleanAsBoolean() interface{}
	SetMapBooleanAsBoolean(val interface{})
	MapBooleanAsBooleanInput() interface{}
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	SlotName() *string
	SetSlotName(val *string)
	SlotNameInput() *string
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
	ResetAfterConnectScript()
	ResetBabelfishDatabaseName()
	ResetCaptureDdls()
	ResetDatabaseMode()
	ResetDdlArtifactsSchema()
	ResetExecuteTimeout()
	ResetFailTasksOnLobTruncation()
	ResetHeartbeatEnable()
	ResetHeartbeatFrequency()
	ResetHeartbeatSchema()
	ResetMapBooleanAsBoolean()
	ResetMaxFileSize()
	ResetPluginName()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	ResetSlotName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointPostgreSqlSettingsOutputReference
type jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) BabelfishDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) BabelfishDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) CaptureDdls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) CaptureDdlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) DatabaseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) DatabaseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) DdlArtifactsSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) DdlArtifactsSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) FailTasksOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) FailTasksOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) HeartbeatSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) MapBooleanAsBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) MapBooleanAsBooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBooleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) SlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointPostgreSqlSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointPostgreSqlSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointPostgreSqlSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointPostgreSqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointPostgreSqlSettingsOutputReference_Override(d DmsEndpointPostgreSqlSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointPostgreSqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetBabelfishDatabaseName(val *string) {
	if err := j.validateSetBabelfishDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"babelfishDatabaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetCaptureDdls(val interface{}) {
	if err := j.validateSetCaptureDdlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captureDdls",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetDatabaseMode(val *string) {
	if err := j.validateSetDatabaseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetDdlArtifactsSchema(val *string) {
	if err := j.validateSetDdlArtifactsSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddlArtifactsSchema",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetFailTasksOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTasksOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTasksOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetHeartbeatEnable(val interface{}) {
	if err := j.validateSetHeartbeatEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatEnable",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetHeartbeatFrequency(val *float64) {
	if err := j.validateSetHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetHeartbeatSchema(val *string) {
	if err := j.validateSetHeartbeatSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatSchema",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetMapBooleanAsBoolean(val interface{}) {
	if err := j.validateSetMapBooleanAsBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBooleanAsBoolean",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetSlotName(val *string) {
	if err := j.validateSetSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		d,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetBabelfishDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetBabelfishDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetCaptureDdls() {
	_jsii_.InvokeVoid(
		d,
		"resetCaptureDdls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetDatabaseMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetDdlArtifactsSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetDdlArtifactsSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetFailTasksOnLobTruncation() {
	_jsii_.InvokeVoid(
		d,
		"resetFailTasksOnLobTruncation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetHeartbeatEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatFrequency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetHeartbeatSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetMapBooleanAsBoolean() {
	_jsii_.InvokeVoid(
		d,
		"resetMapBooleanAsBoolean",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ResetSlotName() {
	_jsii_.InvokeVoid(
		d,
		"resetSlotName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointPostgreSqlSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

