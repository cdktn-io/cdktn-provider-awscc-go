// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewayv2apigatewaymanagedoverrides/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	IntegrationMethod() *string
	SetIntegrationMethod(val *string)
	IntegrationMethodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PayloadFormatVersion() *string
	SetPayloadFormatVersion(val *string)
	PayloadFormatVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutInMillis() *float64
	SetTimeoutInMillis(val *float64)
	TimeoutInMillisInput() *float64
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
	ResetDescription()
	ResetIntegrationMethod()
	ResetPayloadFormatVersion()
	ResetTimeoutInMillis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference
type jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) IntegrationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) IntegrationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) PayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) PayloadFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) TimeoutInMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) TimeoutInMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMillisInput",
		&returns,
	)
	return returns
}


func NewApigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference {
	_init_.Initialize()

	if err := validateNewApigatewayv2ApiGatewayManagedOverridesIntegrationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference_Override(a Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetIntegrationMethod(val *string) {
	if err := j.validateSetIntegrationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationMethod",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetPayloadFormatVersion(val *string) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)SetTimeoutInMillis(val *float64) {
	if err := j.validateSetTimeoutInMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMillis",
		val,
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ResetIntegrationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ResetPayloadFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ResetTimeoutInMillis() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInMillis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

