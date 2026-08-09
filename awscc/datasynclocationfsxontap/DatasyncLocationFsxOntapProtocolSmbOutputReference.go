// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datasynclocationfsxontap/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncLocationFsxOntapProtocolSmbOutputReference interface {
	cdktn.ComplexObject
	CmkSecretConfig() DatasyncLocationFsxOntapProtocolSmbCmkSecretConfigOutputReference
	CmkSecretConfigInput() interface{}
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
	CustomSecretConfig() DatasyncLocationFsxOntapProtocolSmbCustomSecretConfigOutputReference
	CustomSecretConfigInput() interface{}
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedSecretConfig() DatasyncLocationFsxOntapProtocolSmbManagedSecretConfigOutputReference
	MountOptions() DatasyncLocationFsxOntapProtocolSmbMountOptionsOutputReference
	MountOptionsInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutCmkSecretConfig(value *DatasyncLocationFsxOntapProtocolSmbCmkSecretConfig)
	PutCustomSecretConfig(value *DatasyncLocationFsxOntapProtocolSmbCustomSecretConfig)
	PutMountOptions(value *DatasyncLocationFsxOntapProtocolSmbMountOptions)
	ResetCmkSecretConfig()
	ResetCustomSecretConfig()
	ResetDomain()
	ResetMountOptions()
	ResetPassword()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatasyncLocationFsxOntapProtocolSmbOutputReference
type jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) CmkSecretConfig() DatasyncLocationFsxOntapProtocolSmbCmkSecretConfigOutputReference {
	var returns DatasyncLocationFsxOntapProtocolSmbCmkSecretConfigOutputReference
	_jsii_.Get(
		j,
		"cmkSecretConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) CmkSecretConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmkSecretConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) CustomSecretConfig() DatasyncLocationFsxOntapProtocolSmbCustomSecretConfigOutputReference {
	var returns DatasyncLocationFsxOntapProtocolSmbCustomSecretConfigOutputReference
	_jsii_.Get(
		j,
		"customSecretConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) CustomSecretConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSecretConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ManagedSecretConfig() DatasyncLocationFsxOntapProtocolSmbManagedSecretConfigOutputReference {
	var returns DatasyncLocationFsxOntapProtocolSmbManagedSecretConfigOutputReference
	_jsii_.Get(
		j,
		"managedSecretConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) MountOptions() DatasyncLocationFsxOntapProtocolSmbMountOptionsOutputReference {
	var returns DatasyncLocationFsxOntapProtocolSmbMountOptionsOutputReference
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) MountOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatasyncLocationFsxOntapProtocolSmbOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatasyncLocationFsxOntapProtocolSmbOutputReference {
	_init_.Initialize()

	if err := validateNewDatasyncLocationFsxOntapProtocolSmbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datasyncLocationFsxOntap.DatasyncLocationFsxOntapProtocolSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatasyncLocationFsxOntapProtocolSmbOutputReference_Override(d DatasyncLocationFsxOntapProtocolSmbOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datasyncLocationFsxOntap.DatasyncLocationFsxOntapProtocolSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) PutCmkSecretConfig(value *DatasyncLocationFsxOntapProtocolSmbCmkSecretConfig) {
	if err := d.validatePutCmkSecretConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCmkSecretConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) PutCustomSecretConfig(value *DatasyncLocationFsxOntapProtocolSmbCustomSecretConfig) {
	if err := d.validatePutCustomSecretConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomSecretConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) PutMountOptions(value *DatasyncLocationFsxOntapProtocolSmbMountOptions) {
	if err := d.validatePutMountOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMountOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetCmkSecretConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetCmkSecretConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetCustomSecretConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSecretConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetMountOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatasyncLocationFsxOntapProtocolSmbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

