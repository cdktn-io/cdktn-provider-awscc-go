// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftcontainerfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference interface {
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
	ScaleInAfterInactivityMinutes() *float64
	SetScaleInAfterInactivityMinutes(val *float64)
	ScaleInAfterInactivityMinutesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ZeroCapacityStrategy() *string
	SetZeroCapacityStrategy(val *string)
	ZeroCapacityStrategyInput() *string
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
	ResetScaleInAfterInactivityMinutes()
	ResetZeroCapacityStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference
type jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ScaleInAfterInactivityMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInAfterInactivityMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ScaleInAfterInactivityMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInAfterInactivityMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ZeroCapacityStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroCapacityStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ZeroCapacityStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroCapacityStrategyInput",
		&returns,
	)
	return returns
}


func NewGameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference_Override(g GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetScaleInAfterInactivityMinutes(val *float64) {
	if err := j.validateSetScaleInAfterInactivityMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInAfterInactivityMinutes",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetZeroCapacityStrategy(val *string) {
	if err := j.validateSetZeroCapacityStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroCapacityStrategy",
		val,
	)
}

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ResetScaleInAfterInactivityMinutes() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleInAfterInactivityMinutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ResetZeroCapacityStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroCapacityStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

