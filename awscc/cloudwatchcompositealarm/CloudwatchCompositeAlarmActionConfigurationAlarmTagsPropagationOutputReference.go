// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchcompositealarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/cloudwatchcompositealarm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludedTagKeys() *[]*string
	SetExcludedTagKeys(val *[]*string)
	ExcludedTagKeysInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedTagKeys() *[]*string
	SetIncludedTagKeys(val *[]*string)
	IncludedTagKeysInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetEnabled()
	ResetExcludedTagKeys()
	ResetIncludedTagKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference
type jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ExcludedTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ExcludedTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedTagKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) IncludedTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) IncludedTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTagKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchCompositeAlarm.CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference_Override(c CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchCompositeAlarm.CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetExcludedTagKeys(val *[]*string) {
	if err := j.validateSetExcludedTagKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedTagKeys",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetIncludedTagKeys(val *[]*string) {
	if err := j.validateSetIncludedTagKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedTagKeys",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ResetExcludedTagKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedTagKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ResetIncludedTagKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedTagKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudwatchCompositeAlarmActionConfigurationAlarmTagsPropagationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

