package main

import "testing"

var listFriendsTests = []struct {
	p1, p2, p3 person
	out        string
}{
	{p1: person{name: "James"}, p2: person{name: "Alex"}, out: "James's friends are Alex"},
	{p1: person{name: "Peter"}, p2: person{name: "Phillip"}, out: "Peter's friends are Phillip"},
	{p1: person{name: "Donald"}, p2: person{name: "Hillary"}, p3: person{name: "Melania"}, out: "Donald's friends are Hillary, Melania"},
}

func TestFriendMaker(t *testing.T) {
	for _, test := range listFriendsTests {
		test.p1, test.p2 = makeFriends(test.p1, test.p2)

		// for testing the example in the test table with a third person
		if test.p3.name != "" {
			test.p1, test.p3 = makeFriends(test.p1, test.p3)
		}

		if listFriends(test.p1) != test.out {
			t.Errorf("Expected %s but received %s", listFriends(test.p1), test.out)
		}
	}
}
