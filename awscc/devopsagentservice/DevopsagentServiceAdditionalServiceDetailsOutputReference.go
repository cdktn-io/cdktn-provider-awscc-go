// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/devopsagentservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentServiceAdditionalServiceDetailsOutputReference interface {
	cdktn.ComplexObject
	AzureIdentity() DevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference
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
	Dynatrace() DevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference
	// Experimental.
	Fqn() *string
	GitLab() DevopsagentServiceAdditionalServiceDetailsGitLabOutputReference
	InternalValue() *DevopsagentServiceAdditionalServiceDetails
	SetInternalValue(val *DevopsagentServiceAdditionalServiceDetails)
	McpServer() DevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference
	McpServerGrafana() DevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference
	McpServerNewRelic() DevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference
	McpServerSigV4() DevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference
	McpServerSplunk() DevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference
	PagerDuty() DevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference
	ServiceNow() DevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference
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

// The jsii proxy struct for DevopsagentServiceAdditionalServiceDetailsOutputReference
type jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) AzureIdentity() DevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference
	_jsii_.Get(
		j,
		"azureIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) Dynatrace() DevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GitLab() DevopsagentServiceAdditionalServiceDetailsGitLabOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsGitLabOutputReference
	_jsii_.Get(
		j,
		"gitLab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) InternalValue() *DevopsagentServiceAdditionalServiceDetails {
	var returns *DevopsagentServiceAdditionalServiceDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) McpServer() DevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerGrafana() DevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference
	_jsii_.Get(
		j,
		"mcpServerGrafana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerNewRelic() DevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference
	_jsii_.Get(
		j,
		"mcpServerNewRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerSigV4() DevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference
	_jsii_.Get(
		j,
		"mcpServerSigV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerSplunk() DevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference
	_jsii_.Get(
		j,
		"mcpServerSplunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) PagerDuty() DevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference
	_jsii_.Get(
		j,
		"pagerDuty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) ServiceNow() DevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference {
	var returns DevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDevopsagentServiceAdditionalServiceDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DevopsagentServiceAdditionalServiceDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDevopsagentServiceAdditionalServiceDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceAdditionalServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevopsagentServiceAdditionalServiceDetailsOutputReference_Override(d DevopsagentServiceAdditionalServiceDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.devopsagentService.DevopsagentServiceAdditionalServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference)SetInternalValue(val *DevopsagentServiceAdditionalServiceDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DevopsagentServiceAdditionalServiceDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

