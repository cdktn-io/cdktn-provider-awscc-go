// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediapackagePackagingConfigurationDashPackageOutputReference interface {
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
	DashManifests() MediapackagePackagingConfigurationDashPackageDashManifestsList
	DashManifestsInput() interface{}
	Encryption() MediapackagePackagingConfigurationDashPackageEncryptionOutputReference
	EncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeEncoderConfigurationInSegments() interface{}
	SetIncludeEncoderConfigurationInSegments(val interface{})
	IncludeEncoderConfigurationInSegmentsInput() interface{}
	IncludeIframeOnlyStream() interface{}
	SetIncludeIframeOnlyStream(val interface{})
	IncludeIframeOnlyStreamInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PeriodTriggers() *[]*string
	SetPeriodTriggers(val *[]*string)
	PeriodTriggersInput() *[]*string
	SegmentDurationSeconds() *float64
	SetSegmentDurationSeconds(val *float64)
	SegmentDurationSecondsInput() *float64
	SegmentTemplateFormat() *string
	SetSegmentTemplateFormat(val *string)
	SegmentTemplateFormatInput() *string
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
	PutDashManifests(value interface{})
	PutEncryption(value *MediapackagePackagingConfigurationDashPackageEncryption)
	ResetDashManifests()
	ResetEncryption()
	ResetIncludeEncoderConfigurationInSegments()
	ResetIncludeIframeOnlyStream()
	ResetPeriodTriggers()
	ResetSegmentDurationSeconds()
	ResetSegmentTemplateFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackagePackagingConfigurationDashPackageOutputReference
type jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) DashManifests() MediapackagePackagingConfigurationDashPackageDashManifestsList {
	var returns MediapackagePackagingConfigurationDashPackageDashManifestsList
	_jsii_.Get(
		j,
		"dashManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) DashManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) Encryption() MediapackagePackagingConfigurationDashPackageEncryptionOutputReference {
	var returns MediapackagePackagingConfigurationDashPackageEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) IncludeEncoderConfigurationInSegments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeEncoderConfigurationInSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) IncludeEncoderConfigurationInSegmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeEncoderConfigurationInSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) IncludeIframeOnlyStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) IncludeIframeOnlyStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) PeriodTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) PeriodTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) SegmentDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) SegmentTemplateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) SegmentTemplateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationDashPackageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediapackagePackagingConfigurationDashPackageOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationDashPackageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationDashPackageOutputReference_Override(m MediapackagePackagingConfigurationDashPackageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetIncludeEncoderConfigurationInSegments(val interface{}) {
	if err := j.validateSetIncludeEncoderConfigurationInSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeEncoderConfigurationInSegments",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetIncludeIframeOnlyStream(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStream",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetPeriodTriggers(val *[]*string) {
	if err := j.validateSetPeriodTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodTriggers",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetSegmentDurationSeconds(val *float64) {
	if err := j.validateSetSegmentDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetSegmentTemplateFormat(val *string) {
	if err := j.validateSetSegmentTemplateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentTemplateFormat",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) PutDashManifests(value interface{}) {
	if err := m.validatePutDashManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDashManifests",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) PutEncryption(value *MediapackagePackagingConfigurationDashPackageEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetDashManifests() {
	_jsii_.InvokeVoid(
		m,
		"resetDashManifests",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetIncludeEncoderConfigurationInSegments() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeEncoderConfigurationInSegments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetIncludeIframeOnlyStream() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStream",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetPeriodTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetPeriodTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetSegmentDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ResetSegmentTemplateFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentTemplateFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

