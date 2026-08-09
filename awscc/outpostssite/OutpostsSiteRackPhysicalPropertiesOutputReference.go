// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostssite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/outpostssite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OutpostsSiteRackPhysicalPropertiesOutputReference interface {
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
	FiberOpticCableType() *string
	SetFiberOpticCableType(val *string)
	FiberOpticCableTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumSupportedWeightLbs() *string
	SetMaximumSupportedWeightLbs(val *string)
	MaximumSupportedWeightLbsInput() *string
	OpticalStandard() *string
	SetOpticalStandard(val *string)
	OpticalStandardInput() *string
	PowerConnector() *string
	SetPowerConnector(val *string)
	PowerConnectorInput() *string
	PowerDrawKva() *string
	SetPowerDrawKva(val *string)
	PowerDrawKvaInput() *string
	PowerFeedDrop() *string
	SetPowerFeedDrop(val *string)
	PowerFeedDropInput() *string
	PowerPhase() *string
	SetPowerPhase(val *string)
	PowerPhaseInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UplinkCount() *string
	SetUplinkCount(val *string)
	UplinkCountInput() *string
	UplinkGbps() *string
	SetUplinkGbps(val *string)
	UplinkGbpsInput() *string
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
	ResetFiberOpticCableType()
	ResetMaximumSupportedWeightLbs()
	ResetOpticalStandard()
	ResetPowerConnector()
	ResetPowerDrawKva()
	ResetPowerFeedDrop()
	ResetPowerPhase()
	ResetUplinkCount()
	ResetUplinkGbps()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OutpostsSiteRackPhysicalPropertiesOutputReference
type jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) FiberOpticCableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fiberOpticCableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) FiberOpticCableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fiberOpticCableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) MaximumSupportedWeightLbs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumSupportedWeightLbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) MaximumSupportedWeightLbsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumSupportedWeightLbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) OpticalStandard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opticalStandard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) OpticalStandardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opticalStandardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerDrawKva() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerDrawKva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerDrawKvaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerDrawKvaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerFeedDrop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerFeedDrop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerFeedDropInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerFeedDropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerPhase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) PowerPhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) UplinkCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uplinkCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) UplinkCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uplinkCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) UplinkGbps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uplinkGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) UplinkGbpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uplinkGbpsInput",
		&returns,
	)
	return returns
}


func NewOutpostsSiteRackPhysicalPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OutpostsSiteRackPhysicalPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOutpostsSiteRackPhysicalPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.outpostsSite.OutpostsSiteRackPhysicalPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOutpostsSiteRackPhysicalPropertiesOutputReference_Override(o OutpostsSiteRackPhysicalPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.outpostsSite.OutpostsSiteRackPhysicalPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetFiberOpticCableType(val *string) {
	if err := j.validateSetFiberOpticCableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fiberOpticCableType",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetMaximumSupportedWeightLbs(val *string) {
	if err := j.validateSetMaximumSupportedWeightLbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumSupportedWeightLbs",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetOpticalStandard(val *string) {
	if err := j.validateSetOpticalStandardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opticalStandard",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetPowerConnector(val *string) {
	if err := j.validateSetPowerConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerConnector",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetPowerDrawKva(val *string) {
	if err := j.validateSetPowerDrawKvaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerDrawKva",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetPowerFeedDrop(val *string) {
	if err := j.validateSetPowerFeedDropParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerFeedDrop",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetPowerPhase(val *string) {
	if err := j.validateSetPowerPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerPhase",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetUplinkCount(val *string) {
	if err := j.validateSetUplinkCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkCount",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference)SetUplinkGbps(val *string) {
	if err := j.validateSetUplinkGbpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkGbps",
		val,
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetFiberOpticCableType() {
	_jsii_.InvokeVoid(
		o,
		"resetFiberOpticCableType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetMaximumSupportedWeightLbs() {
	_jsii_.InvokeVoid(
		o,
		"resetMaximumSupportedWeightLbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetOpticalStandard() {
	_jsii_.InvokeVoid(
		o,
		"resetOpticalStandard",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetPowerConnector() {
	_jsii_.InvokeVoid(
		o,
		"resetPowerConnector",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetPowerDrawKva() {
	_jsii_.InvokeVoid(
		o,
		"resetPowerDrawKva",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetPowerFeedDrop() {
	_jsii_.InvokeVoid(
		o,
		"resetPowerFeedDrop",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetPowerPhase() {
	_jsii_.InvokeVoid(
		o,
		"resetPowerPhase",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetUplinkCount() {
	_jsii_.InvokeVoid(
		o,
		"resetUplinkCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ResetUplinkGbps() {
	_jsii_.InvokeVoid(
		o,
		"resetUplinkGbps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteRackPhysicalPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

