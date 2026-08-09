// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eksaddon

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksAddonPodIdentityAssociationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksAddonPodIdentityAssociationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksAddonPodIdentityAssociationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksAddonPodIdentityAssociationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

