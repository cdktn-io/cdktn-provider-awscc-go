// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connecthoursofoperation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectHoursOfOperationHoursOfOperationOverridesOutputReference interface {
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
	EffectiveFrom() *string
	SetEffectiveFrom(val *string)
	EffectiveFromInput() *string
	EffectiveTill() *string
	SetEffectiveTill(val *string)
	EffectiveTillInput() *string
	// Experimental.
	Fqn() *string
	HoursOfOperationOverrideId() *string
	SetHoursOfOperationOverrideId(val *string)
	HoursOfOperationOverrideIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OverrideConfig() ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigList
	OverrideConfigInput() interface{}
	OverrideDescription() *string
	SetOverrideDescription(val *string)
	OverrideDescriptionInput() *string
	OverrideName() *string
	SetOverrideName(val *string)
	OverrideNameInput() *string
	OverrideType() *string
	SetOverrideType(val *string)
	OverrideTypeInput() *string
	RecurrenceConfig() ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfigOutputReference
	RecurrenceConfigInput() interface{}
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
	PutOverrideConfig(value interface{})
	PutRecurrenceConfig(value *ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfig)
	ResetEffectiveFrom()
	ResetEffectiveTill()
	ResetHoursOfOperationOverrideId()
	ResetOverrideConfig()
	ResetOverrideDescription()
	ResetOverrideName()
	ResetOverrideType()
	ResetRecurrenceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectHoursOfOperationHoursOfOperationOverridesOutputReference
type jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) EffectiveFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) EffectiveFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) EffectiveTill() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveTill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) EffectiveTillInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveTillInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) HoursOfOperationOverrideId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hoursOfOperationOverrideId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) HoursOfOperationOverrideIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hoursOfOperationOverrideIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideConfig() ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigList {
	var returns ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigList
	_jsii_.Get(
		j,
		"overrideConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) OverrideTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) RecurrenceConfig() ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfigOutputReference {
	var returns ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfigOutputReference
	_jsii_.Get(
		j,
		"recurrenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) RecurrenceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recurrenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectHoursOfOperationHoursOfOperationOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConnectHoursOfOperationHoursOfOperationOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewConnectHoursOfOperationHoursOfOperationOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectHoursOfOperation.ConnectHoursOfOperationHoursOfOperationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConnectHoursOfOperationHoursOfOperationOverridesOutputReference_Override(c ConnectHoursOfOperationHoursOfOperationOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectHoursOfOperation.ConnectHoursOfOperationHoursOfOperationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetEffectiveFrom(val *string) {
	if err := j.validateSetEffectiveFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveFrom",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetEffectiveTill(val *string) {
	if err := j.validateSetEffectiveTillParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveTill",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetHoursOfOperationOverrideId(val *string) {
	if err := j.validateSetHoursOfOperationOverrideIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfOperationOverrideId",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetOverrideDescription(val *string) {
	if err := j.validateSetOverrideDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideDescription",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetOverrideName(val *string) {
	if err := j.validateSetOverrideNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideName",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetOverrideType(val *string) {
	if err := j.validateSetOverrideTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideType",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) PutOverrideConfig(value interface{}) {
	if err := c.validatePutOverrideConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOverrideConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) PutRecurrenceConfig(value *ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfig) {
	if err := c.validatePutRecurrenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRecurrenceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetEffectiveFrom() {
	_jsii_.InvokeVoid(
		c,
		"resetEffectiveFrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetEffectiveTill() {
	_jsii_.InvokeVoid(
		c,
		"resetEffectiveTill",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetHoursOfOperationOverrideId() {
	_jsii_.InvokeVoid(
		c,
		"resetHoursOfOperationOverrideId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetOverrideConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetOverrideDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetOverrideName() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetOverrideType() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ResetRecurrenceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRecurrenceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectHoursOfOperationHoursOfOperationOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

