package components

import "testing"
import . "github.com/smartystreets/goconvey/convey"

var procedureName = "testA"

func TestProcedureManager_ConfigureMainOps(t *testing.T) {
	Convey("test normal functions", t, func() {
		procM := NewProcedureManager()
		Convey("test main configure", func() {
			procM.CreateNewProcedure(procedureName, MainOperationA)
			states := procM.procedures[procedureName]
			So(states.initial.Name(), ShouldEqual, MainOperationA)

			procM.ConfigureMainOps(procedureName, MainOperationB, MainOperationC, MainOperationD)
			next := states.initial.Next()
			So(next.Name(), ShouldEqual, MainOperationB)
			next = next.Next()
			So(next.Name(), ShouldEqual, MainOperationC)
			next = next.Next()
			So(next.Name(), ShouldEqual, MainOperationD)
			next = next.Next()
			So(next, ShouldEqual, nil)

			procM.ConfigureSubs(procedureName, MainOperationA, SubOperationA, SubOperationB)
			procM.ConfigureSubs(procedureName, MainOperationC, SubOperationD, SubOperationE)

			ma := states.states[MainOperationA]
			maSubs := ma.Subs()
			So(len(maSubs), ShouldEqual, 2)

			mc := states.states[MainOperationC]
			mcSubs := mc.Subs()
			So(len(mcSubs), ShouldEqual, 2)

		})
	})
}
