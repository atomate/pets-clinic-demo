package clinic

type Visit struct {
	Id     int       `jsonapi:"primary,visits"`
	Pet    *Pet      `jsonapi:"relation,pet" orm:"column(pet_id);rel(fk)"`
    Result string    `jsonapi:"attr,result"`
}

func (v *Visit) TableName() string {
    return "visit"
}
