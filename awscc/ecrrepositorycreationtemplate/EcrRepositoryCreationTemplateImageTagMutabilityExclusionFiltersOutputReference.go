// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecrrepositorycreationtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference interface {
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
	ImageTagMutabilityExclusionFilterType() *string
	SetImageTagMutabilityExclusionFilterType(val *string)
	ImageTagMutabilityExclusionFilterTypeInput() *string
	ImageTagMutabilityExclusionFilterValue() *string
	SetImageTagMutabilityExclusionFilterValue(val *string)
	ImageTagMutabilityExclusionFilterValueInput() *string
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
	ResetImageTagMutabilityExclusionFilterType()
	ResetImageTagMutabilityExclusionFilterValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference
type jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ImageTagMutabilityExclusionFilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutabilityExclusionFilterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ImageTagMutabilityExclusionFilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutabilityExclusionFilterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ImageTagMutabilityExclusionFilterValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutabilityExclusionFilterValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ImageTagMutabilityExclusionFilterValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutabilityExclusionFilterValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewEcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference_Override(e EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetImageTagMutabilityExclusionFilterType(val *string) {
	if err := j.validateSetImageTagMutabilityExclusionFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageTagMutabilityExclusionFilterType",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetImageTagMutabilityExclusionFilterValue(val *string) {
	if err := j.validateSetImageTagMutabilityExclusionFilterValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageTagMutabilityExclusionFilterValue",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ResetImageTagMutabilityExclusionFilterType() {
	_jsii_.InvokeVoid(
		e,
		"resetImageTagMutabilityExclusionFilterType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ResetImageTagMutabilityExclusionFilterValue() {
	_jsii_.InvokeVoid(
		e,
		"resetImageTagMutabilityExclusionFilterValue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplateImageTagMutabilityExclusionFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

