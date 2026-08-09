// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinessdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference interface {
	cdktn.ComplexObject
	AttributeValueOperator() *string
	SetAttributeValueOperator(val *string)
	AttributeValueOperatorInput() *string
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
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValueOutputReference
	ValueInput() interface{}
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
	PutValue(value *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValue)
	ResetAttributeValueOperator()
	ResetKey()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference
type jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) AttributeValueOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeValueOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) AttributeValueOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeValueOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) Value() QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValueOutputReference {
	var returns QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewQbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference_Override(q QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetAttributeValueOperator(val *string) {
	if err := j.validateSetAttributeValueOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeValueOperator",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) PutValue(value *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValue) {
	if err := q.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ResetAttributeValueOperator() {
	_jsii_.InvokeVoid(
		q,
		"resetAttributeValueOperator",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		q,
		"resetKey",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		q,
		"resetValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

