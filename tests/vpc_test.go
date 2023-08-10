package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../infrastructure",
		NoColor: true,  // Disable colors in Terraform commands
	}

	defer terraform.Destroy(t, terraformOptions)

	// Deploy the infrastructure with Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Fetch the VPC ID and tags from Terraform outputs
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	vpcTags := terraform.OutputMap(t, terraformOptions, "vpc_tags")

	// Assert that the VPC ID is not empty, implying it has been created
	assert.NotEmpty(t, vpcID, "VPC ID should not be empty")

	// Assert that the VPC has a specific tag. For example, a tag "Name" with value "MyVPC".
	expectedVPCName := "MyVPC"
	if vpcName, exists := vpcTags["Name"]; exists {
	    assert.Equal(t, expectedVPCName, vpcName)
	} else {
	    assert.Fail(t, "Tag 'Name' not found in VPC tags")
	}

	// Additional assertions can be added based on other tags or outputs as required
}
