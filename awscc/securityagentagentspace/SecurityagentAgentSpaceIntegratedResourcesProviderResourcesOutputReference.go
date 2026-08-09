// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securityagentagentspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference interface {
	cdktn.ComplexObject
	BitbucketCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference
	BitbucketCapabilitiesInput() interface{}
	BitbucketRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference
	BitbucketRepositoryInput() interface{}
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
	ConfluenceCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference
	ConfluenceCapabilitiesInput() interface{}
	ConfluenceDocument() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference
	ConfluenceDocumentInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GitHubCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference
	GitHubCapabilitiesInput() interface{}
	GitHubRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference
	GitHubRepositoryInput() interface{}
	GitLabCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference
	GitLabCapabilitiesInput() interface{}
	GitLabRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference
	GitLabRepositoryInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutBitbucketCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilities)
	PutBitbucketRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepository)
	PutConfluenceCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilities)
	PutConfluenceDocument(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocument)
	PutGitHubCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilities)
	PutGitHubRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepository)
	PutGitLabCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilities)
	PutGitLabRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepository)
	ResetBitbucketCapabilities()
	ResetBitbucketRepository()
	ResetConfluenceCapabilities()
	ResetConfluenceDocument()
	ResetGitHubCapabilities()
	ResetGitHubRepository()
	ResetGitLabCapabilities()
	ResetGitLabRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference
type jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"bitbucketCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitbucketCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference
	_jsii_.Get(
		j,
		"bitbucketRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitbucketRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"confluenceCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceDocument() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference
	_jsii_.Get(
		j,
		"confluenceDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceDocumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"gitHubCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitHubCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitHubRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitHubRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabCapabilities() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"gitLabCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitLabCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabRepository() SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference {
	var returns SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitLabRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitLabRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference_Override(s SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutBitbucketCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilities) {
	if err := s.validatePutBitbucketCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBitbucketCapabilities",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutBitbucketRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepository) {
	if err := s.validatePutBitbucketRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBitbucketRepository",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutConfluenceCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilities) {
	if err := s.validatePutConfluenceCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfluenceCapabilities",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutConfluenceDocument(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocument) {
	if err := s.validatePutConfluenceDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfluenceDocument",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutGitHubCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilities) {
	if err := s.validatePutGitHubCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitHubCapabilities",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutGitHubRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepository) {
	if err := s.validatePutGitHubRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitHubRepository",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutGitLabCapabilities(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilities) {
	if err := s.validatePutGitLabCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitLabCapabilities",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) PutGitLabRepository(value *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepository) {
	if err := s.validatePutGitLabRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitLabRepository",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetBitbucketCapabilities() {
	_jsii_.InvokeVoid(
		s,
		"resetBitbucketCapabilities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetBitbucketRepository() {
	_jsii_.InvokeVoid(
		s,
		"resetBitbucketRepository",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetConfluenceCapabilities() {
	_jsii_.InvokeVoid(
		s,
		"resetConfluenceCapabilities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetConfluenceDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetConfluenceDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetGitHubCapabilities() {
	_jsii_.InvokeVoid(
		s,
		"resetGitHubCapabilities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetGitHubRepository() {
	_jsii_.InvokeVoid(
		s,
		"resetGitHubRepository",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetGitLabCapabilities() {
	_jsii_.InvokeVoid(
		s,
		"resetGitLabCapabilities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ResetGitLabRepository() {
	_jsii_.InvokeVoid(
		s,
		"resetGitLabRepository",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

