// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizerecommender

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/personalizerecommender/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeRecommenderRecommenderConfigOutputReference interface {
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
	EnableMetadataWithRecommendations() interface{}
	SetEnableMetadataWithRecommendations(val interface{})
	EnableMetadataWithRecommendationsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ItemExplorationConfig() *map[string]*string
	SetItemExplorationConfig(val *map[string]*string)
	ItemExplorationConfigInput() *map[string]*string
	MinRecommendationRequestsPerSecond() *float64
	SetMinRecommendationRequestsPerSecond(val *float64)
	MinRecommendationRequestsPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrainingDataConfig() PersonalizeRecommenderRecommenderConfigTrainingDataConfigOutputReference
	TrainingDataConfigInput() interface{}
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
	PutTrainingDataConfig(value *PersonalizeRecommenderRecommenderConfigTrainingDataConfig)
	ResetEnableMetadataWithRecommendations()
	ResetItemExplorationConfig()
	ResetMinRecommendationRequestsPerSecond()
	ResetTrainingDataConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersonalizeRecommenderRecommenderConfigOutputReference
type jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) EnableMetadataWithRecommendations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMetadataWithRecommendations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) EnableMetadataWithRecommendationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMetadataWithRecommendationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ItemExplorationConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"itemExplorationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ItemExplorationConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"itemExplorationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) MinRecommendationRequestsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendationRequestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) MinRecommendationRequestsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendationRequestsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) TrainingDataConfig() PersonalizeRecommenderRecommenderConfigTrainingDataConfigOutputReference {
	var returns PersonalizeRecommenderRecommenderConfigTrainingDataConfigOutputReference
	_jsii_.Get(
		j,
		"trainingDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) TrainingDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingDataConfigInput",
		&returns,
	)
	return returns
}


func NewPersonalizeRecommenderRecommenderConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PersonalizeRecommenderRecommenderConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPersonalizeRecommenderRecommenderConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeRecommender.PersonalizeRecommenderRecommenderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersonalizeRecommenderRecommenderConfigOutputReference_Override(p PersonalizeRecommenderRecommenderConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeRecommender.PersonalizeRecommenderRecommenderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetEnableMetadataWithRecommendations(val interface{}) {
	if err := j.validateSetEnableMetadataWithRecommendationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMetadataWithRecommendations",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetItemExplorationConfig(val *map[string]*string) {
	if err := j.validateSetItemExplorationConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"itemExplorationConfig",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetMinRecommendationRequestsPerSecond(val *float64) {
	if err := j.validateSetMinRecommendationRequestsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRecommendationRequestsPerSecond",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) PutTrainingDataConfig(value *PersonalizeRecommenderRecommenderConfigTrainingDataConfig) {
	if err := p.validatePutTrainingDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTrainingDataConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ResetEnableMetadataWithRecommendations() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableMetadataWithRecommendations",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ResetItemExplorationConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetItemExplorationConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ResetMinRecommendationRequestsPerSecond() {
	_jsii_.InvokeVoid(
		p,
		"resetMinRecommendationRequestsPerSecond",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ResetTrainingDataConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetTrainingDataConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeRecommenderRecommenderConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

