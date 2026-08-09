// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointMicrosoftSqlServerSettingsOutputReference interface {
	cdktn.ComplexObject
	BcpPacketSize() *float64
	SetBcpPacketSize(val *float64)
	BcpPacketSizeInput() *float64
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
	ControlTablesFileGroup() *string
	SetControlTablesFileGroup(val *string)
	ControlTablesFileGroupInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	ForceLobLookup() interface{}
	SetForceLobLookup(val interface{})
	ForceLobLookupInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	QuerySingleAlwaysOnNode() interface{}
	SetQuerySingleAlwaysOnNode(val interface{})
	QuerySingleAlwaysOnNodeInput() interface{}
	ReadBackupOnly() interface{}
	SetReadBackupOnly(val interface{})
	ReadBackupOnlyInput() interface{}
	SafeguardPolicy() *string
	SetSafeguardPolicy(val *string)
	SafeguardPolicyInput() *string
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlogAccessMode() *string
	SetTlogAccessMode(val *string)
	TlogAccessModeInput() *string
	TrimSpaceInChar() interface{}
	SetTrimSpaceInChar(val interface{})
	TrimSpaceInCharInput() interface{}
	UseBcpFullLoad() interface{}
	SetUseBcpFullLoad(val interface{})
	UseBcpFullLoadInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UseThirdPartyBackupDevice() interface{}
	SetUseThirdPartyBackupDevice(val interface{})
	UseThirdPartyBackupDeviceInput() interface{}
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
	ResetBcpPacketSize()
	ResetControlTablesFileGroup()
	ResetDatabaseName()
	ResetForceLobLookup()
	ResetPassword()
	ResetPort()
	ResetQuerySingleAlwaysOnNode()
	ResetReadBackupOnly()
	ResetSafeguardPolicy()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	ResetServerName()
	ResetTlogAccessMode()
	ResetTrimSpaceInChar()
	ResetUseBcpFullLoad()
	ResetUsername()
	ResetUseThirdPartyBackupDevice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointMicrosoftSqlServerSettingsOutputReference
type jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) BcpPacketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bcpPacketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) BcpPacketSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bcpPacketSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ControlTablesFileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlTablesFileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ControlTablesFileGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlTablesFileGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ForceLobLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceLobLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ForceLobLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceLobLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) QuerySingleAlwaysOnNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"querySingleAlwaysOnNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) QuerySingleAlwaysOnNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"querySingleAlwaysOnNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ReadBackupOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readBackupOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ReadBackupOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readBackupOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SafeguardPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safeguardPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SafeguardPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safeguardPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TlogAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlogAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TlogAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlogAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TrimSpaceInChar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpaceInChar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) TrimSpaceInCharInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpaceInCharInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) UseBcpFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBcpFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) UseBcpFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBcpFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) UseThirdPartyBackupDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useThirdPartyBackupDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) UseThirdPartyBackupDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useThirdPartyBackupDeviceInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointMicrosoftSqlServerSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointMicrosoftSqlServerSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointMicrosoftSqlServerSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointMicrosoftSqlServerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointMicrosoftSqlServerSettingsOutputReference_Override(d DmsEndpointMicrosoftSqlServerSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointMicrosoftSqlServerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetBcpPacketSize(val *float64) {
	if err := j.validateSetBcpPacketSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bcpPacketSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetControlTablesFileGroup(val *string) {
	if err := j.validateSetControlTablesFileGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlTablesFileGroup",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetForceLobLookup(val interface{}) {
	if err := j.validateSetForceLobLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceLobLookup",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetQuerySingleAlwaysOnNode(val interface{}) {
	if err := j.validateSetQuerySingleAlwaysOnNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"querySingleAlwaysOnNode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetReadBackupOnly(val interface{}) {
	if err := j.validateSetReadBackupOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readBackupOnly",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetSafeguardPolicy(val *string) {
	if err := j.validateSetSafeguardPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safeguardPolicy",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetTlogAccessMode(val *string) {
	if err := j.validateSetTlogAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlogAccessMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetTrimSpaceInChar(val interface{}) {
	if err := j.validateSetTrimSpaceInCharParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpaceInChar",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetUseBcpFullLoad(val interface{}) {
	if err := j.validateSetUseBcpFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBcpFullLoad",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference)SetUseThirdPartyBackupDevice(val interface{}) {
	if err := j.validateSetUseThirdPartyBackupDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useThirdPartyBackupDevice",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetBcpPacketSize() {
	_jsii_.InvokeVoid(
		d,
		"resetBcpPacketSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetControlTablesFileGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetControlTablesFileGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetForceLobLookup() {
	_jsii_.InvokeVoid(
		d,
		"resetForceLobLookup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetQuerySingleAlwaysOnNode() {
	_jsii_.InvokeVoid(
		d,
		"resetQuerySingleAlwaysOnNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetReadBackupOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetReadBackupOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetSafeguardPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSafeguardPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetTlogAccessMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTlogAccessMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetTrimSpaceInChar() {
	_jsii_.InvokeVoid(
		d,
		"resetTrimSpaceInChar",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetUseBcpFullLoad() {
	_jsii_.InvokeVoid(
		d,
		"resetUseBcpFullLoad",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ResetUseThirdPartyBackupDevice() {
	_jsii_.InvokeVoid(
		d,
		"resetUseThirdPartyBackupDevice",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointMicrosoftSqlServerSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

