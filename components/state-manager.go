package components

import "fmt"

const (
	MainOperationA string = "MainOperationA"
	MainOperationB string = "MainOperationB"
	MainOperationC string = "MainOperationC"
	MainOperationD string = "MainOperationD"

	SubOperationA string = "SubOperationA"
	SubOperationB string = "SubOperationB"
	SubOperationC string = "SubOperationC"
	SubOperationD string = "SubOperationD"
	SubOperationE string = "SubOperationE"
)

var (
	mainOpList = []string{MainOperationA, MainOperationB, MainOperationC, MainOperationD}
	subOpList = []string{SubOperationA, SubOperationB, SubOperationC, SubOperationD, SubOperationE}
	DefaultManager = NewProcedureManager()
)

type  procedure struct {
	initial *StateTemplate
	states map[string]*StateTemplate
}

func newProcedure(op string) *procedure  {
	s := NewStateTemplate(op, -1)
	sts := make(map[string]*StateTemplate)
	sts[op] = s
	return &procedure{
		initial: s,
		states: sts,
	}

}

type ProcedureManager struct {
	procedures map[string]*procedure
}


func NewProcedureManager() *ProcedureManager {
	return &ProcedureManager{
		procedures: make(map[string]*procedure),
	}
}

func (m *ProcedureManager)CreateNewProcedure(procedureName, initialOp string) *ProcedureManager {
	if _, ok := m.procedures[procedureName]; ok {
		fmt.Println(procedureName, "already exist")
		return m
	}
	if !isvalidMainOp(initialOp) {
		fmt.Println(initialOp, "is not a valid operation")
		return m
	}
	proc := newProcedure(initialOp)
	m.procedures[procedureName] = proc
	return m
}

func (m *ProcedureManager)ConfigureMainOps(procedureName string, ops ...string) *ProcedureManager {
	// check parameters before doing anything
	p, ok := m.procedures[procedureName]
	if !ok {
		fmt.Println(procedureName, " :does not exist! ")
		return m
	}
	if len(ops) < 1 {
		fmt.Println("no operation contained in the transaction")
		return m
	}
	for _, o := range ops {
		if !isvalidMainOp(o){
			fmt.Println(o, ": invalid main operation, nothing configured")
			return m
		}
	}

	initial := p.initial
	curOp := ops[0]
	CurState := NewStateTemplate(curOp, -1)
	p.states[curOp] = CurState
	initial.SetNext(CurState)
	for i := 1; i < len(ops); i++ {
		nextOp := ops[i]
		nextState := NewStateTemplate(nextOp, -1)
		CurState.SetNext(nextState)
		p.states[nextOp] = nextState
		CurState = nextState
	}
	return m
}

func (m *ProcedureManager)ConfigureSubs(procedureName, mainOpName string, subs ...string) *ProcedureManager {
	// check parameters before doing anything
	for _, op := range subs {
		if !isValidSubOp(op){
			fmt.Println(op,": is not a valid sub operation")
			return m
		}
	}

	p, ok := m.procedures[procedureName]
	if !ok {
		fmt.Println(procedureName, " :does not exist! ")
		return m
	}

	mSt, ok := p.states[mainOpName]
	if !ok {
		fmt.Println(mainOpName,": does not exist in specified procedure")
		return m
	}
	for _, op := range subs {
		st := NewStateTemplate(op, -1)
		mSt.AddSubs(st)
		p.states[op] = st
	}
	return m
}

func (m *ProcedureManager)SetThreshHold(procedureName, opName string, threshHold int) *ProcedureManager {
	proc, ok := m.procedures[procedureName]
	if !ok {
		fmt.Println(procedureName, ": is not valid procedure!")
		return m
	}
	st, ok := proc.states[opName]
	if !ok {
		fmt.Println(opName, ": is not valid operation!")
		return m
	}
	st.SetThreshHold(threshHold)
	return m
}


func isvalidMainOp(op string) bool {
	if len(mainOpList) == 0 {
		return false
	}
	for _, m := range mainOpList{
		if m == op{
			return true
		}
	}
	return false
}

func isValidSubOp(op string) bool {
	if len(subOpList) == 0 {
		return false
	}
	for _, s := range subOpList {
		if s == op{
			return true
		}
	}
	return false
}

