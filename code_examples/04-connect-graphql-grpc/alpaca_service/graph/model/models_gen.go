// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Alpaca struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Mutation struct {
}

type NewAlpaca struct {
	Name             string `json:"name"`
	FaceSquishRating int    `json:"faceSquishRating"`
}

type Query struct {
}