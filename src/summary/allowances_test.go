package summary

import (
	"linty/src/config"
	"testing"
)

func TestGetFailureAllowance(t *testing.T) {
	tests := []struct {
		name     string
		config   *config.Config
		expected Summary
	}{
		{
			name: "All allowances true",
			config: &config.Config{
				FailureAllowances: config.FailureAllowances{
					Readme:    true,
					License:   true,
					Gitignore: true,
					Workflow:  true,
					Tests:     true,
				},
			},
			expected: Summary{
				Readme:    true,
				License:   true,
				Gitignore: true,
				Workflow:  true,
				Tests:     true,
			},
		},
		{
			name: "All allowances false",
			config: &config.Config{
				FailureAllowances: config.FailureAllowances{
					Readme:    false,
					License:   false,
					Gitignore: false,
					Workflow:  false,
					Tests:     false,
				},
			},
			expected: Summary{
				Readme:    false,
				License:   false,
				Gitignore: false,
				Workflow:  false,
				Tests:     false,
			},
		},
		{
			name: "Mixed allowances",
			config: &config.Config{
				FailureAllowances: config.FailureAllowances{
					Readme:    true,
					License:   false,
					Gitignore: true,
					Workflow:  false,
					Tests:     true,
				},
			},
			expected: Summary{
				Readme:    true,
				License:   false,
				Gitignore: true,
				Workflow:  false,
				Tests:     true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := &Summary{}
			GetFailureAllowance(tt.config, actual)

			if actual.Readme != tt.expected.Readme ||
				actual.License != tt.expected.License ||
				actual.Gitignore != tt.expected.Gitignore ||
				actual.Workflow != tt.expected.Workflow ||
				actual.Tests != tt.expected.Tests {
				t.Errorf("GetFailureAllowance() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestIsFailureAllowed(t *testing.T) {
	summary := Summary{
		Readme:    true,
		License:   false,
		Gitignore: true,
		Workflow:  false,
		Tests:     true,
	}

	tests := []struct {
		name     string
		category string
		want     bool
	}{
		{
			name:     "Readme allowed",
			category: "readme",
			want:     true,
		},
		{
			name:     "License not allowed",
			category: "license",
			want:     false,
		},
		{
			name:     "Gitignore allowed",
			category: "gitignore",
			want:     true,
		},
		{
			name:     "Workflow not allowed",
			category: "workflow",
			want:     false,
		},
		{
			name:     "Tests allowed",
			category: "tests",
			want:     true,
		},
		{
			name:     "Invalid category",
			category: "invalid",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFailureAllowed(tt.category, summary); got != tt.want {
				t.Errorf("IsFailureAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
