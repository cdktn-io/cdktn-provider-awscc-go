// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccnetworkfirewallfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccnetworkfirewallfirewallpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference interface {
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
	EnableTlsSessionHolding() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccNetworkfirewallFirewallPolicyFirewallPolicy
	SetInternalValue(val *DataAwsccNetworkfirewallFirewallPolicyFirewallPolicy)
	PolicyVariables() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesOutputReference
	StatefulDefaultActions() *[]*string
	StatefulEngineOptions() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference
	StatefulRuleGroupReferences() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferencesList
	StatelessCustomActions() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionsList
	StatelessDefaultActions() *[]*string
	StatelessFragmentDefaultActions() *[]*string
	StatelessRuleGroupReferences() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferencesList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsInspectionConfigurationArn() *string
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

// The jsii proxy struct for DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference
type jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) EnableTlsSessionHolding() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableTlsSessionHolding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InternalValue() *DataAwsccNetworkfirewallFirewallPolicyFirewallPolicy {
	var returns *DataAwsccNetworkfirewallFirewallPolicyFirewallPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) PolicyVariables() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesOutputReference {
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesOutputReference
	_jsii_.Get(
		j,
		"policyVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulEngineOptions() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference {
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference
	_jsii_.Get(
		j,
		"statefulEngineOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulRuleGroupReferences() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferencesList {
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferencesList
	_jsii_.Get(
		j,
		"statefulRuleGroupReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessCustomActions() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionsList {
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionsList
	_jsii_.Get(
		j,
		"statelessCustomActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessFragmentDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessRuleGroupReferences() DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferencesList {
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferencesList
	_jsii_.Get(
		j,
		"statelessRuleGroupReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) TlsInspectionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationArn",
		&returns,
	)
	return returns
}


func NewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccNetworkfirewallFirewallPolicy.DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference_Override(d DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccNetworkfirewallFirewallPolicy.DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetInternalValue(val *DataAwsccNetworkfirewallFirewallPolicyFirewallPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

