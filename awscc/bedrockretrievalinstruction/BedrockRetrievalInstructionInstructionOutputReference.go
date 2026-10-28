// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockretrievalinstruction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/bedrockretrievalinstruction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockRetrievalInstructionInstructionOutputReference interface {
	cdktn.ComplexObject
	AnswerFormat() *string
	SetAnswerFormat(val *string)
	AnswerFormatInput() *string
	AnswerPolicy() *string
	SetAnswerPolicy(val *string)
	AnswerPolicyInput() *string
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
	Persona() *string
	SetPersona(val *string)
	PersonaInput() *string
	RetrievalGuidance() *string
	SetRetrievalGuidance(val *string)
	RetrievalGuidanceInput() *string
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
	ResetAnswerFormat()
	ResetAnswerPolicy()
	ResetPersona()
	ResetRetrievalGuidance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockRetrievalInstructionInstructionOutputReference
type jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) AnswerFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) AnswerFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) AnswerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) AnswerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) Persona() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persona",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) PersonaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) RetrievalGuidance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalGuidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) RetrievalGuidanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalGuidanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockRetrievalInstructionInstructionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockRetrievalInstructionInstructionOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockRetrievalInstructionInstructionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockRetrievalInstruction.BedrockRetrievalInstructionInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockRetrievalInstructionInstructionOutputReference_Override(b BedrockRetrievalInstructionInstructionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockRetrievalInstruction.BedrockRetrievalInstructionInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetAnswerFormat(val *string) {
	if err := j.validateSetAnswerFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerFormat",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetAnswerPolicy(val *string) {
	if err := j.validateSetAnswerPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerPolicy",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetPersona(val *string) {
	if err := j.validateSetPersonaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persona",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetRetrievalGuidance(val *string) {
	if err := j.validateSetRetrievalGuidanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalGuidance",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ResetAnswerFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetAnswerFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ResetAnswerPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetAnswerPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ResetPersona() {
	_jsii_.InvokeVoid(
		b,
		"resetPersona",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ResetRetrievalGuidance() {
	_jsii_.InvokeVoid(
		b,
		"resetRetrievalGuidance",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockRetrievalInstructionInstructionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

