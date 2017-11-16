package main
import ("fmt";"os")

// raw label description to be read from JSON
type SrcLabel struct { 
	name string;
	isa []string;
	examples []string;
	part_of []string;
	has []string;
	smaller_than []string;
	bigger_than []string;
	found_in []string;
	states []string;
	
}

// compiled label, with complete links
type Label struct {
	name string;
	isa []*Label;
	examples []*Label;
	has []*Label;
	part_of []*Label;
	smaller_than []*Label;
	bigger_than []*Label;
	states []string;
	minDistFromRoot int;
	minDistFromLeaf int;
}

func appendLabelPtrList(ls *[]*Label,l *Label){
	*ls = append(*ls, l)
}

var(g_srcLabels=[]SrcLabel{
	{
		name:"person",
		isa:[]string{"human"},
		examples:[]string{"man","woman","child","boy","girl","baby","police officer","soldier","workman","pedestrian","guard"},
	},
	{
		name:"metalabel",
		examples:[]string{"military","urban","industrial","domestic","natural","artificial","scientific","medical","law enforcement","professional","trade","aquatic","airborn","academic","educational","transport","mechanical","organic"},
	},
	{
		name:"human",
		isa:[]string{"mammal"},
		has:[]string{"head","arm","leg","torso","neck"},
		states:[]string{"standing","walking","running","sitting","kneeling","reclining","sleeping"},
	},
	{
		name:"soldier",
		isa:[]string{"person","military"},
	},
	{
		name:"weapon",
		isa:[]string{"military"},
		examples:[]string{"firearm","combat knife","sword","rocket launcher","flame thrower","grenade launcher",},
	},
	{
		name:"machine",
		examples:[]string{
			"vehicle","agricultural equipment","factory equipment","power tools","weapon","electrical equipment","electrical applicance","construction machinery","manufacturing tools",
		},
	},
	{
		name:"electrical",
		examples:[]string{"battery"},
	},
	{
		name:"battery",
		examples:[]string{"AA batter","AAA battery","button cell"},
	},
	{
		name:"button cell",
		examples:[]string{"LR44","CR2032","SR516","LR1154"},
	},
	{
		name:"generic object",
		examples:[]string{"barrel","cylinder","box","tray","wall","roof","bin","bottle","tub","bag","clothing","fabrics","sports equipment","mechanism","desktop object","household object","agricultural object","urban object","military","ornament","painting","photograph","container","cleaning tool","barrier","razor wire","barbed wire","spikes","peice of art","pylon","post","beam","bracket","shelter","electrical"},
	},
	{
		name:"generic object",
		examples:[]string{"tent"},
	},
	{
		name:"peice of art",
		examples:[]string{"sculpture","painting","engraving"},
	},
	{
		name:"urban feature",
		examples:[]string{"T junction","fork (road)","hairpin bend","cul-du-sac","dual carriage way","traffic island","round-a-bout","junction (road)","intersection (road)","flyover (road)","bypass (road)"},	
	},
	{
		name:"barrier",
		examples:[]string{"fence","railing","wall","low wall","level crossing barrier"},
	},
	{
		name:"cleaning tool",
		examples:[]string{"brush","broom","cloth","dustpan","vacuum cleaner","mop","chamois","feather duster"},
	},
	{
		name:"brush",
		examples:[]string{"broom","bike cleaning brush","toothbrush","hairbrush"},
	},
	{
		name:"household object",
		examples:[]string{"furniture","kitchen appliance","kitchenware","ash tray","mirror"},
	},
	{
		name:"sports equipment",
		examples:[]string{"skis","ski pole","skateboard","football","tennis ball","shuttlecock","tennis raquet","badminton racket","hocket stick","cricket bat","baseball bat","snooker cue","roller scate"},
	},
	{
		name:"personal transport",
		isa:[]string{"generic object"},
		examples:[]string{"scooter","self balancing scooter","bicycle","skateboard"},
	},
	{
		name:"clothing",
		examples:[]string{"jacket","trousers","skirt","jumper","dress","tracksuit","swimwear","hat","coat","winter coat","waterproof clothing","sportswear","footwear"},
	},
	{
		name:"footwear",
		examples:[]string{"shoes","flip flops","sandals","cycling shoes","clogs","slippers","trainers (footwear)"},
	},
	{
		name:"excercise equipment",
		isa:[]string{"generic object"},
		examples:[]string{"dumbell","barbell","kettlebell","excercise bike","treadmill","rowing machine","weighted vest","parallel bars","pullup bar"},
	},
	{
		name:"police box",
		isa:[]string{"urban object"},
	},
	{
		name:"telephone box",
		isa:[]string{"urban object"},
	},
	{
		name:"hat",
		examples:[]string{"party hat","peaked cap","baseball cap","beanie","flat cap","mortar board","hard hat"},
	},
	{
		name:"bag",
		examples:[]string{"rucksack","sports bag","handbag","courier bag"},
	},
	{
		name:"component",
		examples:[]string{"room","building part","electronic component","vehicle component","bicycle component","mechanical component","car parts","aircraft component","weapon component","bodypart","lever","wings","wheel","trunk","handgrip","domestic fitting","corridor","hallway","metal component"},
	},
	{
		name:"metal component",
		examples:[]string{"nut (metal)","bolt","nail (metal)"},
	},
	{
		name:"gargoyle",
		isa:[]string{"building part"},
	},
	{
		name:"mechanical component",
		isa:[]string{"mechanical","component"},
		examples:[]string{"hydaulic ram","gearwheel","crankshaft","drive shaft","drive belt","conveyor belt","gearbox","turbine","spring","hinge"},
	},
	{
		name:"room",
		examples:[]string{"board room","office","atrium","domestic room"},
	},
	{
		name:"trunk",
		examples:[]string{"trunk (elephant)","trunk (tree)","trunk (car)"},
	},
	{
		name:"building part",
		examples:[]string{"door","window (building)","wall","buttress","archway","pillar","chimney"},
	},
	{	name:"arch",
		isa:[]string{"building part"},
		examples:[]string{"pointed arch","round arch","parabolic arch","lancet arch","trefoil arch","horseshoe arch","three centred arch","ogee arch","tudor arch","inflex arch","reverse ogee arch","trefoil arch","shouldered flat arch","equilateral pointed arch"},
	},
	{
		name:"window",
		isa:[]string{"generic object"},
		examples:[]string{"window (building)","window (vehicle)"},
	},
	{
		name:"window (building)",
		isa:[]string{"building part"},
		examples:[]string{"stained glass window","glass window","lattice window","decorative window","casement window","awning window","skylight","pivot window","casement window"},
	},
	{
		name:"vehicle component",
		examples:[]string{"land vehicle component","engine","cabin","turret","window (vehicle)"},
		
	},
	{
		name:"window (vehicle)",
		examples:[]string{"windscreen","passenger window","cockpit window","observation dome (vehicle)"},
	},
	{
		name:"wheel",
		examples:[]string{"wheel (bicycle)","wheel (tractor)","wheel (car)","castor wheel"},
	},
	{
		name:"land vehicle component",
		examples:[]string{"bonnet","windscreen","wheel","license plate","headlight","tail light","steering wheel","joystick","caterpillar tracks","hydraulic ram","exhaust pipe","wing mirror","license plate","indicator","differential gear","suspension","brake disk","tire","wheel hub"},
	},
	{
		name:"weapon component",
		examples:[]string{"muzzle","gun barrel","pistol grip", "stock","sights","charging handle","gas tube","foregrip","picitany rail","laser sight","box magazine","stripper clip","ammunition belt"},
	},
	{
		name:"aircraft component",
		examples:[]string{"wing","control column","tail boom","tail rotor","tail fin","cockpit","aileron","propeller","jet engine","cabin","landing gear","rotor blades"},
	},
	{
		name:"bicycle component",
		examples:[]string{"derailleur","bicycle frame","handlebars (bicycle)","bicycle wheel","brake lever","gear lever","integrated shifter","saddle","mudguard","chain","chainset","casette (bicycle)","pedal","clipless pedal","disc brake (bicycle)"},
	},
	{
		name:"handlebars (bicycle)",
		examples:[]string{"flat bars","drop handlebars","tribars","bullhorn bars"},	
	},
	{
		name:"manufacturing tools",
		examples:[]string{"3d printer","CNC machine","lathe","press","injection moulding machine"},
	},
	{
		name:"tool",
		examples:[]string{"shovel","drill","hand drill","dentist drill","multitool","swiss army knife","tweasers"},
	},
	{
		name:"hand tool",
		examples:[]string{"hammer","spanner","screwdriver","chisel","saw","mallet","crowbar","hacksaw","wood saw","shovel","spade","axe"},
	},
	{
		name:"workshop items",
		isa:[]string{"generic object"},
		examples:[]string{"vice","lathe","clamp","spirit level","toolbox","drill bit","adjustable spanner","pliers"},
	},
	{
		name:"firearm",
		isa:[]string{"weapon"},
		has:[]string{"gun barrel","stock","handgrip","sights"},
		examples:[]string{"gun","canon","rifle","pistol","revolver","handgun","machine gun","automatic weapon"},
	},
	{
		name:"magazine fed firearm",
		isa:[]string{"firearm"},
		examples:[]string{"assault rifle"},
		has:[]string{"box magazine"},
	},
	{
		name:"battle rifle",
		isa:[]string{"firearm"},
		examples:[]string{"FN FAL","HK G3","M14","M1 Garand"},	
	},
	{
		name:"assault rifle",
		isa:[]string{"weapon","assault rifle","magazine fed firearm","automatic weapon","firearm"},
		examples:[]string{"M16 variant","g36","G3"},
		smaller_than:[]string{"machine gun"},
	},
	{
		name:"kalashnikov assault rifle",
		isa:[]string{"assault rifle"},
		examples:[]string{"AK47","AK47M","AK74","AK103"},
	},
	{
		name:"M16 variant",
		examples:[]string{"M4 carbine","M16A1","M16A2","M16A4","AR15"},
	},
	{
		name:"bullpup assault rifle",
		isa:[]string{"assault rifle"},
		examples:[]string{"IMI Tavor","IMI X95","SA80","Steyr AUG"},
	},
	{
		name:"full length rifle",
		isa:[]string{"rifle"},
		examples:[]string{"M16A2,AK47,AK74","FN FAL","HK G3"},
	},
	{
		name:"carbine",
		isa:[]string{"assault rifle","firearm"},
		examples:[]string{"m4","micro tavor","g36k","ak74su"},
		smaller_than:[]string{"full length rifle"},
		bigger_than:[]string{"pistol"},
	},
	{
		name:"SMG",
		isa:[]string{"automatic weapon","firearm"},
		examples:[]string{"uzi","mac10","mp5"},
		smaller_than:[]string{"carbine"},
		bigger_than:[]string{"pistol"},
	},
	{
		name:"tank",
		isa:[]string{"vehicle","military"},
		has:[]string{"turret","gun","caterpillar tracks"},
	},
	{
		name:"canon",
		isa:[]string{"weapon"},
	},
	{
		name:"vehicle",
		isa:[]string{"machine"},
		examples:[]string{"aircraft","ship","boat","bicycle","motorbike","semi trailer","caravan","trailer"},
	},
	{
		name:"pose",
		isa:[]string{"metalabel"},
		examples:[]string{"sitting","kneeling","standing","walking","reclining","prone","crawling","waving"},
	},
	{
		name:"wheeled powered vehicle",
		isa:[]string{"vehicle"},
		has:[]string{"wheel (car)","windscreen","license plate","headlight","wing mirror","tail light","indicator","bonnet"},
		examples:[]string{"car","truck","van","bus","semi truck"},
	},
	{
		name:"bicycle",
		has:[]string{"bicycle component"},
		examples:[]string{"mountain bike","city bike","touring bike","BMX","road bike","triathalon bike"},
	},
	{
		name:"aircraft",
		isa:[]string{"vehicle","flying object"},
		examples:[]string{"helicopter","light aircraft","glider","autogyro","jet aircraft"},
		has:[]string{"landing gear"},
	},
	{
		name:"rotorcraft",
		isa:[]string{"aircraft"},
		examples:[]string{"quadcopter","helicopter","autogryo"},
	},
	{
		name:"helicopter",
		isa:[]string{"rotorcraft"},
		has:[]string{"rotor blades","tail rotor","tail boom","cockpit"},
		examples:[]string{"helicopter gunship","transport helicopter","search and rescue helicopter","air ambulance","police helicopter"},
	},
	{
		name:"jet aircraft",
		has:[]string{"jet engine","wings"},
		examples:[]string{"fighter jet","bomber","jet airliner","private jet"},
	},
	{
		name:"fighter jet",
		examples:[]string{"F16","F14","F15","Panavia Tornado","MIG","Mirage","Eurofighter Typhoon"},
	},
	{
		name:"bird",
		isa:[]string{"vertebrate","flying object"},
		has:[]string{"wings"},
		examples:[]string{"chicken","seagul","vulture","stalk","ostrich","duck","swan","bird of prey"},
	},
	{
		name:"organism",
		isa:[]string{"natural"},
		examples:[]string{"photosynthesier","chemosynthesier","consuming organism","plant","animal","fungus","edible"},
	},
	{
		name:"mushroom",
		isa:[]string{"fungus"},
		examples:[]string{
			"edible mushroom","hallucinogenic mushroom","poisonous mushroom",
		},
	},
	{
		name:"edible mushroom",
		isa:[]string{"mushroom","edible"},
		examples:[]string{
			"Boletus edulis","Cantharellus cibarius","Cantharellus tubaeformis","Clitocybe nuda","Cortinarius caperatus","Craterellus cornucopioides","Grifola frondosa","Hericium erinaceus","Hydnum repandum","Lactarius deliciosus","Pleurotus ostreatus","Tricholoma matsutake","truffle","white mushroom",
		},
	},
	{
		name:"poisonous mushroom",
		isa:[]string{"mushroom","poisonous"},
		examples:[]string{
			"Gyromitra esculenta","Morchella",
		},
	},
	{
		name:"consuming organism",
		isa:[]string{"organism"},
		examples:[]string{"herbivore","carnivore","predator","omnivore","predator"},
	},
	{
		name:"bird of prey",
		isa:[]string{"predator","carnivore"},
		examples:[]string{"eagle","falcon"},
	},
	{
		name:"car",
		isa:[]string{"vehicle"},
		has:[]string{"wheel","bonnet","license plate","windscreen","headlight","tail light","exhaust pipe"},
		examples:[]string{"hatchback","SUV","minivan","pickup truck","sedan","coupe","sportscar","convertible"},
	},	
	{
		name:"animal",
		isa:[]string{"organism"},
		examples:[]string{"land animal","marine animal","vertebrate","invertebrate"},
	},
	{
		name:"construction machinery",
		examples:[]string{"bulldozer","excavator","mini excavator","road roller","wrecking ball","pile driver","digger","crane","tower crane"},
	},
	{
		name:"bulldozer",
		has:[]string{"bucket (bulldozer)","caterpillar tracks"},
	},
	{
		name:"machine gun",
		isa:[]string{"firearm","automatic weapon"},
	},
	{
		name:"belt fed machine gun",
		isa:[]string{"machine gun"},
		has:[]string{"ammunition belt"},
		examples:[]string{"light machine gun","heavy machine gun","GPMG"},
	},
	{
		name:"light machine gun",
		examples:[]string{"RPK","bren light machine gun","browning automatic rifle","HK MG4","FN minimi","ultimax 100","stoner 63","L86 LSW","steyr AUG hbar","lewis gun"},
	},
	{
		name:"heavy machine gun",
		examples:[]string{"M2 Browning machine gun","gatling gun","maxim gun","vickers machine gun"},
	},
	{
		name:"GPMG",
		examples:[]string{"M60","PK machine gun","MG34","MG42","FN MAG"},
	},
	{
		name:"excavator",
		has:[]string{"bucket (excavator)","caterpillar tracks"},
	},
	{
		name:"construction equipment parts",
		isa:[]string{"mechanical component"},
		examples:[]string{"bucket (excavator)","bucket (bulldozer)"},
	},
	{
		name:"mini excavator",
		isa:[]string{"excavator"},
		smaller_than:[]string{"excavator"},
	},
	{
		name:"quadruped",
		isa:[]string{"animal"},
		has:[]string{"head","body","tail","leg"},
		examples:[]string{"dog","cat","horse","cow","sheep","lion","tiger","elephant"},
	},
	{
		name:"animal",
		examples:[]string{"wild animal","domesticated animal"},
	},
	{
		name:"farm animal",
		isa:[]string{"domesticated animal"},
		examples:[]string{"sheep","pig","cow","chicken","bull"},
	},
	{
		name:"bodypart",
		examples:[]string{"ear","eye","eyebrow","cheek","neck","nose","mouth","chin","elbow","foot","hand","snout","tail","leg","arm","torso","body","shoulder","hips","knee","ankle"},
	},
	{	name:"head",
		isa:[]string{"bodypart"},
		has:[]string{"eye","ear","nose","mouth"},
	},	
	{	name:"arm",
		isa:[]string{"bodypart"},
		has:[]string{"hand","elbow"},
	},	
	{	name:"leg",
		isa:[]string{"bodypart"},
		has:[]string{"knee","foot"},
	},	
	{	name:"elephant",
		isa:[]string{"quadruped"},
		has:[]string{"trunk (elephant)"},
	},
	{
		name:"plant",
		isa:[]string{"organism"},
		examples:[]string{"tree","bush","flower","hedge","shrub","vines"},
	},
	{
		name:"rodent",
		isa:[]string{"mammal"},
		examples:[]string{"mouse","rat","shrew"},
	},
	{
		name:"fruit",
		isa:[]string{"food"},
		part_of:[]string{"plant"},
	},
	{
		name:"vegtable",
		isa:[]string{"plant"},
	},
	{	name:"food",
		examples:[]string{"vegtable","fruit","nuts","meat","cereal","egg","salad","soup","sandwich","junk food","confectionary","hot dog","desert","pie","pastry","garnish","fast food","snack","meal"},
	},
	{
		name:"nuts",
		examples:[]string{"wallnuts","hazelnuts","pecans","almonds","peanuts","cashew nuts","pistachio nuts"},
	},
	{
		name:"desert",
		examples:[]string{"cake","ice cream","blancmange","jelly","custard"},
	},
	{
		name:"junk food",
		examples:[]string{"hamburger","french fries","battered fish","potato chips (crisps)"},
	},
	{
		name:"shopping mall",
		isa:[]string{"building"},
	},
	// TODO .. is it a building, or a part, or an area????
	{
		name:"shopping arcade",
		isa:[]string{"building"},
	},
	{
		name:"confectionary",
		examples:[]string{"chocolate bar"},
	},
	{	name:"vegtable",
		examples:[]string{"brocoli","peas","carrots","spinach","cellery","beansprouts","brussel sprouts","cauliflower","mushroom","peppers","courgette","leak","cabbage","onion","beans","tomato","lentils","tomato"},
	},
	{
		name:"grains",
		isa:[]string{"food"},
		examples:[]string{"rice","wheat","oats","barley"},
	},
	{
		name:"rice",
		examples:[]string{"white rice","brown rice","long grain rice","wild rice"},
	},
	{
		name:"oats",
		examples:[]string{"rolled oats"},
	},
	{	name:"furniture",
		examples:[]string{"table","chair","bed","cupboard","desk","park bench","public bench","bench","workbench","dinner table","round table","shelf"},
	},
	{
		name:"enclosure",
		isa:[]string{"metalabel"},
		examples:[]string{"cubicle","cell","housing","casing","fence"},
	},
	{
		name:"toxic substance",
		isa:[]string{"substance"},
		examples:[]string{"radioactive waste","chlorine gas","acid","bleach"},
	},
	{
		name:"water",
		isa:[]string{"substance"},
		examples:[]string{"fresh water","drinking water","mineral water","lake","freshwater","salt water","sea","river","waterfall"},
	},
	{
		name:"agricultural tool",
		isa:[]string{"tool","agricultural"},
		examples:[]string{"rake","sheers","plough","scythe"},
	},
	{	name:"agricultural equipment",
		isa:[]string{"agricultural","mechanical"},
		examples:[]string{"tractor","combine harvester","crop duster"},
	},
	{	name:"tractor",
		isa:[]string{"vehicle"},
		has:[]string{"wheel (tractor)","cabin","engine"},
	},
	{	name:"domestic room",
		examples:[]string{"kitchen","dining room","bedroom","living room","study","store room","garage"},
		part_of:[]string{"house"},
	},
	{	name:"office building",
		isa:[]string{"building"},
		has:[]string{"atrium","board room","office"},
	},
	{
		name:"agricultural object",
		examples:[]string{"hay bail","farm animal","agricultural equipment"},
	},
	{	name:"urban object",
		examples:[]string{"street bin","wheeliebin","skip","lamp post","utility pole","electricity pylon","telegraph pole","traffic lights","sign post","traffic sign","radio tower","satelite dish","bottle bank","plant pot","hanging basket","flower pot","metal cover","drain pipe","roadworks","bollard","traffic cone","statue","monument","bus shelter","bus stop","pedestrian crossing","fountain","water feature"},
	},
	{
		name:"basket",
		isa:[]string{"container"},
		examples:[]string{"wicker basket","wire basket","metal basket","plastic basket"},
	},
	{
		name:"pallet",
		isa:[]string{"platform"},
		examples:[]string{"wooden pallet","plastic skid","steel pallet"},
	},
	{
		name:"container",
		isa:[]string{"generic object"},
		examples:[]string{"drum","barrel","cylinder","box","tray","basket","bag","shipping container"},
	},
	{
		name:"traffic sign",
		examples:[]string{"stop sign","no entry sign","no parking sign","speed limit","roadworks sign"},
	},
	{
		name:"cutting tool",
		isa:[]string{"tool"},
		examples:[]string{"knife","sword","craft knife","scalpel","stanley knife","boxcutter","machete","meat cleaver","circular saw","chainsaw","axe","wood axe","pen knife"},
	},
	{
		name:"bin",
		examples:[]string{"street bin","wheeliebin","wastepaper basket"},
	},
	{
		name:"infrastructure",
		examples:[]string{"road","bridge","dam","resevoir"},
	},
	{
		name:"renewable energy system",
		examples:[]string{"wind turbine","solar panel","solar concentrator","hydroelectric dam","geothermal power station","wave power device"},
	},
	{	name:"building",
		examples:[]string{"church","house","tower block","factory","warehouse","cathederal","terminal building","train station","skyscraper","tower","tall building","stadium","log cabin","castle","fortress"},
	},
	{
		name:"urban area",
		examples:[]string{"building site","financial district","town centre","park","suburb","residential area","shopping centre"},
	},
	{
		name:"power tool",
		isa:[]string{"tool"},
		examples:[]string{"chainsaw","powerdrill"},
	},
	{
		name:"complex",
		examples:[]string{"power station","military base","industrial site","airport","harbour","docks","shipyard"},
	},
	{
		name:"arthropod",
		isa:[]string{"animal"},
		examples:[]string{"insect","arachnid","crustacean"},
	},
	{
		name:"invertebrate",
		isa:[]string{"animal"},
		examples:[]string{"arthropod","mollusc","worm"},
	},
	{
		name:"mollusc",
		examples:[]string{"snail","slug","octopus","squid"},
	},
	{
		name:"marine animal",
		examples:[]string{"fish","octopus","squid","jellyfish","shrimp","lobster","crab"},
	},
	{
		name:"vertebrate",
		isa:[]string{"animal"},
		examples:[]string{"mammal","fish","reptile","amphibian"},
	},
	{
		name:"lizard",
		isa:[]string{"vertebrate"},
		examples:[]string{"snake","quadrupedal lizard","quadrupedal amphibian"},
	},
	{
		name:"quadrupedal lizard",
		isa:[]string{"lizard"},
		examples:[]string{"gecko","iguana","crocodile","alligator","dinosaur","chameleon","komodo dragon"},
	},
	{
		name:"quadrupedal amphibian",
		isa:[]string{"amphibian"},
		examples:[]string{"frog","salamander","toad"},
	},
	{
		name:"tree",
		isa:[]string{"plant"},
		examples:[]string{"palm tree","fern","oak tree","conifer","evergreen","small tree","large tree"},
		has:[]string{"trunk (tree)","foilage"},
	},
	{
		name:"bush",
		isa:[]string{"plant"},
	},	
	{
		name:"grass",
		isa:[]string{"plant"},
	},
	{
		name:"fruit",
		examples:[]string{"apple","banana","orange","grapefruit","strawberry","blackberry","tangerine","peach",},
	},
	{
		name:"cutlery",
		isa:[]string{"tool","kitchenware"},
		examples:[]string{"knife","fork","spoon","glass","mug","plate","serving bowl","serving dish","saucepan","frying pan","pot","wok","steamer"},
	},
	{
		name:"kitchen appliance",
		isa:[]string{"electrical applicance"},
		examples:[]string{"microwave oven","toaster","kettle","coffee machine","blender","electric cooker",},
		found_in:[]string{"kitchen",},
	},
	{
		name:"domestic fittings",
		examples:[]string{"electric socket","light switch","air vent","airconditioning unit","tap","toilet"},
	},
	{
		name:"desktop object",
		examples:[]string{"intray","pen holder","stapler","drawing pins","paper clips","pen","desklamp","desktop PC"},
	},
	{
		name:"electrical applicance",
		examples:[]string{"kitchen applicance","consumer electronics","lamp","desk lamp","light bulb","ceiling light","lantern","security camera"},
	},
	{
		name:"consumer electronics",
		isa:[]string{"electrical applicance"},
		examples:[]string{"TV","monitor","PC","laptop","tablet computer","smartphone","telephone","radio","game console","sound system","speakers","network switch","network hub","camera","cam corder","remote control handset","electric torch"},
	},
	{
		name:"mounted object",
		isa:[]string{"generic object"},
		examples:[]string{"ceiling mounted","wall mounted","ground mounted"},
	},
	{
		name:"lighting",
		isa:[]string{"generic object"},
		examples:[]string{"candle","light bulb","flourescent light","LED light","torch","electric torch","burning torch","lantern","lamp","gas lamp","floodlight"},
	},
	{
		name:"chandelier",
		isa:[]string{"ornament","light fitting","ceiling mounted"},
	},
	{
		name:"computer perhipheral",
		isa:[]string{"consumer electronics"},		
		examples:[]string{"computer mouse","computer keyboard","joystick","gamepad","webcam","microphone"},
	},
	{	name:"TV",
		examples:[]string{"flatscreen TV","LCD TV","plasma TV","LED TV","OLED TV","curved TV","CRT TV"},
	},
	{
		name:"geographic feature",
		examples:[]string{"mountain","hill","coastline","volcano","plain","valley","cave","forest"},
	},
	{
		name:"surface material",
		examples:[]string{"fur","feathers","wood","plastic","stone","sand","dirt","mud","soil","vegetation","grass","tiles","paving stones","bricks","concrete","corrugated metal","metal","rusted metal","plastic sheets","rubber","foilage"},
	},
	{
		name:"grass",
		examples:[]string{"dry grass","long grass","cut grass","wild grass"},
	},
	{	name:"vegetation",
		isa:[]string{"plant"},
	},
	{
		name:"ground",
		examples:[]string{"soil","grass","park","lawn","field","sidewalk","pavement","road","runway","path","footpath"},
	},
	{
		name:"road",
		examples:[]string{"cobbled road","tarmac road","brick road","dirt road"},
	},
})


