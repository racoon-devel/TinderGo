package main

import (
	"fmt"

	"racoondev.tk/gitea/racoon/tindergo"
)

func main() {

	t := tindergo.New()

	err := t.Authenticate("EAAGm0PX4ZCpsBAEFvY9nSirc95Sa2tAvawZCwqngZAOsrFMhZC4TArPu7ZCc3fABZCscpJ0ZAdfa7uCdk5MBkQnl7jitptZBMyFkJnOH8cMEka8ZBpp11ZAhI6N5SEvAyRLrP9NOdmi3H5kQpIZALi5GheyGuLzIG57S18ZA26WEm9LsX97DrknK3TDf2COYGOWjNsYZD")
	if err != nil {
		panic(err)
	}

	fmt.Println("API token", t.APIToken())

	err = t.UpdateLocation(55.741676, 37.624928)
	if err != nil {
		panic(err)
	}

	pref := tindergo.SearchPreferences{AgeFilterMin: 20, AgeFilterMax: 29,
		DistanceFilter: 10, GenderFilter: 1}

	err = t.UpdateSearchPreferences(pref)
	if err != nil {
		panic(err)
	}
}
