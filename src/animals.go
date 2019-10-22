package k8szoo

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type AnimalData struct {
	AnimalName  string
	AnimalSound string
	PictureURL  string
}

var Animals = []AnimalData{
	{"Alligator", "bellowed", "https://upload.wikimedia.org/wikipedia/commons/6/65/AmericanAlligator.JPG"},
	{"Antelope", "snorted", "https://upload.wikimedia.org/wikipedia/commons/b/bb/Sable_bull.jpg"},
	{"Badger", "growled", "https://upload.wikimedia.org/wikipedia/commons/1/10/Badger-badger.jpg"},
	{"Bat", "screeched", "https://upload.wikimedia.org/wikipedia/commons/6/66/Short-nosed_Fruit_Bat_%28Cynopterus_sphinx%29_Photograph_By_Shantanu_Kuveskar.jpg"},
	{"Bear", "roared", "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Kamchatka_Brown_Bear_near_Dvuhyurtochnoe_on_2015-07-23.jpg/1200px-Kamchatka_Brown_Bear_near_Dvuhyurtochnoe_on_2015-07-23.jpg"},
	{"Bee", "buzzed", "https://upload.wikimedia.org/wikipedia/commons/7/76/Tetragonula_carbonaria_%2814521993792%29.jpg"},
	{"Tiger", "snarled", "https://upload.wikimedia.org/wikipedia/commons/d/da/Panthera_tigris_tigris_Tidoba_20150306.jpg"},
	{"Lion", "roared", "https://upload.wikimedia.org/wikipedia/commons/d/da/Panthera_tigris_tigris_Tidoba_20150306.jpg"},
	{"Panther", "growled", "https://upload.wikimedia.org/wikipedia/commons/f/fc/Jaguar.jpg"},
	{"Bittern", "boomed", "https://upload.wikimedia.org/wikipedia/commons/4/45/American_Bittern_Seney_NWR_3.jpg"},
	{"Cat", "meowed", "https://upload.wikimedia.org/wikipedia/commons/1/1e/Calico_cat_-_bright.jpg"},
	{"Chicken", "clucked", "https://upload.wikimedia.org/wikipedia/commons/b/b8/Cascais_Costa_do_Esteril_52_%2836583204550%29.jpg"},
	{"Rooster", "crowed", "https://en.wikipedia.org/wiki/Chicken#/media/File:Rooster_portrait2.jpg"},
	{"Chimpanzee", "screamed", "https://upload.wikimedia.org/wikipedia/commons/6/68/Eastern_Chimpanzee_%28Pan_troglodytes_schweinfurthii%29_%287068198095%29_%28cropped%29.jpg"},
	{"Chinchilla", "squeaked", "https://upload.wikimedia.org/wikipedia/commons/1/18/Chinchilla_lanigera_%28Wroclaw_zoo%29-2.JPG"},
	{"Cicada", "chirped", "https://upload.wikimedia.org/wikipedia/commons/f/fb/Tibicen_linnei.jpg"},
	{"Cow", " mooed", "https://upload.wikimedia.org/wikipedia/commons/d/d4/CH_cow_2_cropped.jpg"},
	{"Cricket", "chirped", "https://upload.wikimedia.org/wikipedia/commons/b/b5/Jiminy_Cricket.png"},
	{"Crow", "cawed", "https://upload.wikimedia.org/wikipedia/commons/a/a9/Corvus_corone_-near_Canford_Cliffs%2C_Poole%2C_England-8.jpg"},
	{"Curlew", "piped", "https://upload.wikimedia.org/wikipedia/commons/0/08/Curlew_-_natures_pics.jpg"},
	{"Deer", "bleated", "https://upload.wikimedia.org/wikipedia/commons/6/62/Chital_%288458215435%29.jpg"},
	{"Dog", "barked", "https://upload.wikimedia.org/wikipedia/commons/9/93/Golden_Retriever_Carlos_%2810581910556%29.jpg"},
	{"Wolf", "howled", "https://upload.wikimedia.org/wikipedia/commons/5/5f/Kolm%C3%A5rden_Wolf.jpg"},
	{"Dolphin", "clicked", "https://upload.wikimedia.org/wikipedia/commons/1/10/Tursiops_truncatus_01.jpg"},
	{"Donkey", "brayed", "https://upload.wikimedia.org/wikipedia/commons/7/7b/Donkey_1_arp_750px.jpg"},
	{"Duck", "quacked", "https://upload.wikimedia.org/wikipedia/commons/b/bf/Bucephala-albeola-010.jpg"},
	{"Eagle", "screeched", "https://upload.wikimedia.org/wikipedia/commons/d/d6/Golden_Eagle_in_flight_-_5.jpg"},
	{"Elephant", "trumpeted", "https://upload.wikimedia.org/wikipedia/commons/9/91/African_Elephant_%28Loxodonta_africana%29_bull_%2831100819046%29.jpg"},
	{"Elk", "bleated", "https://upload.wikimedia.org/wikipedia/commons/5/55/Rocky_Mountain_Bull_Elk.jpg"},
	{"Fox", "yiffed", "https://upload.wikimedia.org/wikipedia/commons/0/00/Cerdocyon_thous_MG_9503.jpg"},
	{"Ferret", "dooked", "https://upload.wikimedia.org/wikipedia/commons/3/32/Ferret_2008.png"},
	{"Toad", "croaked", "https://upload.wikimedia.org/wikipedia/commons/2/2b/Crinia_signifera.jpg"},
	{"Frog", "ribbitted", "https://upload.wikimedia.org/wikipedia/commons/c/c1/Variegated_golden_frog_%28Mantella_baroni%29_Ranomafana.jpg"},
	{"Giraffe", "bleated", "https://upload.wikimedia.org/wikipedia/commons/9/9e/Giraffe_Mikumi_National_Park.jpg"},
	{"Goose", "honked", "https://upload.wikimedia.org/wikipedia/commons/3/34/Anser_anser_1_%28Piotr_Kuczynski%29.jpg"},
	{"Grasshopper", "chirped", "https://upload.wikimedia.org/wikipedia/commons/3/37/Heupferd_fg01.jpg"},
	{"Guinea pig", "squeaked", "https://upload.wikimedia.org/wikipedia/commons/8/88/Guinea_Pig_eating_apple.JPG"},
	{"Hamster", "squeaked", "https://en.wikipedia.org/wiki/Hamster#/media/File:Pearl_Winter_White_Russian_Dwarf_Hamster_-_Front.jpg"},
	{"Hermit crab", "chirped", "https://upload.wikimedia.org/wikipedia/commons/8/8e/Calliactis_and_Dardanus_001.JPG"},
	{"Horse", "neighed", "https://upload.wikimedia.org/wikipedia/commons/a/aa/Przewalski%27s_Horse_%2802710137%29.jpg"},
	{"Hippo", "growled", "https://upload.wikimedia.org/wikipedia/commons/b/b3/Hipop%C3%B3tamo_%28Hippopotamus_amphibius%29%2C_parque_nacional_de_Chobe%2C_Botsuana%2C_2018-07-28%2C_DD_82.jpg"},
	{"Hyena", "laughed", "https://upload.wikimedia.org/wikipedia/commons/d/dc/Spotted_hyena_%28Crocuta_crocuta%29.jpg"},
	{"Linnet", "chuckled", "https://upload.wikimedia.org/wikipedia/commons/c/cf/Carduelis_cannabina_-England_-male-8.jpg"},
	{"Magpie", "chattered", "https://upload.wikimedia.org/wikipedia/commons/c/c2/Birds_of_Sweden_2016_35.jpg"},
	{"Mouse", "squeaked", "https://upload.wikimedia.org/wikipedia/commons/1/11/ApodemusSylvaticus.jpg"},
	{"Monkey", "chattered", "https://upload.wikimedia.org/wikipedia/commons/4/47/Squirrel_monkey_in_Ubon_Zoo%2CThailand.jpg"},
	{"Moose", "bellowed", "https://upload.wikimedia.org/wikipedia/commons/8/8b/Moose_superior.jpg"},
	{"Mosquito", "buzzed", "https://upload.wikimedia.org/wikipedia/commons/d/d0/Aedes_aegypti.jpg"},
	{"Okapi", "coughed", "https://upload.wikimedia.org/wikipedia/commons/3/3a/Saint-Aignan_%28Loir-et-Cher%29._Okapi.jpg"},
	{"Ox", "mooed", "https://upload.wikimedia.org/wikipedia/commons/d/dd/Traditional_Farming_Methods_and_Equipments.jpg"},
	{"Owl", "hooted", "https://upload.wikimedia.org/wikipedia/commons/d/d3/Ural_Owl_%28Strix_uralensis%29_in_Ljubljana%2C_Slovenia.jpg"},
	{"Parrot", "squawked", "https://upload.wikimedia.org/wikipedia/commons/8/88/Eclectus_roratus-20030511.jpg"},
	{"Peacock", "screamed", "https://upload.wikimedia.org/wikipedia/commons/c/c5/Peacock_Plumage.jpg"},
	{"Pig", "oinked", "https://upload.wikimedia.org/wikipedia/commons/2/27/Sus_scrofa_domesticus%2C_miniature_pig%2C_juvenile.jpg"},
	{"Pigeon", "cooed", "https://upload.wikimedia.org/wikipedia/commons/f/ff/Treron_vernans_male_-_Kent_Ridge_Park.jpg"},
	{"Prairie dog", "barked", "https://upload.wikimedia.org/wikipedia/commons/c/c7/Cynomys_ludovicianus_-Paignton_Zoo%2C_Devon%2C_England-8a.jpg"},
	{"Rabbit", "squeaked", "https://upload.wikimedia.org/wikipedia/commons/5/50/Sylvilagus_bachmani_01035t.JPG"},
	{"Raccoon", "trilled", "https://upload.wikimedia.org/wikipedia/commons/c/cd/Common_Raccoon_%28Procyon_lotor%29_in_Northwest_Indiana.jpg"},
	{"Raven", "cawed", "https://upload.wikimedia.org/wikipedia/commons/3/31/3782_Common_Raven_in_flight.jpg"},
	{"Rhinoceros", "bellowed", "https://upload.wikimedia.org/wikipedia/commons/6/63/Diceros_bicornis.jpg"},
	{"Rook", "cawed", "https://upload.wikimedia.org/wikipedia/commons/6/63/Rook_at_Slimbridge_Wetland_Centre%2C_Gloucestershire%2C_England_22May2019_arp.jpg"},
	{"Seal", "barked", "https://upload.wikimedia.org/wikipedia/commons/7/7d/Seehund.jpg"},
	{"Sheep", "bleated", "https://upload.wikimedia.org/wikipedia/commons/d/d7/Lička_pramenka.jpg"},
	{"Snake", "hissed", "https://upload.wikimedia.org/wikipedia/commons/8/8c/Milksnake2.jpg"},
	{"Songbird", "chirrupped", "https://upload.wikimedia.org/wikipedia/commons/d/d0/Eastern_yellow_robin.jpg"},
	{"Swan", "cried", "https://upload.wikimedia.org/wikipedia/commons/d/df/Black-necked_swan_745r.jpg"},
	{"Tapir", "squeaked", "https://upload.wikimedia.org/wikipedia/commons/3/36/South_American_tapir_%28Tapirus_terrestris%29.JPG"},
	{"Tarantula", "hissed", "https://upload.wikimedia.org/wikipedia/commons/9/98/Brachypelma_smithi_2009_G03.jpg"},
	{"Gecko", "croaked", "https://upload.wikimedia.org/wikipedia/commons/a/ae/Mediterranean_house_gecko.JPG"},
	{"Turkey", "gobbled", "https://upload.wikimedia.org/wikipedia/commons/c/cf/Meleagris_ocellata_-Guatemala-8a.jpg"},
	{"Vulture", "screamed", "https://upload.wikimedia.org/wikipedia/commons/5/5d/Vulture_19o05.jpg"},
	{"Walrus", "groaned", "https://upload.wikimedia.org/wikipedia/commons/2/22/Pacific_Walrus_-_Bull_%288247646168%29.jpg"},
	{"Whale", "sang", "https://upload.wikimedia.org/wikipedia/commons/e/e2/Delphinapterus_leucas_in_shallows.jpg"},
	{"Zebra", "brayed", "https://upload.wikimedia.org/wikipedia/commons/e/e3/Blondzebra.jpg"},
}

