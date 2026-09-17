// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/chimemediainsightspipelineconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueDetectionConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfigurationOutputReference
	IssueDetectionConfigurationInput() interface{}
	KeywordMatchConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfigurationOutputReference
	KeywordMatchConfigurationInput() interface{}
	SentimentConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfigurationOutputReference
	SentimentConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutIssueDetectionConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfiguration)
	PutKeywordMatchConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfiguration)
	PutSentimentConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration)
	ResetIssueDetectionConfiguration()
	ResetKeywordMatchConfiguration()
	ResetSentimentConfiguration()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference
type jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) IssueDetectionConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfigurationOutputReference {
	var returns ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfigurationOutputReference
	_jsii_.Get(
		j,
		"issueDetectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) IssueDetectionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueDetectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) KeywordMatchConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfigurationOutputReference {
	var returns ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"keywordMatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) KeywordMatchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordMatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) SentimentConfiguration() ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfigurationOutputReference {
	var returns ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfigurationOutputReference
	_jsii_.Get(
		j,
		"sentimentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) SentimentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference_Override(c ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) PutIssueDetectionConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfiguration) {
	if err := c.validatePutIssueDetectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIssueDetectionConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) PutKeywordMatchConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfiguration) {
	if err := c.validatePutKeywordMatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKeywordMatchConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) PutSentimentConfiguration(value *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration) {
	if err := c.validatePutSentimentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSentimentConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ResetIssueDetectionConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetIssueDetectionConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ResetKeywordMatchConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetKeywordMatchConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ResetSentimentConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSentimentConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

