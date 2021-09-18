package award

type Planner struct {
    Id string
    ProductLines []string
    Awards []Award
}
func (s Planner) GetId() string {return s.Id}
func (s Planner) GetAwards() []Award {return s.Awards}
