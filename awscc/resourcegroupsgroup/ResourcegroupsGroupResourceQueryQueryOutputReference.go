// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/resourcegroupsgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourcegroupsGroupResourceQueryQueryOutputReference interface {
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
	ResourceTypeFilters() *[]*string
	SetResourceTypeFilters(val *[]*string)
	ResourceTypeFiltersInput() *[]*string
	StackIdentifier() *string
	SetStackIdentifier(val *string)
	StackIdentifierInput() *string
	TagFilters() ResourcegroupsGroupResourceQueryQueryTagFiltersList
	TagFiltersInput() interface{}
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
	PutTagFilters(value interface{})
	ResetResourceTypeFilters()
	ResetStackIdentifier()
	ResetTagFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResourcegroupsGroupResourceQueryQueryOutputReference
type jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ResourceTypeFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ResourceTypeFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) StackIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) StackIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) TagFilters() ResourcegroupsGroupResourceQueryQueryTagFiltersList {
	var returns ResourcegroupsGroupResourceQueryQueryTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResourcegroupsGroupResourceQueryQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ResourcegroupsGroupResourceQueryQueryOutputReference {
	_init_.Initialize()

	if err := validateNewResourcegroupsGroupResourceQueryQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.resourcegroupsGroup.ResourcegroupsGroupResourceQueryQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResourcegroupsGroupResourceQueryQueryOutputReference_Override(r ResourcegroupsGroupResourceQueryQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.resourcegroupsGroup.ResourcegroupsGroupResourceQueryQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetResourceTypeFilters(val *[]*string) {
	if err := j.validateSetResourceTypeFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypeFilters",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetStackIdentifier(val *string) {
	if err := j.validateSetStackIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackIdentifier",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) PutTagFilters(value interface{}) {
	if err := r.validatePutTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTagFilters",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ResetResourceTypeFilters() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceTypeFilters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ResetStackIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetStackIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		r,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

