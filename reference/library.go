package reference

import (
	"MonsterMash/network/retroarch"
	"fmt"
	"log"
)

type Library struct {
	client                  *retroarch.Client
	cleanupChannel          chan interface{}
	exitRequestedChannel    chan interface{}
	encounterIndexChannel   chan uint16
	encounterChangedChannel chan Encounter
	enemies                 []Enemy
	items                   []Item
	metamorphSets           [][]Item
	encounterSets           []Encounter
	loadedEncounter         *Encounter
}

func NewLibrary(itemData []byte, enemyData []byte, metaMorphData []byte, encounterData []byte) *Library {
	library := &Library{
		cleanupChannel:          nil,
		exitRequestedChannel:    nil,
		encounterIndexChannel:   nil,
		encounterChangedChannel: nil,
		enemies:                 nil,
		items:                   nil,
		metamorphSets:           nil,
		encounterSets:           nil,
		loadedEncounter:         nil,
	}

	library.setupChannels()
	library.loadData(itemData, enemyData, metaMorphData, encounterData)
	return library
}

func (l *Library) Start() {
	l.client = new(retroarch.Client)
	go l.client.ListenAndServe(l.encounterIndexChannel, l.exitRequestedChannel, l.cleanupChannel)
	go l.loadEncounter()
}

func (l *Library) Stop() {
	close(l.exitRequestedChannel)
	<-l.cleanupChannel
}

func (l *Library) GetClientStatus() interface{} {
	return struct {
		Status string `json:"status"`
	}{
		Status: l.client.Status,
	}
}

func (l *Library) GetCurrentEncounter() FlattenedEncounter {
	if l.loadedEncounter != nil {
		fe := FlattenedEncounter{}
		for _, m := range []*Enemy{
			l.loadedEncounter.Monster1Data,
			l.loadedEncounter.Monster2Data,
			l.loadedEncounter.Monster3Data,
			l.loadedEncounter.Monster4Data,
			l.loadedEncounter.Monster5Data,
			l.loadedEncounter.Monster6Data} {

			if m.Name != "Empty" {
				log.Println("Adding enemy " + m.Name)
				fm := FlattenedEnemy{}
				fm.Init(m, l)
				fe.Monsters = append(fe.Monsters, fm)
			}
		}

		return fe
	} else {
		return FlattenedEncounter{}
	}
}

func (l *Library) loadEncounter() {
	go func() {
		for {
			select {
			case encounterIdx := <-l.encounterIndexChannel:
				if l.loadedEncounter != &l.encounterSets[encounterIdx] {
					log.Printf("Loading new encounter %d\n", encounterIdx)
					l.loadedEncounter = &l.encounterSets[encounterIdx]
					log.Println(l.loadedEncounter)

					//log.Printf("Got encounter: %d\nEnemy 1: %s\nEnemy 2: %s\nEnemy 3: %s\nEnemy 4: %s\nEnemy 5: %s\nEnemy 6: %s\n\n",
					//	encounterIdx,
					//	l.loadedEncounter.Monster1Data.Name,
					//	l.loadedEncounter.Monster2Data.Name,
					//	l.loadedEncounter.Monster3Data.Name,
					//	l.loadedEncounter.Monster4Data.Name,
					//	l.loadedEncounter.Monster5Data.Name,
					//	l.loadedEncounter.Monster6Data.Name,
					//)
				}
			case <-l.exitRequestedChannel:
				log.Println("Encounter loader exit requested")
				return
			}
		}
	}()
}

// setupChannels builds the communication channels for sending/receiving the loaded encounter from SNES RAM and shutting
// down the program cleanly
func (l *Library) setupChannels() {
	encounterIndex := make(chan uint16)
	cleanup := make(chan interface{})
	exitRequested := make(chan interface{})
	encounterChanged := make(chan Encounter)

	l.encounterIndexChannel = encounterIndex
	l.cleanupChannel = cleanup
	l.exitRequestedChannel = exitRequested
	l.encounterChangedChannel = encounterChanged
}

// loadData calls the parses for the embedded YAML files that contains the encounter, metamorph, item, and formation data
func (l *Library) loadData(itemData []byte, monsterData []byte, morphData []byte, encounterData []byte) {
	items := ParseItems(itemData)
	monsters := ParseEnemyData(monsterData)
	monsters = append(monsters, Enemy{Name: "Empty"})

	l.enemies = monsters
	l.items = items
	l.metamorphSets = ParseMetamorphSets(morphData, items)
	l.encounterSets = ParseEncounterData(encounterData, l.enemies)

	fmt.Println("Loaded encounter sets in the library...")
	fmt.Println(l.encounterSets)
}
