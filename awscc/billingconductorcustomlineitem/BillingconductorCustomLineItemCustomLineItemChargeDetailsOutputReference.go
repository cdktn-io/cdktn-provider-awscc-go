// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorcustomlineitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/billingconductorcustomlineitem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference interface {
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
	Flat() BillingconductorCustomLineItemCustomLineItemChargeDetailsFlatOutputReference
	FlatInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LineItemFilters() BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersList
	LineItemFiltersInput() interface{}
	Percentage() BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentageOutputReference
	PercentageInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutFlat(value *BillingconductorCustomLineItemCustomLineItemChargeDetailsFlat)
	PutLineItemFilters(value interface{})
	PutPercentage(value *BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentage)
	ResetFlat()
	ResetLineItemFilters()
	ResetPercentage()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference
type jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) Flat() BillingconductorCustomLineItemCustomLineItemChargeDetailsFlatOutputReference {
	var returns BillingconductorCustomLineItemCustomLineItemChargeDetailsFlatOutputReference
	_jsii_.Get(
		j,
		"flat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) FlatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) LineItemFilters() BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersList {
	var returns BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFiltersList
	_jsii_.Get(
		j,
		"lineItemFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) LineItemFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lineItemFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) Percentage() BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentageOutputReference {
	var returns BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentageOutputReference
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) PercentageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewBillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorCustomLineItem.BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference_Override(b BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorCustomLineItem.BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) PutFlat(value *BillingconductorCustomLineItemCustomLineItemChargeDetailsFlat) {
	if err := b.validatePutFlatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFlat",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) PutLineItemFilters(value interface{}) {
	if err := b.validatePutLineItemFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLineItemFilters",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) PutPercentage(value *BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentage) {
	if err := b.validatePutPercentageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPercentage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ResetFlat() {
	_jsii_.InvokeVoid(
		b,
		"resetFlat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ResetLineItemFilters() {
	_jsii_.InvokeVoid(
		b,
		"resetLineItemFilters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		b,
		"resetPercentage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

