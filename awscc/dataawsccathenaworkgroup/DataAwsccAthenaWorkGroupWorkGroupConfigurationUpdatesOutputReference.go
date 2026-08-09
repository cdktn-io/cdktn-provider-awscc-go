// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccathenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccathenaworkgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference interface {
	cdktn.ComplexObject
	AdditionalConfiguration() *string
	BytesScannedCutoffPerQuery() *float64
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
	CustomerContentEncryptionConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference
	EnforceWorkGroupConfiguration() cdktn.IResolvable
	EngineConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference
	EngineVersion() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference
	ExecutionRole() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdates
	SetInternalValue(val *DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdates)
	ManagedQueryResultsConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference
	MonitoringConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference
	PublishCloudwatchMetricsEnabled() cdktn.IResolvable
	RemoveBytesScannedCutoffPerQuery() cdktn.IResolvable
	RemoveCustomerContentEncryptionConfiguration() cdktn.IResolvable
	RequesterPaysEnabled() cdktn.IResolvable
	ResultConfigurationUpdates() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference
type jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) AdditionalConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) CustomerContentEncryptionConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EnforceWorkGroupConfiguration() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enforceWorkGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference
	_jsii_.Get(
		j,
		"engineConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineVersion() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InternalValue() *DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdates {
	var returns *DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdates
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ManagedQueryResultsConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) MonitoringConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PublishCloudwatchMetricsEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveBytesScannedCutoffPerQuery() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"removeBytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveCustomerContentEncryptionConfiguration() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"removeCustomerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RequesterPaysEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResultConfigurationUpdates() DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference
	_jsii_.Get(
		j,
		"resultConfigurationUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAthenaWorkGroup.DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference_Override(d DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAthenaWorkGroup.DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetInternalValue(val *DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdates) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

