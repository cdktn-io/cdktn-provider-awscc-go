// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewayv2apigatewaymanagedoverrides/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(key *string) Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
type jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap {
	_init_.Initialize()

	if err := validateNewApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap_Override(a Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Get(key *string) Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference {
	if err := a.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

