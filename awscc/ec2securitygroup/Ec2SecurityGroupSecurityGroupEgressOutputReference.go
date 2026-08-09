// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2securitygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2securitygroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SecurityGroupSecurityGroupEgressOutputReference interface {
	cdktn.ComplexObject
	CidrIp() *string
	SetCidrIp(val *string)
	CidrIpInput() *string
	CidrIpv6() *string
	SetCidrIpv6(val *string)
	CidrIpv6Input() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationPrefixListId() *string
	SetDestinationPrefixListId(val *string)
	DestinationPrefixListIdInput() *string
	DestinationSecurityGroupId() *string
	SetDestinationSecurityGroupId(val *string)
	DestinationSecurityGroupIdInput() *string
	// Experimental.
	Fqn() *string
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToPort() *float64
	SetToPort(val *float64)
	ToPortInput() *float64
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
	ResetCidrIp()
	ResetCidrIpv6()
	ResetDescription()
	ResetDestinationPrefixListId()
	ResetDestinationSecurityGroupId()
	ResetFromPort()
	ResetIpProtocol()
	ResetToPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2SecurityGroupSecurityGroupEgressOutputReference
type jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) CidrIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) CidrIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) CidrIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) CidrIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) DestinationPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) DestinationSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) DestinationSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}


func NewEc2SecurityGroupSecurityGroupEgressOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2SecurityGroupSecurityGroupEgressOutputReference {
	_init_.Initialize()

	if err := validateNewEc2SecurityGroupSecurityGroupEgressOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2SecurityGroup.Ec2SecurityGroupSecurityGroupEgressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2SecurityGroupSecurityGroupEgressOutputReference_Override(e Ec2SecurityGroupSecurityGroupEgressOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2SecurityGroup.Ec2SecurityGroupSecurityGroupEgressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetCidrIp(val *string) {
	if err := j.validateSetCidrIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrIp",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetCidrIpv6(val *string) {
	if err := j.validateSetCidrIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrIpv6",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetDestinationPrefixListId(val *string) {
	if err := j.validateSetDestinationPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixListId",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetDestinationSecurityGroupId(val *string) {
	if err := j.validateSetDestinationSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetFromPort(val *float64) {
	if err := j.validateSetFromPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference)SetToPort(val *float64) {
	if err := j.validateSetToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetCidrIp() {
	_jsii_.InvokeVoid(
		e,
		"resetCidrIp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetCidrIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetCidrIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetDestinationPrefixListId() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPrefixListId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetDestinationSecurityGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationSecurityGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetFromPort() {
	_jsii_.InvokeVoid(
		e,
		"resetFromPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ResetToPort() {
	_jsii_.InvokeVoid(
		e,
		"resetToPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

