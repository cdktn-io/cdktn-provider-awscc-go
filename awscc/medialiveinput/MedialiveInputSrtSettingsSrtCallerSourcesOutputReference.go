// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/medialiveinput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveInputSrtSettingsSrtCallerSourcesOutputReference interface {
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
	Decryption() MedialiveInputSrtSettingsSrtCallerSourcesDecryptionOutputReference
	DecryptionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumLatency() *float64
	SetMinimumLatency(val *float64)
	MinimumLatencyInput() *float64
	SrtListenerAddress() *string
	SetSrtListenerAddress(val *string)
	SrtListenerAddressInput() *string
	SrtListenerPort() *string
	SetSrtListenerPort(val *string)
	SrtListenerPortInput() *string
	StreamId() *string
	SetStreamId(val *string)
	StreamIdInput() *string
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
	PutDecryption(value *MedialiveInputSrtSettingsSrtCallerSourcesDecryption)
	ResetDecryption()
	ResetMinimumLatency()
	ResetSrtListenerAddress()
	ResetSrtListenerPort()
	ResetStreamId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveInputSrtSettingsSrtCallerSourcesOutputReference
type jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) Decryption() MedialiveInputSrtSettingsSrtCallerSourcesDecryptionOutputReference {
	var returns MedialiveInputSrtSettingsSrtCallerSourcesDecryptionOutputReference
	_jsii_.Get(
		j,
		"decryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) DecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) MinimumLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) MinimumLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) SrtListenerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srtListenerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) SrtListenerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srtListenerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) SrtListenerPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srtListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) SrtListenerPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srtListenerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveInputSrtSettingsSrtCallerSourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MedialiveInputSrtSettingsSrtCallerSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveInputSrtSettingsSrtCallerSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.medialiveInput.MedialiveInputSrtSettingsSrtCallerSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMedialiveInputSrtSettingsSrtCallerSourcesOutputReference_Override(m MedialiveInputSrtSettingsSrtCallerSourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.medialiveInput.MedialiveInputSrtSettingsSrtCallerSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetMinimumLatency(val *float64) {
	if err := j.validateSetMinimumLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLatency",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetSrtListenerAddress(val *string) {
	if err := j.validateSetSrtListenerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srtListenerAddress",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetSrtListenerPort(val *string) {
	if err := j.validateSetSrtListenerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srtListenerPort",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) PutDecryption(value *MedialiveInputSrtSettingsSrtCallerSourcesDecryption) {
	if err := m.validatePutDecryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDecryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ResetDecryption() {
	_jsii_.InvokeVoid(
		m,
		"resetDecryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ResetMinimumLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ResetSrtListenerAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtListenerAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ResetSrtListenerPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtListenerPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveInputSrtSettingsSrtCallerSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

