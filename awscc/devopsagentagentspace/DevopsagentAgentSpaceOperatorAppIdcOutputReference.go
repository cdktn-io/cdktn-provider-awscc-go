// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/devopsagentagentspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentAgentSpaceOperatorAppIdcOutputReference interface {
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdcApplicationArn() *string
	IdcInstanceArn() *string
	SetIdcInstanceArn(val *string)
	IdcInstanceArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperatorAppRoleArn() *string
	SetOperatorAppRoleArn(val *string)
	OperatorAppRoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdatedAt() *string
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
	ResetIdcInstanceArn()
	ResetOperatorAppRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevopsagentAgentSpaceOperatorAppIdcOutputReference
type jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) IdcApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) IdcInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) IdcInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) OperatorAppRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorAppRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) OperatorAppRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorAppRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDevopsagentAgentSpaceOperatorAppIdcOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DevopsagentAgentSpaceOperatorAppIdcOutputReference {
	_init_.Initialize()

	if err := validateNewDevopsagentAgentSpaceOperatorAppIdcOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentAgentSpace.DevopsagentAgentSpaceOperatorAppIdcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevopsagentAgentSpaceOperatorAppIdcOutputReference_Override(d DevopsagentAgentSpaceOperatorAppIdcOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentAgentSpace.DevopsagentAgentSpaceOperatorAppIdcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetIdcInstanceArn(val *string) {
	if err := j.validateSetIdcInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idcInstanceArn",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetOperatorAppRoleArn(val *string) {
	if err := j.validateSetOperatorAppRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatorAppRoleArn",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ResetIdcInstanceArn() {
	_jsii_.InvokeVoid(
		d,
		"resetIdcInstanceArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ResetOperatorAppRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetOperatorAppRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DevopsagentAgentSpaceOperatorAppIdcOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

