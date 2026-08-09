// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/b2bipartnership/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference interface {
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
	ControlNumbers() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference
	ControlNumbersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delimiters() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimitersOutputReference
	DelimitersInput() interface{}
	// Experimental.
	Fqn() *string
	FunctionalGroupHeaders() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeadersOutputReference
	FunctionalGroupHeadersInput() interface{}
	Gs05TimeFormat() *string
	SetGs05TimeFormat(val *string)
	Gs05TimeFormatInput() *string
	InterchangeControlHeaders() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeadersOutputReference
	InterchangeControlHeadersInput() interface{}
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
	ValidateEdi() interface{}
	SetValidateEdi(val interface{})
	ValidateEdiInput() interface{}
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
	PutControlNumbers(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbers)
	PutDelimiters(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimiters)
	PutFunctionalGroupHeaders(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeaders)
	PutInterchangeControlHeaders(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeaders)
	ResetControlNumbers()
	ResetDelimiters()
	ResetFunctionalGroupHeaders()
	ResetGs05TimeFormat()
	ResetInterchangeControlHeaders()
	ResetValidateEdi()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference
type jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ControlNumbers() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference {
	var returns B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference
	_jsii_.Get(
		j,
		"controlNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ControlNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) Delimiters() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimitersOutputReference {
	var returns B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimitersOutputReference
	_jsii_.Get(
		j,
		"delimiters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) DelimitersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delimitersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) FunctionalGroupHeaders() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeadersOutputReference {
	var returns B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeadersOutputReference
	_jsii_.Get(
		j,
		"functionalGroupHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) FunctionalGroupHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionalGroupHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) Gs05TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gs05TimeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) Gs05TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gs05TimeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) InterchangeControlHeaders() B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeadersOutputReference {
	var returns B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeadersOutputReference
	_jsii_.Get(
		j,
		"interchangeControlHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) InterchangeControlHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interchangeControlHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ValidateEdi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateEdi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ValidateEdiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateEdiInput",
		&returns,
	)
	return returns
}


func NewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference {
	_init_.Initialize()

	if err := validateNewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiPartnership.B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference_Override(b B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiPartnership.B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetGs05TimeFormat(val *string) {
	if err := j.validateSetGs05TimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gs05TimeFormat",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference)SetValidateEdi(val interface{}) {
	if err := j.validateSetValidateEdiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateEdi",
		val,
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) PutControlNumbers(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbers) {
	if err := b.validatePutControlNumbersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putControlNumbers",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) PutDelimiters(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimiters) {
	if err := b.validatePutDelimitersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDelimiters",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) PutFunctionalGroupHeaders(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeaders) {
	if err := b.validatePutFunctionalGroupHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFunctionalGroupHeaders",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) PutInterchangeControlHeaders(value *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeaders) {
	if err := b.validatePutInterchangeControlHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInterchangeControlHeaders",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetControlNumbers() {
	_jsii_.InvokeVoid(
		b,
		"resetControlNumbers",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetDelimiters() {
	_jsii_.InvokeVoid(
		b,
		"resetDelimiters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetFunctionalGroupHeaders() {
	_jsii_.InvokeVoid(
		b,
		"resetFunctionalGroupHeaders",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetGs05TimeFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetGs05TimeFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetInterchangeControlHeaders() {
	_jsii_.InvokeVoid(
		b,
		"resetInterchangeControlHeaders",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ResetValidateEdi() {
	_jsii_.InvokeVoid(
		b,
		"resetValidateEdi",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

