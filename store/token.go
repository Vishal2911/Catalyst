package store

import (
	
	"magantas/catalyst/model"
	"time"
)

func (store Store) SaveToken(tokenSchema model.TokenSchema) {

	store.Token[tokenSchema.Token] = tokenSchema

}

func (store Store) GetToken(token string) (model.TokenSchema, bool) {

	tokenschema, ok := store.Token[token]
	return tokenschema, ok

}

func (store Store) RevokeToken(token string) bool {

	
	tokenschema, ok := store.Token[token]
	if ok {
		tokenschema.Active = false
		store.Token[tokenschema.Token] = tokenschema
		
		return ok
	} else {
		return false
	}

}

func (store Store) SetTokenUsed(token string) bool {

	
	tokenschema, ok := store.Token[token]
	if ok {
		tokenschema.Used = true
		store.Token[tokenschema.Token] = tokenschema
		return ok
	} else {
		return false
	}

}

func (store Store) DayPassedSinceTokenCreated(token string) (int , bool ,bool ){

	
	tokenschema, ok := store.Token[token]
	if ok {
		currentday := time.Now()
		diff := currentday.Sub(tokenschema.CreatedAT)
		return int(diff.Hours()/24) , ok , tokenschema.Used
	} else {
		return 0 , false , false
	}

}
