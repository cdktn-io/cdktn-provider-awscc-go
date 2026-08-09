// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointMySqlSettingsOutputReference interface {
	cdktn.ComplexObject
	AfterConnectScript() *string
	SetAfterConnectScript(val *string)
	AfterConnectScriptInput() *string
	CleanSourceMetadataOnMismatch() interface{}
	SetCleanSourceMetadataOnMismatch(val interface{})
	CleanSourceMetadataOnMismatchInput() interface{}
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
	EventsPollInterval() *float64
	SetEventsPollInterval(val *float64)
	EventsPollIntervalInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	ParallelLoadThreads() *float64
	SetParallelLoadThreads(val *float64)
	ParallelLoadThreadsInput() *float64
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	ServerTimezone() *string
	SetServerTimezone(val *string)
	ServerTimezoneInput() *string
	TargetDbType() *string
	SetTargetDbType(val *string)
	TargetDbTypeInput() *string
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
	ResetCleanSourceMetadataOnMismatch()
	ResetEventsPollInterval()
	ResetMaxFileSize()
	ResetParallelLoadThreads()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	ResetServerTimezone()
	ResetTargetDbType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointMySqlSettingsOutputReference
type jsiiProxy_DmsEndpointMySqlSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) CleanSourceMetadataOnMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) CleanSourceMetadataOnMismatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) EventsPollInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) EventsPollIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ParallelLoadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ParallelLoadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ServerTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ServerTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) TargetDbType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) TargetDbTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointMySqlSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointMySqlSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointMySqlSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointMySqlSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointMySqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointMySqlSettingsOutputReference_Override(d DmsEndpointMySqlSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointMySqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetCleanSourceMetadataOnMismatch(val interface{}) {
	if err := j.validateSetCleanSourceMetadataOnMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanSourceMetadataOnMismatch",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetEventsPollInterval(val *float64) {
	if err := j.validateSetEventsPollIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsPollInterval",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetParallelLoadThreads(val *float64) {
	if err := j.validateSetParallelLoadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelLoadThreads",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetServerTimezone(val *string) {
	if err := j.validateSetServerTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimezone",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetTargetDbType(val *string) {
	if err := j.validateSetTargetDbTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMySqlSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		d,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetCleanSourceMetadataOnMismatch() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanSourceMetadataOnMismatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetEventsPollInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetEventsPollInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetParallelLoadThreads() {
	_jsii_.InvokeVoid(
		d,
		"resetParallelLoadThreads",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetServerTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetServerTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ResetTargetDbType() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetDbType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointMySqlSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

