// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdeadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdeadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference interface {
	cdktn.ComplexObject
	AutoScalingConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference
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
	InstanceCapabilities() DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference
	InstanceMarketOptions() DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference
	InternalValue() *DataAwsccDeadlineFleetConfigurationServiceManagedEc2
	SetInternalValue(val *DataAwsccDeadlineFleetConfigurationServiceManagedEc2)
	PersistentVolumeConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference
	StorageProfileId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference
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

// The jsii proxy struct for DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference
type jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) AutoScalingConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference {
	var returns DataAwsccDeadlineFleetConfigurationServiceManagedEc2AutoScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceCapabilities() DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference {
	var returns DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"instanceCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) InstanceMarketOptions() DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference {
	var returns DataAwsccDeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) InternalValue() *DataAwsccDeadlineFleetConfigurationServiceManagedEc2 {
	var returns *DataAwsccDeadlineFleetConfigurationServiceManagedEc2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) PersistentVolumeConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference {
	var returns DataAwsccDeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) StorageProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) VpcConfiguration() DataAwsccDeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference {
	var returns DataAwsccDeadlineFleetConfigurationServiceManagedEc2VpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}


func NewDataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDeadlineFleet.DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference_Override(d DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDeadlineFleet.DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference)SetInternalValue(val *DataAwsccDeadlineFleetConfigurationServiceManagedEc2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDeadlineFleetConfigurationServiceManagedEc2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

