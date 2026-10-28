// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/licensemanagerusersubscriptionslicenseserverendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference interface {
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
	ServerSettings() LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsOutputReference
	ServerSettingsInput() interface{}
	ServerType() *string
	SetServerType(val *string)
	ServerTypeInput() *string
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
	PutServerSettings(value *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettings)
	ResetServerSettings()
	ResetServerType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference
type jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ServerSettings() LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsOutputReference {
	var returns LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsOutputReference
	_jsii_.Get(
		j,
		"serverSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ServerSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerusersubscriptionsLicenseServerEndpoint.LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference_Override(l LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerusersubscriptionsLicenseServerEndpoint.LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetServerType(val *string) {
	if err := j.validateSetServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverType",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) PutServerSettings(value *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettings) {
	if err := l.validatePutServerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putServerSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ResetServerSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetServerSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ResetServerType() {
	_jsii_.InvokeVoid(
		l,
		"resetServerType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

