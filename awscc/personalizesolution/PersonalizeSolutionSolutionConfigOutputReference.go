// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/personalizesolution/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeSolutionSolutionConfigOutputReference interface {
	cdktn.ComplexObject
	AlgorithmHyperParameters() *map[string]*string
	SetAlgorithmHyperParameters(val *map[string]*string)
	AlgorithmHyperParametersInput() *map[string]*string
	AutoMlConfig() PersonalizeSolutionSolutionConfigAutoMlConfigOutputReference
	AutoMlConfigInput() interface{}
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
	EventValueThreshold() *string
	SetEventValueThreshold(val *string)
	EventValueThresholdInput() *string
	FeatureTransformationParameters() *map[string]*string
	SetFeatureTransformationParameters(val *map[string]*string)
	FeatureTransformationParametersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	HpoConfig() PersonalizeSolutionSolutionConfigHpoConfigOutputReference
	HpoConfigInput() interface{}
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
	PutAutoMlConfig(value *PersonalizeSolutionSolutionConfigAutoMlConfig)
	PutHpoConfig(value *PersonalizeSolutionSolutionConfigHpoConfig)
	ResetAlgorithmHyperParameters()
	ResetAutoMlConfig()
	ResetEventValueThreshold()
	ResetFeatureTransformationParameters()
	ResetHpoConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersonalizeSolutionSolutionConfigOutputReference
type jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) AlgorithmHyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"algorithmHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) AlgorithmHyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"algorithmHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) AutoMlConfig() PersonalizeSolutionSolutionConfigAutoMlConfigOutputReference {
	var returns PersonalizeSolutionSolutionConfigAutoMlConfigOutputReference
	_jsii_.Get(
		j,
		"autoMlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) AutoMlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) EventValueThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventValueThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) EventValueThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventValueThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) FeatureTransformationParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"featureTransformationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) FeatureTransformationParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"featureTransformationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) HpoConfig() PersonalizeSolutionSolutionConfigHpoConfigOutputReference {
	var returns PersonalizeSolutionSolutionConfigHpoConfigOutputReference
	_jsii_.Get(
		j,
		"hpoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) HpoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hpoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPersonalizeSolutionSolutionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PersonalizeSolutionSolutionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPersonalizeSolutionSolutionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeSolution.PersonalizeSolutionSolutionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersonalizeSolutionSolutionConfigOutputReference_Override(p PersonalizeSolutionSolutionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeSolution.PersonalizeSolutionSolutionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetAlgorithmHyperParameters(val *map[string]*string) {
	if err := j.validateSetAlgorithmHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithmHyperParameters",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetEventValueThreshold(val *string) {
	if err := j.validateSetEventValueThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventValueThreshold",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetFeatureTransformationParameters(val *map[string]*string) {
	if err := j.validateSetFeatureTransformationParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureTransformationParameters",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) PutAutoMlConfig(value *PersonalizeSolutionSolutionConfigAutoMlConfig) {
	if err := p.validatePutAutoMlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAutoMlConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) PutHpoConfig(value *PersonalizeSolutionSolutionConfigHpoConfig) {
	if err := p.validatePutHpoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHpoConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ResetAlgorithmHyperParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetAlgorithmHyperParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ResetAutoMlConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoMlConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ResetEventValueThreshold() {
	_jsii_.InvokeVoid(
		p,
		"resetEventValueThreshold",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ResetFeatureTransformationParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetFeatureTransformationParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ResetHpoConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetHpoConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

