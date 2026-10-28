// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/sagemakercomputequota/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference interface {
	cdktn.ComplexObject
	AbsoluteBorrowLimits() SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsList
	AbsoluteBorrowLimitsInput() interface{}
	BorrowLimit() *float64
	SetBorrowLimit(val *float64)
	BorrowLimitInput() *float64
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
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	PutAbsoluteBorrowLimits(value interface{})
	ResetAbsoluteBorrowLimits()
	ResetBorrowLimit()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference
type jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) AbsoluteBorrowLimits() SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsList {
	var returns SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsList
	_jsii_.Get(
		j,
		"absoluteBorrowLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) AbsoluteBorrowLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"absoluteBorrowLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) BorrowLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"borrowLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) BorrowLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"borrowLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerComputeQuota.SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference_Override(s SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerComputeQuota.SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetBorrowLimit(val *float64) {
	if err := j.validateSetBorrowLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"borrowLimit",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) PutAbsoluteBorrowLimits(value interface{}) {
	if err := s.validatePutAbsoluteBorrowLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbsoluteBorrowLimits",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ResetAbsoluteBorrowLimits() {
	_jsii_.InvokeVoid(
		s,
		"resetAbsoluteBorrowLimits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ResetBorrowLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetBorrowLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

