// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appflowflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppflowFlowTasksConnectorOperatorOutputReference interface {
	cdktn.ComplexObject
	Amplitude() *string
	SetAmplitude(val *string)
	AmplitudeInput() *string
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
	CustomConnector() *string
	SetCustomConnector(val *string)
	CustomConnectorInput() *string
	Datadog() *string
	SetDatadog(val *string)
	DatadogInput() *string
	Dynatrace() *string
	SetDynatrace(val *string)
	DynatraceInput() *string
	// Experimental.
	Fqn() *string
	GoogleAnalytics() *string
	SetGoogleAnalytics(val *string)
	GoogleAnalyticsInput() *string
	InforNexus() *string
	SetInforNexus(val *string)
	InforNexusInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Marketo() *string
	SetMarketo(val *string)
	MarketoInput() *string
	Pardot() *string
	SetPardot(val *string)
	PardotInput() *string
	S3() *string
	SetS3(val *string)
	S3Input() *string
	Salesforce() *string
	SetSalesforce(val *string)
	SalesforceInput() *string
	SapoData() *string
	SetSapoData(val *string)
	SapoDataInput() *string
	ServiceNow() *string
	SetServiceNow(val *string)
	ServiceNowInput() *string
	Singular() *string
	SetSingular(val *string)
	SingularInput() *string
	Slack() *string
	SetSlack(val *string)
	SlackInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Trendmicro() *string
	SetTrendmicro(val *string)
	TrendmicroInput() *string
	Veeva() *string
	SetVeeva(val *string)
	VeevaInput() *string
	Zendesk() *string
	SetZendesk(val *string)
	ZendeskInput() *string
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
	ResetAmplitude()
	ResetCustomConnector()
	ResetDatadog()
	ResetDynatrace()
	ResetGoogleAnalytics()
	ResetInforNexus()
	ResetMarketo()
	ResetPardot()
	ResetS3()
	ResetSalesforce()
	ResetSapoData()
	ResetServiceNow()
	ResetSingular()
	ResetSlack()
	ResetTrendmicro()
	ResetVeeva()
	ResetZendesk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowTasksConnectorOperatorOutputReference
type jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Amplitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) AmplitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) CustomConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) CustomConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Datadog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) DatadogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Dynatrace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) DynatraceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GoogleAnalytics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GoogleAnalyticsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) InforNexus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) InforNexusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Marketo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) MarketoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Pardot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pardot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) PardotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pardotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) S3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) S3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Salesforce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) SalesforceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) SapoData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) SapoDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ServiceNow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ServiceNowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Singular() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) SingularInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Slack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) SlackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Trendmicro() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) TrendmicroInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Veeva() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) VeevaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Zendesk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ZendeskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowTasksConnectorOperatorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppflowFlowTasksConnectorOperatorOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowTasksConnectorOperatorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appflowFlow.AppflowFlowTasksConnectorOperatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowTasksConnectorOperatorOutputReference_Override(a AppflowFlowTasksConnectorOperatorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appflowFlow.AppflowFlowTasksConnectorOperatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetAmplitude(val *string) {
	if err := j.validateSetAmplitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amplitude",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetCustomConnector(val *string) {
	if err := j.validateSetCustomConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customConnector",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetDatadog(val *string) {
	if err := j.validateSetDatadogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadog",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetDynatrace(val *string) {
	if err := j.validateSetDynatraceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynatrace",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetGoogleAnalytics(val *string) {
	if err := j.validateSetGoogleAnalyticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAnalytics",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetInforNexus(val *string) {
	if err := j.validateSetInforNexusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inforNexus",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetMarketo(val *string) {
	if err := j.validateSetMarketoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketo",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetPardot(val *string) {
	if err := j.validateSetPardotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pardot",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetS3(val *string) {
	if err := j.validateSetS3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetSalesforce(val *string) {
	if err := j.validateSetSalesforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforce",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetSapoData(val *string) {
	if err := j.validateSetSapoDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapoData",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetServiceNow(val *string) {
	if err := j.validateSetServiceNowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNow",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetSingular(val *string) {
	if err := j.validateSetSingularParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singular",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetSlack(val *string) {
	if err := j.validateSetSlackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slack",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetTrendmicro(val *string) {
	if err := j.validateSetTrendmicroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trendmicro",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetVeeva(val *string) {
	if err := j.validateSetVeevaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"veeva",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference)SetZendesk(val *string) {
	if err := j.validateSetZendeskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zendesk",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetPardot() {
	_jsii_.InvokeVoid(
		a,
		"resetPardot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowTasksConnectorOperatorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

