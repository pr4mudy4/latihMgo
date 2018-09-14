package repository

import "latihMgo/src/modules/profile/model"

//ProfileRepository Interface
type ProfileRepository interface{
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete (string) error
	FindByID (string) (*model.Profile, error)
	FindAll()(model.Profiles, error)
}