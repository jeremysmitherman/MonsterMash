package ff6library

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

func Has(b, flag uint32) bool { return b&flag != 0 }

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
