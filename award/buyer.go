package award

type Buyer struct {
    Id string
    Suppliers []string
    Awards []Award
}
func (s Buyer) GetId() string {return s.Id}
func (s Buyer) GetAwards() []Award {return s.Awards}
