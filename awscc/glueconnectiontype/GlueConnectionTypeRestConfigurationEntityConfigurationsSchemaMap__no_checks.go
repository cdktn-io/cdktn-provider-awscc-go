// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueconnectiontype

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateGetParameters(key *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewGlueConnectionTypeRestConfigurationEntityConfigurationsSchemaMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

