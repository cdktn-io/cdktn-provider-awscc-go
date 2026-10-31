// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/iototaupdate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference interface {
	cdktn.ComplexObject
	BaseRatePerMinute() *float64
	SetBaseRatePerMinute(val *float64)
	BaseRatePerMinuteInput() *float64
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
	IncrementFactor() *float64
	SetIncrementFactor(val *float64)
	IncrementFactorInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RateIncreaseCriteria() IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference
	RateIncreaseCriteriaInput() interface{}
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
	PutRateIncreaseCriteria(value *IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria)
	ResetBaseRatePerMinute()
	ResetIncrementFactor()
	ResetRateIncreaseCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference
type jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) BaseRatePerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseRatePerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) BaseRatePerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseRatePerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) IncrementFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) IncrementFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) RateIncreaseCriteria() IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference {
	var returns IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference
	_jsii_.Get(
		j,
		"rateIncreaseCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) RateIncreaseCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateIncreaseCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference {
	_init_.Initialize()

	if err := validateNewIotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference_Override(i IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetBaseRatePerMinute(val *float64) {
	if err := j.validateSetBaseRatePerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseRatePerMinute",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetIncrementFactor(val *float64) {
	if err := j.validateSetIncrementFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementFactor",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) PutRateIncreaseCriteria(value *IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria) {
	if err := i.validatePutRateIncreaseCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRateIncreaseCriteria",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ResetBaseRatePerMinute() {
	_jsii_.InvokeVoid(
		i,
		"resetBaseRatePerMinute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ResetIncrementFactor() {
	_jsii_.InvokeVoid(
		i,
		"resetIncrementFactor",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ResetRateIncreaseCriteria() {
	_jsii_.InvokeVoid(
		i,
		"resetRateIncreaseCriteria",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

