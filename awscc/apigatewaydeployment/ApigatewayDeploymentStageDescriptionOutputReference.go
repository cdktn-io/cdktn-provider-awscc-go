// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewaydeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayDeploymentStageDescriptionOutputReference interface {
	cdktn.ComplexObject
	AccessLogSetting() ApigatewayDeploymentStageDescriptionAccessLogSettingOutputReference
	AccessLogSettingInput() interface{}
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	CacheClusterEnabledInput() interface{}
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	CacheClusterSizeInput() *string
	CacheDataEncrypted() interface{}
	SetCacheDataEncrypted(val interface{})
	CacheDataEncryptedInput() interface{}
	CacheTtlInSeconds() *float64
	SetCacheTtlInSeconds(val *float64)
	CacheTtlInSecondsInput() *float64
	CachingEnabled() interface{}
	SetCachingEnabled(val interface{})
	CachingEnabledInput() interface{}
	CanarySetting() ApigatewayDeploymentStageDescriptionCanarySettingOutputReference
	CanarySettingInput() interface{}
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	ClientCertificateIdInput() *string
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
	DataTraceEnabled() interface{}
	SetDataTraceEnabled(val interface{})
	DataTraceEnabledInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentationVersion() *string
	SetDocumentationVersion(val *string)
	DocumentationVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
	MethodSettings() ApigatewayDeploymentStageDescriptionMethodSettingsList
	MethodSettingsInput() interface{}
	MetricsEnabled() interface{}
	SetMetricsEnabled(val interface{})
	MetricsEnabledInput() interface{}
	Tags() ApigatewayDeploymentStageDescriptionTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThrottlingBurstLimit() *float64
	SetThrottlingBurstLimit(val *float64)
	ThrottlingBurstLimitInput() *float64
	ThrottlingRateLimit() *float64
	SetThrottlingRateLimit(val *float64)
	ThrottlingRateLimitInput() *float64
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
	TracingEnabledInput() interface{}
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
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
	PutAccessLogSetting(value *ApigatewayDeploymentStageDescriptionAccessLogSetting)
	PutCanarySetting(value *ApigatewayDeploymentStageDescriptionCanarySetting)
	PutMethodSettings(value interface{})
	PutTags(value interface{})
	ResetAccessLogSetting()
	ResetCacheClusterEnabled()
	ResetCacheClusterSize()
	ResetCacheDataEncrypted()
	ResetCacheTtlInSeconds()
	ResetCachingEnabled()
	ResetCanarySetting()
	ResetClientCertificateId()
	ResetDataTraceEnabled()
	ResetDescription()
	ResetDocumentationVersion()
	ResetLoggingLevel()
	ResetMethodSettings()
	ResetMetricsEnabled()
	ResetTags()
	ResetThrottlingBurstLimit()
	ResetThrottlingRateLimit()
	ResetTracingEnabled()
	ResetVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigatewayDeploymentStageDescriptionOutputReference
type jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) AccessLogSetting() ApigatewayDeploymentStageDescriptionAccessLogSettingOutputReference {
	var returns ApigatewayDeploymentStageDescriptionAccessLogSettingOutputReference
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) AccessLogSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheClusterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheClusterSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheDataEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheDataEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CacheTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CachingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CachingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CanarySetting() ApigatewayDeploymentStageDescriptionCanarySettingOutputReference {
	var returns ApigatewayDeploymentStageDescriptionCanarySettingOutputReference
	_jsii_.Get(
		j,
		"canarySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CanarySettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canarySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ClientCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) DocumentationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) DocumentationVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) MethodSettings() ApigatewayDeploymentStageDescriptionMethodSettingsList {
	var returns ApigatewayDeploymentStageDescriptionMethodSettingsList
	_jsii_.Get(
		j,
		"methodSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) MethodSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) Tags() ApigatewayDeploymentStageDescriptionTagsList {
	var returns ApigatewayDeploymentStageDescriptionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) TracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) TracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


func NewApigatewayDeploymentStageDescriptionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApigatewayDeploymentStageDescriptionOutputReference {
	_init_.Initialize()

	if err := validateNewApigatewayDeploymentStageDescriptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayDeployment.ApigatewayDeploymentStageDescriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayDeploymentStageDescriptionOutputReference_Override(a ApigatewayDeploymentStageDescriptionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayDeployment.ApigatewayDeploymentStageDescriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetCacheClusterEnabled(val interface{}) {
	if err := j.validateSetCacheClusterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetCacheClusterSize(val *string) {
	if err := j.validateSetCacheClusterSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetCacheDataEncrypted(val interface{}) {
	if err := j.validateSetCacheDataEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDataEncrypted",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetCacheTtlInSeconds(val *float64) {
	if err := j.validateSetCacheTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetCachingEnabled(val interface{}) {
	if err := j.validateSetCachingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetClientCertificateId(val *string) {
	if err := j.validateSetClientCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetDataTraceEnabled(val interface{}) {
	if err := j.validateSetDataTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetDocumentationVersion(val *string) {
	if err := j.validateSetDocumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentationVersion",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetThrottlingBurstLimit(val *float64) {
	if err := j.validateSetThrottlingBurstLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetThrottlingRateLimit(val *float64) {
	if err := j.validateSetThrottlingRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetTracingEnabled(val interface{}) {
	if err := j.validateSetTracingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference)SetVariables(val *map[string]*string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) PutAccessLogSetting(value *ApigatewayDeploymentStageDescriptionAccessLogSetting) {
	if err := a.validatePutAccessLogSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLogSetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) PutCanarySetting(value *ApigatewayDeploymentStageDescriptionCanarySetting) {
	if err := a.validatePutCanarySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanarySetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) PutMethodSettings(value interface{}) {
	if err := a.validatePutMethodSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMethodSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetAccessLogSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCacheClusterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCacheClusterSize() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCacheDataEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheDataEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCacheTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCachingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetCanarySetting() {
	_jsii_.InvokeVoid(
		a,
		"resetCanarySetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetClientCertificateId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetDocumentationVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentationVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetMethodSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMethodSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetTracingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTracingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ResetVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

