// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccredshiftredshiftidcapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccredshiftredshiftidcapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList interface {
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
	Get(index *float64) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList
type jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList {
	_init_.Initialize()

	if err := validateNewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftRedshiftIdcApplication.DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList_Override(d DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftRedshiftIdcApplication.DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) Get(index *float64) DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccRedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

