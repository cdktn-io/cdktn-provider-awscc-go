// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagertransitgatewayroutetableattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/networkmanagertransitgatewayroutetableattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference interface {
	cdktn.ComplexObject
	AttachmentPolicyRuleNumber() *float64
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
	InternalValue() *NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange
	SetInternalValue(val *NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange)
	NetworkFunctionGroupName() *string
	Tags() NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeTagsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference
type jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) InternalValue() *NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange {
	var returns *NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) NetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) Tags() NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeTagsList {
	var returns NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerTransitGatewayRouteTableAttachment.NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference_Override(n NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerTransitGatewayRouteTableAttachment.NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetInternalValue(val *NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

