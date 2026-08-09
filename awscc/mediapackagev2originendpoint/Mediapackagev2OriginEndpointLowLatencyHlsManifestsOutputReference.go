// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference interface {
	cdktn.ComplexObject
	ChildManifestName() *string
	SetChildManifestName(val *string)
	ChildManifestNameInput() *string
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
	FilterConfiguration() Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference
	FilterConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	ManifestWindowSeconds() *float64
	SetManifestWindowSeconds(val *float64)
	ManifestWindowSecondsInput() *float64
	ProgramDateTimeIntervalSeconds() *float64
	SetProgramDateTimeIntervalSeconds(val *float64)
	ProgramDateTimeIntervalSecondsInput() *float64
	ScteHls() Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHlsOutputReference
	ScteHlsInput() interface{}
	StartTag() Mediapackagev2OriginEndpointLowLatencyHlsManifestsStartTagOutputReference
	StartTagInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UriPathType() *string
	SetUriPathType(val *string)
	UriPathTypeInput() *string
	Url() *string
	SetUrl(val *string)
	UrlEncodeChildManifest() interface{}
	SetUrlEncodeChildManifest(val interface{})
	UrlEncodeChildManifestInput() interface{}
	UrlInput() *string
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
	PutFilterConfiguration(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfiguration)
	PutScteHls(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHls)
	PutStartTag(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsStartTag)
	ResetChildManifestName()
	ResetFilterConfiguration()
	ResetManifestName()
	ResetManifestWindowSeconds()
	ResetProgramDateTimeIntervalSeconds()
	ResetScteHls()
	ResetStartTag()
	ResetUriPathType()
	ResetUrl()
	ResetUrlEncodeChildManifest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ChildManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childManifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ChildManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childManifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) FilterConfiguration() Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference {
	var returns Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ManifestWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ManifestWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ProgramDateTimeIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ProgramDateTimeIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ScteHls() Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHlsOutputReference {
	var returns Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHlsOutputReference
	_jsii_.Get(
		j,
		"scteHls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ScteHlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scteHlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) StartTag() Mediapackagev2OriginEndpointLowLatencyHlsManifestsStartTagOutputReference {
	var returns Mediapackagev2OriginEndpointLowLatencyHlsManifestsStartTagOutputReference
	_jsii_.Get(
		j,
		"startTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) StartTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) UriPathType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriPathType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) UriPathTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriPathTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) UrlEncodeChildManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlEncodeChildManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) UrlEncodeChildManifestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlEncodeChildManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference_Override(m Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetChildManifestName(val *string) {
	if err := j.validateSetChildManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childManifestName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetManifestWindowSeconds(val *float64) {
	if err := j.validateSetManifestWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetProgramDateTimeIntervalSeconds(val *float64) {
	if err := j.validateSetProgramDateTimeIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetUriPathType(val *string) {
	if err := j.validateSetUriPathTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriPathType",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference)SetUrlEncodeChildManifest(val interface{}) {
	if err := j.validateSetUrlEncodeChildManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlEncodeChildManifest",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) PutFilterConfiguration(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfiguration) {
	if err := m.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) PutScteHls(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsScteHls) {
	if err := m.validatePutScteHlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScteHls",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) PutStartTag(value *Mediapackagev2OriginEndpointLowLatencyHlsManifestsStartTag) {
	if err := m.validatePutStartTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStartTag",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetChildManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetChildManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetManifestWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetProgramDateTimeIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimeIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetScteHls() {
	_jsii_.InvokeVoid(
		m,
		"resetScteHls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetStartTag() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTag",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetUriPathType() {
	_jsii_.InvokeVoid(
		m,
		"resetUriPathType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ResetUrlEncodeChildManifest() {
	_jsii_.InvokeVoid(
		m,
		"resetUrlEncodeChildManifest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

