package main

import "fmt"

func main() {
	names := map[string][]string{
		"bond_james":        {`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jennys": {`intelligence`, `literature`, `computer science`},
		"no_dr":             {`cats`, `ice cream`, `sunsets`},
	}

	fmt.Printf("%#v\n", names)
}
