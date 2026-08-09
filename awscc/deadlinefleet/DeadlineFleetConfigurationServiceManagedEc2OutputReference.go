// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineFleetConfigurationServiceManagedEc2OutputReference interface {
	cdktn.ComplexObject
	AutoScalingConfiguration() DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference
	AutoScalingConfigurationInput() interface{}
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
	InstanceCapabilities() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference
	InstanceCapabilitiesInput() interface{}
	InstanceMarketOptions() DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference
	InstanceMarketOptionsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PersistentVolumeConfiguration() DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference
	PersistentVolumeConfigurationInput() interface{}
	StorageProfileId() *string
	SetStorageProfileId(val *string)
	StorageProfileIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcConfiguration() DeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference
	VpcConfigurationInput() interface{}
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
	PutAutoScalingConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfiguration)
	PutInstanceCapabilities(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities)
	PutInstanceMarketOptions(value *DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptions)
	PutPersistentVolumeConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfiguration)
	PutVpcConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2VpcConfiguration)
	ResetAutoScalingConfiguration()
	ResetInstanceCapabilities()
	ResetInstanceMarketOptions()
	ResetPersistentVolumeConfiguration()
	ResetStorageProfileId()
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeadlineFleetConfigurationServiceManagedEc2OutputReference
type jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) AutoScalingConfiguration() DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) AutoScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceCapabilities() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"instanceCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceMarketOptions() DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceMarketOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PersistentVolumeConfiguration() DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PersistentVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistentVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) StorageProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) StorageProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) VpcConfiguration() DeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) VpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewDeadlineFleetConfigurationServiceManagedEc2OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineFleetConfigurationServiceManagedEc2OutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineFleetConfigurationServiceManagedEc2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineFleetConfigurationServiceManagedEc2OutputReference_Override(d DeadlineFleetConfigurationServiceManagedEc2OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetStorageProfileId(val *string) {
	if err := j.validateSetStorageProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfileId",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PutAutoScalingConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfiguration) {
	if err := d.validatePutAutoScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoScalingConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PutInstanceCapabilities(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities) {
	if err := d.validatePutInstanceCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstanceCapabilities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PutInstanceMarketOptions(value *DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptions) {
	if err := d.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PutPersistentVolumeConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfiguration) {
	if err := d.validatePutPersistentVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersistentVolumeConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) PutVpcConfiguration(value *DeadlineFleetConfigurationServiceManagedEc2VpcConfiguration) {
	if err := d.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetAutoScalingConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoScalingConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetInstanceCapabilities() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceCapabilities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetPersistentVolumeConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistentVolumeConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetStorageProfileId() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageProfileId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

