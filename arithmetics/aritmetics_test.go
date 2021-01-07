package arithmetics

import "testing"

func TestSumSqrtLoop(t *testing.T) {
	const sumSqrtLoop = 1
	const sumSqrtValue = 4
	const sumSqrtLoopValueExpected = 4.0

	sumSqrtLoopResult := SumSqrtLoop(sumSqrtLoop, sumSqrtValue)

	if sumSqrtLoopResult != sumSqrtLoopValueExpected {
		t.Errorf("Expected sumSqrtValue is %f, but result is %f",
			sumSqrtLoopValueExpected,
			sumSqrtLoopResult)
	}
}
