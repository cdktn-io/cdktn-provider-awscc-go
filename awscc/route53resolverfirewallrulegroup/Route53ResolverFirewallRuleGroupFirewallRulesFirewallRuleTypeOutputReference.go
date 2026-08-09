// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/route53resolverfirewallrulegroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference interface {
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
	FirewallAdvancedContentCategory() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategoryOutputReference
	FirewallAdvancedContentCategoryInput() interface{}
	FirewallAdvancedThreatCategory() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategoryOutputReference
	FirewallAdvancedThreatCategoryInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartnerThreatProtection() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtectionOutputReference
	PartnerThreatProtectionInput() interface{}
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
	PutFirewallAdvancedContentCategory(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategory)
	PutFirewallAdvancedThreatCategory(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategory)
	PutPartnerThreatProtection(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtection)
	ResetFirewallAdvancedContentCategory()
	ResetFirewallAdvancedThreatCategory()
	ResetPartnerThreatProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference
type jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) FirewallAdvancedContentCategory() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategoryOutputReference {
	var returns Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategoryOutputReference
	_jsii_.Get(
		j,
		"firewallAdvancedContentCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) FirewallAdvancedContentCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallAdvancedContentCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) FirewallAdvancedThreatCategory() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategoryOutputReference {
	var returns Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategoryOutputReference
	_jsii_.Get(
		j,
		"firewallAdvancedThreatCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) FirewallAdvancedThreatCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallAdvancedThreatCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) PartnerThreatProtection() Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtectionOutputReference {
	var returns Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtectionOutputReference
	_jsii_.Get(
		j,
		"partnerThreatProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) PartnerThreatProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partnerThreatProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoute53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.route53ResolverFirewallRuleGroup.Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoute53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference_Override(r Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.route53ResolverFirewallRuleGroup.Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) PutFirewallAdvancedContentCategory(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategory) {
	if err := r.validatePutFirewallAdvancedContentCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFirewallAdvancedContentCategory",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) PutFirewallAdvancedThreatCategory(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategory) {
	if err := r.validatePutFirewallAdvancedThreatCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFirewallAdvancedThreatCategory",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) PutPartnerThreatProtection(value *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtection) {
	if err := r.validatePutPartnerThreatProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPartnerThreatProtection",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ResetFirewallAdvancedContentCategory() {
	_jsii_.InvokeVoid(
		r,
		"resetFirewallAdvancedContentCategory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ResetFirewallAdvancedThreatCategory() {
	_jsii_.InvokeVoid(
		r,
		"resetFirewallAdvancedThreatCategory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ResetPartnerThreatProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetPartnerThreatProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

