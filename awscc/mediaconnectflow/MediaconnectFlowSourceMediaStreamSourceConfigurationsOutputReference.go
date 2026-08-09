// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference interface {
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
	EncodingName() *string
	SetEncodingName(val *string)
	EncodingNameInput() *string
	// Experimental.
	Fqn() *string
	InputConfigurations() MediaconnectFlowSourceMediaStreamSourceConfigurationsInputConfigurationsList
	InputConfigurationsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MediaStreamName() *string
	SetMediaStreamName(val *string)
	MediaStreamNameInput() *string
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
	PutInputConfigurations(value interface{})
	ResetEncodingName()
	ResetInputConfigurations()
	ResetMediaStreamName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference
type jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) EncodingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) EncodingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) InputConfigurations() MediaconnectFlowSourceMediaStreamSourceConfigurationsInputConfigurationsList {
	var returns MediaconnectFlowSourceMediaStreamSourceConfigurationsInputConfigurationsList
	_jsii_.Get(
		j,
		"inputConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) InputConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) MediaStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) MediaStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference_Override(m MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetEncodingName(val *string) {
	if err := j.validateSetEncodingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodingName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetMediaStreamName(val *string) {
	if err := j.validateSetMediaStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStreamName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) PutInputConfigurations(value interface{}) {
	if err := m.validatePutInputConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInputConfigurations",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ResetEncodingName() {
	_jsii_.InvokeVoid(
		m,
		"resetEncodingName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ResetInputConfigurations() {
	_jsii_.InvokeVoid(
		m,
		"resetInputConfigurations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ResetMediaStreamName() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaStreamName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectFlowSourceMediaStreamSourceConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

