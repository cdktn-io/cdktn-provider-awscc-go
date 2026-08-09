// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotaccountauditconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotaccountauditconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference interface {
	cdktn.ComplexObject
	AuthenticatedCognitoRoleOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	AuthenticatedCognitoRoleOverlyPermissiveCheckInput() interface{}
	CaCertificateExpiringCheck() IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference
	CaCertificateExpiringCheckInput() interface{}
	CaCertificateKeyQualityCheck() IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference
	CaCertificateKeyQualityCheckInput() interface{}
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
	ConflictingClientIdsCheck() IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference
	ConflictingClientIdsCheckInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceCertificateAgeCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheckOutputReference
	DeviceCertificateAgeCheckInput() interface{}
	DeviceCertificateExpiringCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference
	DeviceCertificateExpiringCheckInput() interface{}
	DeviceCertificateKeyQualityCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference
	DeviceCertificateKeyQualityCheckInput() interface{}
	DeviceCertificateSharedCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference
	DeviceCertificateSharedCheckInput() interface{}
	// Experimental.
	Fqn() *string
	IntermediateCaRevokedForActiveDeviceCertificatesCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference
	IntermediateCaRevokedForActiveDeviceCertificatesCheckInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotPolicyOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference
	IotPolicyOverlyPermissiveCheckInput() interface{}
	IoTPolicyPotentialMisConfigurationCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference
	IoTPolicyPotentialMisConfigurationCheckInput() interface{}
	IotRoleAliasAllowsAccessToUnusedServicesCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference
	IotRoleAliasAllowsAccessToUnusedServicesCheckInput() interface{}
	IotRoleAliasOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference
	IotRoleAliasOverlyPermissiveCheckInput() interface{}
	LoggingDisabledCheck() IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference
	LoggingDisabledCheckInput() interface{}
	RevokedCaCertificateStillActiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference
	RevokedCaCertificateStillActiveCheckInput() interface{}
	RevokedDeviceCertificateStillActiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference
	RevokedDeviceCertificateStillActiveCheckInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnauthenticatedCognitoRoleOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	UnauthenticatedCognitoRoleOverlyPermissiveCheckInput() interface{}
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
	PutAuthenticatedCognitoRoleOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheck)
	PutCaCertificateExpiringCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheck)
	PutCaCertificateKeyQualityCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheck)
	PutConflictingClientIdsCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheck)
	PutDeviceCertificateAgeCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheck)
	PutDeviceCertificateExpiringCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheck)
	PutDeviceCertificateKeyQualityCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheck)
	PutDeviceCertificateSharedCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheck)
	PutIntermediateCaRevokedForActiveDeviceCertificatesCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck)
	PutIotPolicyOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheck)
	PutIoTPolicyPotentialMisConfigurationCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheck)
	PutIotRoleAliasAllowsAccessToUnusedServicesCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheck)
	PutIotRoleAliasOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheck)
	PutLoggingDisabledCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheck)
	PutRevokedCaCertificateStillActiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheck)
	PutRevokedDeviceCertificateStillActiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheck)
	PutUnauthenticatedCognitoRoleOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheck)
	ResetAuthenticatedCognitoRoleOverlyPermissiveCheck()
	ResetCaCertificateExpiringCheck()
	ResetCaCertificateKeyQualityCheck()
	ResetConflictingClientIdsCheck()
	ResetDeviceCertificateAgeCheck()
	ResetDeviceCertificateExpiringCheck()
	ResetDeviceCertificateKeyQualityCheck()
	ResetDeviceCertificateSharedCheck()
	ResetIntermediateCaRevokedForActiveDeviceCertificatesCheck()
	ResetIotPolicyOverlyPermissiveCheck()
	ResetIoTPolicyPotentialMisConfigurationCheck()
	ResetIotRoleAliasAllowsAccessToUnusedServicesCheck()
	ResetIotRoleAliasOverlyPermissiveCheck()
	ResetLoggingDisabledCheck()
	ResetRevokedCaCertificateStillActiveCheck()
	ResetRevokedDeviceCertificateStillActiveCheck()
	ResetUnauthenticatedCognitoRoleOverlyPermissiveCheck()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference
type jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) AuthenticatedCognitoRoleOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"authenticatedCognitoRoleOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) AuthenticatedCognitoRoleOverlyPermissiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticatedCognitoRoleOverlyPermissiveCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateExpiringCheck() IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference
	_jsii_.Get(
		j,
		"caCertificateExpiringCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateExpiringCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caCertificateExpiringCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateKeyQualityCheck() IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference
	_jsii_.Get(
		j,
		"caCertificateKeyQualityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateKeyQualityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caCertificateKeyQualityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ConflictingClientIdsCheck() IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference
	_jsii_.Get(
		j,
		"conflictingClientIdsCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ConflictingClientIdsCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conflictingClientIdsCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateAgeCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateAgeCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateAgeCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceCertificateAgeCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateExpiringCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateExpiringCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateExpiringCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceCertificateExpiringCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateKeyQualityCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateKeyQualityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateKeyQualityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceCertificateKeyQualityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateSharedCheck() IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateSharedCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateSharedCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceCertificateSharedCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IntermediateCaRevokedForActiveDeviceCertificatesCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference
	_jsii_.Get(
		j,
		"intermediateCaRevokedForActiveDeviceCertificatesCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IntermediateCaRevokedForActiveDeviceCertificatesCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intermediateCaRevokedForActiveDeviceCertificatesCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotPolicyOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"iotPolicyOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotPolicyOverlyPermissiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotPolicyOverlyPermissiveCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IoTPolicyPotentialMisConfigurationCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference
	_jsii_.Get(
		j,
		"ioTPolicyPotentialMisConfigurationCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IoTPolicyPotentialMisConfigurationCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ioTPolicyPotentialMisConfigurationCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasAllowsAccessToUnusedServicesCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference
	_jsii_.Get(
		j,
		"iotRoleAliasAllowsAccessToUnusedServicesCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasAllowsAccessToUnusedServicesCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotRoleAliasAllowsAccessToUnusedServicesCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"iotRoleAliasOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasOverlyPermissiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotRoleAliasOverlyPermissiveCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) LoggingDisabledCheck() IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference
	_jsii_.Get(
		j,
		"loggingDisabledCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) LoggingDisabledCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingDisabledCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedCaCertificateStillActiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference
	_jsii_.Get(
		j,
		"revokedCaCertificateStillActiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedCaCertificateStillActiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokedCaCertificateStillActiveCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedDeviceCertificateStillActiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference
	_jsii_.Get(
		j,
		"revokedDeviceCertificateStillActiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedDeviceCertificateStillActiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokedDeviceCertificateStillActiveCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) UnauthenticatedCognitoRoleOverlyPermissiveCheck() IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference {
	var returns IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"unauthenticatedCognitoRoleOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) UnauthenticatedCognitoRoleOverlyPermissiveCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unauthenticatedCognitoRoleOverlyPermissiveCheckInput",
		&returns,
	)
	return returns
}


func NewIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewIotAccountAuditConfigurationAuditCheckConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotAccountAuditConfiguration.IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference_Override(i IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotAccountAuditConfiguration.IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutAuthenticatedCognitoRoleOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheck) {
	if err := i.validatePutAuthenticatedCognitoRoleOverlyPermissiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthenticatedCognitoRoleOverlyPermissiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutCaCertificateExpiringCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheck) {
	if err := i.validatePutCaCertificateExpiringCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCaCertificateExpiringCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutCaCertificateKeyQualityCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheck) {
	if err := i.validatePutCaCertificateKeyQualityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCaCertificateKeyQualityCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutConflictingClientIdsCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheck) {
	if err := i.validatePutConflictingClientIdsCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putConflictingClientIdsCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutDeviceCertificateAgeCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheck) {
	if err := i.validatePutDeviceCertificateAgeCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceCertificateAgeCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutDeviceCertificateExpiringCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheck) {
	if err := i.validatePutDeviceCertificateExpiringCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceCertificateExpiringCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutDeviceCertificateKeyQualityCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheck) {
	if err := i.validatePutDeviceCertificateKeyQualityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceCertificateKeyQualityCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutDeviceCertificateSharedCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheck) {
	if err := i.validatePutDeviceCertificateSharedCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceCertificateSharedCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutIntermediateCaRevokedForActiveDeviceCertificatesCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck) {
	if err := i.validatePutIntermediateCaRevokedForActiveDeviceCertificatesCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIntermediateCaRevokedForActiveDeviceCertificatesCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutIotPolicyOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheck) {
	if err := i.validatePutIotPolicyOverlyPermissiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotPolicyOverlyPermissiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutIoTPolicyPotentialMisConfigurationCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheck) {
	if err := i.validatePutIoTPolicyPotentialMisConfigurationCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIoTPolicyPotentialMisConfigurationCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutIotRoleAliasAllowsAccessToUnusedServicesCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheck) {
	if err := i.validatePutIotRoleAliasAllowsAccessToUnusedServicesCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotRoleAliasAllowsAccessToUnusedServicesCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutIotRoleAliasOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheck) {
	if err := i.validatePutIotRoleAliasOverlyPermissiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotRoleAliasOverlyPermissiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutLoggingDisabledCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheck) {
	if err := i.validatePutLoggingDisabledCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLoggingDisabledCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutRevokedCaCertificateStillActiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheck) {
	if err := i.validatePutRevokedCaCertificateStillActiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRevokedCaCertificateStillActiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutRevokedDeviceCertificateStillActiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheck) {
	if err := i.validatePutRevokedDeviceCertificateStillActiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRevokedDeviceCertificateStillActiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) PutUnauthenticatedCognitoRoleOverlyPermissiveCheck(value *IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheck) {
	if err := i.validatePutUnauthenticatedCognitoRoleOverlyPermissiveCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUnauthenticatedCognitoRoleOverlyPermissiveCheck",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetAuthenticatedCognitoRoleOverlyPermissiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthenticatedCognitoRoleOverlyPermissiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetCaCertificateExpiringCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetCaCertificateExpiringCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetCaCertificateKeyQualityCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetCaCertificateKeyQualityCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetConflictingClientIdsCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetConflictingClientIdsCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetDeviceCertificateAgeCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceCertificateAgeCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetDeviceCertificateExpiringCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceCertificateExpiringCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetDeviceCertificateKeyQualityCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceCertificateKeyQualityCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetDeviceCertificateSharedCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceCertificateSharedCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetIntermediateCaRevokedForActiveDeviceCertificatesCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetIntermediateCaRevokedForActiveDeviceCertificatesCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetIotPolicyOverlyPermissiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetIotPolicyOverlyPermissiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetIoTPolicyPotentialMisConfigurationCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetIoTPolicyPotentialMisConfigurationCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetIotRoleAliasAllowsAccessToUnusedServicesCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetIotRoleAliasAllowsAccessToUnusedServicesCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetIotRoleAliasOverlyPermissiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetIotRoleAliasOverlyPermissiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetLoggingDisabledCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetLoggingDisabledCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetRevokedCaCertificateStillActiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetRevokedCaCertificateStillActiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetRevokedDeviceCertificateStillActiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetRevokedDeviceCertificateStillActiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ResetUnauthenticatedCognitoRoleOverlyPermissiveCheck() {
	_jsii_.InvokeVoid(
		i,
		"resetUnauthenticatedCognitoRoleOverlyPermissiveCheck",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

