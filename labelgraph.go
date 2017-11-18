package main
import ("fmt";"os")

// raw label description to be read from JSON
type SrcLabel struct { 
	isa []string;
	examples []string;
	part_of []string;
	has []string;
	smaller_than []string;
	bigger_than []string;
	found_in []string;
	states []string;
	abstract bool;
	found_on []string;
	translations map[string]string;
}

type Void struct{}

type LabelPtrSet struct {
	items map[*Label]Void;
}

// compiled label, with complete links
type Label struct {
	name string;
	isa LabelPtrSet;
	initialized bool;
	abstract bool;
	examples LabelPtrSet;
	has LabelPtrSet;
	part_of LabelPtrSet;
	smaller_than LabelPtrSet;
	bigger_than LabelPtrSet;
	states []string;
	minDistFromRoot int;
	minDistFromLeaf int;
}

func (self *LabelPtrSet) Insert(ptr *Label){
	self.items[ptr]=Void{}
}
func (self *LabelPtrSet) Init(){
	self.items=make(map[*Label]Void)
}
func (self *LabelPtrSet) len() int{return len(self.items)}

func appendLabelPtrList(ls *[]*Label,l *Label){
	*ls = append(*ls, l)
}


var(g_srcLabels=map[string]SrcLabel{
	"person":{
		isa:[]string{"human"},
		examples:[]string{"man","woman","child","boy","girl","baby","police officer","soldier","workman","pedestrian","guard"},
		
	},
	"category":{
		abstract:true,
		examples:[]string{"military","urban","industrial","domestic","natural","man made","scientific","medical","law enforcement","professional","trade","aquatic","airborn","academic","educational","transport","mechanical","organic","professional"},
	},
	"fork lift truck":{
		isa:[]string{"wheeled powered vehicle","industrial"},
	},
	"human":{
		isa:[]string{"mammal"},
		has:[]string{"head","arm","leg","torso","neck"},
		states:[]string{"standing","walking","running","sitting","kneeling","reclining","sleeping"},
	},
	"soldier":{
		isa:[]string{"person","military","professional"},
	},
	"weapon":{
		isa:[]string{"military","tool"},
		examples:[]string{"firearm","combat knife","sword","rocket launcher","flame thrower","grenade launcher","crossbow","longbow","compound bow","arrow","crossbow bolt","spear","lance","mace","ball and chain","trident","pike"},
	},
	"machine":{
		isa:[]string{"man made"},
		examples:[]string{
			"vehicle","agricultural equipment","factory equipment","power tools","weapon","electrical equipment","electrical applicance","construction machinery","manufacturing tools",
		},
	},
	"electrical":{
		examples:[]string{"battery"},
	},
	"battery":{
		isa:[]string{"electrical"},
		examples:[]string{"AA battery","AAA battery","button cell","car battery","laptop battery","smartphone battery","rechargeable battery","disposable battery"},
	},
	"button cell":{
		examples:[]string{"LR44","CR2032","SR516","LR1154"},
	},
	"generic object":{
		examples:[]string{"barrel","cylinder","box","tray","wall","roof","bin","bottle","tub","bag","clothing","textile","sports equipment","mechanism","desktop object","household object","agricultural object","urban object","military","ornament","painting","photograph","container","cleaning tool","barrier","razor wire","barbed wire","spikes","peice of art","pylon","post","beam","bracket","shelter","electrical","water related object","tube","control","pedal","key","masking tape","desktop object","trash","tent","workshop object","stock"},
	},
	"control device":{
		isa:[]string{"generic object"},
		examples:[]string{"lever","dial","knob","switch","joystick","control pedal"},
	},
	"pedal":{
		examples:[]string{"control pedal","pedal (bicycle)"},
	},
	"water related object":{
		examples:[]string{"tap","basin","plug hole","sink","sink plug","pipe"},
	},
	"textile":{
		examples:[]string{"wool","silk","synthetic fabric","nylon","spandex","polyester fibre"},
	},
	"chickenwire":{
		isa:[]string{"generic object"},
	},
	"peice of art":{
		examples:[]string{"sculpture","painting","engraving"},
	},
	"urban feature":{
		examples:[]string{"road layout feature","road barrier","tunnel entrance","canal","towpath"},	
	},
	"road layout feature":{
		examples:[]string{
			"T junction","fork (road)","hairpin bend","cul-du-sac","dual carriage way","traffic island","round-a-bout","junction (road)","intersection (road)","flyover (road)","bypass (road)","cycle lane","bus lane","hard shoulder","central reservation","road bridge","road marking"},
	},
	"barrier":{
		examples:[]string{"fence","railing","wall","low wall","level crossing barrier"},
	},
	"fence":{
		examples:[]string{"wire fence","wooden fence","metal fence","picket fence","concrete fence","barbed wire fence","palisade","stockade fence","hurdle fence","wattle fence","hedgerow","live fencing","cactus fence","dry stone wall","welded wire mesh fence","brushwood fence","chain-link fencing","woven fence","temporary fencing"},
	},
	"cleaning tool":{
		isa:[]string{"tool"},
		examples:[]string{"brush","broom","cloth","dustpan","vacuum cleaner","mop","chamois","feather duster"},
	},
	"brush":{
		examples:[]string{"broom","bike cleaning brush","toothbrush","hairbrush"},
	},
	"household object":{
		examples:[]string{"furniture","kitchen appliance","kitchenware","ash tray","wall mirror","hand mirror","radiator","fan heater","storage heater","white goods","bathroom object"},
	},
	"bathroom object":{
		examples:[]string{"toothpaste tube","razor","shaving foam dispenser","electric razor","shampoo bottle","bar of soap","soap dispenser","toilet roll"},
	},
	"mirror":{
		isa:[]string{"generic object"},
		examples:[]string{"wing mirror","rear view mirror","security mirror","hand mirror","wall mirror","desk mirror"},
	},
	"domestic appliance":{
		translations:map[string]string{"british english":"white goods"},
		examples:[]string{"washing machine","refridgerator","half height refridgerator","tall refridgerator","freezer","mini refridgerator","tumble drier","oven"},
	},
	"sports equipment":{
		examples:[]string{"skis","ski pole","skateboard","football","tennis ball","shuttlecock","tennis raquet","badminton racket","hocket stick","cricket bat","baseball bat","snooker cue","roller scate"},
	},
	"personal transport":{
		isa:[]string{"generic object"},
		examples:[]string{"scooter","self balancing scooter","bicycle","skateboard"},
	},
	"personal item":{
		isa:[]string{"generic object"},
		examples:[]string{"clothing","eyewear","footwear","wristwatch","jewelry","bag","suitcase","smartphone","mobile phone"},
	},
	"clothing":{
		isa:[]string{"personal item"},
		examples:[]string{"jacket","trousers","skirt","jumper","dress","tracksuit","swimwear","hat","coat","winter coat","waterproof clothing","sportswear","footwear"},
	},
	"eyewear":{
		isa:[]string{"personal item"},
		examples:[]string{"sunglasses","reading glasses","spectacles","monacle","safety glasses","goggles","swimming goggles"},
	},
	"footwear":{
		examples:[]string{"shoes","flip flops","sandals","cycling shoes","clogs","slippers","trainers (footwear)"},
	},
	"timepeice":{
		examples:[]string{"sundial","clock","pendulum clock","digital clock","egg timer"},
	},
	"excercise equipment":{
		isa:[]string{"generic object"},
		examples:[]string{"dumbell","barbell","kettlebell","excercise bike","treadmill","rowing machine","weighted vest","parallel bars","pullup bar"},
	},
	"police box":{
		isa:[]string{"urban object"},
	},
	"telephone box":{
		isa:[]string{"urban object"},
		has:[]string{"telephone"},
	},
	"mammal":{
		examples:[]string{"giraffe","deer","bison","rodent","felinae","dog","wolf","hedgehog","anteater","primate","horse","donkey","oxen","sheep","cow"},
	},
	"primate":{
		examples:[]string{"human","gorilla","chimpanzee","monkey"},
	},
	"felinae":{
		examples:[]string{"cat","lion","tiger","cheetah","panther"},
	},
	"hat":{
		examples:[]string{"party hat","peaked cap","baseball cap","beanie","flat cap","mortar board","hard hat"},
	},
	"bag":{
		examples:[]string{"rucksack","sports bag","handbag","courier bag"},
	},
	"component":{
		examples:[]string{"room","building part","electronic component","vehicle component","bicycle component","mechanical component","car parts","aircraft component","weapon component","bodypart","lever","wings","wheel","trunk","handgrip","domestic fitting","corridor","hallway","metal component","handle","coin slot","keyhole","blade"},
	},
	"electronic component":{
		examples:[]string{"resistor","capacitor","LED","diode","breadboard","printed circuit board"},
	},
	"gargoyle":{
		isa:[]string{"building part"},
	},
	"mechanical component":{
		isa:[]string{"mechanical","component"},
		examples:[]string{"hydaulic ram","gearwheel","crankshaft","drive shaft","drive belt","conveyor belt","gearbox","turbine","spring","hinge"},
	},
	"metal component":{
		examples:[]string{"ankerbolt","bolt","machine screw","nut (metal)","nail (metal)","socket screw","stainless steel screw","setscrew","tek screw","threaded rod","throughbolt","tube connector","tube insert","washer","woodscrew"},
	},
	"room":{
		examples:[]string{"board room","office","atrium","domestic room"},
	},
	"trunk":{
		examples:[]string{"trunk (elephant)","trunk (tree)","trunk (car)"},
	},
	"building part":{
		examples:[]string{"door","window (building)","wall","buttress","archway","pillar","chimney","dome","escalator","elevator"},
	},
	"dome":{
		examples:[]string{"geodesic dome","onion dome","metal dome","stone dome"},
	},
	"arch":{
		isa:[]string{"building part"},
		examples:[]string{"pointed arch","round arch","parabolic arch","lancet arch","trefoil arch","horseshoe arch","three centred arch","ogee arch","tudor arch","inflex arch","reverse ogee arch","trefoil arch","shouldered flat arch","equilateral pointed arch"},
	},
	"window":{
		isa:[]string{"generic object"},
		examples:[]string{"window (building)","window (vehicle)"},
	},
	"window (building)":{
		isa:[]string{"building part"},
		examples:[]string{"stained glass window","glass window","lattice window","decorative window","casement window","awning window","skylight","pivot window","casement window"},
	},
	"coastal object":{
		isa:[]string{"generic object"},
		examples:[]string{"lighthouse","pier","harbour","jetty","beach","cliff","river estuary"},
	},
	"vehicle component":{
		examples:[]string{"land vehicle component","engine","cabin","turret","window (vehicle)"},
		
	},
	"window (vehicle)":{
		examples:[]string{"windscreen","passenger window","cockpit window","observation dome (vehicle)"},
	},
	"wheel":{
		examples:[]string{"wheel (bicycle)","wheel (tractor)","wheel (car)","castor wheel","wheel (mountain bike)","wheel (road bike)","mag wheel","wheel (truck)"},
	},
	"land vehicle component":{
		examples:[]string{"bonnet","windscreen","wheel","license plate","headlight","tail light","steering wheel","joystick","caterpillar tracks","hydraulic ram","exhaust pipe","wing mirror","license plate","indicator","differential gear","suspension","brake disk","tire","wheel hub"},
	},
	"weapon component":{
		examples:[]string{"muzzle","gun barrel","pistol grip", "stock (firearms)","sights","charging handle","gas tube","foregrip","picitany rail","laser sight","box magazine","stripper clip","ammunition belt","cartridge (firearm)","bullet","shotgun shell"},
	},
	"stock (firearms)":{
		examples:[]string{"solid stock","wooden stock","side folding stock","under folding stock","retractable sliding stock","skeletal stock","adjustable stock","M4 stock","bullpup stock","sniper stock"},
	},
	"aircraft component":{
		examples:[]string{"wing","control column","tail boom","tail rotor","tail fin","cockpit","aileron","propeller","jet engine","cabin","landing gear","rotor blades"},
	},
	"bicycle component":{
		examples:[]string{"derailleur","bicycle frame","handlebars (bicycle)","bicycle wheel","brake lever","gear lever","integrated shifter","saddle","mudguard","chain","chainset","casette (bicycle)","pedal (bicycle)","clipless pedal","disc brake (bicycle)"},
	},
	"handlebars (bicycle)":{
		examples:[]string{"flat bars","drop handlebars","tribars","bullhorn bars"},	
	},
	"manufacturing tools":{
		examples:[]string{"3d printer","CNC machine","lathe","press","injection moulding machine"},
	},
	"tool":{
		examples:[]string{"shovel","drill","hand drill","dentist drill","multitool","tweasers"},
	},
	"swiss army knife":{
		isa:[]string{"multitool"},
	},
	"hand tool":{
		has:[]string{"handle"},
		examples:[]string{"hammer","spanner","screwdriver","alan key","wrench","chisel","saw","mallet","crowbar","hacksaw","wood saw","shovel","spade","axe","rake"},
	},
	"workshop items":{
		isa:[]string{"generic object"},
		examples:[]string{"vice","lathe","clamp","spirit level","toolbox","drill bit","adjustable spanner","pliers","WD40","glue"},
	},
	"firearm":{
		isa:[]string{"weapon"},
		has:[]string{"gun barrel","stock","handgrip","sights","receiver","charging handle"},
		examples:[]string{"gun","canon","rifle","pistol","revolver","handgun","machine gun","automatic weapon"},
	},
	"rifle":{
		examples:[]string{"bolt action rifle","semi automatic rifle","automatic rifle","battle rifle","assault rifle","hunting rifle","carbine","sniper rifle"},
	},
	"magazine fed firearm":{
		isa:[]string{"firearm"},
		examples:[]string{"assault rifle"},
		has:[]string{"box magazine"},
	},
	"battle rifle":{
		isa:[]string{"firearm"},
		examples:[]string{"FN FAL","HK G3","M14","M1 Garand"},	
	},
	"assault rifle":{
		isa:[]string{"weapon","assault rifle","magazine fed firearm","automatic weapon","firearm"},
		examples:[]string{"M16 variant","g36","G3"},
		smaller_than:[]string{"machine gun"},
	},
	"kalashnikov assault rifle":{
		isa:[]string{"assault rifle"},
		examples:[]string{"AK47","AK47M","AK74","AK103"},
	},
	"M16 variant":{
		examples:[]string{"M4 carbine","M16A1","M16A2","M16A4","AR15"},
	},
	"bullpup assault rifle":{
		isa:[]string{"assault rifle"},
		examples:[]string{"IMI Tavor","IMI X95","SA80","Steyr AUG","FAMAS"},
	},
	"full length rifle":{
		isa:[]string{"rifle"},
		examples:[]string{"M16A2,AK47,AK74","FN FAL","HK G3"},
	},
	"carbine":{
		isa:[]string{"assault rifle","firearm"},
		examples:[]string{"M4","micro tavor","G36K","AK74SU"},
		smaller_than:[]string{"full length rifle"},
		bigger_than:[]string{"pistol"},
	},
	"SMG":{
		isa:[]string{"automatic weapon","firearm"},
		examples:[]string{"uzi","mac10","mp5"},
		smaller_than:[]string{"carbine"},
		bigger_than:[]string{"pistol"},
	},
	"military vehicle":{
		isa:[]string{"military","vehicle"},
		examples:[]string{"armoured car","tank (vehicle)","armoured personel carrier","MLRS","self propelled artillery","military jeep"},
	},
	"tank (vehicle)":{
		isa:[]string{"military vehicle"},
		has:[]string{"turret","gun","caterpillar tracks","hatch"},
	},
	"tank":{
		isa:[]string{"generic object"},
		examples:[]string{"tank (vehicle)","fuel tank","oxygen tank","liquid tank","container tank","gas tank","fish tank"},
	},
	"canon":{
		isa:[]string{"weapon"},
	},
	"vehicle":{
		isa:[]string{"machine"},
		examples:[]string{"aircraft","aquatic vehicle","bicycle","motorbike","semi trailer","caravan","trailer"},
	},
	"aquatic vehicle":{
		examples:[]string{
			"ship","boat","dinghy","jet ski","canoe","yacht",
			"sailing boat","ocean liner","cruise ship","oil tanker",
			"container ship","catamaran","lifeboat",
			"submarine","submersible","fishing boat","longboat",
			"lifeboat","rowing boat","narrowboat","barge",
			"military vessel"},
	},
	"military vessel":{
		examples:[]string{
			"destroyer","minesweeper","battleship","aircraft carrier",
			"nuclear submarine","gunboat","military submarine"},
	},
	"pose":{
		abstract:true,
		examples:[]string{"sitting","kneeling","standing","walking","reclining","prone","crawling","waving"},
	},
	"wheeled vehicle":{
		isa:[]string{"vehicle"},
	},
	"wheeled powered vehicle":{
		isa:[]string{"wheeled vehicle"},
		has:[]string{"wheel (car)","windscreen","license plate","headlight","wing mirror","tail light","indicator","bonnet"},
		examples:[]string{"car","truck","van","bus","semi truck","road sweeper","snow plough"},
	},
	"bicycle":{
		has:[]string{"bicycle component"},
		examples:[]string{"mountain bike","city bike","touring bike","BMX","road bike","triathalon bike"},
	},
	"aircraft":{
		isa:[]string{"vehicle","flying object"},
		examples:[]string{"helicopter","light aircraft","glider","autogyro","jet aircraft"},
		has:[]string{"landing gear"},
	},
	"rotorcraft":{
		isa:[]string{"aircraft"},
		examples:[]string{"quadcopter","helicopter","autogryo"},
	},
	"helicopter":{
		isa:[]string{"rotorcraft"},
		has:[]string{"rotor blades","tail rotor","tail boom","cockpit"},
		examples:[]string{"helicopter gunship","transport helicopter","search and rescue helicopter","air ambulance","police helicopter"},
	},
	"jet aircraft":{
		has:[]string{"jet engine","wings"},
		examples:[]string{"fighter jet","bomber","jet airliner","private jet"},
	},
	"fighter jet":{
		examples:[]string{"F16","F14","F15","Panavia Tornado","MIG","Mirage","Eurofighter Typhoon"},
	},
	"bird":{
		isa:[]string{"vertebrate","flying object"},
		has:[]string{"wings"},
		examples:[]string{"chicken","seagul","vulture","stalk","ostrich","duck","swan","bird of prey"},
	},
	"road marking":{
		isa:[]string{"road layout feature"},
		examples:[]string{"give way line","parking space","centreline","box junction","pedestrian crossing","striped reservation","lane divider","bus lane","cycle lane","stop line","chevron reservatoin","double white line","yellow lines (restricted parking)","double yellow lines (no parking)","disabled parking","mini roundabout","left turn lane","right turn lane","straight ahead lane","keep clear (road marking)","no entry (road marking)","slow (road marking)"},
	},
	"organism":{
		isa:[]string{"natural"},
		examples:[]string{"photosynthesier","chemosynthesier","consuming organism","plant","animal","fungus","edible"},
	},
	"mushroom":{
		isa:[]string{"fungus"},
		examples:[]string{
			"edible mushroom","hallucinogenic mushroom","poisonous mushroom",
		},
	},
	"edible mushroom":{
		isa:[]string{"mushroom","edible"},
		examples:[]string{
			"Boletus edulis","Cantharellus cibarius","Cantharellus tubaeformis","Clitocybe nuda","Cortinarius caperatus","Craterellus cornucopioides","Grifola frondosa","Hericium erinaceus","Hydnum repandum","Lactarius deliciosus","Pleurotus ostreatus","Tricholoma matsutake","truffle","white mushroom",
		},
	},
	"poisonous mushroom":{
		isa:[]string{"mushroom","poisonous"},
		examples:[]string{
			"Gyromitra esculenta","Morchella",
		},
	},
	"consuming organism":{
		isa:[]string{"organism"},
		examples:[]string{"herbivore","carnivore","predator","omnivore","predator"},
	},
	"bird of prey":{
		isa:[]string{"predator","carnivore","bird"},
		examples:[]string{"eagle","falcon","kestrel","hawk"},
	},
	"car":{
		isa:[]string{"vehicle"},
		has:[]string{"wheel","bonnet","license plate","windscreen","headlight","tail light","exhaust pipe"},
		examples:[]string{"hatchback","SUV","minivan","pickup truck","sedan","coupe","sportscar","convertible","open wheel racing car","gocart","racing car","touring car"},
		found_on:[]string{"road"},
	},
	"racing car":{
		examples:[]string{"open wheel racing car","rally car","dragster","Sports prototype"},
	},
	"open wheel racing car":{
		isa:[]string{"racing car"},
		examples:[]string{"formula one car","indycar"},
	},
	"animal":{
		isa:[]string{"organism"},
		examples:[]string{"land animal","marine animal","vertebrate","invertebrate","wild animal","domesticated animal"},

	},
	"construction machinery":{
		examples:[]string{"bulldozer","excavator","mini excavator","road roller","wrecking ball","pile driver","digger","crane","tower crane","cement mixer"},
	},
	"bulldozer":{
		has:[]string{"bucket (bulldozer)","caterpillar tracks"},
	},
	"machine gun":{
		isa:[]string{"firearm","automatic weapon"},
		examples:[]string{"belt fed machine gun","magazine fed machine gun"},
	},
	"magazine fed machine gun":{
		examples:[]string{"lewes gun","Browning Automatic Rifle","Bren gun"},
	},
	"belt fed machine gun":{
		isa:[]string{"machine gun"},
		has:[]string{"ammunition belt"},
		examples:[]string{"light machine gun","heavy machine gun","GPMG"},
	},
	"light machine gun":{
		examples:[]string{"RPK","bren light machine gun","browning automatic rifle","HK MG4","FN minimi","ultimax 100","stoner 63","L86 LSW","steyr AUG hbar","lewis gun"},
	},
	"heavy machine gun":{
		examples:[]string{"M2 Browning machine gun","gatling gun","maxim gun","vickers machine gun"},
	},
	"GPMG":{
		examples:[]string{"M60","PK machine gun","MG34","MG42","FN MAG"},
	},
	"excavator":{
		has:[]string{"bucket (excavator)","caterpillar tracks"},
	},
	"construction equipment parts":{
		isa:[]string{"mechanical component"},
		examples:[]string{"bucket (excavator)","bucket (bulldozer)"},
	},
	"mini excavator":{
		isa:[]string{"excavator"},
		smaller_than:[]string{"excavator"},
	},
	"quadruped":{
		isa:[]string{"animal"},
		has:[]string{"head","body","tail","leg"},
		examples:[]string{"dog","cat","horse","cow","sheep","lion","tiger","elephant"},
	},
	"door":{
		examples:[]string{"revolving door","double door","wooden door","glass door","metal door","hatch"},
	},
	"farm animal":{
		isa:[]string{"domesticated animal"},
		examples:[]string{"sheep","pig","cow","chicken","bull"},
		found_on:[]string{"farm"},
	},
	"bodypart":{
		examples:[]string{"ear","eye","eyebrow","cheek","neck","nose","mouth","chin","elbow","foot","hand","snout","tail","leg","arm","torso","body","shoulder","hips","knee","ankle","hoof","paw"},
	},
	"head":{
		isa:[]string{"bodypart"},
		has:[]string{"eye","ear","nose","mouth"},
	},	
	"arm":{
		isa:[]string{"bodypart"},
		has:[]string{"hand","elbow"},
	},	
	"leg":{
		isa:[]string{"bodypart"},
		has:[]string{"knee","foot"},
	},	
	"elephant":{
		isa:[]string{"quadruped"},
		has:[]string{"trunk (elephant)"},
	},
	"plant":{
		isa:[]string{"organism"},
		examples:[]string{"tree","bush","flower","hedge","shrub","vines"},
	},
	"rodent":{
		isa:[]string{"mammal"},
		examples:[]string{"mouse","rat","shrew"},
	},
	"fruit":{
		isa:[]string{"food"},
		part_of:[]string{"plant"},
		examples:[]string{"apple","banana","orange","pear","apricot","grapefruit","strawberry","raspberry","tangerine","dragonfruit","pineapple","strawberry","blackberry","peach"},
		states:[]string{"raw","sliced","cooked","diced","peeled"},
	},
	"food":{
		examples:[]string{"vegtable","fruit","nuts","meat","cereal","egg","salad","soup","sandwich","junk food","confectionary","hot dog","desert","pie","pastry","garnish","fast food","snack","meal","berry","beans","salad","stew","burger bun"},
	},
	"nuts":{
		examples:[]string{"wallnuts","hazelnuts","pecans","almonds","peanuts","cashew nuts","pistachio nuts"},
	},
	"desert":{
		examples:[]string{"cake","ice cream","blancmange","jelly","custard"},
	},
	"junk food":{
		examples:[]string{"hamburger","french fries","battered fish","potato chips (crisps)"},
	},
	"shopping mall":{
		isa:[]string{"building","retail area"},
	},
	// TODO .. is it a building, or a part, or an area????
	"shopping arcade":{
		isa:[]string{"building","retail area"},
	},
	"confectionary":{
		examples:[]string{"chocolate bar"},
	},
	"vegtable":{
		isa:[]string{"plant"},
		examples:[]string{"brocoli","peas","carrots","spinach","cellery","beansprouts","brussel sprouts","cauliflower","mushroom","peppers","courgette","leak","cabbage","onion","beans","tomato","lentils","tomato"},
	},
	"grains":{
		isa:[]string{"food"},
		examples:[]string{"rice","wheat","oats","barley"},
	},
	"rice":{
		examples:[]string{"white rice","brown rice","long grain rice","wild rice"},
	},
	"oats":{
		examples:[]string{"rolled oats"},
	},
	"furniture":{
		examples:[]string{"table","chair","bed","cupboard","desk","park bench","public bench","bench","workbench","dinner table","round table","shelf"},
	},
	"enclosure":{
		isa:[]string{"generic object"},
		examples:[]string{"cubicle","housing (mechanical)","casing","fence","electrical enclosure"},
	},
	"substance":{
		examples:[]string{"solid","liquid","emulsion","gas","organic substance","inorganic substance"},
	},
	"toxic substance":{
		isa:[]string{"substance"},
		examples:[]string{"radioactive waste","chlorine gas","acid","bleach"},
	},
	"water":{
		isa:[]string{"liquid"},
		examples:[]string{"fresh water","drinking water","mineral water","lake","freshwater","river water","puddle","fountain spray","polluted water","salt water","sea","river","waterfall"},
	},
	"agricultural tool":{
		isa:[]string{"tool","agricultural"},
		examples:[]string{"rake","sheers","plough","scythe"},
	},
	"agricultural equipment":{
		isa:[]string{"agricultural","mechanical"},
		examples:[]string{"tractor","combine harvester","crop duster"},
	},
	"tractor":{
		isa:[]string{"vehicle"},
		has:[]string{"wheel (tractor)","cabin","engine"},
	},
	"domestic room":{
		examples:[]string{"kitchen","dining room","bedroom","living room","study","store room","garage"},
		part_of:[]string{"house"},
	},
	"office building":{
		isa:[]string{"building"},
		has:[]string{"atrium","board room","office"},
	},
	"agricultural object":{
		examples:[]string{"hay bail","farm animal","agricultural equipment"},
	},
	"urban object":{
		examples:[]string{"street bin","wheeliebin","skip","lamp post","utility pole","electricity pylon","telegraph pole","traffic lights","sign post","traffic sign","radio tower","satelite dish","bottle bank","plant pot","hanging basket","flower pot","metal cover","drain pipe","drain","metal cover","manhole cover","roadworks","bollard","traffic cone","statue","monument","bus shelter","bus stop","pedestrian crossing","fountain","water feature"},
	},
	"material":{
		examples:[]string {"metal","plastic","vegetation","soil","stone","metal","plastic","textile","surface material"},
	},
	"metal":{
		examples:[]string{"metal tube","box section","sheet metal","wire","solder"},
	},
	"stroller":{
		isa:[]string{"urban object"},
		translations:map[string]string{"british english":"pushchair"},
	},
	"urban area":{
		isa:[]string{"area"},
		examples:[]string{"retail area","residential area","parking area","construction site","building site","financial district","town centre","park","suburb","shopping centre"},

	},
	"area":{
		examples:[]string{"urban area","industrial area","rural area","wilderness","desert","coastal area"},
	},
	"marque":{
		isa:[]string{"metalabel"},
		examples:[]string{"BMW","Ferrari","Maserati","Fiat","Ford","General Motors","Renault","Porsche","Mercedes"},
	},
	"metal object":{
		isa:[]string{"generic object"},
		examples:[]string{"bell","anchor","hook","chain"},
	},
	"basket":{
		isa:[]string{"container"},
		examples:[]string{"wicker basket","wire basket","metal basket","plastic basket"},
	},
	"pallet":{
		isa:[]string{"platform","industrial"},
		examples:[]string{"wooden pallet","plastic skid","steel pallet"},
	},
	"container":{
		isa:[]string{"generic object"},
		examples:[]string{"drum","barrel","cylinder","box","tray","basket","bag","shipping container"},
	},
	"traffic sign":{
		examples:[]string{"stop sign","no entry sign","no parking sign","speed limit","roadworks sign","hazard sign"},
	},
	"cutting tool":{
		isa:[]string{"tool"},
		examples:[]string{"knife","sword","craft knife","scalpel","stanley knife","boxcutter","machete","meat cleaver","circular saw","chainsaw","axe","wood axe","scissors","soldering iron"},
	},
	"instrument":{
		isa:[]string{"generic object"},
		examples:[]string{"musical instrument","medical instrument","electrical instrument"},
	},
	"electrical instrument":{
		examples:[]string{"oscilloscope","voltmeter"},
	},
	"musical instrument":{
		examples:[]string{"piano","grand piano","string instrument","wind instrument","electronic musical instrument","keyboard (musical instrument)","sound synthesiser"},
	},
	"wind instrument":{
		examples:[]string{
			"trumpet","trumpbone","flute","clarinet",
			"musical pipe","mouth organ","bagpipes",
		},
	},
	"string instrument":{
		examples:[]string{
			"violin","viola","guitar","electric guitar",
			"harp","harpsicord","banjo",
		},
	},
	"knife":{
		isa:[]string{"cutting tool"},
		has:[]string{"handle","blade"},
		examples:[]string{
			"pen knife","paper knife","kitchen knife","bread knife",
			"serated knife","combat knife","jungle knife","table knife",
			"dagger","survival knife","swiss army knife","butterfly knife",
			"flick knife"},
	},
	"bin":{
		isa:[]string{"container"},
		examples:[]string{"street bin","wheeliebin","wastepaper basket"},
	},
	"trash":{isa:[]string{"waste"}},
	"infrastructure":{
		examples:[]string{"road","bridge","dam","resevoir"},
	},
	"bridge":{
		examples:[]string{"footbridge","stone bridge","metal bridge","suspension bridge"},
	},
	"renewable energy system":{
		examples:[]string{"wind turbine","solar panel","solar concentrator","hydroelectric dam","geothermal power station","wave power device"},
	},
	"building":{
		examples:[]string{"church","house","tower block","factory","warehouse","cathederal","terminal building","train station","skyscraper","tower","tall building","stadium","log cabin","castle","fortress","lighthouse","wooden barn","barn","grainstore"},
	},
	"power tool":{
		isa:[]string{"tool"},
		examples:[]string{"chainsaw","powerdrill"},
	},
	"building complex":{
		examples:[]string{"power station","military base","industrial site","airport","harbour","dockyard","shipyard","university campus","housing estate"},
	},
	"arthropod":{
		isa:[]string{"animal"},
		examples:[]string{"insect","arachnid","crustacean"},
	},
	"invertebrate":{
		isa:[]string{"animal"},
		examples:[]string{"arthropod","mollusc","worm"},
	},
	"mollusc":{
		examples:[]string{"snail","slug","octopus","squid"},
	},
	"marine animal":{
		examples:[]string{"fish","octopus","squid","jellyfish","shrimp","lobster","crab"},
	},
	"vertebrate":{
		isa:[]string{"animal"},
		examples:[]string{"mammal","fish","reptile","amphibian"},
	},
	"lizard":{
		isa:[]string{"vertebrate"},
		examples:[]string{"snake","quadrupedal lizard","quadrupedal amphibian"},
	},
	"quadrupedal lizard":{
		isa:[]string{"lizard"},
		examples:[]string{"gecko","iguana","crocodile","alligator","dinosaur","chameleon","komodo dragon"},
	},
	"quadrupedal amphibian":{
		isa:[]string{"amphibian"},
		examples:[]string{"frog","salamander","toad"},
	},
	"tree":{
		isa:[]string{"plant"},
		examples:[]string{"palm tree","fern","oak tree","conifer","evergreen","small tree","large tree","tree stump"},
		has:[]string{"trunk (tree)","foilage"},
	},
	"bush":{
		isa:[]string{"plant"},
	},	
	"cutlery":{
		isa:[]string{"tool","kitchenware"},
		examples:[]string{"knife","fork","spoon","glass"},
	},
	"kitchen object":{
		isa:[]string{"generic object"},
		examples:[]string{"mug","plate","serving bowl","serving dish","saucepan","frying pan","pot","wok","steamer"},
	},
	"mug":{
		has:[]string{"handle"},
	},
	"kitchen appliance":{
		isa:[]string{"electrical applicance"},
		examples:[]string{"microwave oven","toaster","kettle","coffee machine","blender","electric cooker",},
		found_in:[]string{"kitchen",},
	},
	"domestic fittings":{
		examples:[]string{"electric socket","light switch","air vent","airconditioning unit","tap","toilet"},
	},
	"desktop object":{
		examples:[]string{
			"intray","pen holder","stapler","drawing pins",
			"paper clips","pen","desklamp","desktop PC",
			"sellotape dispenser","sellotape","hole punch",
			"file","ring binder","paper tray","pencil sharpener",
			"eraser","paperweight","notepad","envelope",
		},
	},
	"electrical applicance":{
		examples:[]string{"kitchen applicance","consumer electronics","lamp","desk lamp","light bulb","ceiling light","lantern","security camera","film projector"},
	},
	"consumer electronics":{
		isa:[]string{"electrical applicance"},
		examples:[]string{"TV","monitor","PC","laptop","tablet computer","smartphone","telephone","radio","game console","sound system","speakers","network switch","network hub","camera","cam corder","remote control handset","electric torch","3d printer"},
	},
	"mounted object":{
		isa:[]string{"generic object"},
		examples:[]string{"ceiling mounted","wall mounted","ground mounted"},
	},
	"lighting":{
		isa:[]string{"generic object"},
		examples:[]string{"candle","light bulb","flourescent light","LED light","torch","electric torch","burning torch","lantern","lamp","gas lamp","floodlight"},
	},
	"chandelier":{
		isa:[]string{"ornament","light fitting","ceiling mounted"},
	},
	"computer perhipheral":{
		isa:[]string{"consumer electronics"},		
		examples:[]string{"computer mouse","computer keyboard","joystick","gamepad","webcam","microphone"},
	},
	"TV":{
		examples:[]string{"flatscreen TV","LCD TV","plasma TV","LED TV","OLED TV","curved TV","CRT TV"},
	},
	"geographic feature":{
		examples:[]string{"mountain","hill","coastline","volcano","plain","valley","cave","forest"},
	},
	"surface material":{
		examples:[]string{"fur","feathers","wood","plastic","stone","sand","dirt","mud","soil","vegetation","grass","tiles","paving stones","bricks","concrete","corrugated metal","metal","rusted metal","plastic sheets","rubber","foilage","carpet","rug","porcelain","pottery","perspex","chipboard","paint","ceramic","stone","rock","building material","crystal","dolomite"},
	},
	"stone":{
		examples:[]string{"granite","limestone","sandstone","marble","ingeous rock","sedimentary rock","metamorphic rock"},
	},
	"grass":{
		isa:[]string{"vegetation","plant"},
		examples:[]string{"dry grass","sparse grass","thick grass","long grass","cut grass","wild grass"},
	},
	"vegetation":{
		isa:[]string{"plant"},
	},
	"ground":{
		examples:[]string{"soil","grass","park","lawn","field","sidewalk","pavement","road","runway","path","footpath"},
	},
	"road":{
		examples:[]string{"cobbled road","tarmac road","brick road","dirt road","brick road"},
	},
})


