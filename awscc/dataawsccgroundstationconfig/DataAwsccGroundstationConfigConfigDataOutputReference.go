// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccgroundstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccgroundstationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccGroundstationConfigConfigDataOutputReference interface {
	cdktn.ComplexObject
	AntennaDownlinkConfig() DataAwsccGroundstationConfigConfigDataAntennaDownlinkConfigOutputReference
	AntennaDownlinkDemodDecodeConfig() DataAwsccGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference
	AntennaUplinkConfig() DataAwsccGroundstationConfigConfigDataAntennaUplinkConfigOutputReference
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
	DataflowEndpointConfig() DataAwsccGroundstationConfigConfigDataDataflowEndpointConfigOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccGroundstationConfigConfigData
	SetInternalValue(val *DataAwsccGroundstationConfigConfigData)
	S3RecordingConfig() DataAwsccGroundstationConfigConfigDataS3RecordingConfigOutputReference
	TelemetrySinkConfig() DataAwsccGroundstationConfigConfigDataTelemetrySinkConfigOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrackingConfig() DataAwsccGroundstationConfigConfigDataTrackingConfigOutputReference
	UplinkEchoConfig() DataAwsccGroundstationConfigConfigDataUplinkEchoConfigOutputReference
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

// The jsii proxy struct for DataAwsccGroundstationConfigConfigDataOutputReference
type jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) AntennaDownlinkConfig() DataAwsccGroundstationConfigConfigDataAntennaDownlinkConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataAntennaDownlinkConfigOutputReference
	_jsii_.Get(
		j,
		"antennaDownlinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) AntennaDownlinkDemodDecodeConfig() DataAwsccGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference
	_jsii_.Get(
		j,
		"antennaDownlinkDemodDecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) AntennaUplinkConfig() DataAwsccGroundstationConfigConfigDataAntennaUplinkConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataAntennaUplinkConfigOutputReference
	_jsii_.Get(
		j,
		"antennaUplinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) DataflowEndpointConfig() DataAwsccGroundstationConfigConfigDataDataflowEndpointConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataDataflowEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"dataflowEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) InternalValue() *DataAwsccGroundstationConfigConfigData {
	var returns *DataAwsccGroundstationConfigConfigData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) S3RecordingConfig() DataAwsccGroundstationConfigConfigDataS3RecordingConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataS3RecordingConfigOutputReference
	_jsii_.Get(
		j,
		"s3RecordingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) TelemetrySinkConfig() DataAwsccGroundstationConfigConfigDataTelemetrySinkConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataTelemetrySinkConfigOutputReference
	_jsii_.Get(
		j,
		"telemetrySinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) TrackingConfig() DataAwsccGroundstationConfigConfigDataTrackingConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataTrackingConfigOutputReference
	_jsii_.Get(
		j,
		"trackingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) UplinkEchoConfig() DataAwsccGroundstationConfigConfigDataUplinkEchoConfigOutputReference {
	var returns DataAwsccGroundstationConfigConfigDataUplinkEchoConfigOutputReference
	_jsii_.Get(
		j,
		"uplinkEchoConfig",
		&returns,
	)
	return returns
}


func NewDataAwsccGroundstationConfigConfigDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccGroundstationConfigConfigDataOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGroundstationConfigConfigDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGroundstationConfig.DataAwsccGroundstationConfigConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccGroundstationConfigConfigDataOutputReference_Override(d DataAwsccGroundstationConfigConfigDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGroundstationConfig.DataAwsccGroundstationConfigConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference)SetInternalValue(val *DataAwsccGroundstationConfigConfigData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccGroundstationConfigConfigDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

