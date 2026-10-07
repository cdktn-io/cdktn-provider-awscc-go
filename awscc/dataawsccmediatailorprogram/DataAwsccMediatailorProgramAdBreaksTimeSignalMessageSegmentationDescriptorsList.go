// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccmediatailorprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccmediatailorprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList interface {
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
	Get(index *float64) DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList
type jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList {
	_init_.Initialize()

	if err := validateNewDataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediatailorProgram.DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList_Override(d DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediatailorProgram.DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) Get(index *float64) DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

