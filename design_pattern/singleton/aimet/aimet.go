package aimet

var aimetInstance *aimet = &aimet{}

type aimet struct {
	EmployeeCount int
}

func (a *aimet) Register() {
	a.EmployeeCount = a.EmployeeCount + 1
}

func GetInstance() *aimet {
	if aimetInstance == nil {
		aimetInstance = &aimet{}
	}
	return aimetInstance
}
