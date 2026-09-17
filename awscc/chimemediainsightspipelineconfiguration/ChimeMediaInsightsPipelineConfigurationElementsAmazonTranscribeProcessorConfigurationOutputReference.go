// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/chimemediainsightspipelineconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference interface {
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
	IdentifyLanguage() interface{}
	SetIdentifyLanguage(val interface{})
	IdentifyLanguageInput() interface{}
	IdentifyMultipleLanguages() interface{}
	SetIdentifyMultipleLanguages(val interface{})
	IdentifyMultipleLanguagesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	LanguageModelName() *string
	SetLanguageModelName(val *string)
	LanguageModelNameInput() *string
	LanguageOptions() *string
	SetLanguageOptions(val *string)
	LanguageOptionsInput() *string
	PartialResultsStability() *string
	SetPartialResultsStability(val *string)
	PartialResultsStabilityInput() *string
	PiiEntityTypes() *string
	SetPiiEntityTypes(val *string)
	PiiEntityTypesInput() *string
	PreferredLanguage() *string
	SetPreferredLanguage(val *string)
	PreferredLanguageInput() *string
	ShowSpeakerLabel() interface{}
	SetShowSpeakerLabel(val interface{})
	ShowSpeakerLabelInput() interface{}
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
	VocabularyFilterNames() *string
	SetVocabularyFilterNames(val *string)
	VocabularyFilterNamesInput() *string
	VocabularyName() *string
	SetVocabularyName(val *string)
	VocabularyNameInput() *string
	VocabularyNames() *string
	SetVocabularyNames(val *string)
	VocabularyNamesInput() *string
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
	ResetContentIdentificationType()
	ResetContentRedactionType()
	ResetEnablePartialResultsStabilization()
	ResetFilterPartialResults()
	ResetIdentifyLanguage()
	ResetIdentifyMultipleLanguages()
	ResetLanguageCode()
	ResetLanguageModelName()
	ResetLanguageOptions()
	ResetPartialResultsStability()
	ResetPiiEntityTypes()
	ResetPreferredLanguage()
	ResetShowSpeakerLabel()
	ResetVocabularyFilterMethod()
	ResetVocabularyFilterName()
	ResetVocabularyFilterNames()
	ResetVocabularyName()
	ResetVocabularyNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference
type jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) IdentifyLanguage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) IdentifyLanguageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) IdentifyMultipleLanguages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyMultipleLanguages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) IdentifyMultipleLanguagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyMultipleLanguagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PreferredLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ShowSpeakerLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ShowSpeakerLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterNamesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyNamesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNamesInput",
		&returns,
	)
	return returns
}


func NewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference_Override(c ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeMediaInsightsPipelineConfiguration.ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetIdentifyLanguage(val interface{}) {
	if err := j.validateSetIdentifyLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifyLanguage",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetIdentifyMultipleLanguages(val interface{}) {
	if err := j.validateSetIdentifyMultipleLanguagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifyMultipleLanguages",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetLanguageOptions(val *string) {
	if err := j.validateSetLanguageOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageOptions",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetPreferredLanguage(val *string) {
	if err := j.validateSetPreferredLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredLanguage",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetShowSpeakerLabel(val interface{}) {
	if err := j.validateSetShowSpeakerLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showSpeakerLabel",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyFilterNames(val *string) {
	if err := j.validateSetVocabularyFilterNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterNames",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (j *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyNames(val *string) {
	if err := j.validateSetVocabularyNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyNames",
		val,
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		c,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetIdentifyLanguage() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentifyLanguage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetIdentifyMultipleLanguages() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentifyMultipleLanguages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetLanguageOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		c,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetPreferredLanguage() {
	_jsii_.InvokeVoid(
		c,
		"resetPreferredLanguage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetShowSpeakerLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetShowSpeakerLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyFilterNames() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyNames() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

