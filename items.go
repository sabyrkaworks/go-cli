package main

import "slices"

var items []string = make([]string, 0)

func contains(slice []string, val string) bool {
	return slices.Contains(slice, val)
}

func findItemsToAdd(fields []string) []string {
	newItems := fields[1:]
	var addedItems []string

	for _, item := range newItems {
		if !contains(items, item) {
			items = append(items, item)
			addedItems = append(addedItems, item)
		}
	}

	return addedItems
}

func extractDisallowedItems(fields []string) []string {
	var deletedItems []string

	for i := 1; i < len(fields); i++ {
		valueToRemove := fields[i]

		foundIndex := -1
		for idx, item := range items {
			if item == valueToRemove {
				foundIndex = idx
				break
			}
		}

		if foundIndex != -1 {
			deletedItems = append(deletedItems, valueToRemove)
			items = append(items[:foundIndex], items[foundIndex+1:]...)
		}
	}

	return deletedItems
}
