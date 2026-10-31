// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/iototaupdate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference interface {
	cdktn.ComplexObject
	CertificateChain() IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChainOutputReference
	CertificateChainInput() interface{}
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
	HashAlgorithm() *string
	SetHashAlgorithm(val *string)
	HashAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Signature() IotOtaUpdateFilesCodeSigningCustomCodeSigningSignatureOutputReference
	SignatureAlgorithm() *string
	SetSignatureAlgorithm(val *string)
	SignatureAlgorithmInput() *string
	SignatureInput() interface{}
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
	PutCertificateChain(value *IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChain)
	PutSignature(value *IotOtaUpdateFilesCodeSigningCustomCodeSigningSignature)
	ResetCertificateChain()
	ResetHashAlgorithm()
	ResetSignature()
	ResetSignatureAlgorithm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference
type jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) CertificateChain() IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChainOutputReference {
	var returns IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChainOutputReference
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) CertificateChainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) HashAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) Signature() IotOtaUpdateFilesCodeSigningCustomCodeSigningSignatureOutputReference {
	var returns IotOtaUpdateFilesCodeSigningCustomCodeSigningSignatureOutputReference
	_jsii_.Get(
		j,
		"signature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) SignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) SignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference {
	_init_.Initialize()

	if err := validateNewIotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference_Override(i IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetHashAlgorithm(val *string) {
	if err := j.validateSetHashAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetSignatureAlgorithm(val *string) {
	if err := j.validateSetSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) PutCertificateChain(value *IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChain) {
	if err := i.validatePutCertificateChainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCertificateChain",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) PutSignature(value *IotOtaUpdateFilesCodeSigningCustomCodeSigningSignature) {
	if err := i.validatePutSignatureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSignature",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		i,
		"resetCertificateChain",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ResetHashAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetHashAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ResetSignature() {
	_jsii_.InvokeVoid(
		i,
		"resetSignature",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ResetSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetSignatureAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningCustomCodeSigningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

