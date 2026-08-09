// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccappflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccappflowconnectorprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference interface {
	cdktn.ComplexObject
	Amplitude() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference
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
	CustomConnector() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference
	Datadog() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference
	Dynatrace() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference
	// Experimental.
	Fqn() *string
	GoogleAnalytics() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference
	InforNexus() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference
	InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials
	SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials)
	Marketo() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference
	Pardot() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsPardotOutputReference
	Redshift() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference
	Salesforce() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference
	SapoData() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference
	ServiceNow() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference
	Singular() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference
	Slack() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference
	Snowflake() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Trendmicro() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference
	Veeva() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference
	Zendesk() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference
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

// The jsii proxy struct for DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference
type jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Amplitude() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) CustomConnector() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Datadog() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Dynatrace() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GoogleAnalytics() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InforNexus() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials {
	var returns *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Marketo() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Pardot() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsPardotOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsPardotOutputReference
	_jsii_.Get(
		j,
		"pardot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Redshift() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Salesforce() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SapoData() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ServiceNow() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Singular() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Slack() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Snowflake() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Trendmicro() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Veeva() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Zendesk() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}


func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference_Override(d DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

