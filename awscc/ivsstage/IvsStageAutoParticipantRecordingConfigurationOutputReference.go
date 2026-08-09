// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ivsstage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IvsStageAutoParticipantRecordingConfigurationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HlsConfiguration() IvsStageAutoParticipantRecordingConfigurationHlsConfigurationOutputReference
	HlsConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MediaTypes() *[]*string
	SetMediaTypes(val *[]*string)
	MediaTypesInput() *[]*string
	RecordingReconnectWindowSeconds() *float64
	SetRecordingReconnectWindowSeconds(val *float64)
	RecordingReconnectWindowSecondsInput() *float64
	StorageConfigurationArn() *string
	SetStorageConfigurationArn(val *string)
	StorageConfigurationArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThumbnailConfiguration() IvsStageAutoParticipantRecordingConfigurationThumbnailConfigurationOutputReference
	ThumbnailConfigurationInput() interface{}
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
	PutHlsConfiguration(value *IvsStageAutoParticipantRecordingConfigurationHlsConfiguration)
	PutThumbnailConfiguration(value *IvsStageAutoParticipantRecordingConfigurationThumbnailConfiguration)
	ResetHlsConfiguration()
	ResetMediaTypes()
	ResetRecordingReconnectWindowSeconds()
	ResetStorageConfigurationArn()
	ResetThumbnailConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IvsStageAutoParticipantRecordingConfigurationOutputReference
type jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) HlsConfiguration() IvsStageAutoParticipantRecordingConfigurationHlsConfigurationOutputReference {
	var returns IvsStageAutoParticipantRecordingConfigurationHlsConfigurationOutputReference
	_jsii_.Get(
		j,
		"hlsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) HlsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) MediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) MediaTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mediaTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) RecordingReconnectWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordingReconnectWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) RecordingReconnectWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordingReconnectWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) StorageConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) StorageConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ThumbnailConfiguration() IvsStageAutoParticipantRecordingConfigurationThumbnailConfigurationOutputReference {
	var returns IvsStageAutoParticipantRecordingConfigurationThumbnailConfigurationOutputReference
	_jsii_.Get(
		j,
		"thumbnailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ThumbnailConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thumbnailConfigurationInput",
		&returns,
	)
	return returns
}


func NewIvsStageAutoParticipantRecordingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IvsStageAutoParticipantRecordingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIvsStageAutoParticipantRecordingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ivsStage.IvsStageAutoParticipantRecordingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIvsStageAutoParticipantRecordingConfigurationOutputReference_Override(i IvsStageAutoParticipantRecordingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ivsStage.IvsStageAutoParticipantRecordingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetMediaTypes(val *[]*string) {
	if err := j.validateSetMediaTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaTypes",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetRecordingReconnectWindowSeconds(val *float64) {
	if err := j.validateSetRecordingReconnectWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordingReconnectWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetStorageConfigurationArn(val *string) {
	if err := j.validateSetStorageConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) PutHlsConfiguration(value *IvsStageAutoParticipantRecordingConfigurationHlsConfiguration) {
	if err := i.validatePutHlsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHlsConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) PutThumbnailConfiguration(value *IvsStageAutoParticipantRecordingConfigurationThumbnailConfiguration) {
	if err := i.validatePutThumbnailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putThumbnailConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ResetHlsConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetHlsConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ResetMediaTypes() {
	_jsii_.InvokeVoid(
		i,
		"resetMediaTypes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ResetRecordingReconnectWindowSeconds() {
	_jsii_.InvokeVoid(
		i,
		"resetRecordingReconnectWindowSeconds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ResetStorageConfigurationArn() {
	_jsii_.InvokeVoid(
		i,
		"resetStorageConfigurationArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ResetThumbnailConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetThumbnailConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvsStageAutoParticipantRecordingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

