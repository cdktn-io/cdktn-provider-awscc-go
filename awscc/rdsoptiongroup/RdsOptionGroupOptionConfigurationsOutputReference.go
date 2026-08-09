// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsoptiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsoptiongroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsOptionGroupOptionConfigurationsOutputReference interface {
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
	DbSecurityGroupMemberships() *[]*string
	SetDbSecurityGroupMemberships(val *[]*string)
	DbSecurityGroupMembershipsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OptionName() *string
	SetOptionName(val *string)
	OptionNameInput() *string
	OptionSettings() RdsOptionGroupOptionConfigurationsOptionSettingsList
	OptionSettingsInput() interface{}
	OptionVersion() *string
	SetOptionVersion(val *string)
	OptionVersionInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcSecurityGroupMemberships() *[]*string
	SetVpcSecurityGroupMemberships(val *[]*string)
	VpcSecurityGroupMembershipsInput() *[]*string
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
	PutOptionSettings(value interface{})
	ResetDbSecurityGroupMemberships()
	ResetOptionName()
	ResetOptionSettings()
	ResetOptionVersion()
	ResetPort()
	ResetVpcSecurityGroupMemberships()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsOptionGroupOptionConfigurationsOutputReference
type jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) DbSecurityGroupMemberships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroupMemberships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) DbSecurityGroupMembershipsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroupMembershipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionSettings() RdsOptionGroupOptionConfigurationsOptionSettingsList {
	var returns RdsOptionGroupOptionConfigurationsOptionSettingsList
	_jsii_.Get(
		j,
		"optionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) OptionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) VpcSecurityGroupMemberships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupMemberships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) VpcSecurityGroupMembershipsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupMembershipsInput",
		&returns,
	)
	return returns
}


func NewRdsOptionGroupOptionConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RdsOptionGroupOptionConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewRdsOptionGroupOptionConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsOptionGroup.RdsOptionGroupOptionConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRdsOptionGroupOptionConfigurationsOutputReference_Override(r RdsOptionGroupOptionConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsOptionGroup.RdsOptionGroupOptionConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetDbSecurityGroupMemberships(val *[]*string) {
	if err := j.validateSetDbSecurityGroupMembershipsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSecurityGroupMemberships",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetOptionName(val *string) {
	if err := j.validateSetOptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionName",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetOptionVersion(val *string) {
	if err := j.validateSetOptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionVersion",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference)SetVpcSecurityGroupMemberships(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupMembershipsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupMemberships",
		val,
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) PutOptionSettings(value interface{}) {
	if err := r.validatePutOptionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putOptionSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetDbSecurityGroupMemberships() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSecurityGroupMemberships",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetOptionName() {
	_jsii_.InvokeVoid(
		r,
		"resetOptionName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetOptionSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetOptionSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetOptionVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetOptionVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ResetVpcSecurityGroupMemberships() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupMemberships",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsOptionGroupOptionConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

