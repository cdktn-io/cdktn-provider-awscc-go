// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendraindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference interface {
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	Freshness() interface{}
	SetFreshness(val interface{})
	FreshnessInput() interface{}
	Importance() *float64
	SetImportance(val *float64)
	ImportanceInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RankOrder() *string
	SetRankOrder(val *string)
	RankOrderInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValueImportanceItems() KendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList
	ValueImportanceItemsInput() interface{}
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
	PutValueImportanceItems(value interface{})
	ResetDuration()
	ResetFreshness()
	ResetImportance()
	ResetRankOrder()
	ResetValueImportanceItems()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference
type jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) Freshness() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) FreshnessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) Importance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ImportanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) RankOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) RankOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ValueImportanceItems() KendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList {
	var returns KendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList
	_jsii_.Get(
		j,
		"valueImportanceItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ValueImportanceItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueImportanceItemsInput",
		&returns,
	)
	return returns
}


func NewKendraIndexDocumentMetadataConfigurationsRelevanceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference {
	_init_.Initialize()

	if err := validateNewKendraIndexDocumentMetadataConfigurationsRelevanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraIndex.KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraIndexDocumentMetadataConfigurationsRelevanceOutputReference_Override(k KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraIndex.KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetFreshness(val interface{}) {
	if err := j.validateSetFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freshness",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetImportance(val *float64) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetRankOrder(val *string) {
	if err := j.validateSetRankOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rankOrder",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) PutValueImportanceItems(value interface{}) {
	if err := k.validatePutValueImportanceItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putValueImportanceItems",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ResetFreshness() {
	_jsii_.InvokeVoid(
		k,
		"resetFreshness",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ResetImportance() {
	_jsii_.InvokeVoid(
		k,
		"resetImportance",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ResetRankOrder() {
	_jsii_.InvokeVoid(
		k,
		"resetRankOrder",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ResetValueImportanceItems() {
	_jsii_.InvokeVoid(
		k,
		"resetValueImportanceItems",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsRelevanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

