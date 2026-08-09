// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Mediapackagev2OriginEndpointSegmentEncryptionOutputReference interface {
	cdktn.ComplexObject
	CmafExcludeSegmentDrmMetadata() interface{}
	SetCmafExcludeSegmentDrmMetadata(val interface{})
	CmafExcludeSegmentDrmMetadataInput() interface{}
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
	ConstantInitializationVector() *string
	SetConstantInitializationVector(val *string)
	ConstantInitializationVectorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionMethod() Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethodOutputReference
	EncryptionMethodInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyRotationIntervalSeconds() *float64
	SetKeyRotationIntervalSeconds(val *float64)
	KeyRotationIntervalSecondsInput() *float64
	SpekeKeyProvider() Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProviderOutputReference
	SpekeKeyProviderInput() interface{}
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
	PutEncryptionMethod(value *Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod)
	PutSpekeKeyProvider(value *Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProvider)
	ResetCmafExcludeSegmentDrmMetadata()
	ResetConstantInitializationVector()
	ResetEncryptionMethod()
	ResetKeyRotationIntervalSeconds()
	ResetSpekeKeyProvider()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointSegmentEncryptionOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) CmafExcludeSegmentDrmMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafExcludeSegmentDrmMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) CmafExcludeSegmentDrmMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafExcludeSegmentDrmMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ConstantInitializationVector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantInitializationVector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ConstantInitializationVectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantInitializationVectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) EncryptionMethod() Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethodOutputReference {
	var returns Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethodOutputReference
	_jsii_.Get(
		j,
		"encryptionMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) EncryptionMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) KeyRotationIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyRotationIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) KeyRotationIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyRotationIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) SpekeKeyProvider() Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProviderOutputReference {
	var returns Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProviderOutputReference
	_jsii_.Get(
		j,
		"spekeKeyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) SpekeKeyProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spekeKeyProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointSegmentEncryptionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Mediapackagev2OriginEndpointSegmentEncryptionOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointSegmentEncryptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointSegmentEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointSegmentEncryptionOutputReference_Override(m Mediapackagev2OriginEndpointSegmentEncryptionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointSegmentEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetCmafExcludeSegmentDrmMetadata(val interface{}) {
	if err := j.validateSetCmafExcludeSegmentDrmMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmafExcludeSegmentDrmMetadata",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetConstantInitializationVector(val *string) {
	if err := j.validateSetConstantInitializationVectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constantInitializationVector",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetKeyRotationIntervalSeconds(val *float64) {
	if err := j.validateSetKeyRotationIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRotationIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) PutEncryptionMethod(value *Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod) {
	if err := m.validatePutEncryptionMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionMethod",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) PutSpekeKeyProvider(value *Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProvider) {
	if err := m.validatePutSpekeKeyProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSpekeKeyProvider",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ResetCmafExcludeSegmentDrmMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetCmafExcludeSegmentDrmMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ResetConstantInitializationVector() {
	_jsii_.InvokeVoid(
		m,
		"resetConstantInitializationVector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ResetEncryptionMethod() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionMethod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ResetKeyRotationIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyRotationIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ResetSpekeKeyProvider() {
	_jsii_.InvokeVoid(
		m,
		"resetSpekeKeyProvider",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

