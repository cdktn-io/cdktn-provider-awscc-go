// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspacesconnectionalias

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspacesConnectionAliasTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkspacesConnectionAliasTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspacesConnectionAliasTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspacesConnectionAliasTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

