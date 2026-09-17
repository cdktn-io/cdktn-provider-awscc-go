// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesischannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList interface {
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
	Get(index *float64) KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList
type jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList {
	_init_.Initialize()

	if err := validateNewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList_Override(k KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := k.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		k,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) Get(index *float64) KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

