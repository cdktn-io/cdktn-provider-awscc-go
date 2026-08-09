// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2networkinsightsaccessscope/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference interface {
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
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	DestinationPrefixLists() *[]*string
	SetDestinationPrefixLists(val *[]*string)
	DestinationPrefixListsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourcePorts() *[]*string
	SetSourcePorts(val *[]*string)
	SourcePortsInput() *[]*string
	SourcePrefixLists() *[]*string
	SetSourcePrefixLists(val *[]*string)
	SourcePrefixListsInput() *[]*string
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
	ResetDestinationAddresses()
	ResetDestinationPorts()
	ResetDestinationPrefixLists()
	ResetProtocols()
	ResetSourceAddresses()
	ResetSourcePorts()
	ResetSourcePrefixLists()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference
type jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationPrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) DestinationPrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourcePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourcePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourcePrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) SourcePrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference_Override(e Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetDestinationPrefixLists(val *[]*string) {
	if err := j.validateSetDestinationPrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixLists",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetSourcePorts(val *[]*string) {
	if err := j.validateSetSourcePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePorts",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetSourcePrefixLists(val *[]*string) {
	if err := j.validateSetSourcePrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePrefixLists",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetDestinationPrefixLists() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPrefixLists",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocols",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ResetSourcePrefixLists() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePrefixLists",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

