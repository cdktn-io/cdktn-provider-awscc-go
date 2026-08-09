// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazonepolicygrant/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference interface {
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
	IncludeChildDomainUnits() interface{}
	SetIncludeChildDomainUnits(val interface{})
	IncludeChildDomainUnitsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProjectProfiles() *[]*string
	SetProjectProfiles(val *[]*string)
	ProjectProfilesInput() *[]*string
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
	ResetIncludeChildDomainUnits()
	ResetProjectProfiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference
type jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) IncludeChildDomainUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildDomainUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) IncludeChildDomainUnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildDomainUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ProjectProfiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ProjectProfilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference {
	_init_.Initialize()

	if err := validateNewDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference_Override(d DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetIncludeChildDomainUnits(val interface{}) {
	if err := j.validateSetIncludeChildDomainUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeChildDomainUnits",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetProjectProfiles(val *[]*string) {
	if err := j.validateSetProjectProfilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectProfiles",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ResetIncludeChildDomainUnits() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeChildDomainUnits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ResetProjectProfiles() {
	_jsii_.InvokeVoid(
		d,
		"resetProjectProfiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

