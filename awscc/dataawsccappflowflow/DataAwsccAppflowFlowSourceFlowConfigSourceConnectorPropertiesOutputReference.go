// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccappflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccappflowflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference interface {
	cdktn.ComplexObject
	Amplitude() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference
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
	CustomConnector() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference
	Datadog() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference
	Dynatrace() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference
	// Experimental.
	Fqn() *string
	GoogleAnalytics() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference
	InforNexus() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference
	InternalValue() *DataAwsccAppflowFlowSourceFlowConfigSourceConnectorProperties
	SetInternalValue(val *DataAwsccAppflowFlowSourceFlowConfigSourceConnectorProperties)
	Marketo() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference
	Pardot() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesPardotOutputReference
	S3() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference
	Salesforce() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference
	SapoData() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference
	ServiceNow() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference
	Singular() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference
	Slack() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Trendmicro() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference
	Veeva() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference
	Zendesk() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference
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

// The jsii proxy struct for DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference
type jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Amplitude() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) CustomConnector() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Datadog() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Dynatrace() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GoogleAnalytics() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InforNexus() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InternalValue() *DataAwsccAppflowFlowSourceFlowConfigSourceConnectorProperties {
	var returns *DataAwsccAppflowFlowSourceFlowConfigSourceConnectorProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Marketo() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Pardot() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesPardotOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesPardotOutputReference
	_jsii_.Get(
		j,
		"pardot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) S3() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Salesforce() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SapoData() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ServiceNow() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Singular() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Slack() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Trendmicro() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Veeva() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Zendesk() DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference {
	var returns DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}


func NewDataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppflowFlow.DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference_Override(d DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppflowFlow.DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetInternalValue(val *DataAwsccAppflowFlowSourceFlowConfigSourceConnectorProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

