package ff6library

import (
	"gopkg.in/yaml.v3"
	"log"
)

type Item struct {
	Name string `yaml:"itemName"`
}

func ParseItems(itemData []byte) []Item {
	var items []Item
	err := yaml.Unmarshal(itemData, &items)
	if err != nil {
		log.Fatalln("Failed to read item data: " + err.Error())
	}

	return items
}

func ParseMetamorphSets(morphData []byte, items []Item) [][]Item {
	type rawItemSet struct {
		Item1 uint8
		Item2 uint8
		Item3 uint8
		Item4 uint8
	}

	var rawItemSets []rawItemSet
	err := yaml.Unmarshal(morphData, &rawItemSets)
	if err != nil {
		log.Fatalln("Failed to read morph data: " + err.Error())
	}

	var ItemSets [][]Item
	for _, set := range rawItemSets {
		ItemSets = append(ItemSets, []Item{items[set.Item1], items[set.Item2], items[set.Item3], items[set.Item4]})
	}

	return ItemSets
}
