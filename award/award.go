package award

type Award struct {
    Project string
    ProductLine string
}
func (s Award) PPL() string {return s.Project+s.ProductLine}
