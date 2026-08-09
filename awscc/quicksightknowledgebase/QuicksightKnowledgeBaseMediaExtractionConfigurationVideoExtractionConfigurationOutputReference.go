// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference interface {
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
	VideoExtractionStatus() *string
	SetVideoExtractionStatus(val *string)
	VideoExtractionStatusInput() *string
	VideoExtractionType() *string
	SetVideoExtractionType(val *string)
	VideoExtractionTypeInput() *string
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
	ResetVideoExtractionStatus()
	ResetVideoExtractionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference
type jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) VideoExtractionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoExtractionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) VideoExtractionStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoExtractionStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) VideoExtractionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoExtractionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) VideoExtractionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoExtractionTypeInput",
		&returns,
	)
	return returns
}


func NewQuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference_Override(q QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetVideoExtractionStatus(val *string) {
	if err := j.validateSetVideoExtractionStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoExtractionStatus",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference)SetVideoExtractionType(val *string) {
	if err := j.validateSetVideoExtractionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoExtractionType",
		val,
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ResetVideoExtractionStatus() {
	_jsii_.InvokeVoid(
		q,
		"resetVideoExtractionStatus",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ResetVideoExtractionType() {
	_jsii_.InvokeVoid(
		q,
		"resetVideoExtractionType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

