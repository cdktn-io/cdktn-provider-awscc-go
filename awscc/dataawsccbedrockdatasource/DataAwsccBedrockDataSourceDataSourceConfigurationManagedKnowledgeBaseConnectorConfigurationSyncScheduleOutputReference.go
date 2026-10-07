// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule
	SetInternalValue(val *DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule)
	Monthly() DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Weekly() DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference
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

// The jsii proxy struct for DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference
type jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Daily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InternalValue() *DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule {
	var returns *DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Monthly() DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference {
	var returns DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyOutputReference
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Weekly() DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference {
	var returns DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeeklyOutputReference
	_jsii_.Get(
		j,
		"weekly",
		&returns,
	)
	return returns
}


func NewDataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockDataSource.DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference_Override(d DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockDataSource.DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetInternalValue(val *DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

