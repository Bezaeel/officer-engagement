package main

import (
	"fmt"
	"time"
)

type officer struct {
	name string
	numberAssigned int
	isTagged bool
	taggedAt time.Time
	updatedAt time.Time
	createdAt time.Time
}	

var officers = []officer{
	officer{
		name: "Talabi",
		numberAssigned: 20,
		isTagged: false,
		taggedAt: time.Now().AddDate(0,0,-3),
		updatedAt: time.Now().AddDate(0,0,-3),
		createdAt: time.Now().AddDate(0,0,-9),
	},
	officer{
		name: "Esmeralda",
		numberAssigned: 19,
		isTagged: false,
		taggedAt: time.Now().AddDate(0,0,-2),
		updatedAt: time.Now().AddDate(0,0,-1),
		createdAt: time.Now().AddDate(0,0,-9),
	},
	officer{
		name: "Tito",
		numberAssigned: 15,
		isTagged: false,
		taggedAt: time.Now().AddDate(0,0,-1),
		updatedAt: time.Now().AddDate(0,0,-1),
		createdAt: time.Now().AddDate(0,0,-9),
	},
	officer{
		name: "Ahmed",
		numberAssigned: 14,
		isTagged: true,
		taggedAt: time.Now().AddDate(0,0,-9),
		updatedAt: time.Now().AddDate(0,0,-1),
		createdAt: time.Now().AddDate(0,0,-9),
	},
	officer{
		name: "Yemi",
		numberAssigned: 10,
		isTagged: false,
		taggedAt: time.Now().AddDate(0,0,-10),
		updatedAt: time.Now().AddDate(0,0,-1),
		createdAt: time.Now().AddDate(0,0,-9),
	},
}

func rotationSchedule() ([]officer) {
	var offset int;
	var rotation []officer;

	for i := 0; i < len(officers); i++ {
		if(officers[i].isTagged){
			offset = i;
			break;
		}
	}

	for j := 0 ; j < len(officers); j++ {
		var pointer = (j+offset)%len(officers);
		rotation = append(rotation, officers[pointer]);
	}

	return rotation;
}


func tag() {
	var offset int

	for i := 0; i < len(officers); i++ {
		if(officers[i].isTagged){
			offset = i;
			break;
		}
	}

	for j := 0 ; j < len(officers); j++ {
		var pointer = (j+offset)%len(officers);
		officers[pointer].isTagged = false;
		officers[pointer].updatedAt = time.Now();
		officers[pointer + 1].isTagged = true;
		officers[pointer + 1].taggedAt = time.Now();
		officers[pointer + 1].updatedAt = time.Now();
		break;
	}
	fmt.Println(officers);
}

func main() {
	// fmt.Println("Ask Talabi")
	// fmt.Println(rotationSchedule());
	// fmt.Println("======================================================");
	tag();
}