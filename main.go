package main

import (
	"fmt"
	"latihMgo/config"
	"latihMgo/src/modules/profile/model"
	"latihMgo/src/modules/profile/repository"
	"time"
)

func main() {
	fmt.Println(" Go mongo db")
	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	//saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)
	//getProfile("U1", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U2"
	p.FirstName = "agung"
	p.LastName = "permana"
	p.Email = "agung.permana@iconpln.co.id"
	p.Password = "germopink"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "aviani"
	p.LastName = "susanti"
	p.Email = "aviani.pramudya@gmail.com"
	p.Password = "balabala"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile update..")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {

	err := profileRepository.Delete("U1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile delete berhasil dilakukan..")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, profile := range profiles {
		fmt.Println("******************************")
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}
