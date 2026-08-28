// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/observabilityadminorganizationcentralizationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference interface {
	cdktn.ComplexObject
	BackupConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationBackupConfigurationOutputReference
	BackupConfigurationInput() interface{}
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
	LogGroupNameConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogGroupNameConfigurationOutputReference
	LogGroupNameConfigurationInput() interface{}
	LogsEncryptionConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference
	LogsEncryptionConfigurationInput() interface{}
	TagPropagationConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference
	TagPropagationConfigurationInput() interface{}
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
	PutBackupConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationBackupConfiguration)
	PutLogGroupNameConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration)
	PutLogsEncryptionConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration)
	PutTagPropagationConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration)
	ResetBackupConfiguration()
	ResetLogGroupNameConfiguration()
	ResetLogsEncryptionConfiguration()
	ResetTagPropagationConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference
type jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) BackupConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationBackupConfigurationOutputReference {
	var returns ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationBackupConfigurationOutputReference
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) BackupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) LogGroupNameConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogGroupNameConfigurationOutputReference {
	var returns ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogGroupNameConfigurationOutputReference
	_jsii_.Get(
		j,
		"logGroupNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) LogGroupNameConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logGroupNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) LogsEncryptionConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference {
	var returns ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"logsEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) LogsEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) TagPropagationConfiguration() ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference {
	var returns ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfigurationOutputReference
	_jsii_.Get(
		j,
		"tagPropagationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) TagPropagationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPropagationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference_Override(o ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationCentralizationRule.ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) PutBackupConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationBackupConfiguration) {
	if err := o.validatePutBackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBackupConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) PutLogGroupNameConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration) {
	if err := o.validatePutLogGroupNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogGroupNameConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) PutLogsEncryptionConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration) {
	if err := o.validatePutLogsEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogsEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) PutTagPropagationConfiguration(value *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration) {
	if err := o.validatePutTagPropagationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTagPropagationConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ResetBackupConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ResetLogGroupNameConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetLogGroupNameConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ResetLogsEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetLogsEncryptionConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ResetTagPropagationConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetTagPropagationConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

