// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/acmpcacertificateauthority/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference interface {
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
	DirectoryName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryNameOutputReference
	DirectoryNameInput() interface{}
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	EdiPartyName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyNameOutputReference
	EdiPartyNameInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	OtherName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherNameOutputReference
	OtherNameInput() interface{}
	RegisteredId() *string
	SetRegisteredId(val *string)
	RegisteredIdInput() *string
	Rfc822Name() *string
	SetRfc822Name(val *string)
	Rfc822NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniformResourceIdentifier() *string
	SetUniformResourceIdentifier(val *string)
	UniformResourceIdentifierInput() *string
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
	PutDirectoryName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryName)
	PutEdiPartyName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyName)
	PutOtherName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName)
	ResetDirectoryName()
	ResetDnsName()
	ResetEdiPartyName()
	ResetIpAddress()
	ResetOtherName()
	ResetRegisteredId()
	ResetRfc822Name()
	ResetUniformResourceIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) DirectoryName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryNameOutputReference {
	var returns AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryNameOutputReference
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) DirectoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) EdiPartyName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyNameOutputReference {
	var returns AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyNameOutputReference
	_jsii_.Get(
		j,
		"ediPartyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) EdiPartyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ediPartyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) OtherName() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherNameOutputReference {
	var returns AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherNameOutputReference
	_jsii_.Get(
		j,
		"otherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) OtherNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"otherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) RegisteredId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) RegisteredIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) Rfc822Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc822Name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) Rfc822NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc822NameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) UniformResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) UniformResourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformResourceIdentifierInput",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference_Override(a AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetRegisteredId(val *string) {
	if err := j.validateSetRegisteredIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registeredId",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetRfc822Name(val *string) {
	if err := j.validateSetRfc822NameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc822Name",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference)SetUniformResourceIdentifier(val *string) {
	if err := j.validateSetUniformResourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniformResourceIdentifier",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) PutDirectoryName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryName) {
	if err := a.validatePutDirectoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDirectoryName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) PutEdiPartyName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyName) {
	if err := a.validatePutEdiPartyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEdiPartyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) PutOtherName(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName) {
	if err := a.validatePutOtherNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOtherName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetDnsName() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetEdiPartyName() {
	_jsii_.InvokeVoid(
		a,
		"resetEdiPartyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetOtherName() {
	_jsii_.InvokeVoid(
		a,
		"resetOtherName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetRegisteredId() {
	_jsii_.InvokeVoid(
		a,
		"resetRegisteredId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetRfc822Name() {
	_jsii_.InvokeVoid(
		a,
		"resetRfc822Name",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ResetUniformResourceIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetUniformResourceIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

