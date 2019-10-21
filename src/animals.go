package k8szoo

import (
	"strings"
)

type AnimalData struct {
	AnimalName  string
	AnimalSound string
}

var Animals = []AnimalData {
	{"Alligator","bellowed"},
	{"Antelope","snorted"},
	{"Badger","growled"},
	{"Bat","screeched"},
	{"Bear","roared"},
	{"Bee","buzz"},
	{"Tiger","snarled"},
	{"Lion","roared"},
	{"Panther","growled"},
	{"Bittern","boomed"},
	{"Cat","meowed"},
	{"Chicken","clucked"},
	{"Cockerel","crowed"},
	{"Chimpanzee","screamed"},
	{"Chinchilla","squeaked"},
	{"Cicada","chirped"},
	{"Cow"," mooed"},
	{"Cricket","chirped"},
	{"Crow","cawed"},
	{"Curlew","piped"},
	{"Deer","bleated"},
	{"Dog","barked"},
	{"Wolf","howled"},
	{"Dolphin","clicked"},
	{"Donkey","brayed"},
	{"Duck","quacked"},
	{"Eagle","screeched"},
	{"Elephant","trumpeted"},
	{"Elk","bleated"},
	{"Fox","yiffed"},
	{"Ferret","dooked"},
	{"Toad","croaked"},
	{"Frog","ribbitted"},
	{"Giraffe","bleated"},
	{"Goose","honked"},
	{"Grasshopper","chirped"},
	{"Guinea pig","squeaked"},
	{"Hamster","squeaked"},
	{"Hermit crab","chirped"},
	{"Horse","neighed"},
	{"Hippo","growled"},
	{"Hyena","laughed"},
	{"Linnet","chuckled"},
	{"Magpie","chattered"},
	{"Mouse","squeaked"},
	{"Monkey","chattered"},
	{"Moose","bellowed"},
	{"Mosquito","buzzed"},
	{"Okapi","coughed"},
	{"Ox","mooed"},
	{"Owl","hooted"},
	{"Parrot","squawked"},
	{"Peacock","screamed"},
	{"Pig","oinked"},
	{"Pigeon","cooed"},
	{"Prairie dog","barked"},
	{"Rabbit","squeaked"},
	{"Raccoon","trilled"},
	{"Raven","cawed"},
	{"Rhinoceros","bellowed"},
	{"Rook","cawed"},
	{"Seal","barked"},
	{"Sheep","bleated"},
	{"Snake","hissed"},
	{"Songbird","chirrupped"},
	{"Swan","cried"},
	{"Tapir","squeaked"},
	{"Tarantula","hissed"},
	{"Gecko","croaked"},
	{"Turkey","gobbled"},
	{"Vulture","screamed"},
	{"Walrus","groaned"},
	{"Whale","sang"},
	{"Zebra","brayed"},
}

func FindAnimal(animalName string) AnimalData {
	for _, listItem := range(Animals) {
		if strings.EqualFold(listItem.AnimalName, animalName) {
			return listItem
		}
	}
	var ret AnimalData
	ret.AnimalName = ""
	return ret
}