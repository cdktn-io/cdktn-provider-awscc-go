// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsEndpointIbmDb2SettingsOutputReference interface {
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
	CurrentLsn() *string
	SetCurrentLsn(val *string)
	CurrentLsnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeepCsvFiles() interface{}
	SetKeepCsvFiles(val interface{})
	KeepCsvFilesInput() interface{}
	LoadTimeout() *float64
	SetLoadTimeout(val *float64)
	LoadTimeoutInput() *float64
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	MaxKBytesPerRead() *float64
	SetMaxKBytesPerRead(val *float64)
	MaxKBytesPerReadInput() *float64
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
	SetDataCaptureChanges() interface{}
	SetSetDataCaptureChanges(val interface{})
	SetDataCaptureChangesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetCurrentLsn()
	ResetKeepCsvFiles()
	ResetLoadTimeout()
	ResetMaxFileSize()
	ResetMaxKBytesPerRead()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	ResetSetDataCaptureChanges()
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

// The jsii proxy struct for DmsEndpointIbmDb2SettingsOutputReference
type jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) CurrentLsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentLsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) CurrentLsnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentLsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) KeepCsvFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepCsvFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) KeepCsvFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepCsvFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) LoadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) LoadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) MaxKBytesPerRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxKBytesPerRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) MaxKBytesPerReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxKBytesPerReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SetDataCaptureChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setDataCaptureChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) SetDataCaptureChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setDataCaptureChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) WriteBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) WriteBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBufferSizeInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointIbmDb2SettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DmsEndpointIbmDb2SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointIbmDb2SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointIbmDb2SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointIbmDb2SettingsOutputReference_Override(d DmsEndpointIbmDb2SettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpointIbmDb2SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetCurrentLsn(val *string) {
	if err := j.validateSetCurrentLsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentLsn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetKeepCsvFiles(val interface{}) {
	if err := j.validateSetKeepCsvFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepCsvFiles",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetLoadTimeout(val *float64) {
	if err := j.validateSetLoadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetMaxKBytesPerRead(val *float64) {
	if err := j.validateSetMaxKBytesPerReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxKBytesPerRead",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetSetDataCaptureChanges(val interface{}) {
	if err := j.validateSetSetDataCaptureChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setDataCaptureChanges",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference)SetWriteBufferSize(val *float64) {
	if err := j.validateSetWriteBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeBufferSize",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetCurrentLsn() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrentLsn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetKeepCsvFiles() {
	_jsii_.InvokeVoid(
		d,
		"resetKeepCsvFiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetLoadTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetMaxKBytesPerRead() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxKBytesPerRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetSetDataCaptureChanges() {
	_jsii_.InvokeVoid(
		d,
		"resetSetDataCaptureChanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ResetWriteBufferSize() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteBufferSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointIbmDb2SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

