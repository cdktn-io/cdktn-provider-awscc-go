// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectmetric/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference interface {
	cdktn.ComplexObject
	BooleanCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference
	BooleanConditionInput() interface{}
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
	MetricFilterKey() *string
	SetMetricFilterKey(val *string)
	MetricFilterKeyInput() *string
	Negate() interface{}
	SetNegate(val interface{})
	NegateInput() interface{}
	NumberCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference
	NumberConditionInput() interface{}
	StringCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference
	StringConditionInput() interface{}
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
	PutBooleanCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanCondition)
	PutNumberCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberCondition)
	PutStringCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringCondition)
	ResetBooleanCondition()
	ResetMetricFilterKey()
	ResetNegate()
	ResetNumberCondition()
	ResetStringCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference
type jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) BooleanCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference {
	var returns ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference
	_jsii_.Get(
		j,
		"booleanCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) BooleanConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) MetricFilterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricFilterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) MetricFilterKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricFilterKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Negate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) NegateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) NumberCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference {
	var returns ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference
	_jsii_.Get(
		j,
		"numberCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) NumberConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) StringCondition() ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference {
	var returns ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference
	_jsii_.Get(
		j,
		"stringCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) StringConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectMetric.ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference_Override(c ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectMetric.ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetMetricFilterKey(val *string) {
	if err := j.validateSetMetricFilterKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricFilterKey",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetNegate(val interface{}) {
	if err := j.validateSetNegateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negate",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) PutBooleanCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanCondition) {
	if err := c.validatePutBooleanConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBooleanCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) PutNumberCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberCondition) {
	if err := c.validatePutNumberConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNumberCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) PutStringCondition(value *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringCondition) {
	if err := c.validatePutStringConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStringCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ResetBooleanCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetBooleanCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ResetMetricFilterKey() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricFilterKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ResetNegate() {
	_jsii_.InvokeVoid(
		c,
		"resetNegate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ResetNumberCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNumberCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ResetStringCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetStringCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

