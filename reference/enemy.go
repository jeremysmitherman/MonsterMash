package reference

import (
	"gopkg.in/yaml.v3"
	"log"
)

type Conditions uint32
type Elements uint8
type MiscFlags uint16
type MetamorphRate uint8

const (
	Dark Conditions = 1 << iota
	Zombie
	PoisonStatus
	Magitek
	Vanish
	Imp
	Petrify
	Death
	Condemned
	Kneeling
	Blink
	Silence
	Berserk
	Confusion
	HpDrain
	Sleep
	Flying
	Regen
	Slow
	Haste
	Stop
	Shell
	Safe
	Reflect
	Rage
	Frozen
	DeathProtection
	Morph
	Casting
	Removed
	Interceptor
	Float
)

const (
	Fire Elements = 1 << iota
	Ice
	Lightning
	PoisonElement
	Wind
	Pearl
	Earth
	Water
)

const (
	Mystic MiscFlags = 1 << iota
	HideName
	Human        = 16
	ImpCrit      = 64
	Undead       = 128
	DifficultRun = 256
	FirstStrike  = 512
	CantSuplex   = 1024
	CantRun      = 2048
	CantScan     = 4096
	CantSketch   = 8192
	SpecialEvent = 16384
	CantControl  = 32768
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
	Rage          string   `json:"rage"`
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

	//f.CommonSketch = l.items[e.CommonSketch].Name
	//f.RareSketch = l.items[e.RareSketch].Name
	//f.Control1 = l.items[e.Control1].Name

}

type FlattenedEncounter struct {
	Monsters []FlattenedEnemy `json:"monsters"`
}

type Encounter struct {
	EncounterID int    `json:"encounter_id"`
	Monster1ID  uint16 `yaml:"monster1"`
	Monster2ID  uint16 `yaml:"monster2"`
	Monster3ID  uint16 `yaml:"monster3"`
	Monster4ID  uint16 `yaml:"monster4"`
	Monster5ID  uint16 `yaml:"monster5"`
	Monster6ID  uint16 `yaml:"monster6"`

	Monster1Data *Enemy `json:"monster_1_data"`
	Monster2Data *Enemy `json:"monster_2_data"`
	Monster3Data *Enemy `json:"monster_3_data"`
	Monster4Data *Enemy `json:"monster_4_data"`
	Monster5Data *Enemy `json:"monster_5_data"`
	Monster6Data *Enemy `json:"monster_6_data"`
}

func (e *Encounter) Init(encounterIdx int, enemyList []Enemy) {
	e.EncounterID = encounterIdx

	if e.Monster1ID != 511 {
		e.Monster1Data = &enemyList[e.Monster1ID]
	} else {
		e.Monster1Data = &enemyList[len(enemyList)-1]
	}

	if e.Monster2ID != 511 {
		e.Monster2Data = &enemyList[e.Monster2ID]
	} else {
		e.Monster2Data = &enemyList[len(enemyList)-1]
	}

	if e.Monster3ID != 511 {
		e.Monster3Data = &enemyList[e.Monster3ID]
	} else {
		e.Monster3Data = &enemyList[len(enemyList)-1]
	}

	if e.Monster4ID != 511 {
		e.Monster4Data = &enemyList[e.Monster4ID]
	} else {
		e.Monster4Data = &enemyList[len(enemyList)-1]
	}

	if e.Monster5ID != 511 {
		e.Monster5Data = &enemyList[e.Monster5ID]
	} else {
		e.Monster5Data = &enemyList[len(enemyList)-1]
	}

	if e.Monster6ID != 511 {
		e.Monster6Data = &enemyList[e.Monster6ID]
	} else {
		e.Monster6Data = &enemyList[len(enemyList)-1]
	}
}

func ContainsElements(elements Elements) []string {
	var elementList []string
	if Has(uint32(elements), uint32(Fire)) {
		elementList = append(elementList, "Fire")
	}

	if Has(uint32(elements), uint32(Ice)) {
		elementList = append(elementList, "Ice")
	}

	if Has(uint32(elements), uint32(Lightning)) {
		elementList = append(elementList, "Lightning")
	}

	if Has(uint32(elements), uint32(PoisonElement)) {
		elementList = append(elementList, "Poison")
	}

	if Has(uint32(elements), uint32(Pearl)) {
		elementList = append(elementList, "Holy")
	}

	if Has(uint32(elements), uint32(Wind)) {
		elementList = append(elementList, "Wind")
	}

	if Has(uint32(elements), uint32(Earth)) {
		elementList = append(elementList, "Earth")
	}

	if Has(uint32(elements), uint32(Water)) {
		elementList = append(elementList, "Water")
	}

	return elementList
}

