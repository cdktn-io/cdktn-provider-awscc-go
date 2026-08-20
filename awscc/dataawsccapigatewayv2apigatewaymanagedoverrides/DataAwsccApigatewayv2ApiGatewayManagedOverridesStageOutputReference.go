// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccapigatewayv2apigatewaymanagedoverrides

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccapigatewayv2apigatewaymanagedoverrides/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference interface {
	cdktn.ComplexObject
	AccessLogSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference
	AutoDeploy() cdktn.IResolvable
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
	DefaultRouteSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference
	Description() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccApigatewayv2ApiGatewayManagedOverridesStage
	SetInternalValue(val *DataAwsccApigatewayv2ApiGatewayManagedOverridesStage)
	RouteSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
	StageVariables() cdktn.StringMap
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference
type jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) AccessLogSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference {
	var returns DataAwsccApigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference
	_jsii_.Get(
		j,
		"accessLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) AutoDeploy() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"autoDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) DefaultRouteSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference {
	var returns DataAwsccApigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) InternalValue() *DataAwsccApigatewayv2ApiGatewayManagedOverridesStage {
	var returns *DataAwsccApigatewayv2ApiGatewayManagedOverridesStage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) RouteSettings() DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap {
	var returns DataAwsccApigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) StageVariables() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccApigatewayv2ApiGatewayManagedOverrides.DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference_Override(d DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccApigatewayv2ApiGatewayManagedOverrides.DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetInternalValue(val *DataAwsccApigatewayv2ApiGatewayManagedOverridesStage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccApigatewayv2ApiGatewayManagedOverridesStageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

