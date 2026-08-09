// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference interface {
	cdktn.ComplexObject
	AcceleratorCount() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCountOutputReference
	AcceleratorCountInput() interface{}
	AcceleratorTotalMemoryMiB() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiBOutputReference
	AcceleratorTotalMemoryMiBInput() interface{}
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
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
	CustomAmounts() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAmountsList
	CustomAmountsInput() interface{}
	CustomAttributes() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAttributesList
	CustomAttributesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemoryMiB() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiBOutputReference
	MemoryMiBInput() interface{}
	OsFamily() *string
	SetOsFamily(val *string)
	OsFamilyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VCpuCount() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCountOutputReference
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
	PutAcceleratorCount(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCount)
	PutAcceleratorTotalMemoryMiB(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiB)
	PutCustomAmounts(value interface{})
	PutCustomAttributes(value interface{})
	PutMemoryMiB(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiB)
	PutVCpuCount(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCount)
	ResetAcceleratorCount()
	ResetAcceleratorTotalMemoryMiB()
	ResetAcceleratorTypes()
	ResetCpuArchitectureType()
	ResetCustomAmounts()
	ResetCustomAttributes()
	ResetMemoryMiB()
	ResetOsFamily()
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

// The jsii proxy struct for DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference
type jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorCount() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCountOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorTotalMemoryMiB() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiBOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorTotalMemoryMiBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CpuArchitectureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuArchitectureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CpuArchitectureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuArchitectureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CustomAmounts() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAmountsList {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAmountsList
	_jsii_.Get(
		j,
		"customAmounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CustomAmountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAmountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CustomAttributes() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAttributesList {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAttributesList
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) CustomAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) MemoryMiB() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiBOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) MemoryMiBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryMiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) OsFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) OsFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) VCpuCount() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCountOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) VCpuCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vCpuCountInput",
		&returns,
	)
	return returns
}


func NewDeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference_Override(d DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetCpuArchitectureType(val *string) {
	if err := j.validateSetCpuArchitectureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuArchitectureType",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetOsFamily(val *string) {
	if err := j.validateSetOsFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osFamily",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutAcceleratorCount(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCount) {
	if err := d.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutAcceleratorTotalMemoryMiB(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiB) {
	if err := d.validatePutAcceleratorTotalMemoryMiBParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAcceleratorTotalMemoryMiB",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutCustomAmounts(value interface{}) {
	if err := d.validatePutCustomAmountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomAmounts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutCustomAttributes(value interface{}) {
	if err := d.validatePutCustomAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutMemoryMiB(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiB) {
	if err := d.validatePutMemoryMiBParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMemoryMiB",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) PutVCpuCount(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCount) {
	if err := d.validatePutVCpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVCpuCount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetAcceleratorTotalMemoryMiB() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceleratorTotalMemoryMiB",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetCpuArchitectureType() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuArchitectureType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetCustomAmounts() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAmounts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetMemoryMiB() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryMiB",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetOsFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetOsFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ResetVCpuCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVCpuCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