// ?! c++ address of member is useful for this, how to do?
// generalise leaf/root tracing 'isa'/'examples'

func computeRootDistances(labels map[string]*Label) {
	
	// find each root..
	for _,x :=range labels{
		if len(x.isa)>0{ //is this a root?
			continue;
		}
		floodFillRootDist(x, 0)
	}
}
func computeLeafDistances(labels map[string]*Label) {
	// find each root..
	for _,x :=range labels{
		if len(x.examples)>0{ //is this a leaf?
			continue;
		}
		floodFillLeafDist(x, 0)
	}
}

func setMinInt(p *int,x int){
	if x<*p {*p=x}
}
func setMaxInt(p *int,x int){
	if x>*p {*p=x}
}

func floodFillRootDist(label *Label, dist int){
	if label.minDistFromRoot<=dist{ return }// dont visit again; shorter path found already
	setMinInt(&label.minDistFromRoot,dist)
	for _,x := range label.examples {
		floodFillRootDist(x,dist+1);
	}
}
func floodFillLeafDist(label *Label, dist int){
	if label.minDistFromLeaf<=dist{ return }// dont visit again; shorter path found already
	setMinInt(&label.minDistFromLeaf,dist)
	for _,x := range label.isa {	// go back one
		floodFillLeafDist(x,dist+1);
	}
}

