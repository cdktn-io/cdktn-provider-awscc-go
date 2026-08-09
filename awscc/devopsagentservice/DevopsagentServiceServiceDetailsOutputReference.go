// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/devopsagentservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentServiceServiceDetailsOutputReference interface {
	cdktn.ComplexObject
	AzureIdentity() DevopsagentServiceServiceDetailsAzureIdentityOutputReference
	AzureIdentityInput() interface{}
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
	Dynatrace() DevopsagentServiceServiceDetailsDynatraceOutputReference
	DynatraceInput() interface{}
	// Experimental.
	Fqn() *string
	GitLab() DevopsagentServiceServiceDetailsGitLabOutputReference
	GitLabInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	McpServer() DevopsagentServiceServiceDetailsMcpServerOutputReference
	McpServerGrafana() DevopsagentServiceServiceDetailsMcpServerGrafanaOutputReference
	McpServerGrafanaInput() interface{}
	McpServerInput() interface{}
	McpServerNewRelic() DevopsagentServiceServiceDetailsMcpServerNewRelicOutputReference
	McpServerNewRelicInput() interface{}
	McpServerSigV4() DevopsagentServiceServiceDetailsMcpServerSigV4OutputReference
	McpServerSigV4Input() interface{}
	McpServerSplunk() DevopsagentServiceServiceDetailsMcpServerSplunkOutputReference
	McpServerSplunkInput() interface{}
	PagerDuty() DevopsagentServiceServiceDetailsPagerDutyOutputReference
	PagerDutyInput() interface{}
	ServiceNow() DevopsagentServiceServiceDetailsServiceNowOutputReference
	ServiceNowInput() interface{}
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
	PutAzureIdentity(value *DevopsagentServiceServiceDetailsAzureIdentity)
	PutDynatrace(value *DevopsagentServiceServiceDetailsDynatrace)
	PutGitLab(value *DevopsagentServiceServiceDetailsGitLab)
	PutMcpServer(value *DevopsagentServiceServiceDetailsMcpServer)
	PutMcpServerGrafana(value *DevopsagentServiceServiceDetailsMcpServerGrafana)
	PutMcpServerNewRelic(value *DevopsagentServiceServiceDetailsMcpServerNewRelic)
	PutMcpServerSigV4(value *DevopsagentServiceServiceDetailsMcpServerSigV4)
	PutMcpServerSplunk(value *DevopsagentServiceServiceDetailsMcpServerSplunk)
	PutPagerDuty(value *DevopsagentServiceServiceDetailsPagerDuty)
	PutServiceNow(value *DevopsagentServiceServiceDetailsServiceNow)
	ResetAzureIdentity()
	ResetDynatrace()
	ResetGitLab()
	ResetMcpServer()
	ResetMcpServerGrafana()
	ResetMcpServerNewRelic()
	ResetMcpServerSigV4()
	ResetMcpServerSplunk()
	ResetPagerDuty()
	ResetServiceNow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevopsagentServiceServiceDetailsOutputReference
type jsiiProxy_DevopsagentServiceServiceDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) AzureIdentity() DevopsagentServiceServiceDetailsAzureIdentityOutputReference {
	var returns DevopsagentServiceServiceDetailsAzureIdentityOutputReference
	_jsii_.Get(
		j,
		"azureIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) AzureIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) Dynatrace() DevopsagentServiceServiceDetailsDynatraceOutputReference {
	var returns DevopsagentServiceServiceDetailsDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) DynatraceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GitLab() DevopsagentServiceServiceDetailsGitLabOutputReference {
	var returns DevopsagentServiceServiceDetailsGitLabOutputReference
	_jsii_.Get(
		j,
		"gitLab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GitLabInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitLabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServer() DevopsagentServiceServiceDetailsMcpServerOutputReference {
	var returns DevopsagentServiceServiceDetailsMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerGrafana() DevopsagentServiceServiceDetailsMcpServerGrafanaOutputReference {
	var returns DevopsagentServiceServiceDetailsMcpServerGrafanaOutputReference
	_jsii_.Get(
		j,
		"mcpServerGrafana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerGrafanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerGrafanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerNewRelic() DevopsagentServiceServiceDetailsMcpServerNewRelicOutputReference {
	var returns DevopsagentServiceServiceDetailsMcpServerNewRelicOutputReference
	_jsii_.Get(
		j,
		"mcpServerNewRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerNewRelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerNewRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerSigV4() DevopsagentServiceServiceDetailsMcpServerSigV4OutputReference {
	var returns DevopsagentServiceServiceDetailsMcpServerSigV4OutputReference
	_jsii_.Get(
		j,
		"mcpServerSigV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerSigV4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerSigV4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerSplunk() DevopsagentServiceServiceDetailsMcpServerSplunkOutputReference {
	var returns DevopsagentServiceServiceDetailsMcpServerSplunkOutputReference
	_jsii_.Get(
		j,
		"mcpServerSplunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) McpServerSplunkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerSplunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PagerDuty() DevopsagentServiceServiceDetailsPagerDutyOutputReference {
	var returns DevopsagentServiceServiceDetailsPagerDutyOutputReference
	_jsii_.Get(
		j,
		"pagerDuty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PagerDutyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagerDutyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ServiceNow() DevopsagentServiceServiceDetailsServiceNowOutputReference {
	var returns DevopsagentServiceServiceDetailsServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ServiceNowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDevopsagentServiceServiceDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DevopsagentServiceServiceDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDevopsagentServiceServiceDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsagentServiceServiceDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevopsagentServiceServiceDetailsOutputReference_Override(d DevopsagentServiceServiceDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutAzureIdentity(value *DevopsagentServiceServiceDetailsAzureIdentity) {
	if err := d.validatePutAzureIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureIdentity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutDynatrace(value *DevopsagentServiceServiceDetailsDynatrace) {
	if err := d.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutGitLab(value *DevopsagentServiceServiceDetailsGitLab) {
	if err := d.validatePutGitLabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitLab",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutMcpServer(value *DevopsagentServiceServiceDetailsMcpServer) {
	if err := d.validatePutMcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServer",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutMcpServerGrafana(value *DevopsagentServiceServiceDetailsMcpServerGrafana) {
	if err := d.validatePutMcpServerGrafanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerGrafana",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutMcpServerNewRelic(value *DevopsagentServiceServiceDetailsMcpServerNewRelic) {
	if err := d.validatePutMcpServerNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerNewRelic",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutMcpServerSigV4(value *DevopsagentServiceServiceDetailsMcpServerSigV4) {
	if err := d.validatePutMcpServerSigV4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerSigV4",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutMcpServerSplunk(value *DevopsagentServiceServiceDetailsMcpServerSplunk) {
	if err := d.validatePutMcpServerSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMcpServerSplunk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutPagerDuty(value *DevopsagentServiceServiceDetailsPagerDuty) {
	if err := d.validatePutPagerDutyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPagerDuty",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) PutServiceNow(value *DevopsagentServiceServiceDetailsServiceNow) {
	if err := d.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetAzureIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		d,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetGitLab() {
	_jsii_.InvokeVoid(
		d,
		"resetGitLab",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetMcpServer() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetMcpServerGrafana() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerGrafana",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetMcpServerNewRelic() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerNewRelic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetMcpServerSigV4() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerSigV4",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetMcpServerSplunk() {
	_jsii_.InvokeVoid(
		d,
		"resetMcpServerSplunk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetPagerDuty() {
	_jsii_.InvokeVoid(
		d,
		"resetPagerDuty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DevopsagentServiceServiceDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

