// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/groundstationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroundstationConfigConfigDataOutputReference interface {
	cdktn.ComplexObject
	AntennaDownlinkConfig() GroundstationConfigConfigDataAntennaDownlinkConfigOutputReference
	AntennaDownlinkConfigInput() interface{}
	AntennaDownlinkDemodDecodeConfig() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference
	AntennaDownlinkDemodDecodeConfigInput() interface{}
	AntennaUplinkConfig() GroundstationConfigConfigDataAntennaUplinkConfigOutputReference
	AntennaUplinkConfigInput() interface{}
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
	DataflowEndpointConfig() GroundstationConfigConfigDataDataflowEndpointConfigOutputReference
	DataflowEndpointConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3RecordingConfig() GroundstationConfigConfigDataS3RecordingConfigOutputReference
	S3RecordingConfigInput() interface{}
	TelemetrySinkConfig() GroundstationConfigConfigDataTelemetrySinkConfigOutputReference
	TelemetrySinkConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrackingConfig() GroundstationConfigConfigDataTrackingConfigOutputReference
	TrackingConfigInput() interface{}
	UplinkEchoConfig() GroundstationConfigConfigDataUplinkEchoConfigOutputReference
	UplinkEchoConfigInput() interface{}
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
	PutAntennaDownlinkConfig(value *GroundstationConfigConfigDataAntennaDownlinkConfig)
	PutAntennaDownlinkDemodDecodeConfig(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfig)
	PutAntennaUplinkConfig(value *GroundstationConfigConfigDataAntennaUplinkConfig)
	PutDataflowEndpointConfig(value *GroundstationConfigConfigDataDataflowEndpointConfig)
	PutS3RecordingConfig(value *GroundstationConfigConfigDataS3RecordingConfig)
	PutTelemetrySinkConfig(value *GroundstationConfigConfigDataTelemetrySinkConfig)
	PutTrackingConfig(value *GroundstationConfigConfigDataTrackingConfig)
	PutUplinkEchoConfig(value *GroundstationConfigConfigDataUplinkEchoConfig)
	ResetAntennaDownlinkConfig()
	ResetAntennaDownlinkDemodDecodeConfig()
	ResetAntennaUplinkConfig()
	ResetDataflowEndpointConfig()
	ResetS3RecordingConfig()
	ResetTelemetrySinkConfig()
	ResetTrackingConfig()
	ResetUplinkEchoConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroundstationConfigConfigDataOutputReference
type jsiiProxy_GroundstationConfigConfigDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaDownlinkConfig() GroundstationConfigConfigDataAntennaDownlinkConfigOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkConfigOutputReference
	_jsii_.Get(
		j,
		"antennaDownlinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaDownlinkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antennaDownlinkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaDownlinkDemodDecodeConfig() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigOutputReference
	_jsii_.Get(
		j,
		"antennaDownlinkDemodDecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaDownlinkDemodDecodeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antennaDownlinkDemodDecodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaUplinkConfig() GroundstationConfigConfigDataAntennaUplinkConfigOutputReference {
	var returns GroundstationConfigConfigDataAntennaUplinkConfigOutputReference
	_jsii_.Get(
		j,
		"antennaUplinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) AntennaUplinkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antennaUplinkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) DataflowEndpointConfig() GroundstationConfigConfigDataDataflowEndpointConfigOutputReference {
	var returns GroundstationConfigConfigDataDataflowEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"dataflowEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) DataflowEndpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataflowEndpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) S3RecordingConfig() GroundstationConfigConfigDataS3RecordingConfigOutputReference {
	var returns GroundstationConfigConfigDataS3RecordingConfigOutputReference
	_jsii_.Get(
		j,
		"s3RecordingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) S3RecordingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3RecordingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TelemetrySinkConfig() GroundstationConfigConfigDataTelemetrySinkConfigOutputReference {
	var returns GroundstationConfigConfigDataTelemetrySinkConfigOutputReference
	_jsii_.Get(
		j,
		"telemetrySinkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TelemetrySinkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"telemetrySinkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TrackingConfig() GroundstationConfigConfigDataTrackingConfigOutputReference {
	var returns GroundstationConfigConfigDataTrackingConfigOutputReference
	_jsii_.Get(
		j,
		"trackingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) TrackingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) UplinkEchoConfig() GroundstationConfigConfigDataUplinkEchoConfigOutputReference {
	var returns GroundstationConfigConfigDataUplinkEchoConfigOutputReference
	_jsii_.Get(
		j,
		"uplinkEchoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference) UplinkEchoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uplinkEchoConfigInput",
		&returns,
	)
	return returns
}


func NewGroundstationConfigConfigDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GroundstationConfigConfigDataOutputReference {
	_init_.Initialize()

	if err := validateNewGroundstationConfigConfigDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroundstationConfigConfigDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationConfig.GroundstationConfigConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroundstationConfigConfigDataOutputReference_Override(g GroundstationConfigConfigDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationConfig.GroundstationConfigConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutAntennaDownlinkConfig(value *GroundstationConfigConfigDataAntennaDownlinkConfig) {
	if err := g.validatePutAntennaDownlinkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntennaDownlinkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutAntennaDownlinkDemodDecodeConfig(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfig) {
	if err := g.validatePutAntennaDownlinkDemodDecodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntennaDownlinkDemodDecodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutAntennaUplinkConfig(value *GroundstationConfigConfigDataAntennaUplinkConfig) {
	if err := g.validatePutAntennaUplinkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntennaUplinkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutDataflowEndpointConfig(value *GroundstationConfigConfigDataDataflowEndpointConfig) {
	if err := g.validatePutDataflowEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataflowEndpointConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutS3RecordingConfig(value *GroundstationConfigConfigDataS3RecordingConfig) {
	if err := g.validatePutS3RecordingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putS3RecordingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutTelemetrySinkConfig(value *GroundstationConfigConfigDataTelemetrySinkConfig) {
	if err := g.validatePutTelemetrySinkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTelemetrySinkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutTrackingConfig(value *GroundstationConfigConfigDataTrackingConfig) {
	if err := g.validatePutTrackingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrackingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) PutUplinkEchoConfig(value *GroundstationConfigConfigDataUplinkEchoConfig) {
	if err := g.validatePutUplinkEchoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUplinkEchoConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetAntennaDownlinkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAntennaDownlinkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetAntennaDownlinkDemodDecodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAntennaDownlinkDemodDecodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetAntennaUplinkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAntennaUplinkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetDataflowEndpointConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDataflowEndpointConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetS3RecordingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetS3RecordingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetTelemetrySinkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTelemetrySinkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetTrackingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTrackingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ResetUplinkEchoConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetUplinkEchoConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

