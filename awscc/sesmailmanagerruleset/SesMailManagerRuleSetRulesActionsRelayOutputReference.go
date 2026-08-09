// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesmailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerRuleSetRulesActionsRelayOutputReference interface {
	cdktn.ComplexObject
	ActionFailurePolicy() *string
	SetActionFailurePolicy(val *string)
	ActionFailurePolicyInput() *string
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
	MailFrom() *string
	SetMailFrom(val *string)
	MailFromInput() *string
	Relay() *string
	SetRelay(val *string)
	RelayInput() *string
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
	ResetActionFailurePolicy()
	ResetMailFrom()
	ResetRelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerRuleSetRulesActionsRelayOutputReference
type jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ActionFailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ActionFailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) MailFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) MailFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) Relay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) RelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesMailManagerRuleSetRulesActionsRelayOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SesMailManagerRuleSetRulesActionsRelayOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerRuleSetRulesActionsRelayOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsRelayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesMailManagerRuleSetRulesActionsRelayOutputReference_Override(s SesMailManagerRuleSetRulesActionsRelayOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsRelayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetActionFailurePolicy(val *string) {
	if err := j.validateSetActionFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionFailurePolicy",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetMailFrom(val *string) {
	if err := j.validateSetMailFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailFrom",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetRelay(val *string) {
	if err := j.validateSetRelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relay",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ResetActionFailurePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetActionFailurePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ResetMailFrom() {
	_jsii_.InvokeVoid(
		s,
		"resetMailFrom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ResetRelay() {
	_jsii_.InvokeVoid(
		s,
		"resetRelay",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsRelayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

