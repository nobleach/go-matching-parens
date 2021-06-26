package main

import "testing"

func TestBalanced(t *testing.T)  {
  actual, _ := CheckBalancedBraces("({[[]{}([])]})")
  if actual != true {
    t.Errorf("Expected true but got %v", actual)
  }
}

func TestNotMatching(t *testing.T)  {
  actual, _ := CheckBalancedBraces("({[[{}([])]})")
  if actual != false {
    t.Errorf("Expected false but got %v", actual)
  }
}

func TestTooManyOpening(t *testing.T)  {
  actual, _ := CheckBalancedBraces("(({[[]{}([])]})")
  if actual != false {
    t.Errorf("Expected false but got %v", actual)
  }
}

func TestTooManyClosing(t *testing.T)  {
  actual, _ := CheckBalancedBraces("({[[]{}([])]}))")
  if actual != false {
    t.Errorf("Expected false but got %v", actual)
  }
}

func TestDanglingOpening(t *testing.T)  {
  actual, _ := CheckBalancedBraces("({[[]{}([])]})(")
  if actual != false {
    t.Errorf("Expected false but got %v", actual)
  }
}

func TestRealInput(t *testing.T)  {
  actual, _ := CheckBalancedBraces("(function(){console.log('foo';)})()")
  if actual != true {
    t.Errorf("Expected true but got %v", actual)
  }
}

