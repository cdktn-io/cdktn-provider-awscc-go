// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/scndataintegrationflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	SortOrder() *string
	SetSortOrder(val *string)
	SortOrderInput() *string
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
	ResetName()
	ResetSortOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference
type jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) SortOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) SortOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.scnDataIntegrationFlow.ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference_Override(s ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.scnDataIntegrationFlow.ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetSortOrder(val *string) {
	if err := j.validateSetSortOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortOrder",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ResetSortOrder() {
	_jsii_.InvokeVoid(
		s,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

