// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServicesTimeout struct {
	// The amount of time in seconds a connection will stay active while idle.
	//
	// A value of ``0`` can be set to disable ``idleTimeout``.
	//  The ``idleTimeout`` default for ``HTTP``/``HTTP2``/``GRPC`` is 5 minutes.
	//  The ``idleTimeout`` default for ``TCP`` is 1 hour.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#idle_timeout_seconds EcsService#idle_timeout_seconds}
	IdleTimeoutSeconds *float64 `field:"optional" json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
	// The amount of time waiting for the upstream to respond with a complete response per request.
	//
	// A value of ``0`` can be set to disable ``perRequestTimeout``. ``perRequestTimeout`` can only be set if Service Connect ``appProtocol`` isn't ``TCP``. Only ``idleTimeout`` is allowed for ``TCP````appProtocol``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#per_request_timeout_seconds EcsService#per_request_timeout_seconds}
	PerRequestTimeoutSeconds *float64 `field:"optional" json:"perRequestTimeoutSeconds" yaml:"perRequestTimeoutSeconds"`
}

