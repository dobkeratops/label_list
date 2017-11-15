package main
import ("fmt";"os")

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

func appendLabelList(ls *[]*Label,l *Label){
	*ls = append(*ls, l)
}

var(g_srcLabels=[]SrcLabel{
	{
		name:"person",
		isa:[]string{"human"},
		examples:[]string{"man","woman","child","boy","girl","baby","police officer","soldier","workman"},
	},
	{
		name:"human",
		isa:[]string{"animal"},
		has:[]string{"head","arm","leg","torso","neck"},
		states:[]string{"standing","walking","running","sitting","kneeling","reclining","sleeping"},
	},
	{
		name:"soldier",
		isa:[]string{"person","military objects"},
	},
	{
		name:"weapon",
		isa:[]string{"military objects"},
		examples:[]string{"firearm","combat knife","sword","handgun","rocket launcher","flame thrower","machine gun","pistol","revolver","handgun","grenade launcher",},
	},
	{
		name:"firearm",
		has:[]string{"gun barrel","stock","handgrip","sights"},
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
		isa:[]string{"M4 carbine","M16A1","M16A2","M16A4","AR15"},
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
		isa:[]string{"vehicle","military object"},
		has:[]string{"turret","gun","catepillar tracks"},
	},
	{
		name:"canon",
		isa:[]string{"weapon"},
	},
	{
		name:"vehicle",
		isa:[]string{"machine"},
		examples:[]string{"car","truck","aircraft","ship","bicycle","motorbike","bus","semi truck"},
	},
	{
		name:"car",
		isa:[]string{"vehicle"},
		has:[]string{"wheel","bonnet","license plate","windscreen","headlight","tail light","exhaust pipe"},
		examples:[]string{"hatchback","SUV","pickup truck","sedan","coupe","sportscar","convertible"},
	},	
	{
		name:"animal",
		isa:[]string{"organism"},
	},
	{
		name:"construction machinery",
		examples:[]string{"bulldozer","excavator","mini excavator","road roller","wrecking ball","pile driver","digger","crane","tower crane"},
	},
	{
		name:"bulldozer",
		has:[]string{"shovel","catepillar tracks"},
	},
	{
		name:"machine gun",
		isa:[]string{"firearm","automatic weapon"},
	},
	{
		name:"belt fed machine gun",
		isa:[]string{"machine gun"},
		examples:[]string{"m60","rpk","GPMG","minimi"},
	},
	{
		name:"excavator",
		has:[]string{"shovel","catepillar tracks"},
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
		has:[]string{"trunk"},
	},
	{
		name:"plant",
		isa:[]string{"organism"},
	},
	{
		name:"fruit",
		part_of:[]string{"plant"},
	},
	{
		name:"vegtable",
		isa:[]string{"plant"},
	},
	{	name:"food",
		examples:[]string{"vegtable","fruit","nuts","meat","cereal","egg","salad","soup","sandwich"},
	},
	{	name:"vegtable",
		examples:[]string{"brocoli","peas","carrots","spinach","cellery","beansprouts","brussel sprouts","cauliflower","mushroom","peppers","courgette","leak","cabbage","onion","beans","tomato"},
	},
	{	name:"furniture",
		examples:[]string{"table","chair","bed","cupboard","desk","bench",},
	},
	{	name:"agricultural equipment",
		examples:[]string{"tractor","combine harvester","crop duster"},
	},
	{	name:"tractor",
		isa:[]string{"vehicle"},
		has:[]string{"tractor wheel,cabin,engine"},
	},
	{	name:"domestic room",
		examples:[]string{"kitchen","dining room","bedroom","living room","study","hallway","store room","garage"},
		part_of:[]string{"house"},
	},
	{	name:"office building",
		isa:[]string{"building"},
		has:[]string{"atrium","boardroom","office"},
	},
	{	name:"urban objects",
		examples:[]string{"street bin","wheeliebin","skip","lamp post","utility pole","electricity pylon","telegraph pole","traffic lights","sign post","building","radio tower","satelite dish","bottle bank",},
	},
	{	name:"building",
		examples:[]string{"church","house","tower block","factory","warehouse","cathederal","terminal building","train station","skyscraper","tower",},
	},
	{
		name:"arthropod",
		isa:[]string{"animal"},
		examples:[]string{"insect","arachnid","crustacean"},
	},
	{
		name:"invertebrate",
		isa:[]string{"animal"},
		examples:[]string{"arthropod","molusc","worm"},
	},
	{
		name:"vertebrate",
		isa:[]string{"animal"},
		examples:[]string{"mamal","fish","reptile","amphibian"},
	},
	{
		name:"tree",
		isa:[]string{"plant"},
		has:[]string{"trunk","foilage"},
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
		examples:[]string{"knife","fork","spoon"},
	},
	{
		name:"kitchen appliance",
		isa:[]string{"electrical applicance"},
		examples:[]string{"microwave oven","toaster","kettle","coffee machine","blender","electric cooker",},
		found_in:[]string{"kitchen",},
	},
	{
		name:"electric socket",
	},
	{
		name:"electrical applicance",
		examples:[]string{"kitchen applicance","consumer electronics","lamp","desk lamp","light bulb","ceiling light","lantern"},
	},
	{
		name:"consumer electronics",
		isa:[]string{"electrical applicance"},
		examples:[]string{"TV","monitor","PC","laptop","tablet computer","smartphone","telephone","radio","game console","sound system","speakers","network switch","network hub"},
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
		examples:[]string{"mountain","hill","coastline","volcano","plain","valley","cave"},
	},
	{
		name:"surface material",
		examples:[]string{"fur","feathers","wood","plastic","stone","sand","dirt","mud","soil","vegetation","grass","tiles","paving stones","bricks","concrete","corrugated metal","metal","rusted metal","plastic sheets","rubber",},
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

type LabelList struct{
	all map[string]*Label;
	orphans []*Label; // no 'isa' or 'examples'
	roots []*Label; // no 'isa'
	leaves []*Label; // no 'examples'
	middle []*Label; // both 'isa' and 'examples'
}
func makeLabelList(src map[string]*Label) *LabelList{
	l:=&LabelList{all:src};
	for _,x := range src{
		num_isa:=len(x.isa);
		num_examples:=len(x.examples);
		if num_isa==0 && num_examples==0 {
			appendLabelList(&l.orphans, x);
		} else if num_isa!=0 && num_examples!=0 {
			appendLabelList(&l.middle, x);
		} else if num_isa==0 {
			appendLabelList(&l.roots, x);
		} else if num_examples==0 {
			appendLabelList(&l.leaves, x);
		} else {
			fmt.Printf("fail!\n");
			os.Exit(0)
		}
	}
	return l;
}

func main() {

	// compile labels into a map for access by string, with links

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
	for _,src:= range g_srcLabels {
		this_label:=findOrMakeLabel(src.name)
		// TODO does go have field pointers or
		// any other means to reduce the cut-paste here..
		
		// "isa" and "examples" are reciprocated:-
		for _,isa_name:= range src.isa{
			isa_labelstruct:=findOrMakeLabel(isa_name)
			appendLabelList(&isa_labelstruct.examples, this_label)
			appendLabelList(&this_label.isa,  isa_labelstruct);
		}
		for _,ex:= range src.examples{
			exl:=findOrMakeLabel(ex)
			appendLabelList(&exl.isa, this_label)
			appendLabelList(&this_label.examples, exl);
		}
		// "has" and "partof" are reciprocated
		for _,has:= range src.has{
			x:=findOrMakeLabel(has)
			appendLabelList(&x.part_of, this_label)
			appendLabelList(&this_label.has, x);
		}
		for _,p:= range src.part_of{
			x:=findOrMakeLabel(p)
			appendLabelList(&x.has, this_label)
			appendLabelList(&this_label.part_of, x);
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

	labelList := makeLabelList(labels);

	// Show results:-
	// TODO formalise this as actual JSON
	fmt.Printf("{\n ");
	for name,label :=range labels {
		fmt.Printf("\t\"%v\":{\n ",name);

		printContent:=func(n string,xs[]*Label,postfix string){
			if len(xs)==0 {return}
			fmt.Printf("\t\t\"%s\":[",n);
			for i,x:=range xs{
				fmt.Printf("\"%v\"",x.name)
				if i<len(xs)-1 {fmt.Printf(",");} 
			}
		
			fmt.Printf("]%s\n",postfix);
		}
		fmt.Printf("\t\tminDistFromRoot:%v\n", label.minDistFromRoot);
		fmt.Printf("\t\tminDistFromLeaf:%v\n", label.minDistFromLeaf);
		printContent("isa",label.isa,",");
		printContent("examples",label.examples,",");
		printContent("has",label.has,",");
		printContent("part_of",label.part_of,"");
		fmt.Printf("\t},\n")
	}
	fmt.Printf("}\n ");
	
	fmt.Printf("\"labelList stats\":{\"total\":%v, \"roots(metalabels)\":%v, \"middle(labels\":%v \"leaf examples\":%v,\"orphans\":%v}",
		len(labelList.all),
		len(labelList.roots), len(labelList.middle),len(labelList.leaves), len(labelList.orphans));
	
}




