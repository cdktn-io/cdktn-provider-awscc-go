// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference interface {
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
	Daily() *string
	SetDaily(val *string)
	DailyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Monthly() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference
	MonthlyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Weekly() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference
	WeeklyInput() interface{}
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
	PutMonthly(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthly)
	PutWeekly(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeekly)
	ResetDaily()
	ResetMonthly()
	ResetWeekly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference
type jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Daily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) DailyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Monthly() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) MonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Weekly() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference
	_jsii_.Get(
		j,
		"weekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) WeeklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyInput",
		&returns,
	)
	return returns
}


func NewBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference_Override(b BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetDaily(val *string) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) PutMonthly(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthly) {
	if err := b.validatePutMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMonthly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) PutWeekly(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeekly) {
	if err := b.validatePutWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putWeekly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		b,
		"resetDaily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ResetMonthly() {
	_jsii_.InvokeVoid(
		b,
		"resetMonthly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ResetWeekly() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

