// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectrouteroutput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference interface {
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
	DestinationTransitEncryption() MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryptionOutputReference
	DestinationTransitEncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MediaLiveInputArn() *string
	SetMediaLiveInputArn(val *string)
	MediaLiveInputArnInput() *string
	MediaLivePipelineId() *string
	SetMediaLivePipelineId(val *string)
	MediaLivePipelineIdInput() *string
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
	PutDestinationTransitEncryption(value *MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryption)
	ResetDestinationTransitEncryption()
	ResetMediaLiveInputArn()
	ResetMediaLivePipelineId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference
type jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) DestinationTransitEncryption() MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryptionOutputReference {
	var returns MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryptionOutputReference
	_jsii_.Get(
		j,
		"destinationTransitEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) DestinationTransitEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationTransitEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) MediaLiveInputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaLiveInputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) MediaLiveInputArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaLiveInputArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) MediaLivePipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaLivePipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) MediaLivePipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaLivePipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterOutputConfigurationMediaLiveInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterOutputConfigurationMediaLiveInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterOutputConfigurationMediaLiveInputOutputReference_Override(m MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetMediaLiveInputArn(val *string) {
	if err := j.validateSetMediaLiveInputArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaLiveInputArn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetMediaLivePipelineId(val *string) {
	if err := j.validateSetMediaLivePipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaLivePipelineId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) PutDestinationTransitEncryption(value *MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryption) {
	if err := m.validatePutDestinationTransitEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDestinationTransitEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ResetDestinationTransitEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationTransitEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ResetMediaLiveInputArn() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaLiveInputArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ResetMediaLivePipelineId() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaLivePipelineId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationMediaLiveInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

