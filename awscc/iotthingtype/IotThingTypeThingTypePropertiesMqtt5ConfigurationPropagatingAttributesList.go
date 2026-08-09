// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotthingtype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotthingtype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList interface {
	cdktn.ComplexList
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
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList
type jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList {
	_init_.Initialize()

	if err := validateNewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotThingType.IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewIotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList_Override(i IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotThingType.IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := i.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		i,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) Get(index *float64) IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference {
	if err := i.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotThingTypeThingTypePropertiesMqtt5ConfigurationPropagatingAttributesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

