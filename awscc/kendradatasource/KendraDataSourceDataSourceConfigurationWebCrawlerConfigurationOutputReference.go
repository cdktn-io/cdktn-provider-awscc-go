// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuthenticationConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() interface{}
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
	CrawlDepth() *float64
	SetCrawlDepth(val *float64)
	CrawlDepthInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxContentSizePerPageInMegaBytes() *float64
	SetMaxContentSizePerPageInMegaBytes(val *float64)
	MaxContentSizePerPageInMegaBytesInput() *float64
	MaxLinksPerPage() *float64
	SetMaxLinksPerPage(val *float64)
	MaxLinksPerPageInput() *float64
	MaxUrlsPerMinuteCrawlRate() *float64
	SetMaxUrlsPerMinuteCrawlRate(val *float64)
	MaxUrlsPerMinuteCrawlRateInput() *float64
	ProxyConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference
	ProxyConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UrlExclusionPatterns() *[]*string
	SetUrlExclusionPatterns(val *[]*string)
	UrlExclusionPatternsInput() *[]*string
	UrlInclusionPatterns() *[]*string
	SetUrlInclusionPatterns(val *[]*string)
	UrlInclusionPatternsInput() *[]*string
	Urls() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference
	UrlsInput() interface{}
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
	PutAuthenticationConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration)
	PutProxyConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration)
	PutUrls(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrls)
	ResetAuthenticationConfiguration()
	ResetCrawlDepth()
	ResetMaxContentSizePerPageInMegaBytes()
	ResetMaxLinksPerPage()
	ResetMaxUrlsPerMinuteCrawlRate()
	ResetProxyConfiguration()
	ResetUrlExclusionPatterns()
	ResetUrlInclusionPatterns()
	ResetUrls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) AuthenticationConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) AuthenticationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) CrawlDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) CrawlDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxContentSizePerPageInMegaBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxContentSizePerPageInMegaBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxLinksPerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxLinksPerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxUrlsPerMinuteCrawlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxUrlsPerMinuteCrawlRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ProxyConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ProxyConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlInclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlInclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) Urls() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference {
	var returns KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetCrawlDepth(val *float64) {
	if err := j.validateSetCrawlDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlDepth",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxContentSizePerPageInMegaBytes(val *float64) {
	if err := j.validateSetMaxContentSizePerPageInMegaBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxContentSizePerPageInMegaBytes",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxLinksPerPage(val *float64) {
	if err := j.validateSetMaxLinksPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLinksPerPage",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxUrlsPerMinuteCrawlRate(val *float64) {
	if err := j.validateSetMaxUrlsPerMinuteCrawlRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUrlsPerMinuteCrawlRate",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetUrlExclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlExclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetUrlInclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlInclusionPatterns",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutAuthenticationConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration) {
	if err := k.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutProxyConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration) {
	if err := k.validatePutProxyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProxyConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutUrls(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrls) {
	if err := k.validatePutUrlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putUrls",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetCrawlDepth() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlDepth",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxContentSizePerPageInMegaBytes() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxContentSizePerPageInMegaBytes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxLinksPerPage() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxLinksPerPage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxUrlsPerMinuteCrawlRate() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUrlsPerMinuteCrawlRate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetProxyConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProxyConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetUrlExclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetUrlExclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetUrlInclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetUrlInclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetUrls() {
	_jsii_.InvokeVoid(
		k,
		"resetUrls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

