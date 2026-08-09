// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package route53healthcheck

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53HealthCheckHealthCheckTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

