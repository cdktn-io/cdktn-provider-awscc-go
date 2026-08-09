// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdommessagetemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomMessageTemplateDefaultAttributesOutputReference interface {
	cdktn.ComplexObject
	AgentAttributes() WisdomMessageTemplateDefaultAttributesAgentAttributesOutputReference
	AgentAttributesInput() interface{}
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
	CustomAttributes() *map[string]*string
	SetCustomAttributes(val *map[string]*string)
	CustomAttributesInput() *map[string]*string
	CustomerProfileAttributes() WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference
	CustomerProfileAttributesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SystemAttributes() WisdomMessageTemplateDefaultAttributesSystemAttributesOutputReference
	SystemAttributesInput() interface{}
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
	PutAgentAttributes(value *WisdomMessageTemplateDefaultAttributesAgentAttributes)
	PutCustomerProfileAttributes(value *WisdomMessageTemplateDefaultAttributesCustomerProfileAttributes)
	PutSystemAttributes(value *WisdomMessageTemplateDefaultAttributesSystemAttributes)
	ResetAgentAttributes()
	ResetCustomAttributes()
	ResetCustomerProfileAttributes()
	ResetSystemAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomMessageTemplateDefaultAttributesOutputReference
type jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) AgentAttributes() WisdomMessageTemplateDefaultAttributesAgentAttributesOutputReference {
	var returns WisdomMessageTemplateDefaultAttributesAgentAttributesOutputReference
	_jsii_.Get(
		j,
		"agentAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) AgentAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) CustomerProfileAttributes() WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference {
	var returns WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference
	_jsii_.Get(
		j,
		"customerProfileAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) CustomerProfileAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerProfileAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) SystemAttributes() WisdomMessageTemplateDefaultAttributesSystemAttributesOutputReference {
	var returns WisdomMessageTemplateDefaultAttributesSystemAttributesOutputReference
	_jsii_.Get(
		j,
		"systemAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) SystemAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomMessageTemplateDefaultAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomMessageTemplateDefaultAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomMessageTemplateDefaultAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateDefaultAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomMessageTemplateDefaultAttributesOutputReference_Override(w WisdomMessageTemplateDefaultAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateDefaultAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) PutAgentAttributes(value *WisdomMessageTemplateDefaultAttributesAgentAttributes) {
	if err := w.validatePutAgentAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAgentAttributes",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) PutCustomerProfileAttributes(value *WisdomMessageTemplateDefaultAttributesCustomerProfileAttributes) {
	if err := w.validatePutCustomerProfileAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomerProfileAttributes",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) PutSystemAttributes(value *WisdomMessageTemplateDefaultAttributesSystemAttributes) {
	if err := w.validatePutSystemAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSystemAttributes",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ResetAgentAttributes() {
	_jsii_.InvokeVoid(
		w,
		"resetAgentAttributes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ResetCustomerProfileAttributes() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomerProfileAttributes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ResetSystemAttributes() {
	_jsii_.InvokeVoid(
		w,
		"resetSystemAttributes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

