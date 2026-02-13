/*
Copyright © 2026 Adriano
*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Execute root command without subcommands
func TestRootCmdBasicExecution(t *testing.T) {
	rootCmd.SetArgs([]string{})

	err := rootCmd.Execute()
	assert.NoError(t, err, "root command should execute without error")
}

// Empty args handling
func TestRootCmdEmptyArgs(t *testing.T) {
	err := rootCmd.Execute()

	// Should execute without error with empty args
	assert.NoError(t, err, "should handle empty args")
}

// Verify root command metadata
func TestRootCmdMetadata(t *testing.T) {
	assert.Equal(t, "k8s-inspector", rootCmd.Use)
	assert.Equal(t, "K8s inspector is a CLI tool gather Kubernetes information", rootCmd.Short)
	assert.Contains(t, rootCmd.Long, "inspect Kubernetes")
}

func TestRootCmdStructure(t *testing.T) {
	assert.NotNil(t, rootCmd, "root command should exist")
	assert.NotEmpty(t, rootCmd.Use, "root command should have Use field")
	assert.NotEmpty(t, rootCmd.Short, "root command should have Short description")
}

func TestExecuteRunsWithoutError(t *testing.T) {
	rootCmd.SetArgs([]string{})
	// should return without exiting the test process
	Execute()
	assert.NotNil(t, rootCmd)
}
