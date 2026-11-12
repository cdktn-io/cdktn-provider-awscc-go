// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccmediaconnectdestinationendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccmediaconnectdestinationendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration
	SetInternalValue(val *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration)
	Rist() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRistOutputReference
	Rtp() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference
	SrtCaller() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	SrtListener() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
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

// The jsii proxy struct for DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference
type jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) InternalValue() *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration {
	var returns *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) Rist() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRistOutputReference {
	var returns DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRistOutputReference
	_jsii_.Get(
		j,
		"rist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) Rtp() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference {
	var returns DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference
	_jsii_.Get(
		j,
		"rtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtCaller() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference {
	var returns DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	_jsii_.Get(
		j,
		"srtCaller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtListener() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	var returns DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
	_jsii_.Get(
		j,
		"srtListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediaconnectDestinationEndpoint.DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference_Override(d DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediaconnectDestinationEndpoint.DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference)SetInternalValue(val *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

