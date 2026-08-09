// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomscollaboration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference interface {
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
	ModelInference() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelInferenceOutputReference
	ModelInferenceInput() interface{}
	ModelTraining() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelTrainingOutputReference
	ModelTrainingInput() interface{}
	SyntheticDataGeneration() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
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
	PutModelInference(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelInference)
	PutModelTraining(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelTraining)
	PutSyntheticDataGeneration(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningSyntheticDataGeneration)
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

// The jsii proxy struct for CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference
type jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ModelInference() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelInferenceOutputReference {
	var returns CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelInferenceOutputReference
	_jsii_.Get(
		j,
		"modelInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ModelInferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ModelTraining() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelTrainingOutputReference {
	var returns CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelTrainingOutputReference
	_jsii_.Get(
		j,
		"modelTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ModelTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) SyntheticDataGeneration() CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference {
	var returns CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
	_jsii_.Get(
		j,
		"syntheticDataGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) SyntheticDataGenerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syntheticDataGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference_Override(c CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsCollaboration.CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) PutModelInference(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelInference) {
	if err := c.validatePutModelInferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelInference",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) PutModelTraining(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningModelTraining) {
	if err := c.validatePutModelTrainingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelTraining",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) PutSyntheticDataGeneration(value *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningSyntheticDataGeneration) {
	if err := c.validatePutSyntheticDataGenerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSyntheticDataGeneration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ResetModelInference() {
	_jsii_.InvokeVoid(
		c,
		"resetModelInference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ResetModelTraining() {
	_jsii_.InvokeVoid(
		c,
		"resetModelTraining",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ResetSyntheticDataGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetSyntheticDataGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsCollaborationCreatorPaymentConfigurationMachineLearningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

