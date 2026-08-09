// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference interface {
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
	FalseyCellValue() *string
	SetFalseyCellValue(val *string)
	FalseyCellValueInput() *string
	FalseyCellValueSynonyms() *[]*string
	SetFalseyCellValueSynonyms(val *[]*string)
	FalseyCellValueSynonymsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SubTypeName() *string
	SetSubTypeName(val *string)
	SubTypeNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TruthyCellValue() *string
	SetTruthyCellValue(val *string)
	TruthyCellValueInput() *string
	TruthyCellValueSynonyms() *[]*string
	SetTruthyCellValueSynonyms(val *[]*string)
	TruthyCellValueSynonymsInput() *[]*string
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	TypeParameters() *map[string]*string
	SetTypeParameters(val *map[string]*string)
	TypeParametersInput() *map[string]*string
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
	ResetFalseyCellValue()
	ResetFalseyCellValueSynonyms()
	ResetSubTypeName()
	ResetTruthyCellValue()
	ResetTruthyCellValueSynonyms()
	ResetTypeName()
	ResetTypeParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference
type jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"falseyCellValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"falseyCellValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValueSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"falseyCellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValueSynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"falseyCellValueSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) SubTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) SubTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truthyCellValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truthyCellValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValueSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"truthyCellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValueSynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"truthyCellValueSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"typeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"typeParametersInput",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference_Override(q QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetFalseyCellValue(val *string) {
	if err := j.validateSetFalseyCellValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"falseyCellValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetFalseyCellValueSynonyms(val *[]*string) {
	if err := j.validateSetFalseyCellValueSynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"falseyCellValueSynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetSubTypeName(val *string) {
	if err := j.validateSetSubTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTypeName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTruthyCellValue(val *string) {
	if err := j.validateSetTruthyCellValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truthyCellValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTruthyCellValueSynonyms(val *[]*string) {
	if err := j.validateSetTruthyCellValueSynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truthyCellValueSynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTypeParameters(val *map[string]*string) {
	if err := j.validateSetTypeParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeParameters",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetFalseyCellValue() {
	_jsii_.InvokeVoid(
		q,
		"resetFalseyCellValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetFalseyCellValueSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetFalseyCellValueSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetSubTypeName() {
	_jsii_.InvokeVoid(
		q,
		"resetSubTypeName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetTruthyCellValue() {
	_jsii_.InvokeVoid(
		q,
		"resetTruthyCellValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetTruthyCellValueSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetTruthyCellValueSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		q,
		"resetTypeName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ResetTypeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTypeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