func createLabel(n string) *Label{
	l:=&Label{name:n, minDistFromRoot:0xffff,minDistFromLeaf:0xffff}
	return l
}

type LabelGraph struct{
	all map[string]*Label;
	orphans []*Label; // no 'isa' or 'examples'
	roots []*Label; // no 'isa'
	leaves []*Label; // no 'examples'
	middle []*Label; // both 'isa' and 'examples'
}

func (self LabelGraph) CreateOrFindLabel(newname string) *Label{
	if lbl,ok:=self.all[newname];ok {return lbl;}
	newlbl:=&Label{name:newname}
	self.all[newname]=newlbl;
	return newlbl;
}

func InsertUniqueLabelPtr(list *[]*Label, item *Label) int{
	for i,x :=range *list{if x==item {return i;}}
	*list = append(*list, item);
	return len(*list)-1;
}

func (self *Label) AddExample(other *Label){
	if (self==other) {return;}	// something wrong!
	InsertUniqueLabelPtr(&self.examples,other);
	InsertUniqueLabelPtr(&other.isa,self);
}
func (self *Label) AddPart(other *Label){
	if (self==other) {return;}	// something wrong!
	InsertUniqueLabelPtr(&self.has,other);
	InsertUniqueLabelPtr(&other.part_of,self);
}

