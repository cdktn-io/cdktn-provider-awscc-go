// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference interface {
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

// The jsii proxy struct for GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference
type jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ScaleInAfterInactivityMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInAfterInactivityMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ScaleInAfterInactivityMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInAfterInactivityMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ZeroCapacityStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroCapacityStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ZeroCapacityStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroCapacityStrategyInput",
		&returns,
	)
	return returns
}


func NewGameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftFleet.GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference_Override(g GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftFleet.GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetScaleInAfterInactivityMinutes(val *float64) {
	if err := j.validateSetScaleInAfterInactivityMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInAfterInactivityMinutes",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference)SetZeroCapacityStrategy(val *string) {
	if err := j.validateSetZeroCapacityStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroCapacityStrategy",
		val,
	)
}

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ResetScaleInAfterInactivityMinutes() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleInAfterInactivityMinutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ResetZeroCapacityStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroCapacityStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GameliftFleetLocationsLocationCapacityManagedCapacityConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

