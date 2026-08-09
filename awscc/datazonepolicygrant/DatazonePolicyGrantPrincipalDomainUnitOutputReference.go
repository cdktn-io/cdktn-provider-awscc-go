// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazonepolicygrant/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazonePolicyGrantPrincipalDomainUnitOutputReference interface {
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
	DomainUnitDesignation() *string
	SetDomainUnitDesignation(val *string)
	DomainUnitDesignationInput() *string
	DomainUnitGrantFilter() DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilterOutputReference
	DomainUnitGrantFilterInput() interface{}
	DomainUnitIdentifier() *string
	SetDomainUnitIdentifier(val *string)
	DomainUnitIdentifierInput() *string
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
	PutDomainUnitGrantFilter(value *DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilter)
	ResetDomainUnitDesignation()
	ResetDomainUnitGrantFilter()
	ResetDomainUnitIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazonePolicyGrantPrincipalDomainUnitOutputReference
type jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitDesignation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUnitDesignation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitDesignationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUnitDesignationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitGrantFilter() DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilterOutputReference {
	var returns DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilterOutputReference
	_jsii_.Get(
		j,
		"domainUnitGrantFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitGrantFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainUnitGrantFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUnitIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) DomainUnitIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUnitIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazonePolicyGrantPrincipalDomainUnitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazonePolicyGrantPrincipalDomainUnitOutputReference {
	_init_.Initialize()

	if err := validateNewDatazonePolicyGrantPrincipalDomainUnitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantPrincipalDomainUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazonePolicyGrantPrincipalDomainUnitOutputReference_Override(d DatazonePolicyGrantPrincipalDomainUnitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantPrincipalDomainUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetDomainUnitDesignation(val *string) {
	if err := j.validateSetDomainUnitDesignationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainUnitDesignation",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetDomainUnitIdentifier(val *string) {
	if err := j.validateSetDomainUnitIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainUnitIdentifier",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) PutDomainUnitGrantFilter(value *DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilter) {
	if err := d.validatePutDomainUnitGrantFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDomainUnitGrantFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ResetDomainUnitDesignation() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainUnitDesignation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ResetDomainUnitGrantFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainUnitGrantFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ResetDomainUnitIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainUnitIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantPrincipalDomainUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

