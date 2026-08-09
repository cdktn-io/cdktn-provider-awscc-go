// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/greengrassv2deployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference interface {
	cdktn.ComplexObject
	BaseRatePerMinute() *float64
	SetBaseRatePerMinute(val *float64)
	BaseRatePerMinuteInput() *float64
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
	IncrementFactor() *float64
	SetIncrementFactor(val *float64)
	IncrementFactorInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RateIncreaseCriteria() Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference
	RateIncreaseCriteriaInput() interface{}
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
	PutRateIncreaseCriteria(value *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria)
	ResetBaseRatePerMinute()
	ResetIncrementFactor()
	ResetRateIncreaseCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference
type jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) BaseRatePerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseRatePerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) BaseRatePerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseRatePerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) IncrementFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) IncrementFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) RateIncreaseCriteria() Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference {
	var returns Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteriaOutputReference
	_jsii_.Get(
		j,
		"rateIncreaseCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) RateIncreaseCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateIncreaseCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGreengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGreengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference_Override(g Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetBaseRatePerMinute(val *float64) {
	if err := j.validateSetBaseRatePerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseRatePerMinute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetIncrementFactor(val *float64) {
	if err := j.validateSetIncrementFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementFactor",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) PutRateIncreaseCriteria(value *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria) {
	if err := g.validatePutRateIncreaseCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRateIncreaseCriteria",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ResetBaseRatePerMinute() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseRatePerMinute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ResetIncrementFactor() {
	_jsii_.InvokeVoid(
		g,
		"resetIncrementFactor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ResetRateIncreaseCriteria() {
	_jsii_.InvokeVoid(
		g,
		"resetRateIncreaseCriteria",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

