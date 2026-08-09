// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomscollaboration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference interface {
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
	ModelInference() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInferenceOutputReference
	ModelInferenceInput() interface{}
	ModelTraining() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTrainingOutputReference
	ModelTrainingInput() interface{}
	SyntheticDataGeneration() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
	SyntheticDataGenerationInput() interface{}
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
	PutModelInference(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInference)
	PutModelTraining(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTraining)
	PutSyntheticDataGeneration(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGeneration)
	ResetModelInference()
	ResetModelTraining()
	ResetSyntheticDataGeneration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference
type jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ModelInference() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInferenceOutputReference {
	var returns CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInferenceOutputReference
	_jsii_.Get(
		j,
		"modelInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ModelInferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ModelTraining() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTrainingOutputReference {
	var returns CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTrainingOutputReference
	_jsii_.Get(
		j,
		"modelTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ModelTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) SyntheticDataGeneration() CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference {
	var returns CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
	_jsii_.Get(
		j,
		"syntheticDataGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) SyntheticDataGenerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syntheticDataGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference_Override(c CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) PutModelInference(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInference) {
	if err := c.validatePutModelInferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelInference",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) PutModelTraining(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTraining) {
	if err := c.validatePutModelTrainingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelTraining",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) PutSyntheticDataGeneration(value *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGeneration) {
	if err := c.validatePutSyntheticDataGenerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSyntheticDataGeneration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ResetModelInference() {
	_jsii_.InvokeVoid(
		c,
		"resetModelInference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ResetModelTraining() {
	_jsii_.InvokeVoid(
		c,
		"resetModelTraining",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ResetSyntheticDataGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetSyntheticDataGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationMembersPaymentConfigurationMachineLearningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

