package plan

// New creates a Plan given a Config
func New(config *Config) *Plan {
	plan := &Plan{
		Config: config,
	}
	plan.TestCases = buildTestCases(plan)
	return plan
}

func buildTestCases(plan *Plan) []TestCase {
	var testCases []TestCase
	for _, client := range plan.Config.Clients {
		combos := recurseCombinations(plan, client, plan.Config.Axes, make(map[string]string))
		testCases = append(testCases, combos...)
	}
	return testCases
}

func recurseCombinations(plan *Plan, client string, axes []Axis, args Arguments) []TestCase {
	if len(axes) == 0 {
		return []TestCase{{
			Plan:      plan,
			Client:    client,
			Arguments: copyArgs(args),
		}}
	}
	var testCases []TestCase
	axis := axes[0]
	for _, p := range axis.Values {
		args[axis.Name] = p
		testCases = append(testCases, recurseCombinations(plan, client, axes[1:], args)...)
	}
	return testCases
}

func copyArgs(args Arguments) Arguments {
	copy := make(map[string]string)
	for k, v := range args {
		copy[k] = v
	}
	return copy
}
