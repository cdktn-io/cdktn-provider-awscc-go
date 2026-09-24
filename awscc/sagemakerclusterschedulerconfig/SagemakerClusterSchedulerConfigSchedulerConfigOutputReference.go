// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerclusterschedulerconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerclusterschedulerconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterSchedulerConfigSchedulerConfigOutputReference interface {
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
	FairShare() *string
	SetFairShare(val *string)
	FairShareInput() *string
	// Experimental.
	Fqn() *string
	IdleResourceSharing() *string
	SetIdleResourceSharing(val *string)
	IdleResourceSharingInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PriorityClasses() SagemakerClusterSchedulerConfigSchedulerConfigPriorityClassesList
	PriorityClassesInput() interface{}
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
	PutPriorityClasses(value interface{})
	ResetFairShare()
	ResetIdleResourceSharing()
	ResetPriorityClasses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerClusterSchedulerConfigSchedulerConfigOutputReference
type jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) FairShare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fairShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) FairShareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fairShareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) IdleResourceSharing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleResourceSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) IdleResourceSharingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleResourceSharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) PriorityClasses() SagemakerClusterSchedulerConfigSchedulerConfigPriorityClassesList {
	var returns SagemakerClusterSchedulerConfigSchedulerConfigPriorityClassesList
	_jsii_.Get(
		j,
		"priorityClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) PriorityClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"priorityClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerClusterSchedulerConfigSchedulerConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerClusterSchedulerConfigSchedulerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterSchedulerConfigSchedulerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerClusterSchedulerConfig.SagemakerClusterSchedulerConfigSchedulerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerClusterSchedulerConfigSchedulerConfigOutputReference_Override(s SagemakerClusterSchedulerConfigSchedulerConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerClusterSchedulerConfig.SagemakerClusterSchedulerConfigSchedulerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetFairShare(val *string) {
	if err := j.validateSetFairShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fairShare",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetIdleResourceSharing(val *string) {
	if err := j.validateSetIdleResourceSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleResourceSharing",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) PutPriorityClasses(value interface{}) {
	if err := s.validatePutPriorityClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPriorityClasses",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ResetFairShare() {
	_jsii_.InvokeVoid(
		s,
		"resetFairShare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ResetIdleResourceSharing() {
	_jsii_.InvokeVoid(
		s,
		"resetIdleResourceSharing",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ResetPriorityClasses() {
	_jsii_.InvokeVoid(
		s,
		"resetPriorityClasses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerClusterSchedulerConfigSchedulerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

