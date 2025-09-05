package ds

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	t.Run("push single item to int stack", func(t *testing.T) {
		var stack Stack[int]
		stack.Push(42)
		
		if stack.IsEmpty() {
			t.Error("Expected stack to not be empty after push")
		}
		
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 42 {
			t.Errorf("Expected top to be 42, got %d", top)
		}
	})
	
	t.Run("push multiple items to string stack", func(t *testing.T) {
		var stack Stack[string]
		items := []string{"first", "second", "third"}
		
		for _, item := range items {
			stack.Push(item)
		}
		
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != "third" {
			t.Errorf("Expected top to be 'third', got %s", top)
		}
	})
}

func TestStackPop(t *testing.T) {
	t.Run("pop from stack with items", func(t *testing.T) {
		var stack Stack[int]
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		
		stack.Pop()
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 2 {
			t.Errorf("Expected top to be 2 after pop, got %d", top)
		}
	})
	
	t.Run("pop from empty stack", func(t *testing.T) {
		var stack Stack[int]
		stack.Pop() // Should not panic
		
		if !stack.IsEmpty() {
			t.Error("Expected stack to remain empty after pop on empty stack")
		}
	})
	
	t.Run("pop until empty", func(t *testing.T) {
		var stack Stack[string]
		stack.Push("a")
		stack.Push("b")
		
		stack.Pop()
		stack.Pop()
		
		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty after popping all items")
		}
	})
}

func TestStackTop(t *testing.T) {
	t.Run("top of empty stack", func(t *testing.T) {
		var stack Stack[int]
		
		_, err := stack.Top()
		if err == nil {
			t.Error("Expected error when calling Top on empty stack")
		}
		if err.Error() != "stack is empty" {
			t.Errorf("Expected 'stack is empty' error, got %v", err)
		}
	})
	
	t.Run("top of stack with one item", func(t *testing.T) {
		var stack Stack[rune]
		stack.Push('x')
		
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 'x' {
			t.Errorf("Expected top to be 'x', got %c", top)
		}
		
		// Verify item is still there after Top()
		top2, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error on second Top call, got %v", err)
		}
		if top2 != 'x' {
			t.Errorf("Expected top to still be 'x' after Top call, got %c", top2)
		}
	})
	
	t.Run("top returns last pushed item", func(t *testing.T) {
		var stack Stack[float64]
		values := []float64{1.1, 2.2, 3.3, 4.4}
		
		for _, v := range values {
			stack.Push(v)
			top, err := stack.Top()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if top != v {
				t.Errorf("Expected top to be %f, got %f", v, top)
			}
		}
	})
}

func TestStackIsEmpty(t *testing.T) {
	t.Run("new stack is empty", func(t *testing.T) {
		var stack Stack[int]
		
		if !stack.IsEmpty() {
			t.Error("Expected new stack to be empty")
		}
	})
	
	t.Run("stack with items is not empty", func(t *testing.T) {
		var stack Stack[string]
		stack.Push("test")
		
		if stack.IsEmpty() {
			t.Error("Expected stack with items to not be empty")
		}
	})
	
	t.Run("stack becomes empty after popping all items", func(t *testing.T) {
		var stack Stack[bool]
		stack.Push(true)
		stack.Push(false)
		
		if stack.IsEmpty() {
			t.Error("Expected stack with items to not be empty")
		}
		
		stack.Pop()
		stack.Pop()
		
		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty after popping all items")
		}
	})
}

func TestStackGenericTypes(t *testing.T) {
	t.Run("int stack", func(t *testing.T) {
		var stack Stack[int]
		stack.Push(100)
		stack.Push(-50)
		
		top, _ := stack.Top()
		if top != -50 {
			t.Errorf("Expected top to be -50, got %d", top)
		}
	})
	
	t.Run("string stack", func(t *testing.T) {
		var stack Stack[string]
		stack.Push("hello")
		stack.Push("world")
		
		top, _ := stack.Top()
		if top != "world" {
			t.Errorf("Expected top to be 'world', got %s", top)
		}
	})
	
	t.Run("rune stack", func(t *testing.T) {
		var stack Stack[rune]
		stack.Push('a')
		stack.Push('z')
		
		top, _ := stack.Top()
		if top != 'z' {
			t.Errorf("Expected top to be 'z', got %c", top)
		}
	})
	
	t.Run("float64 stack", func(t *testing.T) {
		var stack Stack[float64]
		stack.Push(3.14)
		stack.Push(2.71)
		
		top, _ := stack.Top()
		if top != 2.71 {
			t.Errorf("Expected top to be 2.71, got %f", top)
		}
	})
	
	t.Run("bool stack", func(t *testing.T) {
		var stack Stack[bool]
		stack.Push(true)
		stack.Push(false)
		
		top, _ := stack.Top()
		if top != false {
			t.Errorf("Expected top to be false, got %t", top)
		}
	})
}

