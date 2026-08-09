// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcVpcEncryptionControlOutputReference interface {
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
	EgressOnlyInternetGatewayExclusion() *string
	SetEgressOnlyInternetGatewayExclusion(val *string)
	EgressOnlyInternetGatewayExclusionInput() *string
	ElasticFileSystemExclusion() *string
	SetElasticFileSystemExclusion(val *string)
	ElasticFileSystemExclusionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InternetGatewayExclusion() *string
	SetInternetGatewayExclusion(val *string)
	InternetGatewayExclusionInput() *string
	LambdaExclusion() *string
	SetLambdaExclusion(val *string)
	LambdaExclusionInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	NatGatewayExclusion() *string
	SetNatGatewayExclusion(val *string)
	NatGatewayExclusionInput() *string
	ResourceExclusions() Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference
	State() *string
	StateMessage() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VirtualPrivateGatewayExclusion() *string
	SetVirtualPrivateGatewayExclusion(val *string)
	VirtualPrivateGatewayExclusionInput() *string
	VpcEncryptionControlId() *string
	VpcId() *string
	VpcLatticeExclusion() *string
	SetVpcLatticeExclusion(val *string)
	VpcLatticeExclusionInput() *string
	VpcPeeringExclusion() *string
	SetVpcPeeringExclusion(val *string)
	VpcPeeringExclusionInput() *string
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
	ResetEgressOnlyInternetGatewayExclusion()
	ResetElasticFileSystemExclusion()
	ResetInternetGatewayExclusion()
	ResetLambdaExclusion()
	ResetMode()
	ResetNatGatewayExclusion()
	ResetVirtualPrivateGatewayExclusion()
	ResetVpcLatticeExclusion()
	ResetVpcPeeringExclusion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VpcVpcEncryptionControlOutputReference
type jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) EgressOnlyInternetGatewayExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) EgressOnlyInternetGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ElasticFileSystemExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticFileSystemExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ElasticFileSystemExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticFileSystemExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) InternetGatewayExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetGatewayExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) InternetGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) LambdaExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) LambdaExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) NatGatewayExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) NatGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResourceExclusions() Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference
	_jsii_.Get(
		j,
		"resourceExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VirtualPrivateGatewayExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPrivateGatewayExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VirtualPrivateGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPrivateGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcEncryptionControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEncryptionControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcLatticeExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLatticeExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcLatticeExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLatticeExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcPeeringExclusion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringExclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) VpcPeeringExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringExclusionInput",
		&returns,
	)
	return returns
}


func NewEc2VpcVpcEncryptionControlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2VpcVpcEncryptionControlOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VpcVpcEncryptionControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Vpc.Ec2VpcVpcEncryptionControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VpcVpcEncryptionControlOutputReference_Override(e Ec2VpcVpcEncryptionControlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Vpc.Ec2VpcVpcEncryptionControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetEgressOnlyInternetGatewayExclusion(val *string) {
	if err := j.validateSetEgressOnlyInternetGatewayExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressOnlyInternetGatewayExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetElasticFileSystemExclusion(val *string) {
	if err := j.validateSetElasticFileSystemExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticFileSystemExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetInternetGatewayExclusion(val *string) {
	if err := j.validateSetInternetGatewayExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetGatewayExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetLambdaExclusion(val *string) {
	if err := j.validateSetLambdaExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetNatGatewayExclusion(val *string) {
	if err := j.validateSetNatGatewayExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetVirtualPrivateGatewayExclusion(val *string) {
	if err := j.validateSetVirtualPrivateGatewayExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualPrivateGatewayExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetVpcLatticeExclusion(val *string) {
	if err := j.validateSetVpcLatticeExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcLatticeExclusion",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference)SetVpcPeeringExclusion(val *string) {
	if err := j.validateSetVpcPeeringExclusionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringExclusion",
		val,
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetEgressOnlyInternetGatewayExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetEgressOnlyInternetGatewayExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetElasticFileSystemExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticFileSystemExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetInternetGatewayExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetInternetGatewayExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetLambdaExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetLambdaExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		e,
		"resetMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetNatGatewayExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetNatGatewayExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetVirtualPrivateGatewayExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetVirtualPrivateGatewayExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetVpcLatticeExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcLatticeExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ResetVpcPeeringExclusion() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcPeeringExclusion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

