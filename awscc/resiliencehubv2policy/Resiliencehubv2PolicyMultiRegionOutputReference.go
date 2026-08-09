// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/resiliencehubv2policy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2PolicyMultiRegionOutputReference interface {
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
	DisasterRecoveryApproach() *string
	SetDisasterRecoveryApproach(val *string)
	DisasterRecoveryApproachInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RpoInMinutes() *float64
	SetRpoInMinutes(val *float64)
	RpoInMinutesInput() *float64
	RtoInMinutes() *float64
	SetRtoInMinutes(val *float64)
	RtoInMinutesInput() *float64
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
	ResetDisasterRecoveryApproach()
	ResetRpoInMinutes()
	ResetRtoInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Resiliencehubv2PolicyMultiRegionOutputReference
type jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) DisasterRecoveryApproach() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disasterRecoveryApproach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) DisasterRecoveryApproachInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disasterRecoveryApproachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) RpoInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rpoInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) RpoInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rpoInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) RtoInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rtoInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) RtoInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rtoInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResiliencehubv2PolicyMultiRegionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Resiliencehubv2PolicyMultiRegionOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubv2PolicyMultiRegionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.resiliencehubv2Policy.Resiliencehubv2PolicyMultiRegionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResiliencehubv2PolicyMultiRegionOutputReference_Override(r Resiliencehubv2PolicyMultiRegionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.resiliencehubv2Policy.Resiliencehubv2PolicyMultiRegionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetDisasterRecoveryApproach(val *string) {
	if err := j.validateSetDisasterRecoveryApproachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disasterRecoveryApproach",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetRpoInMinutes(val *float64) {
	if err := j.validateSetRpoInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rpoInMinutes",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetRtoInMinutes(val *float64) {
	if err := j.validateSetRtoInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rtoInMinutes",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ResetDisasterRecoveryApproach() {
	_jsii_.InvokeVoid(
		r,
		"resetDisasterRecoveryApproach",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ResetRpoInMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetRpoInMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ResetRtoInMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetRtoInMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_Resiliencehubv2PolicyMultiRegionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

