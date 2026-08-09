// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdevopsagentassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdevopsagentassociation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDevopsagentAssociationConfigurationOutputReference interface {
	cdktn.ComplexObject
	Aws() DataAwsccDevopsagentAssociationConfigurationAwsOutputReference
	Azure() DataAwsccDevopsagentAssociationConfigurationAzureOutputReference
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
	Dynatrace() DataAwsccDevopsagentAssociationConfigurationDynatraceOutputReference
	EventChannel() DataAwsccDevopsagentAssociationConfigurationEventChannelOutputReference
	// Experimental.
	Fqn() *string
	GitHub() DataAwsccDevopsagentAssociationConfigurationGitHubOutputReference
	GitLab() DataAwsccDevopsagentAssociationConfigurationGitLabOutputReference
	InternalValue() *DataAwsccDevopsagentAssociationConfiguration
	SetInternalValue(val *DataAwsccDevopsagentAssociationConfiguration)
	McpServer() DataAwsccDevopsagentAssociationConfigurationMcpServerOutputReference
	McpServerDatadog() DataAwsccDevopsagentAssociationConfigurationMcpServerDatadogOutputReference
	McpServerGrafana() DataAwsccDevopsagentAssociationConfigurationMcpServerGrafanaOutputReference
	McpServerNewRelic() DataAwsccDevopsagentAssociationConfigurationMcpServerNewRelicOutputReference
	McpServerSigV4() DataAwsccDevopsagentAssociationConfigurationMcpServerSigV4OutputReference
	McpServerSplunk() DataAwsccDevopsagentAssociationConfigurationMcpServerSplunkOutputReference
	PagerDuty() DataAwsccDevopsagentAssociationConfigurationPagerDutyOutputReference
	ServiceNow() DataAwsccDevopsagentAssociationConfigurationServiceNowOutputReference
	Slack() DataAwsccDevopsagentAssociationConfigurationSlackOutputReference
	SourceAws() DataAwsccDevopsagentAssociationConfigurationSourceAwsOutputReference
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

// The jsii proxy struct for DataAwsccDevopsagentAssociationConfigurationOutputReference
type jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Aws() DataAwsccDevopsagentAssociationConfigurationAwsOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Azure() DataAwsccDevopsagentAssociationConfigurationAzureOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationAzureOutputReference
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Dynatrace() DataAwsccDevopsagentAssociationConfigurationDynatraceOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) EventChannel() DataAwsccDevopsagentAssociationConfigurationEventChannelOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationEventChannelOutputReference
	_jsii_.Get(
		j,
		"eventChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GitHub() DataAwsccDevopsagentAssociationConfigurationGitHubOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationGitHubOutputReference
	_jsii_.Get(
		j,
		"gitHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GitLab() DataAwsccDevopsagentAssociationConfigurationGitLabOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationGitLabOutputReference
	_jsii_.Get(
		j,
		"gitLab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) InternalValue() *DataAwsccDevopsagentAssociationConfiguration {
	var returns *DataAwsccDevopsagentAssociationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServer() DataAwsccDevopsagentAssociationConfigurationMcpServerOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServerDatadog() DataAwsccDevopsagentAssociationConfigurationMcpServerDatadogOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerDatadogOutputReference
	_jsii_.Get(
		j,
		"mcpServerDatadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServerGrafana() DataAwsccDevopsagentAssociationConfigurationMcpServerGrafanaOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerGrafanaOutputReference
	_jsii_.Get(
		j,
		"mcpServerGrafana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServerNewRelic() DataAwsccDevopsagentAssociationConfigurationMcpServerNewRelicOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerNewRelicOutputReference
	_jsii_.Get(
		j,
		"mcpServerNewRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServerSigV4() DataAwsccDevopsagentAssociationConfigurationMcpServerSigV4OutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerSigV4OutputReference
	_jsii_.Get(
		j,
		"mcpServerSigV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) McpServerSplunk() DataAwsccDevopsagentAssociationConfigurationMcpServerSplunkOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationMcpServerSplunkOutputReference
	_jsii_.Get(
		j,
		"mcpServerSplunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) PagerDuty() DataAwsccDevopsagentAssociationConfigurationPagerDutyOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationPagerDutyOutputReference
	_jsii_.Get(
		j,
		"pagerDuty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) ServiceNow() DataAwsccDevopsagentAssociationConfigurationServiceNowOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Slack() DataAwsccDevopsagentAssociationConfigurationSlackOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) SourceAws() DataAwsccDevopsagentAssociationConfigurationSourceAwsOutputReference {
	var returns DataAwsccDevopsagentAssociationConfigurationSourceAwsOutputReference
	_jsii_.Get(
		j,
		"sourceAws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDevopsagentAssociationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDevopsagentAssociationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDevopsagentAssociationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDevopsagentAssociation.DataAwsccDevopsagentAssociationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDevopsagentAssociationConfigurationOutputReference_Override(d DataAwsccDevopsagentAssociationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDevopsagentAssociation.DataAwsccDevopsagentAssociationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference)SetInternalValue(val *DataAwsccDevopsagentAssociationConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDevopsagentAssociationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

