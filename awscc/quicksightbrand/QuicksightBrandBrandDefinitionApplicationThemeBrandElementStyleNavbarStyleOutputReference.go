// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/quicksightbrand/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference interface {
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
	ContextualNavbar() QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbarOutputReference
	ContextualNavbarInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GlobalNavbar() QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbarOutputReference
	GlobalNavbarInput() interface{}
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
	PutContextualNavbar(value *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbar)
	PutGlobalNavbar(value *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbar)
	ResetContextualNavbar()
	ResetGlobalNavbar()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference
type jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ContextualNavbar() QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbarOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbarOutputReference
	_jsii_.Get(
		j,
		"contextualNavbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ContextualNavbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextualNavbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GlobalNavbar() QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbarOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbarOutputReference
	_jsii_.Get(
		j,
		"globalNavbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GlobalNavbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalNavbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightBrand.QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference_Override(q QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightBrand.QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) PutContextualNavbar(value *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbar) {
	if err := q.validatePutContextualNavbarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putContextualNavbar",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) PutGlobalNavbar(value *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbar) {
	if err := q.validatePutGlobalNavbarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putGlobalNavbar",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ResetContextualNavbar() {
	_jsii_.InvokeVoid(
		q,
		"resetContextualNavbar",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ResetGlobalNavbar() {
	_jsii_.InvokeVoid(
		q,
		"resetGlobalNavbar",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

