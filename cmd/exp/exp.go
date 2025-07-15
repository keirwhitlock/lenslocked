package main

import (
	"html/template"
	"os"
)

type Sibling struct {
	Name string
	Age  int
	Sex  string
}

type User struct {
	Name           string
	Age            int
	FavouriteFoods []string
	Siblings       []Sibling
	//Bio  template.HTML
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Age:  43,
		FavouriteFoods: []string{
			"Curry",
			"Fish & Chips",
			"Eggs",
		},
		Siblings: []Sibling{
			Sibling{
				Name: "Bobby",
				Age:  43,
				Sex:  "male",
			},
			Sibling{
				Name: "Gemma",
				Age:  33,
				Sex:  "female",
			},
		},
		//Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	// is.Stdout - io.Writer
	// user - interface{}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
