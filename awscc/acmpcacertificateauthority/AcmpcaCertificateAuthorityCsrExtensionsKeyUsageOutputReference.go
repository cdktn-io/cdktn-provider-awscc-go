// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/acmpcacertificateauthority/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference interface {
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
	CrlSign() interface{}
	SetCrlSign(val interface{})
	CrlSignInput() interface{}
	DataEncipherment() interface{}
	SetDataEncipherment(val interface{})
	DataEnciphermentInput() interface{}
	DecipherOnly() interface{}
	SetDecipherOnly(val interface{})
	DecipherOnlyInput() interface{}
	DigitalSignature() interface{}
	SetDigitalSignature(val interface{})
	DigitalSignatureInput() interface{}
	EncipherOnly() interface{}
	SetEncipherOnly(val interface{})
	EncipherOnlyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyAgreement() interface{}
	SetKeyAgreement(val interface{})
	KeyAgreementInput() interface{}
	KeyCertSign() interface{}
	SetKeyCertSign(val interface{})
	KeyCertSignInput() interface{}
	KeyEncipherment() interface{}
	SetKeyEncipherment(val interface{})
	KeyEnciphermentInput() interface{}
	NonRepudiation() interface{}
	SetNonRepudiation(val interface{})
	NonRepudiationInput() interface{}
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
	ResetCrlSign()
	ResetDataEncipherment()
	ResetDecipherOnly()
	ResetDigitalSignature()
	ResetEncipherOnly()
	ResetKeyAgreement()
	ResetKeyCertSign()
	ResetKeyEncipherment()
	ResetNonRepudiation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) CrlSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) CrlSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DataEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DataEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DecipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DecipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DigitalSignature() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DigitalSignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) EncipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) EncipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyAgreement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyAgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyCertSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyCertSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyCertSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyCertSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) NonRepudiation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonRepudiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) NonRepudiationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonRepudiationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference_Override(a AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetCrlSign(val interface{}) {
	if err := j.validateSetCrlSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlSign",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetDataEncipherment(val interface{}) {
	if err := j.validateSetDataEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataEncipherment",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetDecipherOnly(val interface{}) {
	if err := j.validateSetDecipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decipherOnly",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetDigitalSignature(val interface{}) {
	if err := j.validateSetDigitalSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalSignature",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetEncipherOnly(val interface{}) {
	if err := j.validateSetEncipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encipherOnly",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetKeyAgreement(val interface{}) {
	if err := j.validateSetKeyAgreementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAgreement",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetKeyCertSign(val interface{}) {
	if err := j.validateSetKeyCertSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCertSign",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetKeyEncipherment(val interface{}) {
	if err := j.validateSetKeyEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyEncipherment",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetNonRepudiation(val interface{}) {
	if err := j.validateSetNonRepudiationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonRepudiation",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetCrlSign() {
	_jsii_.InvokeVoid(
		a,
		"resetCrlSign",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetDataEncipherment() {
	_jsii_.InvokeVoid(
		a,
		"resetDataEncipherment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetDecipherOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetDecipherOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetDigitalSignature() {
	_jsii_.InvokeVoid(
		a,
		"resetDigitalSignature",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetEncipherOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetEncipherOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetKeyAgreement() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyAgreement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetKeyCertSign() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyCertSign",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetKeyEncipherment() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyEncipherment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ResetNonRepudiation() {
	_jsii_.InvokeVoid(
		a,
		"resetNonRepudiation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

