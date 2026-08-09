// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apsanomalydetector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApsAnomalyDetectorConfigurationRandomCutForestOutputReference interface {
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
	IgnoreNearExpectedFromAbove() ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference
	IgnoreNearExpectedFromAboveInput() interface{}
	IgnoreNearExpectedFromBelow() ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference
	IgnoreNearExpectedFromBelowInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	SampleSize() *float64
	SetSampleSize(val *float64)
	SampleSizeInput() *float64
	ShingleSize() *float64
	SetShingleSize(val *float64)
	ShingleSizeInput() *float64
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
	PutIgnoreNearExpectedFromAbove(value *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove)
	PutIgnoreNearExpectedFromBelow(value *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow)
	ResetIgnoreNearExpectedFromAbove()
	ResetIgnoreNearExpectedFromBelow()
	ResetSampleSize()
	ResetShingleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApsAnomalyDetectorConfigurationRandomCutForestOutputReference
type jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromAbove() ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference {
	var returns ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAbove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromAboveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAboveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromBelow() ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference {
	var returns ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromBelowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ShingleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ShingleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApsAnomalyDetectorConfigurationRandomCutForestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApsAnomalyDetectorConfigurationRandomCutForestOutputReference {
	_init_.Initialize()

	if err := validateNewApsAnomalyDetectorConfigurationRandomCutForestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apsAnomalyDetector.ApsAnomalyDetectorConfigurationRandomCutForestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApsAnomalyDetectorConfigurationRandomCutForestOutputReference_Override(a ApsAnomalyDetectorConfigurationRandomCutForestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apsAnomalyDetector.ApsAnomalyDetectorConfigurationRandomCutForestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetSampleSize(val *float64) {
	if err := j.validateSetSampleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetShingleSize(val *float64) {
	if err := j.validateSetShingleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shingleSize",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) PutIgnoreNearExpectedFromAbove(value *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove) {
	if err := a.validatePutIgnoreNearExpectedFromAboveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIgnoreNearExpectedFromAbove",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) PutIgnoreNearExpectedFromBelow(value *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow) {
	if err := a.validatePutIgnoreNearExpectedFromBelowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIgnoreNearExpectedFromBelow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetIgnoreNearExpectedFromAbove() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreNearExpectedFromAbove",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetIgnoreNearExpectedFromBelow() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreNearExpectedFromBelow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetShingleSize() {
	_jsii_.InvokeVoid(
		a,
		"resetShingleSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsAnomalyDetectorConfigurationRandomCutForestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

