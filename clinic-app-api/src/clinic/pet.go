package clinic

type Pet struct {
	Id         int      `jsonapi:"primary,pets"`
	Name       string   `jsonapi:"attr,name"`
	OwnerName  string   `jsonapi:"attr,owner-name"`
	Visits     []*Visit `jsonapi:"relation, visit" orm:"reverse(many)" `
	OwnerPhone string   `jsonapi:"attr,owner-phone"`
}

func (p *Pet) TableName() string {
    return "pet"
}
