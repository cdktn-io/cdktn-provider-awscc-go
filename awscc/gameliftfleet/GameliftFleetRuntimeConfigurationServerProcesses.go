// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetRuntimeConfigurationServerProcesses struct {
	// The number of server processes that use this configuration to run concurrently on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_fleet#concurrent_executions GameliftFleet#concurrent_executions}
	ConcurrentExecutions *float64 `field:"optional" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function.
	//
	// Game builds and Realtime scripts are installed on instances at the root:
	//
	// Windows (for custom game builds only): C:\game. Example: "C:\game\MyGame\server.exe"
	//
	// Linux: /local/game. Examples: "/local/game/MyGame/server.exe" or "/local/game/MyRealtimeScript.js"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_fleet#launch_path GameliftFleet#launch_path}
	LaunchPath *string `field:"optional" json:"launchPath" yaml:"launchPath"`
	// An optional list of parameters to pass to the server executable or Realtime script on launch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_fleet#parameters GameliftFleet#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

