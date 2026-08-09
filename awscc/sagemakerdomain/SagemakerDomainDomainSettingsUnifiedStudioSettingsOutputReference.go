// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerdomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference interface {
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
	DomainAccountId() *string
	SetDomainAccountId(val *string)
	DomainAccountIdInput() *string
	DomainId() *string
	SetDomainId(val *string)
	DomainIdInput() *string
	DomainRegion() *string
	SetDomainRegion(val *string)
	DomainRegionInput() *string
	EnvironmentId() *string
	SetEnvironmentId(val *string)
	EnvironmentIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	ProjectS3Path() *string
	SetProjectS3Path(val *string)
	ProjectS3PathInput() *string
	StudioWebPortalAccess() *string
	SetStudioWebPortalAccess(val *string)
	StudioWebPortalAccessInput() *string
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
	ResetDomainAccountId()
	ResetDomainId()
	ResetDomainRegion()
	ResetEnvironmentId()
	ResetProjectId()
	ResetProjectS3Path()
	ResetStudioWebPortalAccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference
type jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) DomainRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ProjectS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ProjectS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) StudioWebPortalAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) StudioWebPortalAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference_Override(s SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetDomainAccountId(val *string) {
	if err := j.validateSetDomainAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAccountId",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetDomainId(val *string) {
	if err := j.validateSetDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetDomainRegion(val *string) {
	if err := j.validateSetDomainRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainRegion",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetEnvironmentId(val *string) {
	if err := j.validateSetEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetProjectS3Path(val *string) {
	if err := j.validateSetProjectS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectS3Path",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetStudioWebPortalAccess(val *string) {
	if err := j.validateSetStudioWebPortalAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortalAccess",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetDomainAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetDomainId() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetDomainRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetEnvironmentId() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironmentId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		s,
		"resetProjectId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetProjectS3Path() {
	_jsii_.InvokeVoid(
		s,
		"resetProjectS3Path",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ResetStudioWebPortalAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetStudioWebPortalAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

