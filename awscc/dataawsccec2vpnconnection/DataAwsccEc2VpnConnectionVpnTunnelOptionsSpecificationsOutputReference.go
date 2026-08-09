// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2vpnconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2vpnconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference interface {
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
	DpdTimeoutSeconds() *float64
	EnableTunnelLifecycleControl() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	IkeVersions() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList
	InternalValue() *DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecifications
	SetInternalValue(val *DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecifications)
	LogOptions() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference
	Phase1DhGroupNumbers() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList
	Phase1EncryptionAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList
	Phase1IntegrityAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList
	Phase1LifetimeSeconds() *float64
	Phase2DhGroupNumbers() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList
	Phase2EncryptionAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList
	Phase2IntegrityAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList
	Phase2LifetimeSeconds() *float64
	PreSharedKey() *string
	RekeyFuzzPercentage() *float64
	RekeyMarginTimeSeconds() *float64
	ReplayWindowSize() *float64
	StartupAction() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TunnelInsideCidr() *string
	TunnelInsideIpv6Cidr() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference
type jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) EnableTunnelLifecycleControl() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) IkeVersions() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersionsList
	_jsii_.Get(
		j,
		"ikeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InternalValue() *DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecifications {
	var returns *DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) LogOptions() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsOutputReference
	_jsii_.Get(
		j,
		"logOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1DhGroupNumbers() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbersList
	_jsii_.Get(
		j,
		"phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1EncryptionAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithmsList
	_jsii_.Get(
		j,
		"phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1IntegrityAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithmsList
	_jsii_.Get(
		j,
		"phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2DhGroupNumbers() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbersList
	_jsii_.Get(
		j,
		"phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2EncryptionAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithmsList
	_jsii_.Get(
		j,
		"phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2IntegrityAlgorithms() DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList {
	var returns DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithmsList
	_jsii_.Get(
		j,
		"phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) PreSharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) TunnelInsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpv6Cidr",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2VpnConnection.DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference_Override(d DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2VpnConnection.DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetInternalValue(val *DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

