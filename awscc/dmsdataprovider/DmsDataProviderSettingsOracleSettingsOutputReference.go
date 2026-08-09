// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsdataprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsDataProviderSettingsOracleSettingsOutputReference interface {
	cdktn.ComplexObject
	AsmServer() *string
	SetAsmServer(val *string)
	AsmServerInput() *string
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	SecretsManagerOracleAsmAccessRoleArn() *string
	SetSecretsManagerOracleAsmAccessRoleArn(val *string)
	SecretsManagerOracleAsmAccessRoleArnInput() *string
	SecretsManagerOracleAsmSecretId() *string
	SetSecretsManagerOracleAsmSecretId(val *string)
	SecretsManagerOracleAsmSecretIdInput() *string
	SecretsManagerSecurityDbEncryptionAccessRoleArn() *string
	SetSecretsManagerSecurityDbEncryptionAccessRoleArn(val *string)
	SecretsManagerSecurityDbEncryptionAccessRoleArnInput() *string
	SecretsManagerSecurityDbEncryptionSecretId() *string
	SetSecretsManagerSecurityDbEncryptionSecretId(val *string)
	SecretsManagerSecurityDbEncryptionSecretIdInput() *string
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
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
	ResetAsmServer()
	ResetCertificateArn()
	ResetDatabaseName()
	ResetPort()
	ResetSecretsManagerOracleAsmAccessRoleArn()
	ResetSecretsManagerOracleAsmSecretId()
	ResetSecretsManagerSecurityDbEncryptionAccessRoleArn()
	ResetSecretsManagerSecurityDbEncryptionSecretId()
	ResetServerName()
	ResetSslMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsDataProviderSettingsOracleSettingsOutputReference
type jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) AsmServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) AsmServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerOracleAsmAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerOracleAsmAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerOracleAsmSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerOracleAsmSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerSecurityDbEncryptionAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecurityDbEncryptionAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerSecurityDbEncryptionAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecurityDbEncryptionAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerSecurityDbEncryptionSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecurityDbEncryptionSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SecretsManagerSecurityDbEncryptionSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecurityDbEncryptionSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsDataProviderSettingsOracleSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsDataProviderSettingsOracleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsDataProviderSettingsOracleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsDataProvider.DmsDataProviderSettingsOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsDataProviderSettingsOracleSettingsOutputReference_Override(d DmsDataProviderSettingsOracleSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsDataProvider.DmsDataProviderSettingsOracleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetAsmServer(val *string) {
	if err := j.validateSetAsmServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmServer",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetSecretsManagerOracleAsmAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetSecretsManagerOracleAsmSecretId(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetSecretsManagerSecurityDbEncryptionAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerSecurityDbEncryptionAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecurityDbEncryptionAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetSecretsManagerSecurityDbEncryptionSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecurityDbEncryptionSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecurityDbEncryptionSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetAsmServer() {
	_jsii_.InvokeVoid(
		d,
		"resetAsmServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetSecretsManagerOracleAsmAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerOracleAsmAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetSecretsManagerOracleAsmSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerOracleAsmSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetSecretsManagerSecurityDbEncryptionAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecurityDbEncryptionAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetSecretsManagerSecurityDbEncryptionSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecurityDbEncryptionSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ResetSslMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSslMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsDataProviderSettingsOracleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

