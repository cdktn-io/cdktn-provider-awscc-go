// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccustomerprofilescalculatedattributedefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawscccustomerprofilescalculatedattributedefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap interface {
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
	Get(key *string) DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesOutputReference
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

// The jsii proxy struct for DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap
type jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap {
	_init_.Initialize()

	if err := validateNewDataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesCalculatedAttributeDefinition.DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap_Override(d DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesCalculatedAttributeDefinition.DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) Get(key *string) DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

