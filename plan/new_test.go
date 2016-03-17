package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	plan := New(&Config{
		Clients: []string{"alpha", "omega"},
		Axes: []Axis{{
			Name:   "behavior",
			Values: []string{"dance", "sing"},
		}},
	})
	assert.Equal(t,
		[]TestCase{
			{
				Plan:   plan,
				Client: "alpha",
				Arguments: Arguments{
					"behavior": "dance",
				},
			},
			{
				Plan:   plan,
				Client: "alpha",
				Arguments: Arguments{
					"behavior": "sing",
				},
			},
			{
				Plan:   plan,
				Client: "omega",
				Arguments: Arguments{
					"behavior": "dance",
				},
			},
			{
				Plan:   plan,
				Client: "omega",
				Arguments: Arguments{
					"behavior": "sing",
				},
			},
		},
		plan.TestCases,
	)
}
