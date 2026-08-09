// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotdomainconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotDomainConfigurationServerCertificateConfigOutputReference interface {
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
	EnableOcspCheck() interface{}
	SetEnableOcspCheck(val interface{})
	EnableOcspCheckInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OcspAuthorizedResponderArn() *string
	SetOcspAuthorizedResponderArn(val *string)
	OcspAuthorizedResponderArnInput() *string
	OcspLambdaArn() *string
	SetOcspLambdaArn(val *string)
	OcspLambdaArnInput() *string
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
	ResetEnableOcspCheck()
	ResetOcspAuthorizedResponderArn()
	ResetOcspLambdaArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotDomainConfigurationServerCertificateConfigOutputReference
type jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) EnableOcspCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOcspCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) EnableOcspCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOcspCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) OcspAuthorizedResponderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspAuthorizedResponderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) OcspAuthorizedResponderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspAuthorizedResponderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) OcspLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) OcspLambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspLambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotDomainConfigurationServerCertificateConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotDomainConfigurationServerCertificateConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIotDomainConfigurationServerCertificateConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfigurationServerCertificateConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotDomainConfigurationServerCertificateConfigOutputReference_Override(i IotDomainConfigurationServerCertificateConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfigurationServerCertificateConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetEnableOcspCheck(val interface{}) {
	if err := j.validateSetEnableOcspCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOcspCheck",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetOcspAuthorizedResponderArn(val *string) {
	if err := j.validateSetOcspAuthorizedResponderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspAuthorizedResponderArn",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetOcspLambdaArn(val *string) {
	if err := j.validateSetOcspLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspLambdaArn",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ResetEnableOcspCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableOcspCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ResetOcspAuthorizedResponderArn() {
	_jsii_.InvokeVoid(
		i,
		"resetOcspAuthorizedResponderArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ResetOcspLambdaArn() {
	_jsii_.InvokeVoid(
		i,
		"resetOcspLambdaArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotDomainConfigurationServerCertificateConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

