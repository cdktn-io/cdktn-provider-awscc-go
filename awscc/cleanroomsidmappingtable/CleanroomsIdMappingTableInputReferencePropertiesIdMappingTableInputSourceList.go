// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsidmappingtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsidmappingtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList
type jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList {
	_init_.Initialize()

	if err := validateNewCleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIdMappingTable.CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList_Override(c CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIdMappingTable.CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) Get(index *float64) CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIdMappingTableInputReferencePropertiesIdMappingTableInputSourceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

