// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2ec2fleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference interface {
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	Ebs() Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsOutputReference
	EbsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoDevice() *string
	SetNoDevice(val *string)
	NoDeviceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VirtualName() *string
	SetVirtualName(val *string)
	VirtualNameInput() *string
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
	PutEbs(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbs)
	ResetDeviceName()
	ResetEbs()
	ResetNoDevice()
	ResetVirtualName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference
type jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) Ebs() Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) EbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) NoDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) NoDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) VirtualNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNameInput",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference_Override(e Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetNoDevice(val *string) {
	if err := j.validateSetNoDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDevice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference)SetVirtualName(val *string) {
	if err := j.validateSetVirtualNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualName",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) PutEbs(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbs) {
	if err := e.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ResetEbs() {
	_jsii_.InvokeVoid(
		e,
		"resetEbs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ResetNoDevice() {
	_jsii_.InvokeVoid(
		e,
		"resetNoDevice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ResetVirtualName() {
	_jsii_.InvokeVoid(
		e,
		"resetVirtualName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

