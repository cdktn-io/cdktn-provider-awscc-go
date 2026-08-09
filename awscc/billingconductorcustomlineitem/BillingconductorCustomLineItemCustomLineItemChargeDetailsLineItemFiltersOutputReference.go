// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorcustomlineitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/billingconductorcustomlineitem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference interface {
	cdktn.ComplexObject
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
	AttributeValues() *[]*string
	SetAttributeValues(val *[]*string)
	AttributeValuesInput() *[]*string
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
	MatchOption() *string
	SetMatchOption(val *string)
	MatchOptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	ResetAttribute()
	ResetAttributeValues()
	ResetMatchOption()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference
type jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) AttributeValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) AttributeValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) MatchOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) MatchOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewBillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewBillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorCustomLineItem.BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference_Override(b BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorCustomLineItem.BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetAttributeValues(val *[]*string) {
	if err := j.validateSetAttributeValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeValues",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetMatchOption(val *string) {
	if err := j.validateSetMatchOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchOption",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		b,
		"resetAttribute",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ResetAttributeValues() {
	_jsii_.InvokeVoid(
		b,
		"resetAttributeValues",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ResetMatchOption() {
	_jsii_.InvokeVoid(
		b,
		"resetMatchOption",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		b,
		"resetValues",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

