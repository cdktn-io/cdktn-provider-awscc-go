// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinessdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessDataSourceMediaExtractionConfigurationOutputReference interface {
	cdktn.ComplexObject
	AudioExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfigurationOutputReference
	AudioExtractionConfigurationInput() interface{}
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
	ImageExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfigurationOutputReference
	ImageExtractionConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VideoExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfigurationOutputReference
	VideoExtractionConfigurationInput() interface{}
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
	PutAudioExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfiguration)
	PutImageExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfiguration)
	PutVideoExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfiguration)
	ResetAudioExtractionConfiguration()
	ResetImageExtractionConfiguration()
	ResetVideoExtractionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessDataSourceMediaExtractionConfigurationOutputReference
type jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) AudioExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfigurationOutputReference {
	var returns QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfigurationOutputReference
	_jsii_.Get(
		j,
		"audioExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) AudioExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ImageExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfigurationOutputReference {
	var returns QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ImageExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) VideoExtractionConfiguration() QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfigurationOutputReference {
	var returns QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfigurationOutputReference
	_jsii_.Get(
		j,
		"videoExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) VideoExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoExtractionConfigurationInput",
		&returns,
	)
	return returns
}


func NewQbusinessDataSourceMediaExtractionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QbusinessDataSourceMediaExtractionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessDataSourceMediaExtractionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceMediaExtractionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessDataSourceMediaExtractionConfigurationOutputReference_Override(q QbusinessDataSourceMediaExtractionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceMediaExtractionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) PutAudioExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfiguration) {
	if err := q.validatePutAudioExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAudioExtractionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) PutImageExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfiguration) {
	if err := q.validatePutImageExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putImageExtractionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) PutVideoExtractionConfiguration(value *QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfiguration) {
	if err := q.validatePutVideoExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putVideoExtractionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ResetAudioExtractionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAudioExtractionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ResetImageExtractionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetImageExtractionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ResetVideoExtractionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetVideoExtractionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceMediaExtractionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

