// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotmitigationaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotmitigationaction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotMitigationActionActionParamsOutputReference interface {
	cdktn.ComplexObject
	AddThingsToThingGroupParams() IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference
	AddThingsToThingGroupParamsInput() interface{}
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
	EnableIoTLoggingParams() IotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference
	EnableIoTLoggingParamsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PublishFindingToSnsParams() IotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference
	PublishFindingToSnsParamsInput() interface{}
	ReplaceDefaultPolicyVersionParams() IotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference
	ReplaceDefaultPolicyVersionParamsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateCaCertificateParams() IotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference
	UpdateCaCertificateParamsInput() interface{}
	UpdateDeviceCertificateParams() IotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference
	UpdateDeviceCertificateParamsInput() interface{}
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
	PutAddThingsToThingGroupParams(value *IotMitigationActionActionParamsAddThingsToThingGroupParams)
	PutEnableIoTLoggingParams(value *IotMitigationActionActionParamsEnableIoTLoggingParams)
	PutPublishFindingToSnsParams(value *IotMitigationActionActionParamsPublishFindingToSnsParams)
	PutReplaceDefaultPolicyVersionParams(value *IotMitigationActionActionParamsReplaceDefaultPolicyVersionParams)
	PutUpdateCaCertificateParams(value *IotMitigationActionActionParamsUpdateCaCertificateParams)
	PutUpdateDeviceCertificateParams(value *IotMitigationActionActionParamsUpdateDeviceCertificateParams)
	ResetAddThingsToThingGroupParams()
	ResetEnableIoTLoggingParams()
	ResetPublishFindingToSnsParams()
	ResetReplaceDefaultPolicyVersionParams()
	ResetUpdateCaCertificateParams()
	ResetUpdateDeviceCertificateParams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotMitigationActionActionParamsOutputReference
type jsiiProxy_IotMitigationActionActionParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) AddThingsToThingGroupParams() IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference {
	var returns IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference
	_jsii_.Get(
		j,
		"addThingsToThingGroupParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) AddThingsToThingGroupParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addThingsToThingGroupParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) EnableIoTLoggingParams() IotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference {
	var returns IotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference
	_jsii_.Get(
		j,
		"enableIoTLoggingParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) EnableIoTLoggingParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIoTLoggingParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) PublishFindingToSnsParams() IotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference {
	var returns IotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference
	_jsii_.Get(
		j,
		"publishFindingToSnsParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) PublishFindingToSnsParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishFindingToSnsParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) ReplaceDefaultPolicyVersionParams() IotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference {
	var returns IotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference
	_jsii_.Get(
		j,
		"replaceDefaultPolicyVersionParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) ReplaceDefaultPolicyVersionParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceDefaultPolicyVersionParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) UpdateCaCertificateParams() IotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference {
	var returns IotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference
	_jsii_.Get(
		j,
		"updateCaCertificateParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) UpdateCaCertificateParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateCaCertificateParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) UpdateDeviceCertificateParams() IotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference {
	var returns IotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference
	_jsii_.Get(
		j,
		"updateDeviceCertificateParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference) UpdateDeviceCertificateParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDeviceCertificateParamsInput",
		&returns,
	)
	return returns
}


func NewIotMitigationActionActionParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotMitigationActionActionParamsOutputReference {
	_init_.Initialize()

	if err := validateNewIotMitigationActionActionParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotMitigationActionActionParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotMitigationAction.IotMitigationActionActionParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotMitigationActionActionParamsOutputReference_Override(i IotMitigationActionActionParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotMitigationAction.IotMitigationActionActionParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutAddThingsToThingGroupParams(value *IotMitigationActionActionParamsAddThingsToThingGroupParams) {
	if err := i.validatePutAddThingsToThingGroupParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAddThingsToThingGroupParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutEnableIoTLoggingParams(value *IotMitigationActionActionParamsEnableIoTLoggingParams) {
	if err := i.validatePutEnableIoTLoggingParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEnableIoTLoggingParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutPublishFindingToSnsParams(value *IotMitigationActionActionParamsPublishFindingToSnsParams) {
	if err := i.validatePutPublishFindingToSnsParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPublishFindingToSnsParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutReplaceDefaultPolicyVersionParams(value *IotMitigationActionActionParamsReplaceDefaultPolicyVersionParams) {
	if err := i.validatePutReplaceDefaultPolicyVersionParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putReplaceDefaultPolicyVersionParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutUpdateCaCertificateParams(value *IotMitigationActionActionParamsUpdateCaCertificateParams) {
	if err := i.validatePutUpdateCaCertificateParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUpdateCaCertificateParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) PutUpdateDeviceCertificateParams(value *IotMitigationActionActionParamsUpdateDeviceCertificateParams) {
	if err := i.validatePutUpdateDeviceCertificateParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUpdateDeviceCertificateParams",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetAddThingsToThingGroupParams() {
	_jsii_.InvokeVoid(
		i,
		"resetAddThingsToThingGroupParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetEnableIoTLoggingParams() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableIoTLoggingParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetPublishFindingToSnsParams() {
	_jsii_.InvokeVoid(
		i,
		"resetPublishFindingToSnsParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetReplaceDefaultPolicyVersionParams() {
	_jsii_.InvokeVoid(
		i,
		"resetReplaceDefaultPolicyVersionParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetUpdateCaCertificateParams() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdateCaCertificateParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ResetUpdateDeviceCertificateParams() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdateDeviceCertificateParams",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotMitigationActionActionParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

