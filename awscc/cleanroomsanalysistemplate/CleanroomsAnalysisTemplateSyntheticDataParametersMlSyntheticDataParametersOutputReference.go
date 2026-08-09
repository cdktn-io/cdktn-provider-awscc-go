// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsanalysistemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference interface {
	cdktn.ComplexObject
	ColumnClassification() CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassificationOutputReference
	ColumnClassificationInput() interface{}
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
	Epsilon() *float64
	SetEpsilon(val *float64)
	EpsilonInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxMembershipInferenceAttackScore() *float64
	SetMaxMembershipInferenceAttackScore(val *float64)
	MaxMembershipInferenceAttackScoreInput() *float64
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
	PutColumnClassification(value *CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassification)
	ResetColumnClassification()
	ResetEpsilon()
	ResetMaxMembershipInferenceAttackScore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference
type jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ColumnClassification() CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassificationOutputReference {
	var returns CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassificationOutputReference
	_jsii_.Get(
		j,
		"columnClassification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ColumnClassificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnClassificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) Epsilon() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epsilon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) EpsilonInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epsilonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) MaxMembershipInferenceAttackScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMembershipInferenceAttackScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) MaxMembershipInferenceAttackScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMembershipInferenceAttackScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference_Override(c CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetEpsilon(val *float64) {
	if err := j.validateSetEpsilonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epsilon",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetMaxMembershipInferenceAttackScore(val *float64) {
	if err := j.validateSetMaxMembershipInferenceAttackScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMembershipInferenceAttackScore",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) PutColumnClassification(value *CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassification) {
	if err := c.validatePutColumnClassificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putColumnClassification",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ResetColumnClassification() {
	_jsii_.InvokeVoid(
		c,
		"resetColumnClassification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ResetEpsilon() {
	_jsii_.InvokeVoid(
		c,
		"resetEpsilon",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ResetMaxMembershipInferenceAttackScore() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxMembershipInferenceAttackScore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

