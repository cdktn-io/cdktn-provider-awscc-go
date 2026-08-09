// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendraindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraIndexDocumentMetadataConfigurationsSearchOutputReference interface {
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
	Displayable() interface{}
	SetDisplayable(val interface{})
	DisplayableInput() interface{}
	Facetable() interface{}
	SetFacetable(val interface{})
	FacetableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Searchable() interface{}
	SetSearchable(val interface{})
	SearchableInput() interface{}
	Sortable() interface{}
	SetSortable(val interface{})
	SortableInput() interface{}
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
	ResetDisplayable()
	ResetFacetable()
	ResetSearchable()
	ResetSortable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexDocumentMetadataConfigurationsSearchOutputReference
type jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Displayable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"displayable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) DisplayableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"displayableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Facetable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) FacetableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Searchable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) SearchableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Sortable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) SortableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraIndexDocumentMetadataConfigurationsSearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraIndexDocumentMetadataConfigurationsSearchOutputReference {
	_init_.Initialize()

	if err := validateNewKendraIndexDocumentMetadataConfigurationsSearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraIndex.KendraIndexDocumentMetadataConfigurationsSearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraIndexDocumentMetadataConfigurationsSearchOutputReference_Override(k KendraIndexDocumentMetadataConfigurationsSearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraIndex.KendraIndexDocumentMetadataConfigurationsSearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetDisplayable(val interface{}) {
	if err := j.validateSetDisplayableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayable",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetFacetable(val interface{}) {
	if err := j.validateSetFacetableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetable",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetSearchable(val interface{}) {
	if err := j.validateSetSearchableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchable",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetSortable(val interface{}) {
	if err := j.validateSetSortableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortable",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ResetDisplayable() {
	_jsii_.InvokeVoid(
		k,
		"resetDisplayable",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ResetFacetable() {
	_jsii_.InvokeVoid(
		k,
		"resetFacetable",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ResetSearchable() {
	_jsii_.InvokeVoid(
		k,
		"resetSearchable",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ResetSortable() {
	_jsii_.InvokeVoid(
		k,
		"resetSortable",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraIndexDocumentMetadataConfigurationsSearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

