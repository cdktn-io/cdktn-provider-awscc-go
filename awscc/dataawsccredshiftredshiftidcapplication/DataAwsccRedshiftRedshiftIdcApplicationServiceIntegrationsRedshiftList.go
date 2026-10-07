// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccredshiftredshiftidcapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccredshiftredshiftidcapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList interface {
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
	Get(index *float64) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList
type jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList {
	_init_.Initialize()

	if err := validateNewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftRedshiftIdcApplication.DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList_Override(d DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftRedshiftIdcApplication.DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) Get(index *float64) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

