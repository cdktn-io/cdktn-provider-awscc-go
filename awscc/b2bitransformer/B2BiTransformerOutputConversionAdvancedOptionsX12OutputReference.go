// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/b2bitransformer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference interface {
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
	SplitOptions() B2BiTransformerOutputConversionAdvancedOptionsX12SplitOptionsOutputReference
	SplitOptionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValidationOptions() B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptionsOutputReference
	ValidationOptionsInput() interface{}
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
	PutSplitOptions(value *B2BiTransformerOutputConversionAdvancedOptionsX12SplitOptions)
	PutValidationOptions(value *B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptions)
	ResetSplitOptions()
	ResetValidationOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference
type jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) SplitOptions() B2BiTransformerOutputConversionAdvancedOptionsX12SplitOptionsOutputReference {
	var returns B2BiTransformerOutputConversionAdvancedOptionsX12SplitOptionsOutputReference
	_jsii_.Get(
		j,
		"splitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) SplitOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ValidationOptions() B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptionsOutputReference {
	var returns B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptionsOutputReference
	_jsii_.Get(
		j,
		"validationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ValidationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationOptionsInput",
		&returns,
	)
	return returns
}


func NewB2BiTransformerOutputConversionAdvancedOptionsX12OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference {
	_init_.Initialize()

	if err := validateNewB2BiTransformerOutputConversionAdvancedOptionsX12OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiTransformer.B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewB2BiTransformerOutputConversionAdvancedOptionsX12OutputReference_Override(b B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.b2BiTransformer.B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) PutSplitOptions(value *B2BiTransformerOutputConversionAdvancedOptionsX12SplitOptions) {
	if err := b.validatePutSplitOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSplitOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) PutValidationOptions(value *B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptions) {
	if err := b.validatePutValidationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putValidationOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ResetSplitOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetSplitOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ResetValidationOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetValidationOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_B2BiTransformerOutputConversionAdvancedOptionsX12OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

