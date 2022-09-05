package main

// Разбор задачи: https://www.youtube.com/watch?v=xGvQN_g-JCI

func main() {

}

type Person struct {
	friends map[*Person]bool
}

func (p *Person) knows(person *Person) bool {
	_, exists := p.friends[person]

	if exists {
		return true
	} else {
		return false
	}
}

func findCelebrity(crowd []*Person) *Person {
	l := 0
	r := len(crowd) - 1
	for l != r {
		if crowd[l].knows(crowd[r]) {
			l++
		} else {
			r--
		}
	}
	for i, val := range crowd {
		if i != l && (!crowd[i].knows(val) || val.knows(crowd[i])) {
			return nil
		}
	}
	return crowd[l]
}
