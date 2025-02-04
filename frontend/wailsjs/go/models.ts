export namespace ff6library {
	
	export class FlattenedEnemy {
	    name: string;
	    level: number;
	    hp: number;
	    mp: number;
	    attack: number;
	    defense: number;
	    evade: number;
	    hitRate: number;
	    magicPower: number;
	    magicDefense: number;
	    magicBlock: number;
	    experience: number;
	    gp: number;
	    commonDrop: string;
	    rareDrop: string;
	    commonSteal: string;
	    rareSteal: string;
	    sketch1: string;
	    sketch2: string;
	    control1: string;
	    control2: string;
	    control3: string;
	    control4: string;
	    rage1: string;
	    rage2: string;
	    metamorphItems: string[];
	    morphRate: string;
	    elementWeak: string[];
	    elementNull: string[];
	    elementAbsorb: string[];
	    statusSet: string[];
	    statusImmunity: string[];
	    flags: string[];
	
	    static createFrom(source: any = {}) {
	        return new FlattenedEnemy(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.level = source["level"];
	        this.hp = source["hp"];
	        this.mp = source["mp"];
	        this.attack = source["attack"];
	        this.defense = source["defense"];
	        this.evade = source["evade"];
	        this.hitRate = source["hitRate"];
	        this.magicPower = source["magicPower"];
	        this.magicDefense = source["magicDefense"];
	        this.magicBlock = source["magicBlock"];
	        this.experience = source["experience"];
	        this.gp = source["gp"];
	        this.commonDrop = source["commonDrop"];
	        this.rareDrop = source["rareDrop"];
	        this.commonSteal = source["commonSteal"];
	        this.rareSteal = source["rareSteal"];
	        this.sketch1 = source["sketch1"];
	        this.sketch2 = source["sketch2"];
	        this.control1 = source["control1"];
	        this.control2 = source["control2"];
	        this.control3 = source["control3"];
	        this.control4 = source["control4"];
	        this.rage1 = source["rage1"];
	        this.rage2 = source["rage2"];
	        this.metamorphItems = source["metamorphItems"];
	        this.morphRate = source["morphRate"];
	        this.elementWeak = source["elementWeak"];
	        this.elementNull = source["elementNull"];
	        this.elementAbsorb = source["elementAbsorb"];
	        this.statusSet = source["statusSet"];
	        this.statusImmunity = source["statusImmunity"];
	        this.flags = source["flags"];
	    }
	}

}

