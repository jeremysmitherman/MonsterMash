package ff6library

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
