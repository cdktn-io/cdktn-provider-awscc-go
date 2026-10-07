// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/scndataintegrationflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference interface {
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
	FieldPriority() ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyFieldPriorityOutputReference
	FieldPriorityInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutFieldPriority(value *ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyFieldPriority)
	ResetFieldPriority()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference
type jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) FieldPriority() ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyFieldPriorityOutputReference {
	var returns ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyFieldPriorityOutputReference
	_jsii_.Get(
		j,
		"fieldPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) FieldPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.scnDataIntegrationFlow.ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference_Override(s ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.scnDataIntegrationFlow.ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) PutFieldPriority(value *ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyFieldPriority) {
	if err := s.validatePutFieldPriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFieldPriority",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ResetFieldPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_ScnDataIntegrationFlowSourcesDatasetSourceOptionsDedupeStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

