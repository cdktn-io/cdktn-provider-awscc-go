// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/networkmanagervpcattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference interface {
	cdktn.ComplexObject
	AttachmentPolicyRuleNumber() *float64
	SetAttachmentPolicyRuleNumber(val *float64)
	AttachmentPolicyRuleNumberInput() *float64
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
	NetworkFunctionGroupName() *string
	SetNetworkFunctionGroupName(val *string)
	NetworkFunctionGroupNameInput() *string
	Tags() NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeTagsList
	TagsInput() interface{}
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
	PutTags(value interface{})
	ResetAttachmentPolicyRuleNumber()
	ResetNetworkFunctionGroupName()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference
type jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) AttachmentPolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) NetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) NetworkFunctionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) Tags() NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeTagsList {
	var returns NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerVpcAttachment.NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference_Override(n NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerVpcAttachment.NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetAttachmentPolicyRuleNumber(val *float64) {
	if err := j.validateSetAttachmentPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentPolicyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetNetworkFunctionGroupName(val *string) {
	if err := j.validateSetNetworkFunctionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFunctionGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) PutTags(value interface{}) {
	if err := n.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetAttachmentPolicyRuleNumber() {
	_jsii_.InvokeVoid(
		n,
		"resetAttachmentPolicyRuleNumber",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetNetworkFunctionGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkFunctionGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

