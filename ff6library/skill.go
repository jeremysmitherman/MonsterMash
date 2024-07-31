package ff6library

import (
	"gopkg.in/yaml.v3"
	"log"
)

type Skill struct {
	Name           string `yaml:"miscName"`
	ElementFlags   uint8  `yaml:"element"`
	ElementsString []string
}

func ParseSkills(skillData []byte) ([]Skill, error) {
	var skills []Skill
	err := yaml.Unmarshal(skillData, &skills)
	if err != nil {
		log.Fatalln("Failed to parse skill data")
	}

	for i := 0; i < len(skills); i++ {
		if skills[i].ElementFlags != 0 {
			skills[i].ElementsString = ContainsElements(Elements(skills[i].ElementFlags))
		}
	}

	return skills, err
}
