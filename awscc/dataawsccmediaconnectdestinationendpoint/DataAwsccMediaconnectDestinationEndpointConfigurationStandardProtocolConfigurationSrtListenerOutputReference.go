// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccmediaconnectdestinationendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccmediaconnectdestinationendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference interface {
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
	DestinationAddress() *string
	DestinationIngestPort() *float64
	EncryptionPassphrase() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener
	SetInternalValue(val *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener)
	MinimumLatencyMilliseconds() *float64
	StreamId() *string
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

// The jsii proxy struct for DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
type jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationIngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) EncryptionPassphrase() DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference {
	var returns DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	_jsii_.Get(
		j,
		"encryptionPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InternalValue() *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener {
	var returns *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) MinimumLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediaconnectDestinationEndpoint.DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference_Override(d DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediaconnectDestinationEndpoint.DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetInternalValue(val *DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

