// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccecscapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccecscapacityprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference interface {
	cdktn.ComplexObject
	CapacityOptionType() *string
	CapacityReservations() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateCapacityReservationsOutputReference
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
	Ec2InstanceProfileArn() *string
	FipsEnabled() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InstanceMetadataTagsPropagation() cdktn.IResolvable
	InstanceRequirements() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference
	InternalValue() *DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	SetInternalValue(val *DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate)
	LocalStorageConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfigurationOutputReference
	Monitoring() *string
	NetworkConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference
	StorageConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference
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

// The jsii proxy struct for DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference
type jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) CapacityOptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityOptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) CapacityReservations() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateCapacityReservationsOutputReference {
	var returns DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateCapacityReservationsOutputReference
	_jsii_.Get(
		j,
		"capacityReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Ec2InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) FipsEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InstanceMetadataTagsPropagation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"instanceMetadataTagsPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InstanceRequirements() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference {
	var returns DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InternalValue() *DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate {
	var returns *DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) LocalStorageConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfigurationOutputReference {
	var returns DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"localStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Monitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) NetworkConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference {
	var returns DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) StorageConfiguration() DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference {
	var returns DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEcsCapacityProvider.DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference_Override(d DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEcsCapacityProvider.DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetInternalValue(val *DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

