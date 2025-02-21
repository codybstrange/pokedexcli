package main
import "testing"

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input string
    expected []string
  }{
    {
      input: "  hello  world      ",
      expected: []string{"hello", "world"},
    },
    {
      input: "charmander   Bulbasaur PIKACHU",
      expected: []string{"charmander", "bulbasaur", "pikachu"},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)
    if len(actual) != len(c.expected) {
      t.Errorf("Incorrect output word counts. Actual: %v. Expected: %v", len(actual), len(c.expected))
    }
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]
      if word != expectedWord {
        t.Errorf("Incorrect output word. Actual: %s. Expected: %s", word, expectedWord)
      }
    }
  }
}
