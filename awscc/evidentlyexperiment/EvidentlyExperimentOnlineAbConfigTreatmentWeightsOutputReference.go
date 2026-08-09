// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evidentlyexperiment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference interface {
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
	SplitWeight() *float64
	SetSplitWeight(val *float64)
	SplitWeightInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Treatment() *string
	SetTreatment(val *string)
	TreatmentInput() *string
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
	ResetSplitWeight()
	ResetTreatment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference
type jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) SplitWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"splitWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) SplitWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"splitWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) Treatment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) TreatmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatmentInput",
		&returns,
	)
	return returns
}


func NewEvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference {
	_init_.Initialize()

	if err := validateNewEvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyExperiment.EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference_Override(e EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyExperiment.EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetSplitWeight(val *float64) {
	if err := j.validateSetSplitWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitWeight",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference)SetTreatment(val *string) {
	if err := j.validateSetTreatmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatment",
		val,
	)
}

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ResetSplitWeight() {
	_jsii_.InvokeVoid(
		e,
		"resetSplitWeight",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ResetTreatment() {
	_jsii_.InvokeVoid(
		e,
		"resetTreatment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EvidentlyExperimentOnlineAbConfigTreatmentWeightsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

