// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationapikey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/locationapikey/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LocationApiKeyRestrictionsOutputReference interface {
	cdktn.ComplexObject
	AllowActions() *[]*string
	SetAllowActions(val *[]*string)
	AllowActionsInput() *[]*string
	AllowAndroidApps() LocationApiKeyRestrictionsAllowAndroidAppsList
	AllowAndroidAppsInput() interface{}
	AllowAppleApps() LocationApiKeyRestrictionsAllowAppleAppsList
	AllowAppleAppsInput() interface{}
	AllowReferers() *[]*string
	SetAllowReferers(val *[]*string)
	AllowReferersInput() *[]*string
	AllowResources() *[]*string
	SetAllowResources(val *[]*string)
	AllowResourcesInput() *[]*string
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
	PutAllowAndroidApps(value interface{})
	PutAllowAppleApps(value interface{})
	ResetAllowAndroidApps()
	ResetAllowAppleApps()
	ResetAllowReferers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LocationApiKeyRestrictionsOutputReference
type jsiiProxy_LocationApiKeyRestrictionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowAndroidApps() LocationApiKeyRestrictionsAllowAndroidAppsList {
	var returns LocationApiKeyRestrictionsAllowAndroidAppsList
	_jsii_.Get(
		j,
		"allowAndroidApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowAndroidAppsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAndroidAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowAppleApps() LocationApiKeyRestrictionsAllowAppleAppsList {
	var returns LocationApiKeyRestrictionsAllowAppleAppsList
	_jsii_.Get(
		j,
		"allowAppleApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowAppleAppsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAppleAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowReferers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowReferers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowReferersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowReferersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) AllowResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLocationApiKeyRestrictionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LocationApiKeyRestrictionsOutputReference {
	_init_.Initialize()

	if err := validateNewLocationApiKeyRestrictionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LocationApiKeyRestrictionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.locationApiKey.LocationApiKeyRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLocationApiKeyRestrictionsOutputReference_Override(l LocationApiKeyRestrictionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.locationApiKey.LocationApiKeyRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetAllowActions(val *[]*string) {
	if err := j.validateSetAllowActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowActions",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetAllowReferers(val *[]*string) {
	if err := j.validateSetAllowReferersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReferers",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetAllowResources(val *[]*string) {
	if err := j.validateSetAllowResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowResources",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LocationApiKeyRestrictionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) PutAllowAndroidApps(value interface{}) {
	if err := l.validatePutAllowAndroidAppsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAllowAndroidApps",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) PutAllowAppleApps(value interface{}) {
	if err := l.validatePutAllowAppleAppsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAllowAppleApps",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ResetAllowAndroidApps() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowAndroidApps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ResetAllowAppleApps() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowAppleApps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ResetAllowReferers() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowReferers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocationApiKeyRestrictionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

