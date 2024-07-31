package ff6library

import (
	"gopkg.in/yaml.v3"
	"log"
)

type Enemy struct {
	Name          string `yaml:"monsterName"`
	Image         string
	Level         uint8      `yaml:"level"`
	Hp            uint16     `yaml:"hp"`
	Mp            uint16     `yaml:"mp"`
	Atk           uint8      `yaml:"attackPower"`
	Def           uint8      `yaml:"defense"`
	Evd           uint8      `yaml:"evade"`
	HitRate       uint8      `yaml:"hitRate"`
	MAtk          uint8      `yaml:"magicPower"`
	MDef          uint8      `yaml:"magicDefense"`
	MEvd          uint8      `yaml:"magicBlock"`
	Xp            uint16     `yaml:"experience"`
	Gil           uint16     `yaml:"gp"`
	CommonDrop    uint8      `yaml:"commonDrop"`
	RareDrop      uint8      `yaml:"rareDrop"`
	CommonSteal   uint8      `yaml:"commonSteal"`
	RareSteal     uint8      `yaml:"rareSteal"`
	CommonSketch  uint8      `yaml:"sketch1"`
	RareSketch    uint8      `yaml:"sketch2"`
	Control1      uint8      `yaml:"control1"`
	Control2      uint8      `yaml:"control2"`
	Control3      uint8      `yaml:"control3"`
	Control4      uint8      `yaml:"control4"`
	Rage1         uint8      `yaml:"rage1"`
	Rage2         uint8      `yaml:"rage2"`
	MetamorphSet  uint8      `yaml:"metamorphItems"`
	MetamorphRate uint8      `yaml:"morphRate"`
	Weak          Elements   `yaml:"elementWeak"`
	Null          Elements   `yaml:"elementNull"`
	Absorb        Elements   `yaml:"elementAbsorb"`
	InitialStatus Conditions `yaml:"statusSet"`
	LockedStatus  Conditions `yaml:"statusImmunity"`
	Flags         MiscFlags  `yaml:"flags"`
}

type FlattenedEnemy struct {
	Name          string   `json:"name"`
	Level         uint8    `json:"level"`
	Hp            uint16   `json:"hp"`
	Mp            uint16   `json:"mp"`
	Atk           uint8    `json:"attack"`
	Def           uint8    `json:"defense"`
	Evd           uint8    `json:"evade"`
	HitRate       uint8    `json:"hitRate"`
	MAtk          uint8    `json:"magicPower"`
	MDef          uint8    `json:"magicDefense"`
	MEvd          uint8    `json:"magicBlock"`
	Xp            uint16   `json:"experience"`
	Gil           uint16   `json:"gp"`
	CommonDrop    string   `json:"commonDrop"`
	RareDrop      string   `json:"rareDrop"`
	CommonSteal   string   `json:"commonSteal"`
	RareSteal     string   `json:"rareSteal"`
	CommonSketch  string   `json:"sketch1"`
	RareSketch    string   `json:"sketch2"`
	Control1      string   `json:"control1"`
	Control2      string   `json:"control2"`
	Control3      string   `json:"control3"`
	Control4      string   `json:"control4"`
	Rage          string   `json:"rage1"`
	Rage2         string   `json:"rage2"`
	MetamorphSet  []string `json:"metamorphItems"`
	MetamorphRate string   `json:"morphRate"`
	Weak          []string `json:"elementWeak"`
	Null          []string `json:"elementNull"`
	Absorb        []string `json:"elementAbsorb"`
	InitialStatus []string `json:"statusSet"`
	LockedStatus  []string `json:"statusImmunity"`
	Flags         []string `json:"flags"`
}

func (f *FlattenedEnemy) Init(e *Enemy, l *Library) {
	f.Name = e.Name
	f.Level = e.Level
	f.Hp = e.Hp
	f.Mp = e.Mp
	f.Atk = e.Atk
	f.Def = e.Def
	f.Evd = e.Evd
	f.HitRate = e.HitRate
	f.MAtk = e.MAtk
	f.MDef = e.MDef
	f.MEvd = e.MEvd
	f.Xp = e.Xp
	f.Gil = e.Gil
	f.CommonDrop = l.items[e.CommonDrop].Name
	f.RareDrop = l.items[e.RareDrop].Name
	f.CommonSteal = l.items[e.CommonSteal].Name
	f.RareSteal = l.items[e.RareSteal].Name
	f.MetamorphRate = GetMetaMorphRate(e.MetamorphRate)
	f.Weak = ContainsElements(e.Weak)
	f.Null = ContainsElements(e.Null)
	f.Absorb = ContainsElements(e.Absorb)
	f.InitialStatus = ContainsConditions(e.InitialStatus)
	f.LockedStatus = ContainsConditions(e.LockedStatus)
	f.Flags = ContainsFlags(e.Flags)

	for _, i := range l.metamorphSets[e.MetamorphSet] {
		f.MetamorphSet = append(f.MetamorphSet, i.Name)
	}

	f.CommonSketch = l.skills[e.CommonSketch].Name
	f.RareSketch = l.skills[e.RareSketch].Name
	f.Control1 = l.skills[e.Control1].Name
	f.Control2 = l.skills[e.Control2].Name
	f.Control3 = l.skills[e.Control3].Name
	f.Control4 = l.skills[e.Control4].Name
}

func ParseEnemyData(enemyData []byte) []Enemy {
	var enemies []Enemy
	err := yaml.Unmarshal(enemyData, &enemies)
	if err != nil {
		log.Fatalln("Failed to read enemy data: " + err.Error())
	}

	return enemies
}
