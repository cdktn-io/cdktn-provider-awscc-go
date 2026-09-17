// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccglueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccglueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference interface {
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
	DefaultValue() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPage
	SetInternalValue(val *DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPage)
	Key() *string
	PropertyLocation() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference
type jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InternalValue() *DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPage {
	var returns *DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) PropertyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Value() DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference {
	var returns DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewDataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnectionType.DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference_Override(d DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnectionType.DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetInternalValue(val *DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

