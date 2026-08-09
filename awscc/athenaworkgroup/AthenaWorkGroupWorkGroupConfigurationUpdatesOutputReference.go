// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/athenaworkgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference interface {
	cdktn.ComplexObject
	AdditionalConfiguration() *string
	SetAdditionalConfiguration(val *string)
	AdditionalConfigurationInput() *string
	BytesScannedCutoffPerQuery() *float64
	SetBytesScannedCutoffPerQuery(val *float64)
	BytesScannedCutoffPerQueryInput() *float64
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
	CustomerContentEncryptionConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference
	CustomerContentEncryptionConfigurationInput() interface{}
	EnforceWorkGroupConfiguration() interface{}
	SetEnforceWorkGroupConfiguration(val interface{})
	EnforceWorkGroupConfigurationInput() interface{}
	EngineConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference
	EngineConfigurationInput() interface{}
	EngineVersion() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference
	EngineVersionInput() interface{}
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedQueryResultsConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference
	ManagedQueryResultsConfigurationInput() interface{}
	MonitoringConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference
	MonitoringConfigurationInput() interface{}
	PublishCloudwatchMetricsEnabled() interface{}
	SetPublishCloudwatchMetricsEnabled(val interface{})
	PublishCloudwatchMetricsEnabledInput() interface{}
	RemoveBytesScannedCutoffPerQuery() interface{}
	SetRemoveBytesScannedCutoffPerQuery(val interface{})
	RemoveBytesScannedCutoffPerQueryInput() interface{}
	RemoveCustomerContentEncryptionConfiguration() interface{}
	SetRemoveCustomerContentEncryptionConfiguration(val interface{})
	RemoveCustomerContentEncryptionConfigurationInput() interface{}
	RequesterPaysEnabled() interface{}
	SetRequesterPaysEnabled(val interface{})
	RequesterPaysEnabledInput() interface{}
	ResultConfigurationUpdates() AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference
	ResultConfigurationUpdatesInput() interface{}
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
	PutCustomerContentEncryptionConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfiguration)
	PutEngineConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfiguration)
	PutEngineVersion(value *AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersion)
	PutManagedQueryResultsConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfiguration)
	PutMonitoringConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfiguration)
	PutResultConfigurationUpdates(value *AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdates)
	ResetAdditionalConfiguration()
	ResetBytesScannedCutoffPerQuery()
	ResetCustomerContentEncryptionConfiguration()
	ResetEnforceWorkGroupConfiguration()
	ResetEngineConfiguration()
	ResetEngineVersion()
	ResetExecutionRole()
	ResetManagedQueryResultsConfiguration()
	ResetMonitoringConfiguration()
	ResetPublishCloudwatchMetricsEnabled()
	ResetRemoveBytesScannedCutoffPerQuery()
	ResetRemoveCustomerContentEncryptionConfiguration()
	ResetRequesterPaysEnabled()
	ResetResultConfigurationUpdates()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference
type jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) AdditionalConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) AdditionalConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) CustomerContentEncryptionConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) CustomerContentEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EnforceWorkGroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EnforceWorkGroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkGroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference
	_jsii_.Get(
		j,
		"engineConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineVersion() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersionOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) EngineVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ManagedQueryResultsConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ManagedQueryResultsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedQueryResultsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) MonitoringConfiguration() AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) MonitoringConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveBytesScannedCutoffPerQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeBytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveBytesScannedCutoffPerQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeBytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveCustomerContentEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeCustomerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RemoveCustomerContentEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeCustomerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResultConfigurationUpdates() AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesOutputReference
	_jsii_.Get(
		j,
		"resultConfigurationUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResultConfigurationUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resultConfigurationUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference_Override(a AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetAdditionalConfiguration(val *string) {
	if err := j.validateSetAdditionalConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetEnforceWorkGroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkGroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkGroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetRemoveBytesScannedCutoffPerQuery(val interface{}) {
	if err := j.validateSetRemoveBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeBytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetRemoveCustomerContentEncryptionConfiguration(val interface{}) {
	if err := j.validateSetRemoveCustomerContentEncryptionConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeCustomerContentEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutCustomerContentEncryptionConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfiguration) {
	if err := a.validatePutCustomerContentEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerContentEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutEngineConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfiguration) {
	if err := a.validatePutEngineConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutEngineVersion(value *AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersion) {
	if err := a.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutManagedQueryResultsConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfiguration) {
	if err := a.validatePutManagedQueryResultsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedQueryResultsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutMonitoringConfiguration(value *AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfiguration) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) PutResultConfigurationUpdates(value *AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdates) {
	if err := a.validatePutResultConfigurationUpdatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResultConfigurationUpdates",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetAdditionalConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetEnforceWorkGroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceWorkGroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetEngineConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetManagedQueryResultsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedQueryResultsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetRemoveBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoveBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetRemoveCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoveCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ResetResultConfigurationUpdates() {
	_jsii_.InvokeVoid(
		a,
		"resetResultConfigurationUpdates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

