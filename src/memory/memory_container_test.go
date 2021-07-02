package memory

import (
	"testing"
)

func TestMemoryContainerSetName(t *testing.T) {

	memoryObject := CreateMemoryObject("TEST_MO")

	if memoryObject.GetName() != "TEST_MO" {
		t.Errorf("The Name variable has diferent value.")
	}

}

func TestMemoryContainerSetID(t *testing.T) {

	memoryContainer := CreateMemoryContainer("TEST_MO")

	memoryContainer.SetID(2)

	if memoryContainer.GetID() != 2 {
		t.Errorf("The ID variable has diferent value.")
	}

}

func TestMemoryContainerSetI(t *testing.T) {

	memoryContainer := CreateMemoryContainer("TEST_MO")

	memoryContainer.SetI("TEST_I")

	if len(memoryContainer.Memories) == 0 {
		t.Errorf("The internal memory object was not created.")
	}

	if memoryContainer.Memories[0].GetI() != "TEST_I" {
		t.Errorf("The internal memory object value was not assigned.")
	}

}

func TestMemoryContainerSetIEvaluation(t *testing.T) {

	memoryContainer := CreateMemoryContainer("TEST_MO")

	memoryContainer.SetIEvaluation("TEST_I", 0.4)

	if len(memoryContainer.Memories) == 0 {
		t.Errorf("The internal memory object was not created.")
	}

	if memoryContainer.Memories[0].GetI() != "TEST_I" {
		t.Errorf("The internal memory object value was not assigned.")
	}

	if memoryContainer.Memories[0].GetEvaluation() != 0.4 {
		t.Errorf("The internal memory object evaluation was not assigned.")
	}

}
