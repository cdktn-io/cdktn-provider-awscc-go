// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/observabilityadminorganizationcentralizationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference interface {
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
	DestinationRoleArn() *string
	SetDestinationRoleArn(val *string)
	DestinationRoleArnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TagConflictResolutionStrategy() *string
	SetTagConflictResolutionStrategy(val *string)
	TagConflictResolutionStrategyInput() *string
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
	ResetDestinationRoleArn()
	ResetTagConflictResolutionStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference
type jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) DestinationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) DestinationRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) TagConflictResolutionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagConflictResolutionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) TagConflictResolutionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagConflictResolutionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference_Override(o ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetDestinationRoleArn(val *string) {
	if err := j.validateSetDestinationRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRoleArn",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetTagConflictResolutionStrategy(val *string) {
	if err := j.validateSetTagConflictResolutionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagConflictResolutionStrategy",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ResetDestinationRoleArn() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationRoleArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ResetTagConflictResolutionStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetTagConflictResolutionStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

