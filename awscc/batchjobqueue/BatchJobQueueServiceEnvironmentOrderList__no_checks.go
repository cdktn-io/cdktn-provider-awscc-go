// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchjobqueue

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueServiceEnvironmentOrderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchJobQueueServiceEnvironmentOrderListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

