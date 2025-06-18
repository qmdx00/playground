package main

import (
	"fmt"
	"testing"
)

func testSliceLenAndCapByAppend(_ *testing.T, length, initCap int) {
	nums := make([]int, 0, initCap)

	oldCap := cap(nums)
	capSlice := []int{oldCap}

	for range length {
		nums = append(nums, 1)

		newCap := cap(nums)
		if newCap != oldCap {
			capSlice = append(capSlice, newCap)
			oldCap = newCap
		}
	}

	fmt.Println("Capacity:", capSlice)
}

func TestSliceLenAndCapByAppend_10000_0(t *testing.T)  { testSliceLenAndCapByAppend(t, 10000, 0) }
func TestSliceLenAndCapByAppend_10000_1(t *testing.T)  { testSliceLenAndCapByAppend(t, 10000, 1) }
func TestSliceLenAndCapByAppend_10000_3(t *testing.T)  { testSliceLenAndCapByAppend(t, 10000, 3) }
func TestSliceLenAndCapByAppend_10000_5(t *testing.T)  { testSliceLenAndCapByAppend(t, 10000, 5) }
func TestSliceLenAndCapByAppend_10000_7(t *testing.T)  { testSliceLenAndCapByAppend(t, 10000, 7) }
func TestSliceLenAndCapByAppend_10000_11(t *testing.T) { testSliceLenAndCapByAppend(t, 10000, 11) }
