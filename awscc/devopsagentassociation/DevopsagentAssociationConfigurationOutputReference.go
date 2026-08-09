// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/devopsagentassociation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentAssociationConfigurationOutputReference interface {
	cdktn.ComplexObject
	Aws() DevopsagentAssociationConfigurationAwsOutputReference
	AwsInput() interface{}
	Azure() DevopsagentAssociationConfigurationAzureOutputReference
	AzureInput() interface{}
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
	Dynatrace() DevopsagentAssociationConfigurationDynatraceOutputReference
	DynatraceInput() interface{}
	EventChannel() DevopsagentAssociationConfigurationEventChannelOutputReference
	EventChannelInput() interface{}
	// Experimental.
	Fqn() *string
	GitHub() DevopsagentAssociationConfigurationGitHubOutputReference
	GitHubInput() interface{}
	GitLab() DevopsagentAssociationConfigurationGitLabOutputReference
	GitLabInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	McpServer() DevopsagentAssociationConfigurationMcpServerOutputReference
	McpServerDatadog() DevopsagentAssociationConfigurationMcpServerDatadogOutputReference
	McpServerDatadogInput() interface{}
	McpServerGrafana() DevopsagentAssociationConfigurationMcpServerGrafanaOutputReference
	McpServerGrafanaInput() interface{}
	McpServerInput() interface{}
	McpServerNewRelic() DevopsagentAssociationConfigurationMcpServerNewRelicOutputReference
	McpServerNewRelicInput() interface{}
	McpServerSigV4() DevopsagentAssociationConfigurationMcpServerSigV4OutputReference
	McpServerSigV4Input() interface{}
	McpServerSplunk() DevopsagentAssociationConfigurationMcpServerSplunkOutputReference
	McpServerSplunkInput() interface{}
	PagerDuty() DevopsagentAssociationConfigurationPagerDutyOutputReference
	PagerDutyInput() interface{}
	ServiceNow() DevopsagentAssociationConfigurationServiceNowOutputReference
	ServiceNowInput() interface{}
	Slack() DevopsagentAssociationConfigurationSlackOutputReference
	SlackInput() interface{}
	SourceAws() DevopsagentAssociationConfigurationSourceAwsOutputReference
	SourceAwsInput() interface{}
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
	PutAws(value *DevopsagentAssociationConfigurationAws)
	PutAzure(value *DevopsagentAssociationConfigurationAzure)
	PutDynatrace(value *DevopsagentAssociationConfigurationDynatrace)
	PutEventChannel(value *DevopsagentAssociationConfigurationEventChannel)
	PutGitHub(value *DevopsagentAssociationConfigurationGitHub)
	PutGitLab(value *DevopsagentAssociationConfigurationGitLab)
	PutMcpServer(value *DevopsagentAssociationConfigurationMcpServer)
	PutMcpServerDatadog(value *DevopsagentAssociationConfigurationMcpServerDatadog)
	PutMcpServerGrafana(value *DevopsagentAssociationConfigurationMcpServerGrafana)
	PutMcpServerNewRelic(value *DevopsagentAssociationConfigurationMcpServerNewRelic)
	PutMcpServerSigV4(value *DevopsagentAssociationConfigurationMcpServerSigV4)
	PutMcpServerSplunk(value *DevopsagentAssociationConfigurationMcpServerSplunk)
	PutPagerDuty(value *DevopsagentAssociationConfigurationPagerDuty)
	PutServiceNow(value *DevopsagentAssociationConfigurationServiceNow)
	PutSlack(value *DevopsagentAssociationConfigurationSlack)
	PutSourceAws(value *DevopsagentAssociationConfigurationSourceAws)
	ResetAws()
	ResetAzure()
	ResetDynatrace()
	ResetEventChannel()
	ResetGitHub()
	ResetGitLab()
	ResetMcpServer()
	ResetMcpServerDatadog()
	ResetMcpServerGrafana()
	ResetMcpServerNewRelic()
	ResetMcpServerSigV4()
	ResetMcpServerSplunk()
	ResetPagerDuty()
	ResetServiceNow()
	ResetSlack()
	ResetSourceAws()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevopsagentAssociationConfigurationOutputReference
