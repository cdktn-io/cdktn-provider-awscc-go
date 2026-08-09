// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpnconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2vpnconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference interface {
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
	DpdTimeoutAction() *string
	SetDpdTimeoutAction(val *string)
	DpdTimeoutActionInput() *string
	DpdTimeoutSeconds() *float64
	SetDpdTimeoutSeconds(val *float64)
	DpdTimeoutSecondsInput() *float64
	EnableTunnelLifecycleControl() interface{}
	SetEnableTunnelLifecycleControl(val interface{})
	EnableTunnelLifecycleControlInput() interface{}
	// Experimental.
	Fqn() *string
	IkeVersions() Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList
	IkeVersionsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogOptions() Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference
	LogOptionsInput() interface{}
	Phase1DhGroupNumbers() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList
	Phase1DhGroupNumbersInput() interface{}
	Phase1EncryptionAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList
	Phase1EncryptionAlgorithmsInput() interface{}
	Phase1IntegrityAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList
	Phase1IntegrityAlgorithmsInput() interface{}
	Phase1LifetimeSeconds() *float64
	SetPhase1LifetimeSeconds(val *float64)
	Phase1LifetimeSecondsInput() *float64
	Phase2DhGroupNumbers() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList
	Phase2DhGroupNumbersInput() interface{}
	Phase2EncryptionAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList
	Phase2EncryptionAlgorithmsInput() interface{}
	Phase2IntegrityAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList
	Phase2IntegrityAlgorithmsInput() interface{}
	Phase2LifetimeSeconds() *float64
	SetPhase2LifetimeSeconds(val *float64)
	Phase2LifetimeSecondsInput() *float64
	PreSharedKey() *string
	SetPreSharedKey(val *string)
	PreSharedKeyInput() *string
	RekeyFuzzPercentage() *float64
	SetRekeyFuzzPercentage(val *float64)
	RekeyFuzzPercentageInput() *float64
	RekeyMarginTimeSeconds() *float64
	SetRekeyMarginTimeSeconds(val *float64)
	RekeyMarginTimeSecondsInput() *float64
	ReplayWindowSize() *float64
	SetReplayWindowSize(val *float64)
	ReplayWindowSizeInput() *float64
	StartupAction() *string
	SetStartupAction(val *string)
	StartupActionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TunnelInsideCidr() *string
	SetTunnelInsideCidr(val *string)
	TunnelInsideCidrInput() *string
	TunnelInsideIpv6Cidr() *string
	SetTunnelInsideIpv6Cidr(val *string)
	TunnelInsideIpv6CidrInput() *string
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
	PutIkeVersions(value interface{})
	PutLogOptions(value *Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptions)
	PutPhase1DhGroupNumbers(value interface{})
	PutPhase1EncryptionAlgorithms(value interface{})
	PutPhase1IntegrityAlgorithms(value interface{})
	PutPhase2DhGroupNumbers(value interface{})
	PutPhase2EncryptionAlgorithms(value interface{})
	PutPhase2IntegrityAlgorithms(value interface{})
	ResetDpdTimeoutAction()
	ResetDpdTimeoutSeconds()
	ResetEnableTunnelLifecycleControl()
	ResetIkeVersions()
	ResetLogOptions()
	ResetPhase1DhGroupNumbers()
	ResetPhase1EncryptionAlgorithms()
	ResetPhase1IntegrityAlgorithms()
	ResetPhase1LifetimeSeconds()
	ResetPhase2DhGroupNumbers()
	ResetPhase2EncryptionAlgorithms()
	ResetPhase2IntegrityAlgorithms()
	ResetPhase2LifetimeSeconds()
	ResetPreSharedKey()
	ResetRekeyFuzzPercentage()
	ResetRekeyMarginTimeSeconds()
	ResetReplayWindowSize()
	ResetStartupAction()
	ResetTunnelInsideCidr()
	ResetTunnelInsideIpv6Cidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference
type jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) IkeVersions() Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList
	_jsii_.Get(
		j,
		"ikeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) IkeVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ikeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) LogOptions() Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference
	_jsii_.Get(
		j,
		"logOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) LogOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1DhGroupNumbers() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList
	_jsii_.Get(
		j,
		"phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1DhGroupNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1EncryptionAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList
	_jsii_.Get(
		j,
		"phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1EncryptionAlgorithmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1IntegrityAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList
	_jsii_.Get(
		j,
		"phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1IntegrityAlgorithmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2DhGroupNumbers() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList
	_jsii_.Get(
		j,
		"phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2DhGroupNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2EncryptionAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList
	_jsii_.Get(
		j,
		"phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2EncryptionAlgorithmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2IntegrityAlgorithms() Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList {
	var returns Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList
	_jsii_.Get(
		j,
		"phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2IntegrityAlgorithmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PreSharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PreSharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpv6CidrInput",
		&returns,
	)
	return returns
}


func NewEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VpnConnection.Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference_Override(e Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VpnConnection.Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetDpdTimeoutAction(val *string) {
	if err := j.validateSetDpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetDpdTimeoutSeconds(val *float64) {
	if err := j.validateSetDpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetEnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetEnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetPhase1LifetimeSeconds(val *float64) {
	if err := j.validateSetPhase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetPhase2LifetimeSeconds(val *float64) {
	if err := j.validateSetPhase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetPreSharedKey(val *string) {
	if err := j.validateSetPreSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKey",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetRekeyFuzzPercentage(val *float64) {
	if err := j.validateSetRekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetRekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetRekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetReplayWindowSize(val *float64) {
	if err := j.validateSetReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replayWindowSize",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetStartupAction(val *string) {
	if err := j.validateSetStartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupAction",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTunnelInsideCidr(val *string) {
	if err := j.validateSetTunnelInsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelInsideCidr",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTunnelInsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnelInsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelInsideIpv6Cidr",
		val,
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutIkeVersions(value interface{}) {
	if err := e.validatePutIkeVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIkeVersions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutLogOptions(value *Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptions) {
	if err := e.validatePutLogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase1DhGroupNumbers(value interface{}) {
	if err := e.validatePutPhase1DhGroupNumbersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase1DhGroupNumbers",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase1EncryptionAlgorithms(value interface{}) {
	if err := e.validatePutPhase1EncryptionAlgorithmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase1EncryptionAlgorithms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase1IntegrityAlgorithms(value interface{}) {
	if err := e.validatePutPhase1IntegrityAlgorithmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase1IntegrityAlgorithms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase2DhGroupNumbers(value interface{}) {
	if err := e.validatePutPhase2DhGroupNumbersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase2DhGroupNumbers",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase2EncryptionAlgorithms(value interface{}) {
	if err := e.validatePutPhase2EncryptionAlgorithmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase2EncryptionAlgorithms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PutPhase2IntegrityAlgorithms(value interface{}) {
	if err := e.validatePutPhase2IntegrityAlgorithmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPhase2IntegrityAlgorithms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetDpdTimeoutAction() {
	_jsii_.InvokeVoid(
		e,
		"resetDpdTimeoutAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetDpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetDpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetEnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetIkeVersions() {
	_jsii_.InvokeVoid(
		e,
		"resetIkeVersions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetLogOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetLogOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPhase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetPhase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetPreSharedKey() {
	_jsii_.InvokeVoid(
		e,
		"resetPreSharedKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetRekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetRekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetRekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetRekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetReplayWindowSize() {
	_jsii_.InvokeVoid(
		e,
		"resetReplayWindowSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetStartupAction() {
	_jsii_.InvokeVoid(
		e,
		"resetStartupAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetTunnelInsideCidr() {
	_jsii_.InvokeVoid(
		e,
		"resetTunnelInsideCidr",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ResetTunnelInsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		e,
		"resetTunnelInsideIpv6Cidr",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

