// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinessdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference interface {
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
	InlineConfigurations() QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsList
	InlineConfigurationsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PostExtractionHookConfiguration() QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference
	PostExtractionHookConfigurationInput() interface{}
	PreExtractionHookConfiguration() QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfigurationOutputReference
	PreExtractionHookConfigurationInput() interface{}
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
	PutInlineConfigurations(value interface{})
	PutPostExtractionHookConfiguration(value *QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfiguration)
	PutPreExtractionHookConfiguration(value *QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfiguration)
	ResetInlineConfigurations()
	ResetPostExtractionHookConfiguration()
	ResetPreExtractionHookConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference
type jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) InlineConfigurations() QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsList {
	var returns QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsList
	_jsii_.Get(
		j,
		"inlineConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) InlineConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PostExtractionHookConfiguration() QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference {
	var returns QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference
	_jsii_.Get(
		j,
		"postExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PostExtractionHookConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PreExtractionHookConfiguration() QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfigurationOutputReference {
	var returns QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfigurationOutputReference
	_jsii_.Get(
		j,
		"preExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PreExtractionHookConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQbusinessDataSourceDocumentEnrichmentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessDataSourceDocumentEnrichmentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessDataSourceDocumentEnrichmentConfigurationOutputReference_Override(q QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PutInlineConfigurations(value interface{}) {
	if err := q.validatePutInlineConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInlineConfigurations",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PutPostExtractionHookConfiguration(value *QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfiguration) {
	if err := q.validatePutPostExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPostExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) PutPreExtractionHookConfiguration(value *QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfiguration) {
	if err := q.validatePutPreExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPreExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ResetInlineConfigurations() {
	_jsii_.InvokeVoid(
		q,
		"resetInlineConfigurations",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ResetPostExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetPostExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ResetPreExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetPreExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

