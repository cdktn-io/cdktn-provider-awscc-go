// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdseventsubscription

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsEventSubscriptionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsEventSubscriptionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsEventSubscriptionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsEventSubscriptionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsEventSubscriptionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

