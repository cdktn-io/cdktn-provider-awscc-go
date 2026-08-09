// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference interface {
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
	CrawlerLimits() BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimitsOutputReference
	CrawlerLimitsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExclusionFilters() *[]*string
	SetExclusionFilters(val *[]*string)
	ExclusionFiltersInput() *[]*string
	// Experimental.
	Fqn() *string
	InclusionFilters() *[]*string
	SetInclusionFilters(val *[]*string)
	InclusionFiltersInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserAgent() *string
	SetUserAgent(val *string)
	UserAgentHeader() *string
	UserAgentInput() *string
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
	PutCrawlerLimits(value *BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimits)
	ResetCrawlerLimits()
	ResetExclusionFilters()
	ResetInclusionFilters()
	ResetScope()
	ResetUserAgent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference
type jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) CrawlerLimits() BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimitsOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimitsOutputReference
	_jsii_.Get(
		j,
		"crawlerLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) CrawlerLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ExclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ExclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) InclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) InclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) UserAgentHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) UserAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


func NewBedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference_Override(b BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetExclusionFilters(val *[]*string) {
	if err := j.validateSetExclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionFilters",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetInclusionFilters(val *[]*string) {
	if err := j.validateSetInclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionFilters",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference)SetUserAgent(val *string) {
	if err := j.validateSetUserAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) PutCrawlerLimits(value *BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimits) {
	if err := b.validatePutCrawlerLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCrawlerLimits",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ResetCrawlerLimits() {
	_jsii_.InvokeVoid(
		b,
		"resetCrawlerLimits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ResetExclusionFilters() {
	_jsii_.InvokeVoid(
		b,
		"resetExclusionFilters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ResetInclusionFilters() {
	_jsii_.InvokeVoid(
		b,
		"resetInclusionFilters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		b,
		"resetScope",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ResetUserAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

