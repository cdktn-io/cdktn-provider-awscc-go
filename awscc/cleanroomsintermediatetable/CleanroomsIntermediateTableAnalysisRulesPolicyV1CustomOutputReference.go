// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsintermediatetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference interface {
	cdktn.ComplexObject
	AdditionalAnalyses() *string
	SetAdditionalAnalyses(val *string)
	AdditionalAnalysesInput() *string
	AggregationThresholds() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsList
	AggregationThresholdsInput() interface{}
	AllowedAnalyses() *[]*string
	SetAllowedAnalyses(val *[]*string)
	AllowedAnalysesInput() *[]*string
	AllowedAnalysisProviders() *[]*string
	SetAllowedAnalysisProviders(val *[]*string)
	AllowedAnalysisProvidersInput() *[]*string
	AllowedResultReceivers() *[]*string
	SetAllowedResultReceivers(val *[]*string)
	AllowedResultReceiversInput() *[]*string
	ComparisonControls() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference
	ComparisonControlsInput() interface{}
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
	DifferentialPrivacy() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference
	DifferentialPrivacyInput() interface{}
	DisallowedOutputColumns() *[]*string
	SetDisallowedOutputColumns(val *[]*string)
	DisallowedOutputColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutAggregationThresholds(value interface{})
	PutComparisonControls(value *CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControls)
	PutDifferentialPrivacy(value *CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacy)
	ResetAdditionalAnalyses()
	ResetAggregationThresholds()
	ResetAllowedAnalyses()
	ResetAllowedAnalysisProviders()
	ResetAllowedResultReceivers()
	ResetComparisonControls()
	ResetDifferentialPrivacy()
	ResetDisallowedOutputColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference
type jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AdditionalAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AdditionalAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AggregationThresholds() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsList {
	var returns CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsList
	_jsii_.Get(
		j,
		"aggregationThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AggregationThresholdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalyses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalysesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalysisProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalysisProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalysisProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalysisProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedResultReceivers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResultReceivers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) AllowedResultReceiversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResultReceiversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ComparisonControls() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference {
	var returns CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference
	_jsii_.Get(
		j,
		"comparisonControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ComparisonControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"comparisonControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) DifferentialPrivacy() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference {
	var returns CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference
	_jsii_.Get(
		j,
		"differentialPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) DifferentialPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"differentialPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) DisallowedOutputColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedOutputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) DisallowedOutputColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedOutputColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIntermediateTable.CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference_Override(c CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIntermediateTable.CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetAdditionalAnalyses(val *string) {
	if err := j.validateSetAdditionalAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalAnalyses",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetAllowedAnalyses(val *[]*string) {
	if err := j.validateSetAllowedAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAnalyses",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetAllowedAnalysisProviders(val *[]*string) {
	if err := j.validateSetAllowedAnalysisProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAnalysisProviders",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetAllowedResultReceivers(val *[]*string) {
	if err := j.validateSetAllowedResultReceiversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResultReceivers",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetDisallowedOutputColumns(val *[]*string) {
	if err := j.validateSetDisallowedOutputColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowedOutputColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) PutAggregationThresholds(value interface{}) {
	if err := c.validatePutAggregationThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAggregationThresholds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) PutComparisonControls(value *CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControls) {
	if err := c.validatePutComparisonControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putComparisonControls",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) PutDifferentialPrivacy(value *CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacy) {
	if err := c.validatePutDifferentialPrivacyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDifferentialPrivacy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetAdditionalAnalyses() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalAnalyses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetAggregationThresholds() {
	_jsii_.InvokeVoid(
		c,
		"resetAggregationThresholds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetAllowedAnalyses() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedAnalyses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetAllowedAnalysisProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedAnalysisProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetAllowedResultReceivers() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedResultReceivers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetComparisonControls() {
	_jsii_.InvokeVoid(
		c,
		"resetComparisonControls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetDifferentialPrivacy() {
	_jsii_.InvokeVoid(
		c,
		"resetDifferentialPrivacy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ResetDisallowedOutputColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetDisallowedOutputColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

