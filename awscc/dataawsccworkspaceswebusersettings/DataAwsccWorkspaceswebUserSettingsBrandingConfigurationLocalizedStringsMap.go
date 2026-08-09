// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccworkspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccworkspaceswebusersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(key *string) DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap
type jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap {
	_init_.Initialize()

	if err := validateNewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspaceswebUserSettings.DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap_Override(d DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspaceswebUserSettings.DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) Get(key *string) DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

