package database

type Client struct {
	Name string
	Date string
}

type Goods struct {
	ID   string
	Name string
	Sort string
}

type Purchase struct {
	Name   string
	Amount string
}

// Other database models and related functions
