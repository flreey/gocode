package main

import (
	"fmt"
	"strings"
)

//-------------------------------------------------------------------------
// formatter interfaces
//-------------------------------------------------------------------------

type formatter interface {
	write_candidates(candidates []candidate, num int)
}

//-------------------------------------------------------------------------
// nice_formatter (just for testing, simple textual output)
//-------------------------------------------------------------------------

type nice_formatter struct{}

func (*nice_formatter) write_candidates(candidates []candidate, num int) {
	if candidates == nil {
		fmt.Printf("Nothing to complete.\n")
		return
	}

	fmt.Printf("Found %d candidates:\n", len(candidates))
	for _, c := range candidates {
		abbr := fmt.Sprintf("%s %s %s", c.Class, c.Name, c.Type)
		if c.Class == decl_func {
			abbr = fmt.Sprintf("%s %s%s", c.Class, c.Name, c.Type[len("func"):])
		}
		fmt.Printf("  %s\n", abbr)
	}
}

//-------------------------------------------------------------------------
// vim_formatter
//-------------------------------------------------------------------------

type vim_formatter struct{}

func (*vim_formatter) write_candidates(candidates []candidate, num int) {
	if candidates == nil {
		fmt.Print("[0, []]")
		return
	}

	fmt.Printf("[%d, [", num)
	for i, c := range candidates {
		if i != 0 {
			fmt.Printf(", ")
		}

		func_property := c.Type
		abbr := fmt.Sprintf("%s %s %s", c.Class, c.Name, c.Type)
		if c.Class == decl_func {
			func_property = c.Type[len("func"):]
			abbr = fmt.Sprintf("%s %s%s", c.Class, c.Name, func_property)
		}

		word := c.Name
		if c.Class == decl_func {
			if strings.HasPrefix(c.Type, "func()") {
				word += "()"
			} else {
				word += func_property
				sign_word := word
				pair := -1
				start := 0
				for i, n := 0, len(word); i < n; i++ {
					if word[i] == '(' {
						if pair < 0 {
							pair = 0
							start = i + 1
							sign_word = sign_word[0:i+1] + "`<"
						}
						pair += 1
					} else if word[i] == ')' {
						pair -= 1
						if pair == 0 {
							sign_word = sign_word + word[start:i] + ">`)"
							break
						}
					}

					if word[i] == ',' {
						sign_word = sign_word + word[start:i] + ">`, `<"
						start = i + 1
						if word[i+1] == ' ' {
							start += 1
						}
					}
				}
				word = sign_word
			}
		}

		fmt.Printf("{'word': '%s', 'abbr': '%s', 'info': '%s'}", word, abbr, abbr)
	}
	fmt.Printf("]]")
}

//-------------------------------------------------------------------------
// godit_formatter
//-------------------------------------------------------------------------

type godit_formatter struct{}

func (*godit_formatter) write_candidates(candidates []candidate, num int) {
	fmt.Printf("%d,,%d\n", num, len(candidates))
	for _, c := range candidates {
		contents := c.Name
		if c.Class == decl_func {
			contents += "("
			if strings.HasPrefix(c.Type, "func()") {
				contents += ")"
			}
		}

		display := fmt.Sprintf("%s %s %s", c.Class, c.Name, c.Type)
		if c.Class == decl_func {
			display = fmt.Sprintf("%s %s%s", c.Class, c.Name, c.Type[len("func"):])
		}
		fmt.Printf("%s,,%s\n", display, contents)
	}
}

//-------------------------------------------------------------------------
// emacs_formatter
//-------------------------------------------------------------------------

type emacs_formatter struct{}

func (*emacs_formatter) write_candidates(candidates []candidate, num int) {
	for _, c := range candidates {
		hint := c.Class.String() + " " + c.Type
		if c.Class == decl_func {
			hint = c.Type
		}
		fmt.Printf("%s,,%s\n", c.Name, hint)
	}
}

//-------------------------------------------------------------------------
// csv_formatter
//-------------------------------------------------------------------------

type csv_formatter struct{}

func (*csv_formatter) write_candidates(candidates []candidate, num int) {
	for _, c := range candidates {
		fmt.Printf("%s,,%s,,%s\n", c.Class, c.Name, c.Type)
	}
}

//-------------------------------------------------------------------------
// json_formatter
//-------------------------------------------------------------------------

type json_formatter struct{}

func (*json_formatter) write_candidates(candidates []candidate, num int) {
	if candidates == nil {
		fmt.Print("[]")
		return
	}

	fmt.Printf(`[%d, [`, num)
	for i, c := range candidates {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf(`{"class": "%s", "name": "%s", "type": "%s"}`,
			c.Class, c.Name, c.Type)
	}
	fmt.Print("]]")
}

//-------------------------------------------------------------------------

func get_formatter(name string) formatter {
	switch name {
	case "vim":
		return new(vim_formatter)
	case "emacs":
		return new(emacs_formatter)
	case "nice":
		return new(nice_formatter)
	case "csv":
		return new(csv_formatter)
	case "json":
		return new(json_formatter)
	case "godit":
		return new(godit_formatter)
	}
	return new(nice_formatter)
}
