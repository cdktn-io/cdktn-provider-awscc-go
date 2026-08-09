// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/b2bipartnership/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartingFunctionalGroupControlNumber() *float64
	SetStartingFunctionalGroupControlNumber(val *float64)
	StartingFunctionalGroupControlNumberInput() *float64
	StartingInterchangeControlNumber() *float64
	SetStartingInterchangeControlNumber(val *float64)
	StartingInterchangeControlNumberInput() *float64
	StartingTransactionSetControlNumber() *float64
	SetStartingTransactionSetControlNumber(val *float64)
	StartingTransactionSetControlNumberInput() *float64
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
	ResetStartingFunctionalGroupControlNumber()
	ResetStartingInterchangeControlNumber()
	ResetStartingTransactionSetControlNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference
type jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingFunctionalGroupControlNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingFunctionalGroupControlNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingFunctionalGroupControlNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingFunctionalGroupControlNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingInterchangeControlNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingInterchangeControlNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingInterchangeControlNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingInterchangeControlNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingTransactionSetControlNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingTransactionSetControlNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) StartingTransactionSetControlNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingTransactionSetControlNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference {
	_init_.Initialize()

	if err := validateNewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiPartnership.B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewB2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference_Override(b B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiPartnership.B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetStartingFunctionalGroupControlNumber(val *float64) {
	if err := j.validateSetStartingFunctionalGroupControlNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingFunctionalGroupControlNumber",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetStartingInterchangeControlNumber(val *float64) {
	if err := j.validateSetStartingInterchangeControlNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingInterchangeControlNumber",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetStartingTransactionSetControlNumber(val *float64) {
	if err := j.validateSetStartingTransactionSetControlNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingTransactionSetControlNumber",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ResetStartingFunctionalGroupControlNumber() {
	_jsii_.InvokeVoid(
		b,
		"resetStartingFunctionalGroupControlNumber",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ResetStartingInterchangeControlNumber() {
	_jsii_.InvokeVoid(
		b,
		"resetStartingInterchangeControlNumber",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ResetStartingTransactionSetControlNumber() {
	_jsii_.InvokeVoid(
		b,
		"resetStartingTransactionSetControlNumber",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

