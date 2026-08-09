// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelesstaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference interface {
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
	Model() *string
	SetModel(val *string)
	ModelInput() *string
	PackageVersion() *string
	SetPackageVersion(val *string)
	PackageVersionInput() *string
	Station() *string
	SetStation(val *string)
	StationInput() *string
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
	ResetModel()
	ResetPackageVersion()
	ResetStation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference
type jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) Model() *string {
	var returns *string
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) PackageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) PackageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) Station() *string {
	var returns *string
	_jsii_.Get(
		j,
		"station",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) StationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessTaskDefinition.IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference_Override(i IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessTaskDefinition.IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetModel(val *string) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetPackageVersion(val *string) {
	if err := j.validateSetPackageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageVersion",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetStation(val *string) {
	if err := j.validateSetStationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"station",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ResetModel() {
	_jsii_.InvokeVoid(
		i,
		"resetModel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ResetPackageVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetPackageVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ResetStation() {
	_jsii_.InvokeVoid(
		i,
		"resetStation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

