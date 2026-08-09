// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanageringresspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesmailmanageringresspoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference interface {
	cdktn.ComplexObject
	CaContent() *string
	SetCaContent(val *string)
	CaContentInput() *string
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
	CrlContent() *string
	SetCrlContent(val *string)
	CrlContentInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	ResetCaContent()
	ResetCrlContent()
	ResetKmsKeyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference
type jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) CaContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) CaContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) CrlContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crlContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) CrlContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crlContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerIngressPoint.SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference_Override(s SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerIngressPoint.SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetCaContent(val *string) {
	if err := j.validateSetCaContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caContent",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetCrlContent(val *string) {
	if err := j.validateSetCrlContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlContent",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ResetCaContent() {
	_jsii_.InvokeVoid(
		s,
		"resetCaContent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ResetCrlContent() {
	_jsii_.InvokeVoid(
		s,
		"resetCrlContent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerIngressPointIngressPointConfigurationTlsAuthConfigurationTrustStoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

