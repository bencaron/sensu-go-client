package sensu

import "fmt"

// GetAggregates Return the list of Aggregates
func (s *Sensu) GetAggregates(limit int, offset int) ([]interface{}, error) {
	return s.GetList("aggregates", limit, offset)
}

// GetAggregate Return Aggregate info
func (s *Sensu) GetAggregate(check string, age int) ([]interface{}, error) {
	fmt.Printf("FIXME GetAgregate Not handling age %d", age)
	return s.GetList(fmt.Sprintf("aggregate/%s", check), 0, 0)
}

// GetAggregateIssued Return Aggregate history
func (s *Sensu) GetAggregateIssued(check string, issued string, summarize bool, result bool) (map[string]interface{}, error) {
	// FIXME
	fmt.Printf("FIXME Aggregate Not handling summarize/result")
	return s.Get(fmt.Sprintf("aggregate/%s/%s", check, issued))
}

// DeleteAggregate Return the list of Aggregates
func (s *Sensu) DeleteAggregate(aggregate string) (map[string]interface{}, error) {
	return s.Delete(fmt.Sprintf("aggregate/%s", aggregate))
}
