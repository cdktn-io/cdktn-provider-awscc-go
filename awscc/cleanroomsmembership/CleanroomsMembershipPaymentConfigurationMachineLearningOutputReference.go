// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsmembership/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference interface {
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
	ModelInference() CleanroomsMembershipPaymentConfigurationMachineLearningModelInferenceOutputReference
	ModelInferenceInput() interface{}
	ModelTraining() CleanroomsMembershipPaymentConfigurationMachineLearningModelTrainingOutputReference
	ModelTrainingInput() interface{}
	SyntheticDataGeneration() CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
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
	PutModelInference(value *CleanroomsMembershipPaymentConfigurationMachineLearningModelInference)
	PutModelTraining(value *CleanroomsMembershipPaymentConfigurationMachineLearningModelTraining)
	PutSyntheticDataGeneration(value *CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGeneration)
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

// The jsii proxy struct for CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference
type jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ModelInference() CleanroomsMembershipPaymentConfigurationMachineLearningModelInferenceOutputReference {
	var returns CleanroomsMembershipPaymentConfigurationMachineLearningModelInferenceOutputReference
	_jsii_.Get(
		j,
		"modelInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ModelInferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ModelTraining() CleanroomsMembershipPaymentConfigurationMachineLearningModelTrainingOutputReference {
	var returns CleanroomsMembershipPaymentConfigurationMachineLearningModelTrainingOutputReference
	_jsii_.Get(
		j,
		"modelTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ModelTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) SyntheticDataGeneration() CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference {
	var returns CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGenerationOutputReference
	_jsii_.Get(
		j,
		"syntheticDataGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) SyntheticDataGenerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syntheticDataGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsMembershipPaymentConfigurationMachineLearningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsMembershipPaymentConfigurationMachineLearningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsMembership.CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsMembershipPaymentConfigurationMachineLearningOutputReference_Override(c CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsMembership.CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) PutModelInference(value *CleanroomsMembershipPaymentConfigurationMachineLearningModelInference) {
	if err := c.validatePutModelInferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelInference",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) PutModelTraining(value *CleanroomsMembershipPaymentConfigurationMachineLearningModelTraining) {
	if err := c.validatePutModelTrainingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelTraining",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) PutSyntheticDataGeneration(value *CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGeneration) {
	if err := c.validatePutSyntheticDataGenerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSyntheticDataGeneration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ResetModelInference() {
	_jsii_.InvokeVoid(
		c,
		"resetModelInference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ResetModelTraining() {
	_jsii_.InvokeVoid(
		c,
		"resetModelTraining",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ResetSyntheticDataGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetSyntheticDataGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsMembershipPaymentConfigurationMachineLearningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

