// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdevopsagentservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdevopsagentservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference interface {
	cdktn.ComplexObject
	AzureIdentity() DataAwsccDevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference
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
	Dynatrace() DataAwsccDevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference
	// Experimental.
	Fqn() *string
	GitLab() DataAwsccDevopsagentServiceAdditionalServiceDetailsGitLabOutputReference
	InternalValue() *DataAwsccDevopsagentServiceAdditionalServiceDetails
	SetInternalValue(val *DataAwsccDevopsagentServiceAdditionalServiceDetails)
	McpServer() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference
	McpServerGrafana() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference
	McpServerNewRelic() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference
	McpServerSigV4() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference
	McpServerSplunk() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference
	PagerDuty() DataAwsccDevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference
	ServiceNow() DataAwsccDevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference
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

// The jsii proxy struct for DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference
type jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) AzureIdentity() DataAwsccDevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsAzureIdentityOutputReference
	_jsii_.Get(
		j,
		"azureIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) Dynatrace() DataAwsccDevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GitLab() DataAwsccDevopsagentServiceAdditionalServiceDetailsGitLabOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsGitLabOutputReference
	_jsii_.Get(
		j,
		"gitLab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) InternalValue() *DataAwsccDevopsagentServiceAdditionalServiceDetails {
	var returns *DataAwsccDevopsagentServiceAdditionalServiceDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) McpServer() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerOutputReference
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerGrafana() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerGrafanaOutputReference
	_jsii_.Get(
		j,
		"mcpServerGrafana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerNewRelic() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerNewRelicOutputReference
	_jsii_.Get(
		j,
		"mcpServerNewRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerSigV4() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSigV4OutputReference
	_jsii_.Get(
		j,
		"mcpServerSigV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) McpServerSplunk() DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsMcpServerSplunkOutputReference
	_jsii_.Get(
		j,
		"mcpServerSplunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) PagerDuty() DataAwsccDevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsPagerDutyOutputReference
	_jsii_.Get(
		j,
		"pagerDuty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) ServiceNow() DataAwsccDevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference {
	var returns DataAwsccDevopsagentServiceAdditionalServiceDetailsServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDevopsagentService.DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference_Override(d DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDevopsagentService.DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference)SetInternalValue(val *DataAwsccDevopsagentServiceAdditionalServiceDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDevopsagentServiceAdditionalServiceDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

