package main

type Quest struct {
	ID       int    `json:"ID"`
	Question string `json:"question"`
	Solution string `json:"solution"`
}
