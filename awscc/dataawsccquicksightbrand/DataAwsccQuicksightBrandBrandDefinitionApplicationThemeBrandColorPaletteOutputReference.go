// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksightbrand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccquicksightbrand/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference interface {
	cdktn.ComplexObject
	Accent() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference
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
	Danger() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference
	Dimension() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference
	// Experimental.
	Fqn() *string
	Info() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference
	InternalValue() *DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette
	SetInternalValue(val *DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette)
	Measure() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference
	Primary() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference
	Secondary() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference
	Success() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Warning() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference
type jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Accent() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Danger() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Dimension() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Info() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference
	_jsii_.Get(
		j,
		"info",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InternalValue() *DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette {
	var returns *DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Measure() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Primary() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Secondary() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference
	_jsii_.Get(
		j,
		"secondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Success() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Warning() DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference {
	var returns DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightBrand.DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference_Override(d DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightBrand.DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetInternalValue(val *DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

