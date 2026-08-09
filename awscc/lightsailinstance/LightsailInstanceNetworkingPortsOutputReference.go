// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lightsailinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailInstanceNetworkingPortsOutputReference interface {
	cdktn.ComplexObject
	AccessDirection() *string
	SetAccessDirection(val *string)
	AccessDirectionInput() *string
	AccessFrom() *string
	SetAccessFrom(val *string)
	AccessFromInput() *string
	AccessType() *string
	SetAccessType(val *string)
	AccessTypeInput() *string
	CidrListAliases() *[]*string
	SetCidrListAliases(val *[]*string)
	CidrListAliasesInput() *[]*string
	Cidrs() *[]*string
	SetCidrs(val *[]*string)
	CidrsInput() *[]*string
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6Cidrs() *[]*string
	SetIpv6Cidrs(val *[]*string)
	Ipv6CidrsInput() *[]*string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	ResetAccessDirection()
	ResetAccessFrom()
	ResetAccessType()
	ResetCidrListAliases()
	ResetCidrs()
	ResetCommonName()
	ResetFromPort()
	ResetIpv6Cidrs()
	ResetProtocol()
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

// The jsii proxy struct for LightsailInstanceNetworkingPortsOutputReference
type jsiiProxy_LightsailInstanceNetworkingPortsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CidrListAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrListAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CidrListAliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrListAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Ipv6Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Ipv6CidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}


func NewLightsailInstanceNetworkingPortsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LightsailInstanceNetworkingPortsOutputReference {
	_init_.Initialize()

	if err := validateNewLightsailInstanceNetworkingPortsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailInstanceNetworkingPortsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lightsailInstance.LightsailInstanceNetworkingPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLightsailInstanceNetworkingPortsOutputReference_Override(l LightsailInstanceNetworkingPortsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lightsailInstance.LightsailInstanceNetworkingPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetAccessDirection(val *string) {
	if err := j.validateSetAccessDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessDirection",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetAccessFrom(val *string) {
	if err := j.validateSetAccessFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessFrom",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetAccessType(val *string) {
	if err := j.validateSetAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessType",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetCidrListAliases(val *[]*string) {
	if err := j.validateSetCidrListAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrListAliases",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetCidrs(val *[]*string) {
	if err := j.validateSetCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrs",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetFromPort(val *float64) {
	if err := j.validateSetFromPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetIpv6Cidrs(val *[]*string) {
	if err := j.validateSetIpv6CidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Cidrs",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference)SetToPort(val *float64) {
	if err := j.validateSetToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetAccessDirection() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessDirection",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetAccessFrom() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessFrom",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetAccessType() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetCidrListAliases() {
	_jsii_.InvokeVoid(
		l,
		"resetCidrListAliases",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetCidrs() {
	_jsii_.InvokeVoid(
		l,
		"resetCidrs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		l,
		"resetCommonName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetFromPort() {
	_jsii_.InvokeVoid(
		l,
		"resetFromPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetIpv6Cidrs() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv6Cidrs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ResetToPort() {
	_jsii_.InvokeVoid(
		l,
		"resetToPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceNetworkingPortsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

