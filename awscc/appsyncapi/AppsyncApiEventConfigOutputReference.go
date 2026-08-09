// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appsyncapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncApiEventConfigOutputReference interface {
	cdktn.ComplexObject
	AuthProviders() AppsyncApiEventConfigAuthProvidersList
	AuthProvidersInput() interface{}
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
	ConnectionAuthModes() AppsyncApiEventConfigConnectionAuthModesList
	ConnectionAuthModesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPublishAuthModes() AppsyncApiEventConfigDefaultPublishAuthModesList
	DefaultPublishAuthModesInput() interface{}
	DefaultSubscribeAuthModes() AppsyncApiEventConfigDefaultSubscribeAuthModesList
	DefaultSubscribeAuthModesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogConfig() AppsyncApiEventConfigLogConfigOutputReference
	LogConfigInput() interface{}
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
	PutAuthProviders(value interface{})
	PutConnectionAuthModes(value interface{})
	PutDefaultPublishAuthModes(value interface{})
	PutDefaultSubscribeAuthModes(value interface{})
	PutLogConfig(value *AppsyncApiEventConfigLogConfig)
	ResetAuthProviders()
	ResetConnectionAuthModes()
	ResetDefaultPublishAuthModes()
	ResetDefaultSubscribeAuthModes()
	ResetLogConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncApiEventConfigOutputReference
type jsiiProxy_AppsyncApiEventConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) AuthProviders() AppsyncApiEventConfigAuthProvidersList {
	var returns AppsyncApiEventConfigAuthProvidersList
	_jsii_.Get(
		j,
		"authProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) AuthProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ConnectionAuthModes() AppsyncApiEventConfigConnectionAuthModesList {
	var returns AppsyncApiEventConfigConnectionAuthModesList
	_jsii_.Get(
		j,
		"connectionAuthModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ConnectionAuthModesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionAuthModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultPublishAuthModes() AppsyncApiEventConfigDefaultPublishAuthModesList {
	var returns AppsyncApiEventConfigDefaultPublishAuthModesList
	_jsii_.Get(
		j,
		"defaultPublishAuthModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultPublishAuthModesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPublishAuthModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultSubscribeAuthModes() AppsyncApiEventConfigDefaultSubscribeAuthModesList {
	var returns AppsyncApiEventConfigDefaultSubscribeAuthModesList
	_jsii_.Get(
		j,
		"defaultSubscribeAuthModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultSubscribeAuthModesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSubscribeAuthModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) LogConfig() AppsyncApiEventConfigLogConfigOutputReference {
	var returns AppsyncApiEventConfigLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) LogConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncApiEventConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppsyncApiEventConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncApiEventConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncApiEventConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncApi.AppsyncApiEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncApiEventConfigOutputReference_Override(a AppsyncApiEventConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncApi.AppsyncApiEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutAuthProviders(value interface{}) {
	if err := a.validatePutAuthProvidersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthProviders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutConnectionAuthModes(value interface{}) {
	if err := a.validatePutConnectionAuthModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionAuthModes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutDefaultPublishAuthModes(value interface{}) {
	if err := a.validatePutDefaultPublishAuthModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultPublishAuthModes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutDefaultSubscribeAuthModes(value interface{}) {
	if err := a.validatePutDefaultSubscribeAuthModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultSubscribeAuthModes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutLogConfig(value *AppsyncApiEventConfigLogConfig) {
	if err := a.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetAuthProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetConnectionAuthModes() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionAuthModes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetDefaultPublishAuthModes() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultPublishAuthModes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetDefaultSubscribeAuthModes() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSubscribeAuthModes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

