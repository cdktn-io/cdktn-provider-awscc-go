// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference interface {
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
	PiiEntityTypes() *[]*string
	SetPiiEntityTypes(val *[]*string)
	PiiEntityTypesInput() *[]*string
	RedactionMaskMode() *string
	SetRedactionMaskMode(val *string)
	RedactionMaskModeInput() *string
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
	ResetPiiEntityTypes()
	ResetRedactionMaskMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference
type jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) PiiEntityTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) PiiEntityTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) RedactionMaskMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionMaskMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) RedactionMaskModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionMaskModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference_Override(b BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetPiiEntityTypes(val *[]*string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetRedactionMaskMode(val *string) {
	if err := j.validateSetRedactionMaskModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactionMaskMode",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ResetRedactionMaskMode() {
	_jsii_.InvokeVoid(
		b,
		"resetRedactionMaskMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

