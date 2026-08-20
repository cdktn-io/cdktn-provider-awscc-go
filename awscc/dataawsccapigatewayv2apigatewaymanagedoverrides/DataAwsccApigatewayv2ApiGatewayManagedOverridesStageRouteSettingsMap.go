// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccapigatewayv2apigatewaymanagedoverrides

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccapigatewayv2apigatewaymanagedoverrides/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(key *string) DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference
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

// The jsii proxy struct for DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
type jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap {
	_init_.Initialize()

	if err := validateNewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccApigatewayv2ApiGatewayManagedOverrides.DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap_Override(d DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccApigatewayv2ApiGatewayManagedOverrides.DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Get(key *string) DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

