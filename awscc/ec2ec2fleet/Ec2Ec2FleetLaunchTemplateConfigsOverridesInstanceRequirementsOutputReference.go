// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2ec2fleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference interface {
	cdktn.ComplexObject
	AcceleratorCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorCountInput() interface{}
	AcceleratorManufacturers() *[]*string
	SetAcceleratorManufacturers(val *[]*string)
	AcceleratorManufacturersInput() *[]*string
	AcceleratorNames() *[]*string
	SetAcceleratorNames(val *[]*string)
	AcceleratorNamesInput() *[]*string
	AcceleratorTotalMemoryMiB() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	AcceleratorTotalMemoryMiBInput() interface{}
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
	AllowedInstanceTypes() *[]*string
	SetAllowedInstanceTypes(val *[]*string)
	AllowedInstanceTypesInput() *[]*string
	BareMetal() *string
	SetBareMetal(val *string)
	BareMetalInput() *string
	BaselineEbsBandwidthMbps() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselineEbsBandwidthMbpsInput() interface{}
	BaselinePerformanceFactors() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference
	BaselinePerformanceFactorsInput() interface{}
	BurstablePerformance() *string
	SetBurstablePerformance(val *string)
	BurstablePerformanceInput() *string
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
	SetCpuManufacturers(val *[]*string)
	CpuManufacturersInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedInstanceTypes() *[]*string
	SetExcludedInstanceTypes(val *[]*string)
	ExcludedInstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InstanceGenerations() *[]*string
	SetInstanceGenerations(val *[]*string)
	InstanceGenerationsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalStorage() *string
	SetLocalStorage(val *string)
	LocalStorageInput() *string
	LocalStorageTypes() *[]*string
	SetLocalStorageTypes(val *[]*string)
	LocalStorageTypesInput() *[]*string
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64)
	MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64
	MemoryGiBPerVCpu() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference
	MemoryGiBPerVCpuInput() interface{}
	MemoryMiB() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiBOutputReference
	MemoryMiBInput() interface{}
	NetworkBandwidthGbps() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkBandwidthGbpsInput() interface{}
	NetworkInterfaceCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
	NetworkInterfaceCountInput() interface{}
	OnDemandMaxPricePercentageOverLowestPrice() *float64
	SetOnDemandMaxPricePercentageOverLowestPrice(val *float64)
	OnDemandMaxPricePercentageOverLowestPriceInput() *float64
	RequireEncryptionInTransit() interface{}
	SetRequireEncryptionInTransit(val interface{})
	RequireEncryptionInTransitInput() interface{}
	RequireHibernateSupport() interface{}
	SetRequireHibernateSupport(val interface{})
	RequireHibernateSupportInput() interface{}
	SpotMaxPricePercentageOverLowestPrice() *float64
	SetSpotMaxPricePercentageOverLowestPrice(val *float64)
	SpotMaxPricePercentageOverLowestPriceInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalLocalStorageGb() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	TotalLocalStorageGbInput() interface{}
	VCpuCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference
	VCpuCountInput() interface{}
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
	PutAcceleratorCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount)
	PutAcceleratorTotalMemoryMiB(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorTotalMemoryMiB)
	PutBaselineEbsBandwidthMbps(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselineEbsBandwidthMbps)
	PutBaselinePerformanceFactors(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors)
	PutMemoryGiBPerVCpu(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryGiBPerVCpu)
	PutMemoryMiB(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiB)
	PutNetworkBandwidthGbps(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbps)
	PutNetworkInterfaceCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkInterfaceCount)
	PutTotalLocalStorageGb(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsTotalLocalStorageGb)
	PutVCpuCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount)
	ResetAcceleratorCount()
	ResetAcceleratorManufacturers()
	ResetAcceleratorNames()
	ResetAcceleratorTotalMemoryMiB()
	ResetAcceleratorTypes()
	ResetAllowedInstanceTypes()
	ResetBareMetal()
	ResetBaselineEbsBandwidthMbps()
	ResetBaselinePerformanceFactors()
	ResetBurstablePerformance()
	ResetCpuManufacturers()
	ResetExcludedInstanceTypes()
	ResetInstanceGenerations()
	ResetLocalStorage()
	ResetLocalStorageTypes()
	ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice()
	ResetMemoryGiBPerVCpu()
	ResetMemoryMiB()
	ResetNetworkBandwidthGbps()
	ResetNetworkInterfaceCount()
	ResetOnDemandMaxPricePercentageOverLowestPrice()
	ResetRequireEncryptionInTransit()
	ResetRequireHibernateSupport()
	ResetSpotMaxPricePercentageOverLowestPrice()
	ResetTotalLocalStorageGb()
	ResetVCpuCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference
type jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorTotalMemoryMiB() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorTotalMemoryMiBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) AllowedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BaselineEbsBandwidthMbpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BaselinePerformanceFactors() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference
	_jsii_.Get(
		j,
		"baselinePerformanceFactors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BaselinePerformanceFactorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baselinePerformanceFactorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MemoryGiBPerVCpu() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference
	_jsii_.Get(
		j,
		"memoryGiBPerVCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MemoryGiBPerVCpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryGiBPerVCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MemoryMiB() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiBOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) MemoryMiBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryMiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) NetworkBandwidthGbps() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) NetworkBandwidthGbpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkBandwidthGbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) NetworkInterfaceCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) NetworkInterfaceCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) RequireEncryptionInTransit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) RequireEncryptionInTransitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEncryptionInTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) TotalLocalStorageGb() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) TotalLocalStorageGbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) VCpuCount() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) VCpuCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vCpuCountInput",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference_Override(e Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetAllowedInstanceTypes(val *[]*string) {
	if err := j.validateSetAllowedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64) {
	if err := j.validateSetMaxSpotPriceAsPercentageOfOptimalOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetRequireEncryptionInTransit(val interface{}) {
	if err := j.validateSetRequireEncryptionInTransitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEncryptionInTransit",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutAcceleratorCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount) {
	if err := e.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutAcceleratorTotalMemoryMiB(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorTotalMemoryMiB) {
	if err := e.validatePutAcceleratorTotalMemoryMiBParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAcceleratorTotalMemoryMiB",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutBaselineEbsBandwidthMbps(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselineEbsBandwidthMbps) {
	if err := e.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutBaselinePerformanceFactors(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors) {
	if err := e.validatePutBaselinePerformanceFactorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBaselinePerformanceFactors",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutMemoryGiBPerVCpu(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryGiBPerVCpu) {
	if err := e.validatePutMemoryGiBPerVCpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMemoryGiBPerVCpu",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutMemoryMiB(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiB) {
	if err := e.validatePutMemoryMiBParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMemoryMiB",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutNetworkBandwidthGbps(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbps) {
	if err := e.validatePutNetworkBandwidthGbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkBandwidthGbps",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutNetworkInterfaceCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsNetworkInterfaceCount) {
	if err := e.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutTotalLocalStorageGb(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsTotalLocalStorageGb) {
	if err := e.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) PutVCpuCount(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount) {
	if err := e.validatePutVCpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVCpuCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAcceleratorTotalMemoryMiB() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorTotalMemoryMiB",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetAllowedInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowedInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		e,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		e,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetBaselinePerformanceFactors() {
	_jsii_.InvokeVoid(
		e,
		"resetBaselinePerformanceFactors",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		e,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetMemoryGiBPerVCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryGiBPerVCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetMemoryMiB() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryMiB",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetNetworkBandwidthGbps() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkBandwidthGbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetRequireEncryptionInTransit() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireEncryptionInTransit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		e,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ResetVCpuCount() {
	_jsii_.InvokeVoid(
		e,
		"resetVCpuCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

