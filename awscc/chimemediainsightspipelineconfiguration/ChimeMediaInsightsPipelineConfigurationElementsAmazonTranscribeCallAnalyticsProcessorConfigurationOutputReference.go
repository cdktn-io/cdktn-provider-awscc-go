// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/chimemediainsightspipelineconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference interface {
	cdktn.ComplexObject
	CallAnalyticsStreamCategories() *[]*string
	SetCallAnalyticsStreamCategories(val *[]*string)
	CallAnalyticsStreamCategoriesInput() *[]*string
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
	ContentIdentificationType() *string
	SetContentIdentificationType(val *string)
	ContentIdentificationTypeInput() *string
	ContentRedactionType() *string
	SetContentRedactionType(val *string)
	ContentRedactionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnablePartialResultsStabilization() interface{}
	SetEnablePartialResultsStabilization(val interface{})
	EnablePartialResultsStabilizationInput() interface{}
	FilterPartialResults() interface{}
	SetFilterPartialResults(val interface{})
	FilterPartialResultsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	LanguageModelName() *string
	SetLanguageModelName(val *string)
	LanguageModelNameInput() *string
	PartialResultsStability() *string
	SetPartialResultsStability(val *string)
	PartialResultsStabilityInput() *string
	PiiEntityTypes() *string
	SetPiiEntityTypes(val *string)
	PiiEntityTypesInput() *string
	PostCallAnalyticsSettings() ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsOutputReference
	PostCallAnalyticsSettingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VocabularyFilterMethod() *string
	SetVocabularyFilterMethod(val *string)
	VocabularyFilterMethodInput() *string
	VocabularyFilterName() *string
	SetVocabularyFilterName(val *string)
	VocabularyFilterNameInput() *string
	VocabularyName() *string
	SetVocabularyName(val *string)
	VocabularyNameInput() *string
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
	PutPostCallAnalyticsSettings(value *ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings)
	ResetCallAnalyticsStreamCategories()
	ResetContentIdentificationType()
	ResetContentRedactionType()
	ResetEnablePartialResultsStabilization()
	ResetFilterPartialResults()
	ResetLanguageCode()
	ResetLanguageModelName()
	ResetPartialResultsStability()
	ResetPiiEntityTypes()
	ResetPostCallAnalyticsSettings()
	ResetVocabularyFilterMethod()
	ResetVocabularyFilterName()
	ResetVocabularyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference
type jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) CallAnalyticsStreamCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) CallAnalyticsStreamCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PostCallAnalyticsSettings() ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsOutputReference {
	var returns ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsOutputReference
	_jsii_.Get(
		j,
		"postCallAnalyticsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PostCallAnalyticsSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postCallAnalyticsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}


func NewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference_Override(c ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetCallAnalyticsStreamCategories(val *[]*string) {
	if err := j.validateSetCallAnalyticsStreamCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callAnalyticsStreamCategories",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) PutPostCallAnalyticsSettings(value *ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings) {
	if err := c.validatePutPostCallAnalyticsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostCallAnalyticsSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetCallAnalyticsStreamCategories() {
	_jsii_.InvokeVoid(
		c,
		"resetCallAnalyticsStreamCategories",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		c,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		c,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetPostCallAnalyticsSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPostCallAnalyticsSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

