package main

import "fmt"

func main() {
	names := map[string][]string{
		"bond_james":        {`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jennys": {`intelligence`, `literature`, `computer science`},
		"no_dr":             {`cats`, `ice cream`, `sunsets`},
	}

	names["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}

	for key, value := range names {
		fmt.Printf("%v : %v\n", key, value)
	}
}
