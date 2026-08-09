// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotthingtype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotthingtype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference interface {
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
	ConnectionAttribute() *string
	SetConnectionAttribute(val *string)
	ConnectionAttributeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThingAttribute() *string
	SetThingAttribute(val *string)
	ThingAttributeInput() *string
	UserPropertyKey() *string
	SetUserPropertyKey(val *string)
	UserPropertyKeyInput() *string
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
	ResetConnectionAttribute()
	ResetThingAttribute()
	ResetUserPropertyKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference
type jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ConnectionAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ConnectionAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ThingAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ThingAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) UserPropertyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPropertyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) UserPropertyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPropertyKeyInput",
		&returns,
	)
	return returns
}


func NewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotThingType.IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference_Override(i IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotThingType.IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetConnectionAttribute(val *string) {
	if err := j.validateSetConnectionAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttribute",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetThingAttribute(val *string) {
	if err := j.validateSetThingAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingAttribute",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference)SetUserPropertyKey(val *string) {
	if err := j.validateSetUserPropertyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPropertyKey",
		val,
	)
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ResetConnectionAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetConnectionAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ResetThingAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetThingAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ResetUserPropertyKey() {
	_jsii_.InvokeVoid(
		i,
		"resetUserPropertyKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

