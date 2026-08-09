// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsecurityagentagentspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsecurityagentagentspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference interface {
	cdktn.ComplexObject
	BitbucketCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference
	BitbucketRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference
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
	ConfluenceCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference
	ConfluenceDocument() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GitHubCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference
	GitHubRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference
	GitLabCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference
	GitLabRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference
	InternalValue() *DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResources
	SetInternalValue(val *DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResources)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference
type jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"bitbucketCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) BitbucketRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepositoryOutputReference
	_jsii_.Get(
		j,
		"bitbucketRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"confluenceCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ConfluenceDocument() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference
	_jsii_.Get(
		j,
		"confluenceDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"gitHubCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitHubRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitHubRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabCapabilities() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"gitLabCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GitLabRepository() DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference {
	var returns DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitLabRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InternalValue() *DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResources {
	var returns *DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSecurityagentAgentSpace.DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference_Override(d DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSecurityagentAgentSpace.DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetInternalValue(val *DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccSecurityagentAgentSpaceIntegratedResourcesProviderResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