// ?! c++ address of member is useful for this, how to do?
// generalise leaf/root tracing 'isa'/'examples'


func setMinInt(p *int,x int){
	if x<*p {*p=x}
}
func setMaxInt(p *int,x int){
	if x>*p {*p=x}
}


func createLabel(n string) *Label{
	l:=&Label{name:n, minDistFromRoot:0xffff,minDistFromLeaf:0xffff}
	// todo - can Go avoid this? - c++ constructors
	l.initialized=true;
	l.isa.Init();
	l.examples.Init();
	l.has.Init();
	l.part_of.Init();
	l.has.Init();
	l.bigger_than.Init();
	l.smaller_than.Init();
	l.abstract=false;
	
	return l
}

type LabelGraph struct{
	all map[string]*Label;
	orphans LabelPtrSet; // no 'isa' or 'examples'
	roots LabelPtrSet; // no 'isa'
	leaves LabelPtrSet; // no 'examples'
	middle LabelPtrSet; // both 'isa' and 'examples'
}

func (self LabelGraph) CreateOrFindLabel(newname string) *Label{
	if lbl,ok:=self.all[newname];ok {return lbl;}
	newlbl:=createLabel(newname);
	self.all[newname]=newlbl;
	return newlbl;
}

func (self *Label) AddExample(other *Label){
	if (self==other) {return;}	// something wrong!
	self.examples.Insert(other);
	other.isa.Insert(self);
}
func (self *Label) AddPart(other *Label){
	if (self==other) {return;}	// something wrong!
	self.has.Insert(other);
	other.part_of.Insert(self);
}

