# Learnings

This file is a rough draft. You can find the complete trail on my Twitter.
## Task 3

1. `var allResults []CheckResult`
→ This gives you a nil slice. In Go, a nil slice is special: it takes zero memory, but still behaves just like an empty slice when you use append

2. `allResults := make([]CheckResult, 0)`
→ This creates an empty slice with backing array allocated, but still of length 0. Functionally the same, but you pay the tiny overhead of allocation upfront, even if you never append.

3. Print using `verb`
    `%v` the value in a default format
    `%+v` the plus flag adds field names
    `%#v` a Go-syntax representation of the value