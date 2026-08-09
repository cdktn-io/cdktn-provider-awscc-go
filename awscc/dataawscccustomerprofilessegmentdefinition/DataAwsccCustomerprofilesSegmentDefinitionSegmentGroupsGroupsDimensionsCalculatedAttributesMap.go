// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccustomerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscccustomerprofilessegmentdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap interface {
	cdktn.ComplexMap
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
	ComputeFqn() *string
	Get(key *string) DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap
type jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap {
	_init_.Initialize()

	if err := validateNewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap_Override(d DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) Get(key *string) DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

