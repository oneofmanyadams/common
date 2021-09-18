package award

type Supplier struct {
    Id string
    Awards []Award
}
func (s Supplier) GetId() string {return s.Id}
func (s Supplier) GetAwards() []Award {return s.Awards}