func makeLabelGraph(srcLabels map[string]SrcLabel) *LabelGraph{

	var labels=make(map[string]*Label);

	findOrMakeLabel:=func (n string) *Label{
		if l,ok:=labels[n]; ok {
			return l;
		} else {
			l=createLabel(n);
			labels[n]=l;
			return l
		}
	}
	for name,src:= range srcLabels {
		this_label:=findOrMakeLabel(name)
		// TODO does go have field pointers or
		// any other means to reduce the cut-paste here..
		
		// "isa" and "examples" are reciprocated:-
		for _,isa_name:= range src.isa {
			findOrMakeLabel(isa_name).AddExample(this_label);
		}
		for _,ex:= range src.examples{
			this_label.AddExample(findOrMakeLabel(ex));
		}
		// "has" and "partof" are reciprocated
		for _,has:= range src.has{
			this_label.AddPart(findOrMakeLabel(has));
		}
		for _,p:= range src.part_of{
			findOrMakeLabel(p).AddPart(this_label);
		}

		// "bigger than" and "smaller than" are reciprocated
		for _,it:= range src.smaller_than{
			x:=findOrMakeLabel(it)
			x.smaller_than.Insert(this_label);
			this_label.bigger_than.Insert(x);
		}
		for _,it:= range src.bigger_than{
			x:=findOrMakeLabel(it)
			x.bigger_than.Insert(this_label);
			this_label.smaller_than.Insert(x);
		}
	}
	// 'orphans'
	// collect them under 'uncategorized objects'

	

	// final collection
	l:=&LabelGraph{all:labels};
	l.orphans.Init();
	l.roots.Init();
	l.middle.Init();
	l.leaves.Init();
	for _,x := range l.all{
		num_isa:=x.isa.len();
		num_examples:=x.examples.len();
		if num_isa==0 && num_examples==0 {
			l.orphans.Insert(x);
			l.CreateOrFindLabel("uncategorized item").AddExample(x);
		} else if num_isa!=0 && num_examples!=0 {
			l.middle.Insert(x);
		} else if num_isa==0 {
			l.roots.Insert(x);		
		} else if num_examples==0 {
			l.leaves.Insert(x);
		} else {
			fmt.Printf("fail!\n");
			os.Exit(0)
		}
	}
	
	return l;
}

	// Show results:-
	// TODO formalise this as actual JSON

