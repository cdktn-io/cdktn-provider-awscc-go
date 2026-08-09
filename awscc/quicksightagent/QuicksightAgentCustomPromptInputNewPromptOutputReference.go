// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightAgentCustomPromptInputNewPromptOutputReference interface {
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
	CustomInstructions() *string
	SetCustomInstructions(val *string)
	CustomInstructionsInput() *string
	// Experimental.
	Fqn() *string
	Identity() *string
	SetIdentity(val *string)
	IdentityInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputStyle() *string
	SetOutputStyle(val *string)
	OutputStyleInput() *string
	ResponseLength() *string
	SetResponseLength(val *string)
	ResponseLengthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tone() *string
	SetTone(val *string)
	ToneInput() *string
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
	ResetCustomInstructions()
	ResetIdentity()
	ResetOutputStyle()
	ResetResponseLength()
	ResetTone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightAgentCustomPromptInputNewPromptOutputReference
type jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) CustomInstructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) CustomInstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) IdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) OutputStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) OutputStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResponseLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResponseLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ToneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toneInput",
		&returns,
	)
	return returns
}


func NewQuicksightAgentCustomPromptInputNewPromptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightAgentCustomPromptInputNewPromptOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightAgentCustomPromptInputNewPromptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightAgent.QuicksightAgentCustomPromptInputNewPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightAgentCustomPromptInputNewPromptOutputReference_Override(q QuicksightAgentCustomPromptInputNewPromptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightAgent.QuicksightAgentCustomPromptInputNewPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetCustomInstructions(val *string) {
	if err := j.validateSetCustomInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customInstructions",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetIdentity(val *string) {
	if err := j.validateSetIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetOutputStyle(val *string) {
	if err := j.validateSetOutputStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputStyle",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetResponseLength(val *string) {
	if err := j.validateSetResponseLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseLength",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference)SetTone(val *string) {
	if err := j.validateSetToneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResetCustomInstructions() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomInstructions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResetIdentity() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentity",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResetOutputStyle() {
	_jsii_.InvokeVoid(
		q,
		"resetOutputStyle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResetResponseLength() {
	_jsii_.InvokeVoid(
		q,
		"resetResponseLength",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ResetTone() {
	_jsii_.InvokeVoid(
		q,
		"resetTone",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAgentCustomPromptInputNewPromptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

