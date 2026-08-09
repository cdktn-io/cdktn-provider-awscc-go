// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftcontainerfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerFleetDeploymentConfigurationOutputReference interface {
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
	ImpairmentStrategy() *string
	SetImpairmentStrategy(val *string)
	ImpairmentStrategyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumHealthyPercentage() *float64
	SetMinimumHealthyPercentage(val *float64)
	MinimumHealthyPercentageInput() *float64
	ProtectionStrategy() *string
	SetProtectionStrategy(val *string)
	ProtectionStrategyInput() *string
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
	ResetImpairmentStrategy()
	ResetMinimumHealthyPercentage()
	ResetProtectionStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftContainerFleetDeploymentConfigurationOutputReference
type jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ImpairmentStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impairmentStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ImpairmentStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impairmentStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) MinimumHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) MinimumHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ProtectionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ProtectionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGameliftContainerFleetDeploymentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GameliftContainerFleetDeploymentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftContainerFleetDeploymentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGameliftContainerFleetDeploymentConfigurationOutputReference_Override(g GameliftContainerFleetDeploymentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetImpairmentStrategy(val *string) {
	if err := j.validateSetImpairmentStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impairmentStrategy",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetMinimumHealthyPercentage(val *float64) {
	if err := j.validateSetMinimumHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetProtectionStrategy(val *string) {
	if err := j.validateSetProtectionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectionStrategy",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ResetImpairmentStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetImpairmentStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ResetMinimumHealthyPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumHealthyPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ResetProtectionStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetProtectionStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GameliftContainerFleetDeploymentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

