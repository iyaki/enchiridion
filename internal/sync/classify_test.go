package sync

import (
	"testing"

	"github.com/iyaki/enchiridion/internal/model"
)

func TestClass(t *testing.T) {
	cases := []struct {
		name string
		tags []string
		want string
	}{
		{"pure tool", []string{"Tool", "English"}, dirTools},
		{"service counts as tool", []string{"Service", "Untried"}, dirTools},
		{"website counts as tool", []string{"Website", "CSS"}, dirTools},
		{"knowledge wins over tool", []string{"Article", "Tool", "Databases"}, dirKnowledge},
		{"note wins over service", []string{"Note", "Service"}, dirKnowledge},
		{"no categories defaults to knowledge", []string{"English", "Databases"}, dirKnowledge},
		{"no tags defaults to knowledge", nil, dirKnowledge},
	}
	for _, tc := range cases {
		if got := Class(model.PageMeta{Tags: tc.tags}); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, got, tc.want)
		}
	}
}
