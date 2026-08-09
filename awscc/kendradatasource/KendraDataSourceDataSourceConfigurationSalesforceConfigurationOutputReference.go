// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference interface {
	cdktn.ComplexObject
	ChatterFeedConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfigurationOutputReference
	ChatterFeedConfigurationInput() interface{}
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
	CrawlAttachments() interface{}
	SetCrawlAttachments(val interface{})
	CrawlAttachmentsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeAttachmentFilePatterns() *[]*string
	SetExcludeAttachmentFilePatterns(val *[]*string)
	ExcludeAttachmentFilePatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeAttachmentFilePatterns() *[]*string
	SetIncludeAttachmentFilePatterns(val *[]*string)
	IncludeAttachmentFilePatternsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KnowledgeArticleConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference
	KnowledgeArticleConfigurationInput() interface{}
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
	ServerUrl() *string
	SetServerUrl(val *string)
	ServerUrlInput() *string
	StandardObjectAttachmentConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfigurationOutputReference
	StandardObjectAttachmentConfigurationInput() interface{}
	StandardObjectConfigurations() KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectConfigurationsList
	StandardObjectConfigurationsInput() interface{}
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
	PutChatterFeedConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfiguration)
	PutKnowledgeArticleConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfiguration)
	PutStandardObjectAttachmentConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfiguration)
	PutStandardObjectConfigurations(value interface{})
	ResetChatterFeedConfiguration()
	ResetCrawlAttachments()
	ResetExcludeAttachmentFilePatterns()
	ResetIncludeAttachmentFilePatterns()
	ResetKnowledgeArticleConfiguration()
	ResetSecretArn()
	ResetServerUrl()
	ResetStandardObjectAttachmentConfiguration()
	ResetStandardObjectConfigurations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ChatterFeedConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfigurationOutputReference
	_jsii_.Get(
		j,
		"chatterFeedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ChatterFeedConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatterFeedConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) CrawlAttachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) CrawlAttachmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ExcludeAttachmentFilePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeAttachmentFilePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ExcludeAttachmentFilePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeAttachmentFilePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) IncludeAttachmentFilePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeAttachmentFilePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) IncludeAttachmentFilePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeAttachmentFilePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) KnowledgeArticleConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference
	_jsii_.Get(
		j,
		"knowledgeArticleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) KnowledgeArticleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeArticleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ServerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) StandardObjectAttachmentConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfigurationOutputReference
	_jsii_.Get(
		j,
		"standardObjectAttachmentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) StandardObjectAttachmentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"standardObjectAttachmentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) StandardObjectConfigurations() KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectConfigurationsList {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectConfigurationsList
	_jsii_.Get(
		j,
		"standardObjectConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) StandardObjectConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"standardObjectConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetCrawlAttachments(val interface{}) {
	if err := j.validateSetCrawlAttachmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlAttachments",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetExcludeAttachmentFilePatterns(val *[]*string) {
	if err := j.validateSetExcludeAttachmentFilePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeAttachmentFilePatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetIncludeAttachmentFilePatterns(val *[]*string) {
	if err := j.validateSetIncludeAttachmentFilePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAttachmentFilePatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetServerUrl(val *string) {
	if err := j.validateSetServerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverUrl",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) PutChatterFeedConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfiguration) {
	if err := k.validatePutChatterFeedConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putChatterFeedConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) PutKnowledgeArticleConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfiguration) {
	if err := k.validatePutKnowledgeArticleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKnowledgeArticleConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) PutStandardObjectAttachmentConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfiguration) {
	if err := k.validatePutStandardObjectAttachmentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStandardObjectAttachmentConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) PutStandardObjectConfigurations(value interface{}) {
	if err := k.validatePutStandardObjectConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStandardObjectConfigurations",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetChatterFeedConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetChatterFeedConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetCrawlAttachments() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlAttachments",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetExcludeAttachmentFilePatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetExcludeAttachmentFilePatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetIncludeAttachmentFilePatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetIncludeAttachmentFilePatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetKnowledgeArticleConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetKnowledgeArticleConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetServerUrl() {
	_jsii_.InvokeVoid(
		k,
		"resetServerUrl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetStandardObjectAttachmentConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetStandardObjectAttachmentConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ResetStandardObjectConfigurations() {
	_jsii_.InvokeVoid(
		k,
		"resetStandardObjectConfigurations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