var AvailableAnimals = make([]AnimalData, len(Animals))

func FindAnimal(animalName string) AnimalData {
	for _, listItem := range Animals {
		if strings.EqualFold(listItem.AnimalName, animalName) {
			return listItem
		}
	}
	var ret AnimalData
	ret.AnimalName = ""
	return ret
}

func FindAnimalID(animalName string) int {
	for i, listItem := range Animals {
		if strings.EqualFold(listItem.AnimalName, animalName) {
			return i
		}
	}
	return -1
}

func ReserveRandomAnimal() (int, AnimalData) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	availID := r.Intn(len(AvailableAnimals))
	fmt.Printf("Decided on %d ", availID)
	animalData := AvailableAnimals[availID]
	fmt.Printf("which is %s ", animalData.AnimalName)
	fullID := FindAnimalID(animalData.AnimalName)
	fmt.Printf(", that's actually %d as an ID\n", fullID)
	copy(AvailableAnimals[availID:], AvailableAnimals[availID+1:]) // Shift AvailableAnimals[i+1:] left one index.
	AvailableAnimals = AvailableAnimals[:len(AvailableAnimals)-1]  // Truncate slice.
	fmt.Println("Global list updated.")
	fmt.Println("========================================")
	return fullID, animalData
}

func IsAnimalReserved(animalName string) bool {
	for _, listItem := range AvailableAnimals {
		if strings.EqualFold(listItem.AnimalName, animalName) {
			return false
		}
	}
	return true
}

func ReserveAnimalByName(animalName string) error {
	var availID int = -1
	for i, listItem := range AvailableAnimals {
		if listItem.AnimalName == animalName {
			availID = i
		}
	}
	if availID == -1 {
		return errors.New("No animal by that name found!")
	}
	copy(AvailableAnimals[availID:], AvailableAnimals[availID+1:]) // Shift AvailableAnimals[i+1:] left one index.
	AvailableAnimals = AvailableAnimals[:len(AvailableAnimals)-1]  // Truncate slice.
	return nil
}

func ReleaseAnimalByName(animalName string) error {
	if !IsAnimalReserved(animalName) {
		return errors.New(fmt.Sprintf("That animal (%s) is not reserved!", animalName))
	}
	animalData := FindAnimal(animalName)
	AvailableAnimals = append(AvailableAnimals, animalData)
	return nil
}
