package ff6library

import (
	"MonsterMash/ff6library/network/retroarch"
	"fmt"
	"log"
	"sync"
)

type Library struct {
	client               *retroarch.Client
	dataChannel          chan []uint16
	cleanupChannel       chan interface{}
	exitRequestedChannel chan interface{}
	enemies              []Enemy
	currentEncounter     []Enemy
	items                []Item
	skills               []Skill
	metamorphSets        [][]Item
	lock                 sync.Mutex
}

func NewLibrary(itemData []byte, enemyData []byte, metaMorphData []byte, skillData []byte) *Library {
	library := &Library{}
	library.setupChannels()
	library.loadData(itemData, enemyData, metaMorphData, skillData)
	return library
}

func (l *Library) Start() {
	// Client will reach out to the emulator periodically and send an updated enemy list to the dataChannel
	l.client = new(retroarch.Client)
	go l.client.ListenAndServe(l.dataChannel, l.exitRequestedChannel, l.cleanupChannel)

	// Listen to the aforementioned datachannel to update the current list of enemies.  This list of enemies is what
	// powers GetCurrentEncounter which is what the frontend calls for display
	go func() {
		for {
			log.Println("Asking for new indexes")
			enemyIndexes := <-l.dataChannel
			log.Printf("Got indexes: %v\n", enemyIndexes)
			l.lock.Lock()
			log.Println("Locked current encounter")
			l.currentEncounter = []Enemy{}
			for _, e := range enemyIndexes {
				l.currentEncounter = append(l.currentEncounter, l.enemies[e])
			}
			l.lock.Unlock()
			log.Println("Unlocked current encounter")
			fmt.Println(l.currentEncounter)
		}
	}()
}

func (l *Library) Stop() {
	close(l.exitRequestedChannel)
	<-l.cleanupChannel
}

func (l *Library) GetClientStatus() interface{} {
	return struct {
		Status  string `json:"status"`
		Message string `json:"status_message"`
	}{
		Status:  l.client.Status,
		Message: l.client.StatusMessage,
	}
}

func (l *Library) GetCurrentEncounter() []FlattenedEnemy {
	var flattenedEnemies []FlattenedEnemy
	l.lock.Lock()
	for _, e := range l.currentEncounter {
		fe := FlattenedEnemy{}
		fe.Init(&e, l)
		flattenedEnemies = append(flattenedEnemies, fe)
	}
	l.lock.Unlock()

	return flattenedEnemies
}

// setupChannels builds the communication channels for sending/receiving the loaded encounter from SNES RAM and shutting
// down the program cleanly
func (l *Library) setupChannels() {
	l.cleanupChannel = make(chan interface{})
	l.exitRequestedChannel = make(chan interface{})
	l.dataChannel = make(chan []uint16)
}

// loadData calls the parses for the embedded YAML files that contains the encounter, metamorph, item, and formation data
func (l *Library) loadData(itemData []byte, monsterData []byte, morphData []byte, skillData []byte) {
	skills, _ := ParseSkills(skillData)
	items := ParseItems(itemData)
	monsters := ParseEnemyData(monsterData)
	monsters = append(monsters, Enemy{Name: "Empty"})

	l.enemies = monsters
	l.items = items
	l.metamorphSets = ParseMetamorphSets(morphData, items)
	l.skills = skills
}
