// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference interface {
	cdktn.ComplexObject
	AcceleratorCapabilities() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilitiesOutputReference
	AcceleratorCapabilitiesInput() interface{}
	AllowedInstanceTypes() *[]*string
	SetAllowedInstanceTypes(val *[]*string)
	AllowedInstanceTypesInput() *[]*string
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
	CpuArchitectureType() *string
	SetCpuArchitectureType(val *string)
	CpuArchitectureTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAmounts() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmountsList
	CustomAmountsInput() interface{}
	CustomAttributes() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributesList
	CustomAttributesInput() interface{}
	ExcludedInstanceTypes() *[]*string
	SetExcludedInstanceTypes(val *[]*string)
	ExcludedInstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemoryMiB() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiBOutputReference
	MemoryMiBInput() interface{}
	OsFamily() *string
	SetOsFamily(val *string)
	OsFamilyInput() *string
	RootEbsVolume() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolumeOutputReference
	RootEbsVolumeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VCpuCount() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCountOutputReference
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
	PutAcceleratorCapabilities(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilities)
	PutCustomAmounts(value interface{})
	PutCustomAttributes(value interface{})
	PutMemoryMiB(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiB)
	PutRootEbsVolume(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolume)
	PutVCpuCount(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCount)
	ResetAcceleratorCapabilities()
	ResetAllowedInstanceTypes()
	ResetCpuArchitectureType()
	ResetCustomAmounts()
	ResetCustomAttributes()
	ResetExcludedInstanceTypes()
	ResetMemoryMiB()
	ResetOsFamily()
	ResetRootEbsVolume()
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

// The jsii proxy struct for DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference
type jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) AcceleratorCapabilities() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilitiesOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"acceleratorCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) AcceleratorCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) AllowedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CpuArchitectureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuArchitectureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CpuArchitectureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuArchitectureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CustomAmounts() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmountsList {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmountsList
	_jsii_.Get(
		j,
		"customAmounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CustomAmountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAmountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CustomAttributes() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributesList {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributesList
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) CustomAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) MemoryMiB() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiBOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) MemoryMiBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryMiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) OsFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) OsFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) RootEbsVolume() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolumeOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolumeOutputReference
	_jsii_.Get(
		j,
		"rootEbsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) RootEbsVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootEbsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) VCpuCount() DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCountOutputReference {
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) VCpuCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vCpuCountInput",
		&returns,
	)
	return returns
}


func NewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference_Override(d DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetAllowedInstanceTypes(val *[]*string) {
	if err := j.validateSetAllowedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetCpuArchitectureType(val *string) {
	if err := j.validateSetCpuArchitectureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuArchitectureType",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetOsFamily(val *string) {
	if err := j.validateSetOsFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osFamily",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutAcceleratorCapabilities(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilities) {
	if err := d.validatePutAcceleratorCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAcceleratorCapabilities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutCustomAmounts(value interface{}) {
	if err := d.validatePutCustomAmountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomAmounts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutCustomAttributes(value interface{}) {
	if err := d.validatePutCustomAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutMemoryMiB(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiB) {
	if err := d.validatePutMemoryMiBParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMemoryMiB",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutRootEbsVolume(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolume) {
	if err := d.validatePutRootEbsVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRootEbsVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) PutVCpuCount(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCount) {
	if err := d.validatePutVCpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVCpuCount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetAcceleratorCapabilities() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceleratorCapabilities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetAllowedInstanceTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedInstanceTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetCpuArchitectureType() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuArchitectureType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetCustomAmounts() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAmounts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetMemoryMiB() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryMiB",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetOsFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetOsFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetRootEbsVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetRootEbsVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ResetVCpuCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVCpuCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