func makeLabelGraph(srcLabels []SrcLabel) *LabelGraph{

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
	for _,src:= range srcLabels {
		this_label:=findOrMakeLabel(src.name)
		// TODO does go have field pointers or
		// any other means to reduce the cut-paste here..
		
		// "isa" and "examples" are reciprocated:-
		for _,isa_name:= range src.isa{
			isa_labelstruct:=findOrMakeLabel(isa_name)
			appendLabelPtrList(&isa_labelstruct.examples, this_label)
			appendLabelPtrList(&this_label.isa,  isa_labelstruct);
		}
		for _,ex:= range src.examples{
			exl:=findOrMakeLabel(ex)
			appendLabelPtrList(&exl.isa, this_label)
			appendLabelPtrList(&this_label.examples, exl);
		}
		// "has" and "partof" are reciprocated
		for _,has:= range src.has{
			x:=findOrMakeLabel(has)
			appendLabelPtrList(&x.part_of, this_label)
			appendLabelPtrList(&this_label.has, x);
		}
		for _,p:= range src.part_of{
			x:=findOrMakeLabel(p)
			appendLabelPtrList(&x.has, this_label)
			appendLabelPtrList(&this_label.part_of, x);
		}

		// "bigger than" and "smaller than" are reciprocated
		for _,it:= range src.smaller_than{
			x:=findOrMakeLabel(it)
			x.smaller_than = append(x.smaller_than,this_label)
			this_label.bigger_than = append(this_label.bigger_than, x)
		}
		for _,it:= range src.bigger_than{
			x:=findOrMakeLabel(it)
			x.bigger_than = append(x.bigger_than,this_label)
			this_label.smaller_than = append(this_label.smaller_than, x)
		}
	}
	computeLeafDistances(labels);
	computeRootDistances(labels);

	// 'orphans'
	// collect them under 'uncategorized objects'

	

	// final collection
	l:=&LabelGraph{all:labels};
	for _,x := range l.all{
		num_isa:=len(x.isa);
		num_examples:=len(x.examples);
		if num_isa==0 && num_examples==0 {
			appendLabelPtrList(&l.orphans, x);
			l.CreateOrFindLabel("uncategorized item").AddExample(x);
		} else if num_isa!=0 && num_examples!=0 {
			appendLabelPtrList(&l.middle, x);
		} else if num_isa==0 {
			appendLabelPtrList(&l.roots, x);
		} else if num_examples==0 {
			appendLabelPtrList(&l.leaves, x);
		} else {
			fmt.Printf("fail!\n");
			os.Exit(0)
		}
	}
	
	return l;
}

	// Show results:-
	// TODO formalise this as actual JSON

func printContent(n string,xs[]*Label,postfix string){
	if len(xs)==0 {return}
	fmt.Printf("\t\t\"%s\":[",n);
	for i,x:=range xs{
		fmt.Printf("\"%v\"",x.name)
		if i<len(xs)-1 {fmt.Printf(",");} 
	}
		
	fmt.Printf("]%s\n",postfix);
}

func (self LabelGraph) DumpJSON(verbose bool){

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
		len(self.roots), len(self.middle),len(self.leaves), len(self.orphans));
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
	
	//labelGraph.DumpInfo();

}




