// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/groundstationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference interface {
	cdktn.ComplexObject
	Bandwidth() GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidthOutputReference
	BandwidthInput() interface{}
	CenterFrequency() GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequencyOutputReference
	CenterFrequencyInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Polarization() *string
	SetPolarization(val *string)
	PolarizationInput() *string
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
	PutBandwidth(value *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidth)
	PutCenterFrequency(value *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequency)
	ResetBandwidth()
	ResetCenterFrequency()
	ResetPolarization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference
type jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) Bandwidth() GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidthOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidthOutputReference
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) BandwidthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) CenterFrequency() GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequencyOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequencyOutputReference
	_jsii_.Get(
		j,
		"centerFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) CenterFrequencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"centerFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) Polarization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) PolarizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polarizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationConfig.GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference_Override(g GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationConfig.GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetPolarization(val *string) {
	if err := j.validateSetPolarizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"polarization",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) PutBandwidth(value *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidth) {
	if err := g.validatePutBandwidthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBandwidth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) PutCenterFrequency(value *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequency) {
	if err := g.validatePutCenterFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCenterFrequency",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ResetBandwidth() {
	_jsii_.InvokeVoid(
		g,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ResetCenterFrequency() {
	_jsii_.InvokeVoid(
		g,
		"resetCenterFrequency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ResetPolarization() {
	_jsii_.InvokeVoid(
		g,
		"resetPolarization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

