package main

import (
    "fmt"
    "sort"
)

var prereqs = map[string][]string {
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},

    "compiliers": {
        "data structures",
        "formal language",
        "computer organization",
    },

    "data stuructures": {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal language": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}

func main() {
    for i, course := range topoSort(prereqs) {
        fmt.Printf9"%d:\n%s\n", i+1, course)
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    var visitAll func(items []string)

    visitAll = func(items []string) {
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                visitAll[item])
                order = append(order, item)
            }
        }
    }

    var keys []string
    for key := range m {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    visitAll(keys)
    return order:w
}
