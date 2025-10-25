package aws

import (
	"testing"
)

func TestDefaultSecurityGroupStrategy(t *testing.T) {
	vpcID := "vpc-12345"
	strategy := DefaultSecurityGroupStrategy(vpcID)

	if strategy.PreferExisting != true {
		t.Error("Expected PreferExisting to be true")
	}

	if strategy.DefaultName != "lens-jupyter" {
		t.Errorf("Expected DefaultName 'lens-jupyter', got: %s", strategy.DefaultName)
	}

	if strategy.VpcID != vpcID {
		t.Errorf("Expected VpcID %s, got: %s", vpcID, strategy.VpcID)
	}

	if strategy.ForceCreate != false {
		t.Error("Expected ForceCreate to be false")
	}
}

func TestSecurityGroupStrategy_GetDefaultSecurityGroupName(t *testing.T) {
	tests := []struct {
		name         string
		strategy     SecurityGroupStrategy
		expectedName string
	}{
		{
			name: "default name",
			strategy: SecurityGroupStrategy{
				DefaultName: "lens-jupyter",
			},
			expectedName: "lens-jupyter",
		},
		{
			name: "user specified",
			strategy: SecurityGroupStrategy{
				DefaultName:   "lens-jupyter",
				UserSpecified: "my-custom-sg",
			},
			expectedName: "my-custom-sg",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.strategy.GetDefaultSecurityGroupName()
			if result != test.expectedName {
				t.Errorf("Expected %s, got: %s", test.expectedName, result)
			}
		})
	}
}

func TestIsAwsJupyterSecurityGroup(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"lens-jupyter", true},
		{"lens-jupyter-session-manager", false},
		{"my-custom-sg", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAwsJupyterSecurityGroup(test.name)
			if result != test.expected {
				t.Errorf("IsAwsJupyterSecurityGroup(%s) = %v, expected %v", test.name, result, test.expected)
			}
		})
	}
}

func TestSecurityGroupInfo_Structure(t *testing.T) {
	info := &SecurityGroupInfo{
		ID:          "sg-12345",
		Name:        "lens-jupyter",
		Description: "Test security group",
		VpcID:       "vpc-12345",
		CreatedBy:   "lens-jupyter",
	}

	if info.ID != "sg-12345" {
		t.Errorf("Expected ID sg-12345, got: %s", info.ID)
	}

	if info.Name != "lens-jupyter" {
		t.Errorf("Expected Name lens-jupyter, got: %s", info.Name)
	}

	if info.VpcID != "vpc-12345" {
		t.Errorf("Expected VpcID vpc-12345, got: %s", info.VpcID)
	}

	if info.CreatedBy != "lens-jupyter" {
		t.Errorf("Expected CreatedBy lens-jupyter, got: %s", info.CreatedBy)
	}
}

func TestSecurityGroupStrategy_MultipleVpcs(t *testing.T) {
	vpcs := []string{"vpc-111", "vpc-222", "vpc-333"}

	for _, vpcID := range vpcs {
		t.Run(vpcID, func(t *testing.T) {
			strategy := DefaultSecurityGroupStrategy(vpcID)
			if strategy.VpcID != vpcID {
				t.Errorf("Expected VpcID %s, got: %s", vpcID, strategy.VpcID)
			}
		})
	}
}

func TestSecurityGroupStrategy_ForceCreate(t *testing.T) {
	strategy := SecurityGroupStrategy{
		PreferExisting: false,
		DefaultName:    "test-sg",
		ForceCreate:    true,
		VpcID:          "vpc-12345",
	}

	if strategy.PreferExisting {
		t.Error("Expected PreferExisting to be false")
	}

	if !strategy.ForceCreate {
		t.Error("Expected ForceCreate to be true")
	}
}

func TestPortConstants(t *testing.T) {
	// Test that port constants are defined correctly
	if portSSH != 22 {
		t.Errorf("Expected portSSH to be 22, got: %d", portSSH)
	}

	if portJupyter != 8888 {
		t.Errorf("Expected portJupyter to be 8888, got: %d", portJupyter)
	}
}

func TestSecurityGroupConstants(t *testing.T) {
	if createdByUser != "user" {
		t.Errorf("Expected createdByUser to be 'user', got: %s", createdByUser)
	}

	if createdByAwsJupyter != "lens-jupyter" {
		t.Errorf("Expected createdByAwsJupyter to be 'lens-jupyter', got: %s", createdByAwsJupyter)
	}

	if defaultSGName != "lens-jupyter" {
		t.Errorf("Expected defaultSGName to be 'lens-jupyter', got: %s", defaultSGName)
	}
}
