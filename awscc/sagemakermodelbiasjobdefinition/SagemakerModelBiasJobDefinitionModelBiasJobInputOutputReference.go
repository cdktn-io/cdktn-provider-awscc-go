// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelbiasjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference interface {
	cdktn.ComplexObject
	BatchTransformInput() SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputOutputReference
	BatchTransformInputInput() interface{}
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
	EndpointInput() SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInputOutputReference
	EndpointInputInput() interface{}
	// Experimental.
	Fqn() *string
	GroundTruthS3Input() SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3InputOutputReference
	GroundTruthS3InputInput() interface{}
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
	PutBatchTransformInput(value *SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInput)
	PutEndpointInput(value *SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInput)
	PutGroundTruthS3Input(value *SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3Input)
	ResetBatchTransformInput()
	ResetEndpointInput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference
type jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) BatchTransformInput() SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputOutputReference {
	var returns SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputOutputReference
	_jsii_.Get(
		j,
		"batchTransformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) BatchTransformInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchTransformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) EndpointInput() SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInputOutputReference {
	var returns SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInputOutputReference
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) EndpointInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GroundTruthS3Input() SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3InputOutputReference {
	var returns SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3InputOutputReference
	_jsii_.Get(
		j,
		"groundTruthS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GroundTruthS3InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groundTruthS3InputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelBiasJobDefinitionModelBiasJobInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelBiasJobDefinition.SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference_Override(s SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelBiasJobDefinition.SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) PutBatchTransformInput(value *SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInput) {
	if err := s.validatePutBatchTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBatchTransformInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) PutEndpointInput(value *SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInput) {
	if err := s.validatePutEndpointInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEndpointInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) PutGroundTruthS3Input(value *SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3Input) {
	if err := s.validatePutGroundTruthS3InputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGroundTruthS3Input",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ResetBatchTransformInput() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchTransformInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ResetEndpointInput() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelBiasJobDefinitionModelBiasJobInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

