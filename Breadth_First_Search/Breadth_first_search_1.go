package Breadth_First_Search

//690. 员工的重要性
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	if len(employees) == 0 {
		return 0
	}
	Importances := 0
	index := make(map[int]*Employee,len(employees))
	for i := 0 ; i < len(employees) ; i++{
		index[employees[i].Id] = employees[i]
	}
	queue := make([]*Employee,0)
	queue = append(queue, index[id])
	for len(queue) > 0 {
		a := queue[0]
		queue = queue[1:]
		Importances += a.Importance
		for i := 0 ; i < len(a.Subordinates) ; i++ {
			q := a.Subordinates[i]
			queue = append(queue, index[q])
		}
	}
	return Importances
}