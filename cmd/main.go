package main

import "fmt"

type user struct {
	name  string
	count int
}

func addTo(u *user) { u.count++ }

func main() {
	users := []user{{"alice", 0}, {"bob", 0}}

	alice := &users[0] // risky because the slice may get realoccatted

	amy := user{"amy", 1}

	users = append(users, amy)

	addTo(alice)       // alice is likely a stale pointer
	fmt.Println(users) // so alice's count will be 0
}
