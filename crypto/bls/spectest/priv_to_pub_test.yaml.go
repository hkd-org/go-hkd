// Code generated by yaml_to_go. DO NOT EDIT.
// source: priv_to_pub.yaml

package spectest

type PrivToPubTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Input  []byte `json:"input" ssz:"size=32"`
		Output []byte `json:"output" ssz:"size=48"`
	} `json:"test_cases"`
}
