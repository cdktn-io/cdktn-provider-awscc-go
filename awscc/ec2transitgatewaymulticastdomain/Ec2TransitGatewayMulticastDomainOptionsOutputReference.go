// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymulticastdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2transitgatewaymulticastdomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TransitGatewayMulticastDomainOptionsOutputReference interface {
	cdktn.ComplexObject
	AutoAcceptSharedAssociations() *string
	SetAutoAcceptSharedAssociations(val *string)
	AutoAcceptSharedAssociationsInput() *string
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
	Igmpv2Support() *string
	SetIgmpv2Support(val *string)
	Igmpv2SupportInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StaticSourcesSupport() *string
	SetStaticSourcesSupport(val *string)
	StaticSourcesSupportInput() *string
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
	ResetAutoAcceptSharedAssociations()
	ResetIgmpv2Support()
	ResetStaticSourcesSupport()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2TransitGatewayMulticastDomainOptionsOutputReference
type jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) AutoAcceptSharedAssociations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) AutoAcceptSharedAssociationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) Igmpv2Support() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igmpv2Support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) Igmpv2SupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igmpv2SupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) StaticSourcesSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticSourcesSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) StaticSourcesSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticSourcesSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2TransitGatewayMulticastDomainOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2TransitGatewayMulticastDomainOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2TransitGatewayMulticastDomainOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomainOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2TransitGatewayMulticastDomainOptionsOutputReference_Override(e Ec2TransitGatewayMulticastDomainOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomainOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetAutoAcceptSharedAssociations(val *string) {
	if err := j.validateSetAutoAcceptSharedAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAcceptSharedAssociations",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetIgmpv2Support(val *string) {
	if err := j.validateSetIgmpv2SupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"igmpv2Support",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetStaticSourcesSupport(val *string) {
	if err := j.validateSetStaticSourcesSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticSourcesSupport",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ResetAutoAcceptSharedAssociations() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoAcceptSharedAssociations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ResetIgmpv2Support() {
	_jsii_.InvokeVoid(
		e,
		"resetIgmpv2Support",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ResetStaticSourcesSupport() {
	_jsii_.InvokeVoid(
		e,
		"resetStaticSourcesSupport",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomainOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

