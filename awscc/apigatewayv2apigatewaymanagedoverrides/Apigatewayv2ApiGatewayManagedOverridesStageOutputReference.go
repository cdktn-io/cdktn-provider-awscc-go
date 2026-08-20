// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewayv2apigatewaymanagedoverrides/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ApiGatewayManagedOverridesStageOutputReference interface {
	cdktn.ComplexObject
	AccessLogSettings() Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference
	AccessLogSettingsInput() interface{}
	AutoDeploy() interface{}
	SetAutoDeploy(val interface{})
	AutoDeployInput() interface{}
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
	DefaultRouteSettings() Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference
	DefaultRouteSettingsInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RouteSettings() Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
	RouteSettingsInput() interface{}
	StageVariables() *map[string]*string
	SetStageVariables(val *map[string]*string)
	StageVariablesInput() *map[string]*string
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
	PutAccessLogSettings(value *Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettings)
	PutDefaultRouteSettings(value *Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings)
	PutRouteSettings(value interface{})
	ResetAccessLogSettings()
	ResetAutoDeploy()
	ResetDefaultRouteSettings()
	ResetDescription()
	ResetRouteSettings()
	ResetStageVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Apigatewayv2ApiGatewayManagedOverridesStageOutputReference
type jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) AccessLogSettings() Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference {
	var returns Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference
	_jsii_.Get(
		j,
		"accessLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) AccessLogSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) AutoDeploy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) AutoDeployInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) DefaultRouteSettings() Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference {
	var returns Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) DefaultRouteSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRouteSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) RouteSettings() Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap {
	var returns Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) RouteSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) StageVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) StageVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigatewayv2ApiGatewayManagedOverridesStageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Apigatewayv2ApiGatewayManagedOverridesStageOutputReference {
	_init_.Initialize()

	if err := validateNewApigatewayv2ApiGatewayManagedOverridesStageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayv2ApiGatewayManagedOverridesStageOutputReference_Override(a Apigatewayv2ApiGatewayManagedOverridesStageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetAutoDeploy(val interface{}) {
	if err := j.validateSetAutoDeployParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeploy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetStageVariables(val *map[string]*string) {
	if err := j.validateSetStageVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageVariables",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) PutAccessLogSettings(value *Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettings) {
	if err := a.validatePutAccessLogSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLogSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) PutDefaultRouteSettings(value *Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings) {
	if err := a.validatePutDefaultRouteSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultRouteSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) PutRouteSettings(value interface{}) {
	if err := a.validatePutRouteSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRouteSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetAccessLogSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetAutoDeploy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDeploy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetDefaultRouteSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRouteSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetRouteSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ResetStageVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetStageVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

