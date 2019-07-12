package main

const (
	recentIndex    = `RecentDay`
)

//key is uid
type RecentDay struct {
	//Uid   string `json:"Uid"`
    mostRecentDay int `json:"mostRecentDay"`
}
