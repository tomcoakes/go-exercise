package main

import "testing"

var listFriendsTests = []struct {
	p1, p2 person
	out    string
}{
	{person{name: "James"}, person{name: "Alex"}, "James's friends are Alex"},
	{person{name: "Peter"}, person{name: "Phillip"}, "Peter's friends are Phillip"},
}

func TestFriendMaker(t *testing.T) {
	for _, test := range listFriendsTests {
		test.p1, test.p2 = makeFriends(test.p1, test.p2)
		if listFriends(test.p1) != test.out {
			t.Errorf("Expected %s but received %s", listFriends(test.p1), test.out)
		}
	}
}
