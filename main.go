package main

import (
	"errors"
	"fmt"
)

func CheckBalancedBraces(input string) (bool, error) {
  var symbolStack []rune
  var symbolMap = map[rune]rune{
    '(': ')',
    '{': '}',
    '[': ']',
  }

  lookForSet := map[rune]bool{
    ')': true,
    '}': true,
    ']': true,
  }

  // Iterate input
  for _, c := range input {
    // If symbol matches key in symbolMap, push
    // value onto symbolStack
    if val, ok := symbolMap[c]; ok {
      symbolStack  = append(symbolStack, val)
      continue
    }

    // If symbol matches symbolStack->peek, pop stack
    if len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] == c {
      symbolStack = symbolStack[:len(symbolStack)-1]
      continue
    }

    // If symbol does not match symbolStack->peek
    // see if it is in lookForSet.
    if lookForSet[c] {
      // if it is, return error
      return false, errors.New("Brackets do not match")
    }
  }

  return len(symbolStack) == 0, nil
}

func main() {
  matching := "({[[]{}([])]})"
  notMatching := "({[[{}([])]})"
  tooManyOpening := "(({[[]{}([])]})"
  tooManyClosing := "({[[]{}([])]}))"
  fmt.Println(CheckBalancedBraces(matching))
  fmt.Println(CheckBalancedBraces(notMatching))
  fmt.Println(CheckBalancedBraces(tooManyOpening))
  fmt.Println(CheckBalancedBraces(tooManyClosing))
}
