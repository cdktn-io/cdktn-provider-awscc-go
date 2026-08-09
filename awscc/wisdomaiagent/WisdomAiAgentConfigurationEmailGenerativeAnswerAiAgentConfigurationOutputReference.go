// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference interface {
	cdktn.ComplexObject
	AssociationConfigurations() WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationAssociationConfigurationsList
	AssociationConfigurationsInput() interface{}
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
	EmailGenerativeAnswerAiPromptId() *string
	SetEmailGenerativeAnswerAiPromptId(val *string)
	EmailGenerativeAnswerAiPromptIdInput() *string
	EmailQueryReformulationAiPromptId() *string
	SetEmailQueryReformulationAiPromptId(val *string)
	EmailQueryReformulationAiPromptIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
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
	PutAssociationConfigurations(value interface{})
	ResetAssociationConfigurations()
	ResetEmailGenerativeAnswerAiPromptId()
	ResetEmailQueryReformulationAiPromptId()
	ResetLocale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference
type jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) AssociationConfigurations() WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationAssociationConfigurationsList {
	var returns WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationAssociationConfigurationsList
	_jsii_.Get(
		j,
		"associationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) AssociationConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associationConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) EmailGenerativeAnswerAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailGenerativeAnswerAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) EmailGenerativeAnswerAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailGenerativeAnswerAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) EmailQueryReformulationAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailQueryReformulationAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) EmailQueryReformulationAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailQueryReformulationAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference_Override(w WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetEmailGenerativeAnswerAiPromptId(val *string) {
	if err := j.validateSetEmailGenerativeAnswerAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailGenerativeAnswerAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetEmailQueryReformulationAiPromptId(val *string) {
	if err := j.validateSetEmailQueryReformulationAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailQueryReformulationAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) PutAssociationConfigurations(value interface{}) {
	if err := w.validatePutAssociationConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAssociationConfigurations",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ResetAssociationConfigurations() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationConfigurations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ResetEmailGenerativeAnswerAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailGenerativeAnswerAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ResetEmailQueryReformulationAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailQueryReformulationAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ResetLocale() {
	_jsii_.InvokeVoid(
		w,
		"resetLocale",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

