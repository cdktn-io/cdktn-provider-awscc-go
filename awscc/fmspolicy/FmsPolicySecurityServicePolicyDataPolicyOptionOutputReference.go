// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fmspolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference interface {
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
	NetworkAclCommonPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyOutputReference
	NetworkAclCommonPolicyInput() interface{}
	NetworkFirewallPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicyOutputReference
	NetworkFirewallPolicyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThirdPartyFirewallPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyOutputReference
	ThirdPartyFirewallPolicyInput() interface{}
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
	PutNetworkAclCommonPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy)
	PutNetworkFirewallPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy)
	PutThirdPartyFirewallPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy)
	ResetNetworkAclCommonPolicy()
	ResetNetworkFirewallPolicy()
	ResetThirdPartyFirewallPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference
type jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) NetworkAclCommonPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyOutputReference {
	var returns FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyOutputReference
	_jsii_.Get(
		j,
		"networkAclCommonPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) NetworkAclCommonPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAclCommonPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) NetworkFirewallPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicyOutputReference {
	var returns FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicyOutputReference
	_jsii_.Get(
		j,
		"networkFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) NetworkFirewallPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ThirdPartyFirewallPolicy() FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyOutputReference {
	var returns FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyOutputReference
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ThirdPartyFirewallPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicyInput",
		&returns,
	)
	return returns
}


func NewFmsPolicySecurityServicePolicyDataPolicyOptionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference {
	_init_.Initialize()

	if err := validateNewFmsPolicySecurityServicePolicyDataPolicyOptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fmsPolicy.FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFmsPolicySecurityServicePolicyDataPolicyOptionOutputReference_Override(f FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fmsPolicy.FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) PutNetworkAclCommonPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy) {
	if err := f.validatePutNetworkAclCommonPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNetworkAclCommonPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) PutNetworkFirewallPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy) {
	if err := f.validatePutNetworkFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNetworkFirewallPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) PutThirdPartyFirewallPolicy(value *FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy) {
	if err := f.validatePutThirdPartyFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putThirdPartyFirewallPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ResetNetworkAclCommonPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetNetworkAclCommonPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ResetNetworkFirewallPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetNetworkFirewallPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ResetThirdPartyFirewallPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetThirdPartyFirewallPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataPolicyOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