func printContent(n string,xs LabelPtrSet,postfix string){
	if xs.len()==0 {return}
	fmt.Printf("\t\t\"%s\":[",n);
	i:=len(xs.items);
	for x,_:=range xs.items{
		fmt.Printf("\"%v\"",x.name);
		if i>1{fmt.Printf(",")}
		i-=1;
	}
		
	fmt.Printf("]%s\n",postfix);
}

func (self *LabelGraph) DumpJSON(verbose bool){

	fmt.Printf("{\n ");
	for name,label :=range self.all {
		fmt.Printf("\t\"%v\":{\n ",name);

		if (verbose){
			fmt.Printf("\t\tminDistFromRoot:%v\n", label.minDistFromRoot);
			fmt.Printf("\t\tminDistFromLeaf:%v\n", label.minDistFromLeaf);
		}
		printContent("isa",label.isa,",");
		printContent("examples",label.examples,",");
		printContent("has",label.has,",");
		printContent("part_of",label.part_of,"");
		fmt.Printf("\t},\n")
	}
	fmt.Printf("}\n ");
}
func (self LabelGraph) DumpInfo(){

	fmt.Printf("{\n ");
	fmt.Printf("\"labelList stats\":{\"total\":%v, \"roots(metalabels)\":%v, \"middle(labels)\":%v \"leaf examples\":%v,\"orphans\":%v},\n",
		len(self.all),
		self.roots.len(), self.middle.len(),self.leaves.len(), self.orphans.len());
	printContent("leaves",self.leaves,",");
	printContent("middle",self.middle,",");
	printContent("roots",self.roots,",");
	printContent("orphans",self.orphans,"");
	
	fmt.Printf("}\n ");
}

func (l LabelGraph) Get(n string) *Label{
	return l.all[n]
}

func main() {

	// compile labels into a map for access by string, with links

	labelGraph := makeLabelGraph(g_srcLabels);
	labelGraph.DumpJSON(false);
	
	labelGraph.DumpInfo();

}




