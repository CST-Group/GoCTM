package memory

import (
	"testing"
)

func TestMemoryObjectSetName(t *testing.T) {

	memoryObject := CreateMemoryObject("TEST_MO")

	if memoryObject.GetName() != "TEST_MO" {
		t.Errorf("The Name variable has diferent value.")
	}

}

func TestMemoryObjectSetID(t *testing.T) {

	memoryObject := CreateMemoryObject("TEST_MO")

	memoryObject.SetID(2)

	if memoryObject.GetID() != 2 {
		t.Errorf("The ID variable has diferent value.")
	}

}

func TestMemoryObjectSetEvaluation(t *testing.T) {

	memoryObject := CreateMemoryObject("TEST_MO")

	memoryObject.SetEvaluation(0.54)

	if memoryObject.GetEvaluation() != 0.54 {
		t.Errorf("The Evaluation variable has diferent value.")
	}

}
