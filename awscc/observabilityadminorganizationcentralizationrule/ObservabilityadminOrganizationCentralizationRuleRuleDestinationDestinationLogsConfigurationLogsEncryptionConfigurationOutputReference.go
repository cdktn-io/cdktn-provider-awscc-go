// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/observabilityadminorganizationcentralizationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference interface {
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
	EncryptionConflictResolutionStrategy() *string
	SetEncryptionConflictResolutionStrategy(val *string)
	EncryptionConflictResolutionStrategyInput() *string
	EncryptionScope() *string
	SetEncryptionScope(val *string)
	EncryptionScopeInput() *string
	EncryptionStrategy() *string
	SetEncryptionStrategy(val *string)
	EncryptionStrategyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	ResetEncryptionConflictResolutionStrategy()
	ResetEncryptionScope()
	ResetEncryptionStrategy()
	ResetKmsKeyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference
type jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionConflictResolutionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionConflictResolutionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionConflictResolutionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionConflictResolutionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) EncryptionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference_Override(o ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetEncryptionConflictResolutionStrategy(val *string) {
	if err := j.validateSetEncryptionConflictResolutionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionConflictResolutionStrategy",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetEncryptionScope(val *string) {
	if err := j.validateSetEncryptionScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionScope",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetEncryptionStrategy(val *string) {
	if err := j.validateSetEncryptionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionStrategy",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ResetEncryptionConflictResolutionStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionConflictResolutionStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ResetEncryptionScope() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionScope",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ResetEncryptionStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		o,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