func TestStackComplexOperations(t *testing.T) {
	t.Run("push and pop sequence", func(t *testing.T) {
		var stack Stack[int]
		
		// Push some items
		for i := 1; i <= 5; i++ {
			stack.Push(i)
		}
		
		// Pop some items
		stack.Pop()
		stack.Pop()
		
		// Check top
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 3 {
			t.Errorf("Expected top to be 3, got %d", top)
		}
		
		// Push more items
		stack.Push(10)
		stack.Push(20)
		
		// Check final top
		top, err = stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 20 {
			t.Errorf("Expected top to be 20, got %d", top)
		}
	})
	
	t.Run("stress test with many operations", func(t *testing.T) {
		var stack Stack[int]
		
		// Push 1000 items
		for i := 0; i < 1000; i++ {
			stack.Push(i)
		}
		
		// Verify top
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 999 {
			t.Errorf("Expected top to be 999, got %d", top)
		}
		
		// Pop 500 items
		for i := 0; i < 500; i++ {
			stack.Pop()
		}
		
		// Verify new top
		top, err = stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 499 {
			t.Errorf("Expected top to be 499, got %d", top)
		}
	})
}

func TestStackEdgeCases(t *testing.T) {
	t.Run("empty stack operations", func(t *testing.T) {
		var stack Stack[int]
		
		// Multiple pops on empty stack
		stack.Pop()
		stack.Pop()
		stack.Pop()
		
		if !stack.IsEmpty() {
			t.Error("Expected stack to remain empty")
		}
		
		// Top on empty stack after multiple pops
		_, err := stack.Top()
		if err == nil {
			t.Error("Expected error when calling Top on empty stack")
		}
	})
	
	t.Run("single item operations", func(t *testing.T) {
		var stack Stack[string]
		
		// Push one item
		stack.Push("single")
		
		// Check it's there
		if stack.IsEmpty() {
			t.Error("Expected stack to not be empty")
		}
		
		top, err := stack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != "single" {
			t.Errorf("Expected top to be 'single', got %s", top)
		}
		
		// Pop it
		stack.Pop()
		
		// Check it's empty
		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty after popping single item")
		}
		
		// Top should error
		_, err = stack.Top()
		if err == nil {
			t.Error("Expected error when calling Top on empty stack")
		}
	})
	
	t.Run("zero values", func(t *testing.T) {
		var intStack Stack[int]
		intStack.Push(0)
		
		top, err := intStack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if top != 0 {
			t.Errorf("Expected top to be 0, got %d", top)
		}
		
		var stringStack Stack[string]
		stringStack.Push("")
		
		topStr, err := stringStack.Top()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if topStr != "" {
			t.Errorf("Expected top to be empty string, got '%s'", topStr)
		}
	})
}

func TestStackLIFOBehavior(t *testing.T) {
	t.Run("verify LIFO order", func(t *testing.T) {
		var stack Stack[string]
		items := []string{"first", "second", "third", "fourth"}
		
		// Push all items
		for _, item := range items {
			stack.Push(item)
		}
		
		// Pop and verify reverse order
		for i := len(items) - 1; i >= 0; i-- {
			if stack.IsEmpty() {
				t.Fatalf("Expected stack to not be empty, iteration %d", i)
			}
			
			top, err := stack.Top()
			if err != nil {
				t.Fatalf("Expected no error at iteration %d, got %v", i, err)
			}
			
			expected := items[i]
			if top != expected {
				t.Errorf("Expected top to be '%s' at iteration %d, got '%s'", expected, i, top)
			}
			
			stack.Pop()
		}
		
		// Stack should be empty now
		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty after popping all items")
		}
	})
}