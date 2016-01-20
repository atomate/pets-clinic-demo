package clinic

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:pass@/pets-clinic?charset=utf8")
	orm.RegisterModel(new(Pet))
	orm.RegisterModel(new(Visit))
	orm.RunCommand()
	orm.Debug = true;
}

func ReadPet(pet *Pet) *Pet{
	o := orm.NewOrm()
	o.Read(pet);
	o.LoadRelated(pet, "Visits")
	return pet;
}

func ReadVisit(visit *Visit) *Visit{
	err := orm.NewOrm().Read(visit)
	if(err != nil){
		fmt.Println(err);
	}
	return visit;
}

func NewPet(pet *Pet){
	id, err := orm.NewOrm().Insert(pet)
	if err == nil {
	    fmt.Println(id)
	}
}

func NewVisit(visit *Visit){
	id, err := orm.NewOrm().Insert(visit)
	if err == nil {
	    fmt.Println(id)
	}
}

func UpdatePet(pet *Pet){
	if num, err := orm.NewOrm().Update(pet); err == nil {
	 fmt.Println(num)
	}
}

func UpdateVisit(visit *Visit){
	if num, err := orm.NewOrm().Update(visit); err == nil {
	 fmt.Println(num)
	}
}

func DeletePet(pet *Pet){
    if num, err := orm.NewOrm().Delete(pet); err == nil {
   	 fmt.Println(num)
    }
}

func AllPets() []interface{} {
	var pets []Pet
	all, err := orm.NewOrm().QueryTable("pet").All(&pets)
	fmt.Printf("Returned Rows Num: %s, %s", all, err)
	result := make([]interface{}, len(pets))
	for i, s := range pets {
		r := s;
	    result[i] = &r
	}
	return result;
}

func AllVisits() []interface{} {
	var visits []Visit
	all, err := orm.NewOrm().QueryTable("visit").RelatedSel().All(&visits)
	fmt.Printf("Returned Rows Num: %s, %s", all, err)
	result := make([]interface{}, len(visits))
	for i, s := range visits {
		r := s;
	    result[i] = &r
	}
	return result;
}