func ContainsConditions(conditions Conditions) []string {
	var conditionList []string
	if Has(uint32(conditions), uint32(Dark)) {
		conditionList = append(conditionList, "Dark")
	}

	if Has(uint32(conditions), uint32(Zombie)) {
		conditionList = append(conditionList, "Zombie")
	}

	if Has(uint32(conditions), uint32(PoisonStatus)) {
		conditionList = append(conditionList, "Poison")
	}

	if Has(uint32(conditions), uint32(Magitek)) {
		conditionList = append(conditionList, "Magitek")
	}

	if Has(uint32(conditions), uint32(Vanish)) {
		conditionList = append(conditionList, "Vanish")
	}

	if Has(uint32(conditions), uint32(Imp)) {
		conditionList = append(conditionList, "Imp")
	}

	if Has(uint32(conditions), uint32(Petrify)) {
		conditionList = append(conditionList, "Petrify")
	}

	if Has(uint32(conditions), uint32(Death)) {
		conditionList = append(conditionList, "Death")
	}

	if Has(uint32(conditions), uint32(Condemned)) {
		conditionList = append(conditionList, "Condemned")
	}

	if Has(uint32(conditions), uint32(Kneeling)) {
		conditionList = append(conditionList, "Kneeling")
	}

	if Has(uint32(conditions), uint32(Blink)) {
		conditionList = append(conditionList, "Blink")
	}

	if Has(uint32(conditions), uint32(Silence)) {
		conditionList = append(conditionList, "Silence")
	}

	if Has(uint32(conditions), uint32(Berserk)) {
		conditionList = append(conditionList, "Berserk")
	}

	if Has(uint32(conditions), uint32(Confusion)) {
		conditionList = append(conditionList, "Confusion")
	}

	if Has(uint32(conditions), uint32(HpDrain)) {
		conditionList = append(conditionList, "HP Drain")
	}

	if Has(uint32(conditions), uint32(Sleep)) {
		conditionList = append(conditionList, "Sleep")
	}

	if Has(uint32(conditions), uint32(Flying)) {
		conditionList = append(conditionList, "Flying")
	}

	if Has(uint32(conditions), uint32(Regen)) {
		conditionList = append(conditionList, "Regen")
	}

	if Has(uint32(conditions), uint32(Slow)) {
		conditionList = append(conditionList, "Slow")
	}

	if Has(uint32(conditions), uint32(Haste)) {
		conditionList = append(conditionList, "Haste")
	}

	if Has(uint32(conditions), uint32(Stop)) {
		conditionList = append(conditionList, "Stop")
	}

	if Has(uint32(conditions), uint32(Shell)) {
		conditionList = append(conditionList, "Shell")
	}

	if Has(uint32(conditions), uint32(Safe)) {
		conditionList = append(conditionList, "Safe")
	}

	if Has(uint32(conditions), uint32(Reflect)) {
		conditionList = append(conditionList, "Reflect")
	}

	if Has(uint32(conditions), uint32(Rage)) {
		conditionList = append(conditionList, "Rage")
	}

	if Has(uint32(conditions), uint32(Frozen)) {
		conditionList = append(conditionList, "Frozen")
	}

	if Has(uint32(conditions), uint32(DeathProtection)) {
		conditionList = append(conditionList, "Death Protection")
	}

	if Has(uint32(conditions), uint32(Morph)) {
		conditionList = append(conditionList, "Morph")
	}

	if Has(uint32(conditions), uint32(Casting)) {
		conditionList = append(conditionList, "Casting")
	}

	if Has(uint32(conditions), uint32(Removed)) {
		conditionList = append(conditionList, "Removed")
	}

	if Has(uint32(conditions), uint32(Interceptor)) {
		conditionList = append(conditionList, "Interceptor")
	}

	if Has(uint32(conditions), uint32(Float)) {
		conditionList = append(conditionList, "Float")
	}

	return conditionList
}

func ContainsFlags(flag MiscFlags) []string {
	var conditions []string
	if Has(uint32(Mystic), uint32(flag)) {
		conditions = append(conditions, "Dies at 0MP")
	}

	if Has(uint32(HideName), uint32(flag)) {
		conditions = append(conditions, "Name Hidden")
	}

	if Has(uint32(Human), uint32(flag)) {
		conditions = append(conditions, "Human")
	}

	if Has(uint32(ImpCrit), uint32(flag)) {
		conditions = append(conditions, "Imp Crit")
	}

	if Has(uint32(Undead), uint32(flag)) {
		conditions = append(conditions, "Undead")
	}

	if Has(uint32(DifficultRun), uint32(flag)) {
		conditions = append(conditions, "Hard to Run")
	}

	if Has(uint32(FirstStrike), uint32(flag)) {
		conditions = append(conditions, "First Strike")
	}

	if Has(uint32(CantSuplex), uint32(flag)) {
		conditions = append(conditions, "Can't Suplex")
	}

	if Has(uint32(CantRun), uint32(flag)) {
		conditions = append(conditions, "Can't Run")
	}

	if Has(uint32(CantScan), uint32(flag)) {
		conditions = append(conditions, "Can't Scan")
	}

	if Has(uint32(CantSketch), uint32(flag)) {
		conditions = append(conditions, "Can't Sketch")
	}

	if Has(uint32(SpecialEvent), uint32(flag)) {
		conditions = append(conditions, "Special Event")
	}

	if Has(uint32(CantControl), uint32(flag)) {
		conditions = append(conditions, "Can't Control")
	}

	return conditions
}

func GetMetaMorphRate(flag uint8) string {
	switch flag {
	case 0:
		return "99.6%"
	case 1:
		return "75%"
	case 2:
		return "50%"
	case 3:
		return "25%"
	case 4:
		return "12.5%"
	case 5:
		return "6.25%"
	case 6:
		return "3%"
	case 7:
		return "0%"
	default:
		return "???"
	}
}

func ParseEnemyData(enemyData []byte) []Enemy {
	var enemies []Enemy
	err := yaml.Unmarshal(enemyData, &enemies)
	if err != nil {
		log.Fatalln("Failed to read enemy data: " + err.Error())
	}

	return enemies
}

func ParseEncounterData(encounterData []byte, enemyList []Enemy) []Encounter {
	var encounters []Encounter
	err := yaml.Unmarshal(encounterData, &encounters)
	if err != nil {
		log.Fatalln("Failed to read encounter data: " + err.Error())
	}

	for i := 0; i < len(encounters); i++ {
		encounters[i].Init(i, enemyList)
	}

	return encounters
}
