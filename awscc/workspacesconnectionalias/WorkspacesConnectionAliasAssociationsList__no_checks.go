// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspacesconnectionalias

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasAssociationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspacesConnectionAliasAssociationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

