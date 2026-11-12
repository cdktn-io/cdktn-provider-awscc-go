// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/networkmanagervpcattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference interface {
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
	MatchSourceCidr() *string
	SetMatchSourceCidr(val *string)
	MatchSourceCidrInput() *string
	NatCidr() *string
	SetNatCidr(val *string)
	NatCidrInput() *string
	RuleNumber() *float64
	SetRuleNumber(val *float64)
	RuleNumberInput() *float64
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
	ResetMatchSourceCidr()
	ResetNatCidr()
	ResetRuleNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference
type jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) MatchSourceCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchSourceCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) MatchSourceCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchSourceCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) NatCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) NatCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) RuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) RuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkmanagerVpcAttachmentNatGatewayRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkmanagerVpcAttachmentNatGatewayRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerVpcAttachment.NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkmanagerVpcAttachmentNatGatewayRulesOutputReference_Override(n NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerVpcAttachment.NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetMatchSourceCidr(val *string) {
	if err := j.validateSetMatchSourceCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchSourceCidr",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetNatCidr(val *string) {
	if err := j.validateSetNatCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natCidr",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetRuleNumber(val *float64) {
	if err := j.validateSetRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNumber",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ResetMatchSourceCidr() {
	_jsii_.InvokeVoid(
		n,
		"resetMatchSourceCidr",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ResetNatCidr() {
	_jsii_.InvokeVoid(
		n,
		"resetNatCidr",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ResetRuleNumber() {
	_jsii_.InvokeVoid(
		n,
		"resetRuleNumber",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentNatGatewayRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

