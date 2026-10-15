// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightbrand/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference interface {
	cdktn.ComplexObject
	Accent() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference
	AccentInput() interface{}
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
	Danger() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference
	DangerInput() interface{}
	Dimension() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference
	DimensionInput() interface{}
	// Experimental.
	Fqn() *string
	Info() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference
	InfoInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Measure() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference
	MeasureInput() interface{}
	Primary() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference
	PrimaryInput() interface{}
	Secondary() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference
	SecondaryInput() interface{}
	Success() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference
	SuccessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Warning() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference
	WarningInput() interface{}
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
	PutAccent(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccent)
	PutDanger(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDanger)
	PutDimension(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimension)
	PutInfo(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfo)
	PutMeasure(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasure)
	PutPrimary(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimary)
	PutSecondary(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondary)
	PutSuccess(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccess)
	PutWarning(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarning)
	ResetAccent()
	ResetDanger()
	ResetDimension()
	ResetInfo()
	ResetMeasure()
	ResetPrimary()
	ResetSecondary()
	ResetSuccess()
	ResetWarning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference
type jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Accent() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccentOutputReference
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) AccentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Danger() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDangerOutputReference
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) DangerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dangerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Dimension() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) DimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Info() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfoOutputReference
	_jsii_.Get(
		j,
		"info",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Measure() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasureOutputReference
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) MeasureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Primary() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimaryOutputReference
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Secondary() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondaryOutputReference
	_jsii_.Get(
		j,
		"secondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) SecondaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Success() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccessOutputReference
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) SuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Warning() QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference {
	var returns QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarningOutputReference
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) WarningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


func NewQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightBrand.QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference_Override(q QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightBrand.QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutAccent(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccent) {
	if err := q.validatePutAccentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAccent",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutDanger(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDanger) {
	if err := q.validatePutDangerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDanger",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutDimension(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimension) {
	if err := q.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDimension",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutInfo(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfo) {
	if err := q.validatePutInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInfo",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutMeasure(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasure) {
	if err := q.validatePutMeasureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMeasure",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutPrimary(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimary) {
	if err := q.validatePutPrimaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPrimary",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutSecondary(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondary) {
	if err := q.validatePutSecondaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSecondary",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutSuccess(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccess) {
	if err := q.validatePutSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSuccess",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) PutWarning(value *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarning) {
	if err := q.validatePutWarningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putWarning",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetAccent() {
	_jsii_.InvokeVoid(
		q,
		"resetAccent",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetDanger() {
	_jsii_.InvokeVoid(
		q,
		"resetDanger",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		q,
		"resetDimension",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetInfo() {
	_jsii_.InvokeVoid(
		q,
		"resetInfo",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetMeasure() {
	_jsii_.InvokeVoid(
		q,
		"resetMeasure",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		q,
		"resetPrimary",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetSecondary() {
	_jsii_.InvokeVoid(
		q,
		"resetSecondary",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetSuccess() {
	_jsii_.InvokeVoid(
		q,
		"resetSuccess",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		q,
		"resetWarning",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

