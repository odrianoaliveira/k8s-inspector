/*
Copyright © 2026 Adriano
*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Execute pods command happy path
func TestPodsCommandExecution(t *testing.T) {
	podsCmd.SetArgs([]string{})

	err := podsCmd.Execute()
	assert.NoError(t, err, "pods command should execute without error")
}

// Verify pods command metadata
func TestPodsCommandMetadata(t *testing.T) {
	assert.Equal(t, "pods", podsCmd.Use)
	assert.Equal(t, "List pods in the cluster", podsCmd.Short)
	assert.NotEmpty(t, podsCmd.Long)
	assert.Contains(t, podsCmd.Long, "Kubernetes cluster")
}

// Pods command is registered under root
func TestPodsCommandRegistration(t *testing.T) {
	var podsFound bool
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "pods" {
			podsFound = true
			break
		}
	}
	assert.True(t, podsFound, "pods command should be registered as subcommand of root")
}

// Root command can execute pods subcommand
func TestRootExecutePodsSubcommand(t *testing.T) {
	rootCmd.SetArgs([]string{"pods"})

	err := rootCmd.Execute()
	assert.NoError(t, err, "root should execute pods subcommand without error")
}

// Pods command Run function is callable
func TestPodsCommandRunFunctionCallable(t *testing.T) {
	assert.NotPanics(t, func() {
		podsCmd.Run(podsCmd, []string{})
	}, "Run function should be callable without panic")
}

// Verify pod command's command structure
func TestPodsCommandStructure(t *testing.T) {
	assert.NotNil(t, podsCmd, "pods command should exist")
	assert.NotEmpty(t, podsCmd.Use, "pods command should have Use field")
	assert.NotEmpty(t, podsCmd.Short, "pods command should have Short description")
	assert.NotNil(t, podsCmd.Run, "pods command should have Run function")
}

// Verify root can list pods command in available commands
func TestPodsInRootAvailableCommands(t *testing.T) {
	commands := rootCmd.Commands()
	require.NotEmpty(t, commands, "root should have subcommands")

	commandNames := make([]string, 0)
	for _, cmd := range commands {
		commandNames = append(commandNames, cmd.Use)
	}

	assert.Contains(t, commandNames, "pods", "pods should be in available commands")
}
