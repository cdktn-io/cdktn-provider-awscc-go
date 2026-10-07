// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configorganizationconfigrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/configorganizationconfigrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference interface {
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
	DebugLogDeliveryAccounts() *[]*string
	SetDebugLogDeliveryAccounts(val *[]*string)
	DebugLogDeliveryAccountsInput() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InputParameters() *string
	SetInputParameters(val *string)
	InputParametersInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OrganizationConfigRuleTriggerTypes() *[]*string
	SetOrganizationConfigRuleTriggerTypes(val *[]*string)
	OrganizationConfigRuleTriggerTypesInput() *[]*string
	PolicyText() *string
	SetPolicyText(val *string)
	PolicyTextInput() *string
	ResourceIdScope() *string
	SetResourceIdScope(val *string)
	ResourceIdScopeInput() *string
	ResourceTypesScope() *[]*string
	SetResourceTypesScope(val *[]*string)
	ResourceTypesScopeInput() *[]*string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	TagKeyScope() *string
	SetTagKeyScope(val *string)
	TagKeyScopeInput() *string
	TagValueScope() *string
	SetTagValueScope(val *string)
	TagValueScopeInput() *string
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
	ResetDebugLogDeliveryAccounts()
	ResetDescription()
	ResetInputParameters()
	ResetOrganizationConfigRuleTriggerTypes()
	ResetPolicyText()
	ResetResourceIdScope()
	ResetResourceTypesScope()
	ResetRuntime()
	ResetTagKeyScope()
	ResetTagValueScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference
type jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) DebugLogDeliveryAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"debugLogDeliveryAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) DebugLogDeliveryAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"debugLogDeliveryAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) OrganizationConfigRuleTriggerTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationConfigRuleTriggerTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) OrganizationConfigRuleTriggerTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationConfigRuleTriggerTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) PolicyText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) PolicyTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.configOrganizationConfigRule.ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference_Override(c ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.configOrganizationConfigRule.ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetDebugLogDeliveryAccounts(val *[]*string) {
	if err := j.validateSetDebugLogDeliveryAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugLogDeliveryAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetInputParameters(val *string) {
	if err := j.validateSetInputParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetOrganizationConfigRuleTriggerTypes(val *[]*string) {
	if err := j.validateSetOrganizationConfigRuleTriggerTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationConfigRuleTriggerTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetPolicyText(val *string) {
	if err := j.validateSetPolicyTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyText",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetResourceIdScope(val *string) {
	if err := j.validateSetResourceIdScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetResourceTypesScope(val *[]*string) {
	if err := j.validateSetResourceTypesScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetTagKeyScope(val *string) {
	if err := j.validateSetTagKeyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetTagValueScope(val *string) {
	if err := j.validateSetTagValueScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetDebugLogDeliveryAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetDebugLogDeliveryAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetOrganizationConfigRuleTriggerTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganizationConfigRuleTriggerTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetPolicyText() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyText",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationCustomPolicyRuleMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

