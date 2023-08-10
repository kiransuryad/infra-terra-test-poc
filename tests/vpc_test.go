package test

import (
	"testing"
	"strings"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../infrastructure",
		
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	// Deploy the infrastructure with Terraform
	initAndApply := terraform.InitAndApply(t, terraformOptions)

	// Asserts to make sure the resources are created successfully
	assert.True(t, initAndApply)

	// You can add more assertions based on outputs or any other logic
}
