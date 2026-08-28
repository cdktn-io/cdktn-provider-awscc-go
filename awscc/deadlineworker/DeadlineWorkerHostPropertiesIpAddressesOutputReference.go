// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlineworker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlineworker/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineWorkerHostPropertiesIpAddressesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpV4Addresses() *[]*string
	SetIpV4Addresses(val *[]*string)
	IpV4AddressesInput() *[]*string
	IpV6Addresses() *[]*string
	SetIpV6Addresses(val *[]*string)
	IpV6AddressesInput() *[]*string
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
	ResetIpV4Addresses()
	ResetIpV6Addresses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeadlineWorkerHostPropertiesIpAddressesOutputReference
type jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) IpV4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipV4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) IpV4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipV4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) IpV6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipV6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) IpV6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipV6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeadlineWorkerHostPropertiesIpAddressesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineWorkerHostPropertiesIpAddressesOutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineWorkerHostPropertiesIpAddressesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineWorker.DeadlineWorkerHostPropertiesIpAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineWorkerHostPropertiesIpAddressesOutputReference_Override(d DeadlineWorkerHostPropertiesIpAddressesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineWorker.DeadlineWorkerHostPropertiesIpAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetIpV4Addresses(val *[]*string) {
	if err := j.validateSetIpV4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV4Addresses",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetIpV6Addresses(val *[]*string) {
	if err := j.validateSetIpV6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV6Addresses",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ResetIpV4Addresses() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV4Addresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ResetIpV6Addresses() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV6Addresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeadlineWorkerHostPropertiesIpAddressesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