type jsiiProxy_DevopsagentAssociationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Aws() DevopsagentAssociationConfigurationAwsOutputReference {
	var returns DevopsagentAssociationConfigurationAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) AwsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Azure() DevopsagentAssociationConfigurationAzureOutputReference {
	var returns DevopsagentAssociationConfigurationAzureOutputReference
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Dynatrace() DevopsagentAssociationConfigurationDynatraceOutputReference {
	var returns DevopsagentAssociationConfigurationDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) DynatraceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) EventChannel() DevopsagentAssociationConfigurationEventChannelOutputReference {
	var returns DevopsagentAssociationConfigurationEventChannelOutputReference
	_jsii_.Get(
		j,
		"eventChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) EventChannelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GitHub() DevopsagentAssociationConfigurationGitHubOutputReference {
	var returns DevopsagentAssociationConfigurationGitHubOutputReference
	_jsii_.Get(
		j,
		"gitHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GitHubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitHubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GitLab() DevopsagentAssociationConfigurationGitLabOutputReference {
	var returns DevopsagentAssociationConfigurationGitLabOutputReference
	_jsii_.Get(
		j,
		"gitLab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GitLabInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitLabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServer() DevopsagentAssociationConfigurationMcpServerOutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerDatadog() DevopsagentAssociationConfigurationMcpServerDatadogOutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerDatadogOutputReference
	_jsii_.Get(
		j,
		"mcpServerDatadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerDatadogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerDatadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerGrafana() DevopsagentAssociationConfigurationMcpServerGrafanaOutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerGrafanaOutputReference
	_jsii_.Get(
		j,
		"mcpServerGrafana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerGrafanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerGrafanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerNewRelic() DevopsagentAssociationConfigurationMcpServerNewRelicOutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerNewRelicOutputReference
	_jsii_.Get(
		j,
		"mcpServerNewRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerNewRelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerNewRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerSigV4() DevopsagentAssociationConfigurationMcpServerSigV4OutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerSigV4OutputReference
	_jsii_.Get(
		j,
		"mcpServerSigV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerSigV4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerSigV4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerSplunk() DevopsagentAssociationConfigurationMcpServerSplunkOutputReference {
	var returns DevopsagentAssociationConfigurationMcpServerSplunkOutputReference
	_jsii_.Get(
		j,
		"mcpServerSplunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) McpServerSplunkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerSplunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PagerDuty() DevopsagentAssociationConfigurationPagerDutyOutputReference {
	var returns DevopsagentAssociationConfigurationPagerDutyOutputReference
	_jsii_.Get(
		j,
		"pagerDuty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PagerDutyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagerDutyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ServiceNow() DevopsagentAssociationConfigurationServiceNowOutputReference {
	var returns DevopsagentAssociationConfigurationServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ServiceNowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Slack() DevopsagentAssociationConfigurationSlackOutputReference {
	var returns DevopsagentAssociationConfigurationSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) SlackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) SourceAws() DevopsagentAssociationConfigurationSourceAwsOutputReference {
	var returns DevopsagentAssociationConfigurationSourceAwsOutputReference
	_jsii_.Get(
		j,
		"sourceAws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) SourceAwsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAwsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDevopsagentAssociationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DevopsagentAssociationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDevopsagentAssociationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsagentAssociationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentAssociation.DevopsagentAssociationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevopsagentAssociationConfigurationOutputReference_Override(d DevopsagentAssociationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentAssociation.DevopsagentAssociationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevopsagentAssociationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutAws(value *DevopsagentAssociationConfigurationAws) {
	if err := d.validatePutAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAws",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutAzure(value *DevopsagentAssociationConfigurationAzure) {
	if err := d.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzure",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutDynatrace(value *DevopsagentAssociationConfigurationDynatrace) {
	if err := d.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutEventChannel(value *DevopsagentAssociationConfigurationEventChannel) {
	if err := d.validatePutEventChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventChannel",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutGitHub(value *DevopsagentAssociationConfigurationGitHub) {
	if err := d.validatePutGitHubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitHub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutGitLab(value *DevopsagentAssociationConfigurationGitLab) {
	if err := d.validatePutGitLabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitLab",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServer(value *DevopsagentAssociationConfigurationMcpServer) {
	if err := d.validatePutMcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServer",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServerDatadog(value *DevopsagentAssociationConfigurationMcpServerDatadog) {
	if err := d.validatePutMcpServerDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerDatadog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServerGrafana(value *DevopsagentAssociationConfigurationMcpServerGrafana) {
	if err := d.validatePutMcpServerGrafanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerGrafana",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServerNewRelic(value *DevopsagentAssociationConfigurationMcpServerNewRelic) {
	if err := d.validatePutMcpServerNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerNewRelic",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServerSigV4(value *DevopsagentAssociationConfigurationMcpServerSigV4) {
	if err := d.validatePutMcpServerSigV4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerSigV4",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutMcpServerSplunk(value *DevopsagentAssociationConfigurationMcpServerSplunk) {
	if err := d.validatePutMcpServerSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerSplunk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutPagerDuty(value *DevopsagentAssociationConfigurationPagerDuty) {
	if err := d.validatePutPagerDutyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPagerDuty",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutServiceNow(value *DevopsagentAssociationConfigurationServiceNow) {
	if err := d.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutSlack(value *DevopsagentAssociationConfigurationSlack) {
	if err := d.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSlack",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) PutSourceAws(value *DevopsagentAssociationConfigurationSourceAws) {
	if err := d.validatePutSourceAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceAws",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetAws() {
	_jsii_.InvokeVoid(
		d,
		"resetAws",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		d,
		"resetAzure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		d,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetEventChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetEventChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetGitHub() {
	_jsii_.InvokeVoid(
		d,
		"resetGitHub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetGitLab() {
	_jsii_.InvokeVoid(
		d,
		"resetGitLab",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServer() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServerDatadog() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerDatadog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServerGrafana() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerGrafana",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServerNewRelic() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerNewRelic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServerSigV4() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerSigV4",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetMcpServerSplunk() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerSplunk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetPagerDuty() {
	_jsii_.InvokeVoid(
		d,
		"resetPagerDuty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		d,
		"resetSlack",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ResetSourceAws() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceAws",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DevopsagentAssociationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

