package problems

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	var out []int
	for cur := l; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both nil", nil, nil, nil},
		{"left nil", nil, []int{0, 1}, []int{0, 1}},
		{"right nil", []int{1, 2}, nil, []int{1, 2}},
		{"interleaved", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"duplicates", []int{1, 1, 1}, []int{1, 1}, []int{1, 1, 1, 1, 1}},
		{"negatives", []int{-3, -1, 2}, []int{-2, 0, 3}, []int{-3, -2, -1, 0, 2, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeTwoLists(sliceToList(tc.a), sliceToList(tc.b))
			if !reflect.DeepEqual(listToSlice(got), tc.want) {
				t.Fatalf("MergeTwoLists(%v, %v) = %v, want %v", tc.a, tc.b, listToSlice(got), tc.want)
			}
		})
	}
}
