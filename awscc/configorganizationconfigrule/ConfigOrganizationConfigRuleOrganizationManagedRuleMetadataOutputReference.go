// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configorganizationconfigrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/configorganizationconfigrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference interface {
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
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	MaximumExecutionFrequencyInput() *string
	ResourceIdScope() *string
	SetResourceIdScope(val *string)
	ResourceIdScopeInput() *string
	ResourceTypesScope() *[]*string
	SetResourceTypesScope(val *[]*string)
	ResourceTypesScopeInput() *[]*string
	RuleIdentifier() *string
	SetRuleIdentifier(val *string)
	RuleIdentifierInput() *string
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
	ResetDescription()
	ResetInputParameters()
	ResetMaximumExecutionFrequency()
	ResetResourceIdScope()
	ResetResourceTypesScope()
	ResetRuleIdentifier()
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

// The jsii proxy struct for ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference
type jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) RuleIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) RuleIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.configOrganizationConfigRule.ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference_Override(c ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.configOrganizationConfigRule.ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetInputParameters(val *string) {
	if err := j.validateSetInputParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetMaximumExecutionFrequency(val *string) {
	if err := j.validateSetMaximumExecutionFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetResourceIdScope(val *string) {
	if err := j.validateSetResourceIdScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetResourceTypesScope(val *[]*string) {
	if err := j.validateSetResourceTypesScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetRuleIdentifier(val *string) {
	if err := j.validateSetRuleIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleIdentifier",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetTagKeyScope(val *string) {
	if err := j.validateSetTagKeyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetTagValueScope(val *string) {
	if err := j.validateSetTagValueScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetRuleIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetRuleIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigOrganizationConfigRuleOrganizationManagedRuleMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

