// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritymanagerpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/networksecuritymanagerpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworksecuritymanagerPolicyPolicyConfigurationOutputReference interface {
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
	RemediationEnabled() interface{}
	SetRemediationEnabled(val interface{})
	RemediationEnabledInput() interface{}
	ResourcesCleanUp() interface{}
	SetResourcesCleanUp(val interface{})
	ResourcesCleanUpInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WafConfig() NetworksecuritymanagerPolicyPolicyConfigurationWafConfigOutputReference
	WafConfigInput() interface{}
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
	PutWafConfig(value *NetworksecuritymanagerPolicyPolicyConfigurationWafConfig)
	ResetRemediationEnabled()
	ResetResourcesCleanUp()
	ResetWafConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworksecuritymanagerPolicyPolicyConfigurationOutputReference
type jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) RemediationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) RemediationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ResourcesCleanUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesCleanUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ResourcesCleanUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesCleanUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) WafConfig() NetworksecuritymanagerPolicyPolicyConfigurationWafConfigOutputReference {
	var returns NetworksecuritymanagerPolicyPolicyConfigurationWafConfigOutputReference
	_jsii_.Get(
		j,
		"wafConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) WafConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafConfigInput",
		&returns,
	)
	return returns
}


func NewNetworksecuritymanagerPolicyPolicyConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworksecuritymanagerPolicyPolicyConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNetworksecuritymanagerPolicyPolicyConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.networksecuritymanagerPolicy.NetworksecuritymanagerPolicyPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworksecuritymanagerPolicyPolicyConfigurationOutputReference_Override(n NetworksecuritymanagerPolicyPolicyConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.networksecuritymanagerPolicy.NetworksecuritymanagerPolicyPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetRemediationEnabled(val interface{}) {
	if err := j.validateSetRemediationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remediationEnabled",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetResourcesCleanUp(val interface{}) {
	if err := j.validateSetResourcesCleanUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesCleanUp",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) PutWafConfig(value *NetworksecuritymanagerPolicyPolicyConfigurationWafConfig) {
	if err := n.validatePutWafConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWafConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ResetRemediationEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetRemediationEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ResetResourcesCleanUp() {
	_jsii_.InvokeVoid(
		n,
		"resetResourcesCleanUp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ResetWafConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetWafConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworksecuritymanagerPolicyPolicyConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

