// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelesstaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessTaskDefinitionUpdateLoRaWanOutputReference interface {
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
	CurrentVersion() IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersionOutputReference
	CurrentVersionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SigKeyCrc() *float64
	SetSigKeyCrc(val *float64)
	SigKeyCrcInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateSignature() *string
	SetUpdateSignature(val *string)
	UpdateSignatureInput() *string
	UpdateVersion() IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference
	UpdateVersionInput() interface{}
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
	PutCurrentVersion(value *IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersion)
	PutUpdateVersion(value *IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersion)
	ResetCurrentVersion()
	ResetSigKeyCrc()
	ResetUpdateSignature()
	ResetUpdateVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessTaskDefinitionUpdateLoRaWanOutputReference
type jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) CurrentVersion() IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersionOutputReference {
	var returns IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersionOutputReference
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) CurrentVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"currentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) SigKeyCrc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sigKeyCrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) SigKeyCrcInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sigKeyCrcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) UpdateSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) UpdateSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) UpdateVersion() IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference {
	var returns IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference
	_jsii_.Get(
		j,
		"updateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) UpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateVersionInput",
		&returns,
	)
	return returns
}


func NewIotwirelessTaskDefinitionUpdateLoRaWanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessTaskDefinitionUpdateLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessTaskDefinitionUpdateLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessTaskDefinition.IotwirelessTaskDefinitionUpdateLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessTaskDefinitionUpdateLoRaWanOutputReference_Override(i IotwirelessTaskDefinitionUpdateLoRaWanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessTaskDefinition.IotwirelessTaskDefinitionUpdateLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetSigKeyCrc(val *float64) {
	if err := j.validateSetSigKeyCrcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigKeyCrc",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference)SetUpdateSignature(val *string) {
	if err := j.validateSetUpdateSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSignature",
		val,
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) PutCurrentVersion(value *IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersion) {
	if err := i.validatePutCurrentVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCurrentVersion",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) PutUpdateVersion(value *IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersion) {
	if err := i.validatePutUpdateVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUpdateVersion",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ResetCurrentVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetCurrentVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ResetSigKeyCrc() {
	_jsii_.InvokeVoid(
		i,
		"resetSigKeyCrc",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ResetUpdateSignature() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdateSignature",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ResetUpdateVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdateVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

