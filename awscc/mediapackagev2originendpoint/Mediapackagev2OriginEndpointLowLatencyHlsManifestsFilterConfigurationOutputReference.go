// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference interface {
	cdktn.ComplexObject
	ClipStartTime() *string
	SetClipStartTime(val *string)
	ClipStartTimeInput() *string
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
	DrmSettings() *string
	SetDrmSettings(val *string)
	DrmSettingsInput() *string
	End() *string
	SetEnd(val *string)
	EndInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestFilter() *string
	SetManifestFilter(val *string)
	ManifestFilterInput() *string
	Start() *string
	SetStart(val *string)
	StartInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	TimeDelaySecondsInput() *float64
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
	ResetClipStartTime()
	ResetDrmSettings()
	ResetEnd()
	ResetManifestFilter()
	ResetStart()
	ResetTimeDelaySeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ClipStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ClipStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) DrmSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) DrmSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ManifestFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ManifestFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) TimeDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySecondsInput",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference_Override(m Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetClipStartTime(val *string) {
	if err := j.validateSetClipStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clipStartTime",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetDrmSettings(val *string) {
	if err := j.validateSetDrmSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drmSettings",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetEnd(val *string) {
	if err := j.validateSetEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"end",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetManifestFilter(val *string) {
	if err := j.validateSetManifestFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestFilter",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetStart(val *string) {
	if err := j.validateSetStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference)SetTimeDelaySeconds(val *float64) {
	if err := j.validateSetTimeDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetClipStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetClipStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetDrmSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetDrmSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetManifestFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		m,
		"resetStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ResetTimeDelaySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeDelaySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

