// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagersitetositevpnattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/networkmanagersitetositevpnattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList
type jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList {
	_init_.Initialize()

	if err := validateNewNetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerSiteToSiteVpnAttachment.NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList_Override(n NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.networkmanagerSiteToSiteVpnAttachment.NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := n.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		n,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) Get(index *float64) NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeTagsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

