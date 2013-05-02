package descriptive_statistics

import (
	"fmt"
	"math"
	"testing"
)

var e = Enum{2, 6, 9, 3, 5, 1, 8, 3, 6, 9, 2}

func init() {
	fmt.Println("Testing: ", e)
}

func TestLen(t *testing.T) {
	if e.Len() != 11 {
		t.Error("Len Mismatch")
	}
}

func TestCount(t *testing.T) {
	if e.Count() != 11 {
		t.Error("Count Mismatch")
	}
}

func TestNumber(t *testing.T) {
	if e.Number() != 11.0 {
		t.Error("Number Mismatch")
	}
}

func TestMean(t *testing.T) {
	if math.Trunc(e.Mean()) != 4.0 {
		t.Error("Mean error")
	}
}

func TestMedian(t *testing.T) {
	if e.Median() != 5.0 {
		t.Error("Median error")
	}
}

func TestVariance(t *testing.T) {
	if math.Trunc(e.Variance()) != 7.0 {
		t.Error("Variance error")
	}
}

func TestStandardDeviation(t *testing.T) {
	if math.Trunc(e.StandardDeviation()) != 2.0 {
		t.Error("StandardDeviation mismatch")
	}
}

func TestPercentile(t *testing.T) {
	if e.Percentile(75.0) != 7.0 {
		t.Error("75% Percentile error")
	}

	if e.Percentile(90.0) != 9.0 {
		t.Error("90% Percentile error")
	}

	if e.Percentile(95.0) != 9.0 {
		t.Error("95% Percentile error")
	}
}
