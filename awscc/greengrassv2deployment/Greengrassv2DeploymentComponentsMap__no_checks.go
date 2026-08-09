// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package greengrassv2deployment

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateGetParameters(key *string) error {
	return nil
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewGreengrassv2DeploymentComponentsMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

