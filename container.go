package gollector

// Container is a ...
type Container struct {
	Template Metric   `json:"template,omitempty"`
	Metrics  []Metric `json:"metrics"`
}

// Validate returns an error if the container is invalid
func (c Container) Validate() error {
	if c.Metrics == nil {
		return fmt.Errorf("Missing metrics[] data")
	}
	if len(c.Metrics) <= 0 {
		return fmt.Errorf("Empty metrics[] data")
	}
	for i := 0; i < len(c.Metrics); i++ {
		if c.Metrics[i].Time == nil && c.Template.Time == nil {
			return fmt.Errorf("Missing timestamp in both metric and container")
		}
		err := c.Metrics[i].Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
