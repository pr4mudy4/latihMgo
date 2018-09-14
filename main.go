package main

import (
	"fmt"
	"latihMgo/config"
	"latihMgo/src/modules/profile/repository"
	"latihMgo/src/modules/profile/model"
	"time"
)

func main() {
	fmt.Println(" Go mongo db")
	db, err := config.GetMongoDB()

	if err != nil{
		fmt.Println(err)
	}
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	saveProfile(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository){
	var p model.Profile
	p.ID ="U1"
	p.FirstName = "setha"
	p.LastName = "sakti pramudya"
	p.Email = "pramudya.aviani@gmail.com"
	p.Password = "setha666"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Profile saved..")
	}
}
