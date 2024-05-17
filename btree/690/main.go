package main

func main() {

}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mp = make(map[int]*Employee)
	for _, e := range employees {
		mp[e.Id] = e
	}

	return totalSum(mp[id])
}

var mp map[int]*Employee

func totalSum(emp *Employee) int {
	if emp == nil {
		return 0
	}

	total := emp.Importance
	for _, id := range emp.Subordinates {
		total += totalSum(mp[id])
	}

	return total
}
