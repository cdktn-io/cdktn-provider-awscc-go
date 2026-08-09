// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2spotfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference interface {
	cdktn.ComplexObject
	AcceleratorCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorManufacturers() *[]*string
	AcceleratorNames() *[]*string
	AcceleratorTotalMemoryMiB() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	AcceleratorTypes() *[]*string
	AllowedInstanceTypes() *[]*string
	BareMetal() *string
	BaselineEbsBandwidthMbps() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselinePerformanceFactors() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsOutputReference
	BurstablePerformance() *string
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
	CpuManufacturers() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedInstanceTypes() *[]*string
	// Experimental.
	Fqn() *string
	InstanceGenerations() *[]*string
	InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements
	SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements)
	LocalStorage() *string
	LocalStorageTypes() *[]*string
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	MemoryGiBPerVCpu() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryGiBPerVCpuOutputReference
	MemoryMiB() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference
	NetworkBandwidthGbps() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkInterfaceCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkInterfaceCountOutputReference
	OnDemandMaxPricePercentageOverLowestPrice() *float64
	RequireEncryptionInTransit() cdktn.IResolvable
	RequireHibernateSupport() cdktn.IResolvable
	SpotMaxPricePercentageOverLowestPrice() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalLocalStorageGb() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsTotalLocalStorageGbOutputReference
	VCpuCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsVCpuCountOutputReference
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

// The jsii proxy struct for DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference
type jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AcceleratorCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorCountOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AcceleratorTotalMemoryMiB() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) BaselinePerformanceFactors() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsOutputReference
	_jsii_.Get(
		j,
		"baselinePerformanceFactors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements {
	var returns *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) MemoryGiBPerVCpu() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryGiBPerVCpuOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryGiBPerVCpuOutputReference
	_jsii_.Get(
		j,
		"memoryGiBPerVCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) MemoryMiB() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) NetworkBandwidthGbps() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) NetworkInterfaceCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) RequireEncryptionInTransit() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requireEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) RequireHibernateSupport() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) TotalLocalStorageGb() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) VCpuCount() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsVCpuCountOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference_Override(d DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference)SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

