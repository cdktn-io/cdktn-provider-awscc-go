// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcscluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsClusterSlurmConfigurationOutputReference interface {
	cdktn.ComplexObject
	Accounting() PcsClusterSlurmConfigurationAccountingOutputReference
	AccountingInput() interface{}
	AuthKey() PcsClusterSlurmConfigurationAuthKeyOutputReference
	AuthKeyInput() interface{}
	CgroupCustomSettings() PcsClusterSlurmConfigurationCgroupCustomSettingsList
	CgroupCustomSettingsInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JwtAuth() PcsClusterSlurmConfigurationJwtAuthOutputReference
	JwtAuthInput() interface{}
	ScaleDownIdleTimeInSeconds() *float64
	SetScaleDownIdleTimeInSeconds(val *float64)
	ScaleDownIdleTimeInSecondsInput() *float64
	SlurmCustomSettings() PcsClusterSlurmConfigurationSlurmCustomSettingsList
	SlurmCustomSettingsInput() interface{}
	SlurmdbdCustomSettings() PcsClusterSlurmConfigurationSlurmdbdCustomSettingsList
	SlurmdbdCustomSettingsInput() interface{}
	SlurmRest() PcsClusterSlurmConfigurationSlurmRestOutputReference
	SlurmRestInput() interface{}
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
	PutAccounting(value *PcsClusterSlurmConfigurationAccounting)
	PutAuthKey(value *PcsClusterSlurmConfigurationAuthKey)
	PutCgroupCustomSettings(value interface{})
	PutJwtAuth(value *PcsClusterSlurmConfigurationJwtAuth)
	PutSlurmCustomSettings(value interface{})
	PutSlurmdbdCustomSettings(value interface{})
	PutSlurmRest(value *PcsClusterSlurmConfigurationSlurmRest)
	ResetAccounting()
	ResetAuthKey()
	ResetCgroupCustomSettings()
	ResetJwtAuth()
	ResetScaleDownIdleTimeInSeconds()
	ResetSlurmCustomSettings()
	ResetSlurmdbdCustomSettings()
	ResetSlurmRest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcsClusterSlurmConfigurationOutputReference
type jsiiProxy_PcsClusterSlurmConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) Accounting() PcsClusterSlurmConfigurationAccountingOutputReference {
	var returns PcsClusterSlurmConfigurationAccountingOutputReference
	_jsii_.Get(
		j,
		"accounting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) AccountingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) AuthKey() PcsClusterSlurmConfigurationAuthKeyOutputReference {
	var returns PcsClusterSlurmConfigurationAuthKeyOutputReference
	_jsii_.Get(
		j,
		"authKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) AuthKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) CgroupCustomSettings() PcsClusterSlurmConfigurationCgroupCustomSettingsList {
	var returns PcsClusterSlurmConfigurationCgroupCustomSettingsList
	_jsii_.Get(
		j,
		"cgroupCustomSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) CgroupCustomSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cgroupCustomSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) JwtAuth() PcsClusterSlurmConfigurationJwtAuthOutputReference {
	var returns PcsClusterSlurmConfigurationJwtAuthOutputReference
	_jsii_.Get(
		j,
		"jwtAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) JwtAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ScaleDownIdleTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownIdleTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ScaleDownIdleTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownIdleTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmCustomSettings() PcsClusterSlurmConfigurationSlurmCustomSettingsList {
	var returns PcsClusterSlurmConfigurationSlurmCustomSettingsList
	_jsii_.Get(
		j,
		"slurmCustomSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmCustomSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slurmCustomSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmdbdCustomSettings() PcsClusterSlurmConfigurationSlurmdbdCustomSettingsList {
	var returns PcsClusterSlurmConfigurationSlurmdbdCustomSettingsList
	_jsii_.Get(
		j,
		"slurmdbdCustomSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmdbdCustomSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slurmdbdCustomSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmRest() PcsClusterSlurmConfigurationSlurmRestOutputReference {
	var returns PcsClusterSlurmConfigurationSlurmRestOutputReference
	_jsii_.Get(
		j,
		"slurmRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) SlurmRestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slurmRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcsClusterSlurmConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PcsClusterSlurmConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPcsClusterSlurmConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcsClusterSlurmConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsCluster.PcsClusterSlurmConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcsClusterSlurmConfigurationOutputReference_Override(p PcsClusterSlurmConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsCluster.PcsClusterSlurmConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetScaleDownIdleTimeInSeconds(val *float64) {
	if err := j.validateSetScaleDownIdleTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownIdleTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcsClusterSlurmConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutAccounting(value *PcsClusterSlurmConfigurationAccounting) {
	if err := p.validatePutAccountingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAccounting",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutAuthKey(value *PcsClusterSlurmConfigurationAuthKey) {
	if err := p.validatePutAuthKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuthKey",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutCgroupCustomSettings(value interface{}) {
	if err := p.validatePutCgroupCustomSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCgroupCustomSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutJwtAuth(value *PcsClusterSlurmConfigurationJwtAuth) {
	if err := p.validatePutJwtAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJwtAuth",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutSlurmCustomSettings(value interface{}) {
	if err := p.validatePutSlurmCustomSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSlurmCustomSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutSlurmdbdCustomSettings(value interface{}) {
	if err := p.validatePutSlurmdbdCustomSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSlurmdbdCustomSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) PutSlurmRest(value *PcsClusterSlurmConfigurationSlurmRest) {
	if err := p.validatePutSlurmRestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSlurmRest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetAccounting() {
	_jsii_.InvokeVoid(
		p,
		"resetAccounting",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetAuthKey() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetCgroupCustomSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetCgroupCustomSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetJwtAuth() {
	_jsii_.InvokeVoid(
		p,
		"resetJwtAuth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetScaleDownIdleTimeInSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetScaleDownIdleTimeInSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetSlurmCustomSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetSlurmCustomSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetSlurmdbdCustomSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetSlurmdbdCustomSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ResetSlurmRest() {
	_jsii_.InvokeVoid(
		p,
		"resetSlurmRest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsClusterSlurmConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

