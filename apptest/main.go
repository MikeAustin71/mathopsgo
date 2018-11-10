package main

import (
	"../examples"
	"../mathops"
	"fmt"
	"math/big"
	"time"
)


func main() {
	xNum := big.NewInt(99123456789)
	xNumPrecision := big.NewInt(9)
	m := big.NewInt(56)
	s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
	agMeanMaxInternalPrecision := big.NewInt(2000)
	agMeanMaxOutputPrecision := big.NewInt(200)
	maxPiDividePrecision := big.NewInt(3000)
	maxFinalOutputPrecision := big.NewInt(31)
	expectedValue := "4.5963661115009113152618465511833"

	TestNatLogOfXArithmeticGeometricMean(
		xNum,
		xNumPrecision,
		m,
		s4DivPrecisionMaxInternalPrecision,
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		maxPiDividePrecision,
		maxFinalOutputPrecision,
		expectedValue)

}


/*


func main() {
	ePrefix := "main() "
	m:= big.NewInt(19)
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	twoToM := big.NewInt(0).Exp(big.NewInt(2), m, nil)

	s, sPrecision, err :=

		mathops.BigIntMathMultiply{}.BigIntMultiply(
			twoToM,
			big.NewInt(0),
			xNum,
			xNumPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	biNum_s, err := mathops.BigIntNum{}.NewBigIntPrecision(s, sPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "biNum_s - %v", err.Error())
	}

	fmt.Println("s: ", biNum_s.GetNumStr())
	s4DivPrecisionMaxInternalPrecision := big.NewInt(5000)

	fourDivS, fourDivSPrecision, err :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			big.NewInt(4),
			big.NewInt(0),
			s,
			sPrecision,
			s4DivPrecisionMaxInternalPrecision)

	agMeanMaxInternalPrecision := big.NewInt(2000)
	agMeanMaxOutputPrecision := big.NewInt(36)
	biNumFourDivS, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(fourDivS, fourDivSPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "biNum_s - %v", err.Error())
	}

	fmt.Println("------------------------------------------")
	fmt.Println("        m: ", m.Text(10))
	fmt.Println("aNumStart: ", biNumFourDivS.GetNumStr())
	fmt.Println("gNumStart: ", 1)
	fmt.Println("-------------------------------------------")

	TestArithmeticGeometricMean(
		fourDivS,
		fourDivSPrecision,
		big.NewInt(1),
		big.NewInt(0),
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		"Unknown")

}


func main() {

	expectedAValPrecision := 2000
	aValStr := "0.11330900360312992286990200624763332427225845327360940502460741924614522351041247792648521970154199158848642654123673227698405692477236472266109196855609272279996132044775664194899106480130107243206006289407009947588761939784295206441186148782199150432190980543958227457382182174961811712740701433528373686850028263356557099307036004020933164371912198894494409233163445418417376317656393559773833435477252064643865703650663895342397251556285964069436652804299630604469850708069544255047265264039934028676956874118410251693772860952288744319934443378859418440217202632771387792003068543082133119771217099896211270121999455290118120427592522619824525249658802085920043060341620200644943571223928553385169676661206127780295959875895452944902621783750384761872509593709499712507442643339911217331218902529103017484916036415406219082974093831138573323510989427988426570543476081546195940223729989589515664295959824506634758095696589549485986934194501589507523317392427364481061080375568536164433878391369083559816334849259102152656351101010724903848643939096793498865660722692794439377135544627578480577100152792457039468135539223575725928216460083166185858684621635479934997523865309824523339774679544961927304128807398287330919859061090968731199076238361045602322010557461106631372227462695987267305493264074300692715019370074961863933375164649721161736931517384778690223960165677292101600539431663136607991180754319992352499589564742040879629953144860588530044577249798050488908687425974026755601194108807555732232037113162607491133774684231880310655024678140610027120395529129973130944020983984304236112128468600996090734854590257080091358714366396084073659754003181802344151003799117329501057954146413064013957003659324001825417191221330032309214850869238978019326442986646026286332459412806980978253416786818844691357258907706442687311227960883987776521389458120323999576249584334014488718749628304508268816595222152714681625611765853286050623763801156904144710099157039288903799134369375057122047394"

	binAValue, err := mathops.BigIntNum{}.NewNumStr(aValStr)

	if err != nil {
		fmt.Printf("Error BigIntNum{}.NewNumStr(aValStr). %v", err.Error())
		return
	}

	if expectedAValPrecision != binAValue.GetPrecision() {
		fmt.Printf("Expected  A Value Precision='%v'. "+
			"Instead, aValue Precision='%v'",
			expectedAValPrecision, binAValue.GetPrecision())
		return
	}

	expectedGValPrecision := 2000

	gValStr := "0.11330900348746480628147615244264405605558552604672970617244776385367920335636997475691501899851519927164825231375809128537158894153036855260420834803861407725368815116452036663129860774485542140418105588425062075290539140739080909044297611626701182409635943521425699853227717922266761997242294326526178921645007539820483193369598759098564034831013290217015018242761170435759030447934502181015226720729835026194210239457879112350097415890734843359170683779864453155664211705134256032474687986203379353524028745638186523497898434336788594379567103495054201902509976734654967422207030984920001416420633878572975632015832420974915350514710270182242434611410804705914525434247773754882477549026973957870175695659876609902117314400155470746065383365038786399033430754629318712783549089313365019370389599671071985963538924476037319250456202183638761529904672131902866019208739538654836323906937854959650689943366455579786631690603778040808825283679571628480371481965572073936215766362937178003812567967418512804985741280891290946577576699635058957127931838577497068347437233912255556733603030131142450326962434828290960084589117421305161369165043614264714829591616255307554022039043932015890359001220917013928648380885005797089693225036145152387975250588105259676249397364622276781147085504336099062676739284891364871888393362962228599233969257230355777010926052041257751357613681327395723310749640191836610430073676145171996798317425880087671296619116326988415423067775437040098908315270914651663048931288169497090116592904017234637547086060320314707614811625315186068031596625756862636155412038303207914899562421331409003259769620487877343355689518854368366489924201251941120085496307462701099687245840728564378880934546730990256229146328047391938877814018443894159066810432575868236543399508769428587120398804342938386547088736666001585501188376996938301029321103833176314320370613718334744580518586584935716117488263996883531727528823703974091671360106454501631823565556610032648350605250411796436852596"

	binGValue, err := mathops.BigIntNum{}.NewNumStr(gValStr)

	if err != nil {
		fmt.Printf("Error BigIntNum{}.NewNumStr(gValStr). %v", err.Error())
		return
	}

	if expectedGValPrecision != binGValue.GetPrecision() {
		fmt.Printf("Expected G Value Precision='%v'. "+
			"Instead, G Value Precision='%v'",
			expectedGValPrecision, binGValue.GetPrecision())
		return
	}

	aVal := binAValue.GetIntegerValue()
	aValPrecision := big.NewInt(int64(binAValue.GetPrecision()))

	gVal := binGValue.GetIntegerValue()
	gValPrecision := big.NewInt(int64(binGValue.GetPrecision()))

	result := big.NewInt(0)
	result.Mul(aVal, gVal)
	resultPrecision := big.NewInt(0).Add(aValPrecision, gValPrecision)

	fmt.Println("result: ")
	fmt.Println(result.Text(10))
	fmt.Println()
	fmt.Println("Result Precision: ", resultPrecision.Text(10))

	binResult, err := mathops.BigIntNum{}.NewBigIntPrecision(result, resultPrecision)
	fmt.Println("BigIntNum Result: ")
	fmt.Println(binResult.GetNumStr())


}




func main() {
	ePrefix := "main() "

	aCycleStr :=
		"0.11927263530963808691463201673094572068099721222991782972072370786575221405178644635242043252362343129196567897588090141970629014333813255863515395103834738689955232861812826884188175259705948023847936837884353428848639543000227765437780882270378327335998776854149291326549827016692154975029596220784668398997444742809844142892548289814282346544985535688898912890001088011563398672674006004317139531645133855544934745731651863739329504393090836628592732331593083360839625257793532547495544147798905638487273854106138207503444419760721593521572707908702572559657346446433415388018696811767547205647229026099999564228239556153757495707328100265601028071820239763021010168716232804018086483965522573080417184532579594943489573414711892246163245539963012574052844385229668295788438331839501300808961267931952451645735219237313465575220023450701478272050587811119947598695252871594956095814174498757472065763534419217215408811318776763088356614185731975825688043613308409562431582851127903760923600117482445852361715982559679465997827404074297283956822812519336625395018297687692727352918625962043567261685877066852940641721004333031751722633213055405904396037165657524106638021771570843997949280875989360781899880695238548020996350034931181044606893084939494975324250629566939405248475215390050305403615628058800967379045662633359029090145611377896737994875870446439601292137467871920898265754528071145432439799852022402279950366169621564018326591039594648845861578590789429627381597721081289634833968370384057181185226886351521843886892485310140894298474147693548962027007115282551511130816021049209653402473368785932483225220921117492873622387055218575040951199011809312081669710267183601720579313360880760442266610384912746924265224983617642753734295510507337118095177453586800515447745925474130244142310491961799349276968922536483400008111381809194079226236708122506858112948192811002541140382461942230795974253236648742080776056293251221393159004932065353088690415976141019644738890077812536418160429476237465334236357827056503631275383806959527109338374928913100950316007508324760495171293210669984352353654940375717393590118338529877477700293241103872768644146648581305831031687148751079394834242751178040616159458916162189217786527921522875678452964484786656209100363275553071803345005002762136468091764336757306748920040615274910998823545329638290524541979183544259730579683827539511416676012851409452062950573194951124204827567934581592047267426441438621726734390652109915324367655230371311577414333028786008925676054979476822203528005719187374901740934731552290042281693489471740839659845972562047079275593006629667004855604425759401606019561513228808955139284107719378373432872236297973497004093125652081320642181039633726259260289196555369165358541786458069005572850746118199523777836931944179074998596016416929952997455294258314128382954328800093272000713404248249806917587655873545406254772169121499740764656909446639093848929630141784321548312953333080412669507679685905975856708852855552661388195685320692726113025249222223086813224407137775634319046681059669642706783186172489102157667753173069788171879631721432259327663678730776572611811255379459632821060661228936389962745479569177942747934086342352524700889640311066271358807120195381487944499187545092334922444714767573197162596446532850195509405716882587654506159264936004741848946846399146480862478485171429442279560843972609819012715756807466885141584435660176952387267280360049639921335768032805937804507034588965473259298116328943350141963369668744290437368127471850707916716832257426623509408801783019146521009239736099301365021204494550220546482111427019116693140497844313935534516135322390162253614541364701738779735906808779962870660775153159909476385852303632685302495039093638266975546288022418642400491547021349513455387083908995916582744555112452190185717480936021196869890307656433061370707634987851610231348856626386338405089347732888082974707568798695979482673795971040089419778694403324303912458938468960324314435700721509019912234495238452348481775375215962521914300037415662899512808622938209360450269379416223941232948741522078255447857352445166061722950421826913651167388880124380604538874281502306859258298610893829118588645176441156148539143888730393352184425211103739287928117989781424938473634622309022271385249831558561176969715699864638297917927648713720573069724126091278351002266511758230493809183049175507449772934268278242177947742521666615865380027769230130536767886927464540249643833232714563592309054952047788103466289160425267625001802271373026587303624638777528847530530245795901536309189070678203104131410837058265661943585393959524854709068854879202136520731521842133358430524790566269772187840759138344239779670042527823219644677483582112264431888107562242801031115162606913609771559201115840576435798981537103207545044957650935802643037896695359028517050438517856640667459478964373901936921813319119179206033451404234233934332848425094517019340675750546603422102369109014234947713259758708464"

	gCycleStr :=
		"0.11927263530963808691322000890984734371731922909532817431762409937125526035443121282176268518268461149904673914990681033505845711783999831461448391740280094668985053072933745478250640958887306838460212361995660807381549834833988075727707399656994161114619993010899138015820431304602196011271916487788480387470775801356655804261036075721120218972719676718083679841779523883572213335710987452557031205535874924848848444641977351479448812225547990931796059078476588082041296553486704124004638352438106727703876820445471154878727335725388024957574067370882715030659197858324425641612479999755523810453648624813791175130420658660382860831074498180140230290606271989914188771774707209682538819655077285977894503065048602390002983104677545815742063419098047940340925723687336075631420497148069114866764167529703878866565974559180051669044166932890702740821890863477109651128546810997268772527134525361686848203412777935802208854323033066515306226448857772775801671626361381128047229556924739309639138911634726218182056826231116892410239876753701826593770542398420530428961095849766569304729956296861488652426802173097292225829804447913198388357814535794313727617239007383026861690200740378462090836781591523108785810946203958298486890700427783579985466331256923560331860322677965294018826427969505254438294289721325445945860182135706724553289699481068180321950481860855509426144534984689335461742012181306620785436776494731385642301288062030359991774932245179170417870874521665629673328463352854519361819336075017850611785138770766962217119144209352060285043852433697262882886221558553484792319217368611231273107844826657361126624450818630607252111340200623738369933288500318017032443925033718457110276107249967130671405269839720480816822911616740890725005222983072768405209381761831907451591607950257618507485233259747605842777997243585194613648499602576736332514256932438485719488882144900662651569048989046548713046690690675505835901444735923603252151462083871288514526033758292286504057112847463463292168551962073607010351343105698423459177671920630559606295253608801835033397829313709415116345160735394241272089557256560410144969507094627167717649930711279476943001703816886765906847175439168234405365885410271206654326046252369758402277555630257910584095356462507931150248818686605780573478532534865136935824838152000204139986374962655637834692226570706635603431664904143821304325013193013584510174943250074763475300747002182681587274785251229132138826206779603046000766540206720804388484975752235575584981710395112616431253872126697163888211171618809044851506883554395255408495435297807843098771573378595200703541270663587556334466769577973887702782798849096304261333755386493961263595866854624550076881331872922554295943828275893920289401151242475451184114265509380144322520004423865391861481111867811824326540781045968824474848390647627681435078044591472618964346798292363983753963471715162303927353028230671334763204413192781279331845047207499972853488297796113627639051596239160038710687430200148842189846854552414337313933616829200654831815023254088828791648898977672795426366215276612661682865221443974723883563645795968309770222891311598904034512574817231783370801466285142304716327035709691825621859318113802321707835889673743697659670862050336516328481978747075027492816566369090387646242165817907292604395372148706143379422441555518845053400889265001570398953812804628746720126482273284149091064471457640137599906510054576927827270843894239939376119927959996427093853511004341344117989947161615855702922909128809062795533860423620359357030758185536905880413851605950912666366382633092928940458354698114137889140303249917319585802546717906830266010396569439923501167138900467238120190589117442606687094204737652346381367838511409881313971778352671250150136788558626670182624545944629981044974420839098893642035555319612415247347941226458283452892237138364391750974744532091374787329604914449337398210112336330345694831560880241102978591090267723936287092330512421596123434059288127810950504282579747393139843610048497081619277938873866433306113038029458690436192875157293403672762979089178242146385503529584168001404180848224297464995527317939424872037014365383508776122437442821109694087017254228994824463025181692658179002474269155248761337333716435206804653753516282574004310212002929796366441607485414420104791675173252473436279011514480324658197026041775796089301615260952185432668131167067586437302589148295795142973409962133241135866496853861006058153007725069463336484195840671776676319341575986310745448895516300220997555762129698563256992940103501003219600769527365230366775421030893274624151783186619638850334378787505169909427299103914281824764468652284613119438559020441390287394716154915267615461764552331879153440288962586630033355717301327922786729868401528799779046027481642190455681508476496293090335529785005326671685810744463896579081037298278110604717933875317243875274062948401354195139802756477422145429196757247473551004305283087392743040371337350163305496900039525146"

	expectedValue :=
		"0.01422596153370592614104169235598"

	maxInternalPrecision := big.NewInt(5000)

	binACycle, err := mathops.BigIntNum{}.NewNumStr(aCycleStr)

	if err != nil {
		fmt.Println(ePrefix +
			"binACycle- %v")
		return
	}

	binGCycle, err := mathops.BigIntNum{}.NewNumStr(gCycleStr)
	if err != nil {
		fmt.Println(ePrefix +
			"binGCycle- %v")
		return
	}

	gCom, gComPrecision, err :=
		mathops.BigIntMathMultiply{}.BigIntMultiply(
			binACycle.GetIntegerValue(),
			binACycle.GetPrecisionBigInt(),
			binGCycle.GetIntegerValue(),
			binGCycle.GetPrecisionBigInt())

	if err != nil {
		fmt.Println(ePrefix +
			"BigIntMultiply- %v")
		return
	}

	gComPreRound, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(gCom, gComPrecision)

	if err != nil {
		fmt.Println(ePrefix +
			"gComPreRound- %v")
		return
	}

	biGComPostRound, biGComPostRoundPrecision, err :=
		mathops.BigIntMath{}.RoundToMaxPrecision(gCom, gComPrecision, maxInternalPrecision)

	if err != nil {
		fmt.Printf( "%v",err.Error())
		return
	}

	gComPostRound, err := mathops.BigIntNum{}.NewBigIntPrecision(biGComPostRound, biGComPostRoundPrecision)


	fmt.Println("aCycle: ", binACycle.GetNumStr())
	fmt.Println("--------------------------------------------------")
	fmt.Println("gCycle: ", binGCycle.GetNumStr())
	fmt.Println("--------------------------------------------------")
	fmt.Println(" expected value: ", expectedValue)
	fmt.Println("   gComPreRound: ", gComPreRound.GetNumStr())
	fmt.Println(" expected value: ", expectedValue)
	fmt.Println("  gComPostRound: ", gComPostRound.GetNumStr())
	fmt.Println("==================================================")
	fmt.Println("==================================================")
	gComPreRound.RoundToDecPlace(50)
	gComPostRound.RoundToDecPlace(50)
	fmt.Println("           gComPreRound: ", gComPreRound.GetNumStr())
	fmt.Println("          gComPostRound: ", gComPostRound.GetNumStr())
	fmt.Println(" gComPreRound Precision: ", gComPreRound.GetPrecision())
	fmt.Println("gComPostRound Precision: ", gComPostRound.GetPrecision())
	fmt.Println("==================================================")
	//fmt.Println("   sqrRtG: ", sqrRtG.GetNumStr())
	fmt.Println()

}


func main() {
	ePrefix := "main() "
	m:= big.NewInt(18)
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	twoToM := big.NewInt(0).Exp(big.NewInt(2), m, nil)

	s, sPrecision, err :=

		mathops.BigIntMathMultiply{}.BigIntMultiply(
			twoToM,
			big.NewInt(0),
			xNum,
			xNumPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	biNum_s, err := mathops.BigIntNum{}.NewBigIntPrecision(s, sPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "biNum_s - %v", err.Error())
	}

	fmt.Println("s: ", biNum_s.GetNumStr())
	s4DivPrecisionMaxInternalPrecision := big.NewInt(5000)

	fourDivS, fourDivSPrecision, err :=
		mathops.BigIntMathDivide{}.BigIntFracQuotient(
			big.NewInt(4),
			big.NewInt(0),
			s,
			sPrecision,
			s4DivPrecisionMaxInternalPrecision)

	agMeanMaxInternalPrecision := big.NewInt(5000)
	agMeanMaxOutputPrecision := big.NewInt(32)
	biNumFourDivS, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(fourDivS, fourDivSPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "biNum_s - %v", err.Error())
	}

	fmt.Println("------------------------------------------")
	fmt.Println("        m: ", m.Text(10))
	fmt.Println("aNumStart: ", biNumFourDivS.GetNumStr())
	fmt.Println("gNumStart: ", 1)
	fmt.Println("-------------------------------------------")

	TestArithmeticGeometricMean(
		fourDivS,
		fourDivSPrecision,
		big.NewInt(1),
		big.NewInt(0),
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		"Unknown")

}
*/

func TestArithmeticGeometricMean(
	aNum,
	aNumPrecision,
	gNum,
	gNumPrecision,
	maxInternalPrecision,
	targetPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestArithmeticGeometricMean() "
	fmt.Println(ePrefix)
	// timeStart := time.Now()

	agMean, agMeanPrecision, gValue, gValuePrecision, cycles, err :=
		mathops.BigIntMath{}.ArithmeticGeometricMean(
			aNum,
			aNumPrecision,
			gNum,
			gNumPrecision,
			maxInternalPrecision,
			targetPrecision)

	// timeEnd := time.Now()

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binANum, err := mathops.BigIntNum{}.NewBigIntPrecision(aNum, aNumPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binGNum, err := mathops.BigIntNum{}.NewBigIntPrecision(gNum, gNumPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binAGMean, err := mathops.BigIntNum{}.NewBigIntPrecision(agMean, agMeanPrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binGValue, err := mathops.BigIntNum{}.NewBigIntPrecision(gValue, gValuePrecision)

	if err!=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	//timeDuration := timeEnd.Sub(timeStart)

	//duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("               BigIntMath{}.ArithmeticGeometricMean() ")
	fmt.Println("======================================================================")
	fmt.Println("                aNum: ", aNum.Text(10))
	fmt.Println("       aNumPrecision: ", aNumPrecision.Text(10))
	fmt.Println("           a Num Str: ", binANum.GetNumStr())
	fmt.Println("                gNum: ", gNum.Text(10))
	fmt.Println("       gNumPrecision: ", gNumPrecision.Text(10))
	fmt.Println("           g Num Str: ", binGNum.GetNumStr())
	fmt.Println("maxInternalPrecision: ", maxInternalPrecision.Text(10))
	fmt.Println("     targetPrecision: ", targetPrecision.Text(10))
	fmt.Println("======================================================================")
	fmt.Println("              agMean: ", agMean.Text(10))
	fmt.Println("     agMeanPrecision: ", agMeanPrecision.Text(10))
	fmt.Println("       agMean NumStr: ", binAGMean.GetNumStr())
	fmt.Println("      expected value: ", expectedValue)
	status := "FAILURE - Actual DOES NOT MATCH Expected Value!!"
	if binAGMean.GetNumStr() == expectedValue{
		status = "SUCCESS - Actual Matches Expected Value!"
	}

	fmt.Println("              Status: ", status)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("              gValue: ", gValue.Text(10))
	fmt.Println("     gValuePrecision: ", gValuePrecision.Text(10))
	fmt.Println("       gValue NumStr: ", binGValue.GetNumStr())
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("              Cycles: ", cycles)
	fmt.Println("----------------------------------------------------------------------")
	//fmt.Println("      Execution Time: ", duration)
	fmt.Println("======================================================================")
	fmt.Println()

}


func TestNatLogOfXArithmeticGeometricMean(
	xNum,
	xNumPrecision,
	m,
	s4DivPrecisionMaxInternalPrecision,
	agMeanMaxInternalPrecision,
	agMeanMaxOutputPrecision,
	piDivideMaxPrecision,
	maxFinalResultPrecision *big.Int,
	expectedValue string) {

	/*

		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(144)
		s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
		agMeanMaxInternalPrecision := big.NewInt(2000)
		agMeanMaxOutputPrecision := big.NewInt(800)
		maxPiDividePrecision := big.NewInt(2000)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"


		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(1029)
		s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
		agMeanMaxInternalPrecision := big.NewInt(2000)
		agMeanMaxOutputPrecision := big.NewInt(800)
		maxPiDividePrecision := big.NewInt(2000)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"

		xNum=14
		expectedValue := "2.6390573296152586145225848649014"
		 xNum=256
		expectedValue := "5.5451774444795624753378569716654"


		xNum := big.NewInt(2)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(512)
		agMeanMaxInternalPrecision := big.NewInt(8000)
		agMeanMaxOutputPrecision := big.NewInt(500)
		maxPiDividePrecision := big.NewInt(150)
		maxFinalOutputPrecision := big.NewInt(99)
		expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"



		xNum := big.NewInt(14)
		xNumPrecision := big.NewInt(0)
		m := big.NewInt(256)
		maxInternalPrecision := big.NewInt(1000)
		maxPrecision := big.NewInt(31)
		expectedValue := "2.6390573296152586145225848649014"

	 */



	timeStart := time.Now()
	timeEnd := time.Now()
	ePrefix := "TestNatLogOfXArithmeticGeometricMean() "

	timeStart = time.Now()

	lnOfX, lnOfXPrecision, err :=
		mathops.BigIntMathLogarithms{}.NatLogOfXArithmeticGeometricMean(
			xNum,
			xNumPrecision,
			m,
			s4DivPrecisionMaxInternalPrecision,
			agMeanMaxInternalPrecision,
			agMeanMaxOutputPrecision,
			piDivideMaxPrecision,
			maxFinalResultPrecision)

	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	bINumXNum, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(xNum, xNumPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "bINumXNum Error= %v", err.Error())
		return
	}

	bINumlnOfX, err :=
		mathops.BigIntNum{}.NewBigIntPrecision(lnOfX, lnOfXPrecision)

	if err != nil {
		fmt.Printf(ePrefix + "bINumlnOfX Error= %v", err.Error())
		return
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("        BigIntMathLogarithms{}.NatLogOfXArithmeticGeometricMean() ")
	fmt.Println("======================================================================")
	fmt.Println("                              xNum: ", xNum.Text(10))
	fmt.Println("                     xNumPrecision: ", xNumPrecision.Text(10))
	fmt.Println("                       xNum NumStr: ", bINumXNum.GetNumStr())
	fmt.Println("                                 m: ", m.Text(10))
	fmt.Println("s4DivPrecisionMaxInternalPrecision: ", s4DivPrecisionMaxInternalPrecision.Text(10))
	fmt.Println("        agMeanMaxInternalPrecision: ", agMeanMaxInternalPrecision.Text(10))
	fmt.Println("          agMeanMaxOutputPrecision: ", agMeanMaxOutputPrecision.Text(10))
	fmt.Println("              piDivideMaxPrecision: ", piDivideMaxPrecision.Text(10))
	fmt.Println("           maxFinalResultPrecision: ", maxFinalResultPrecision.Text(10))
	fmt.Println("========================== Result ====================================")
	fmt.Println("                             ln(x): ", lnOfX.Text(10))
	fmt.Println("                   ln(x) Precision: ", lnOfXPrecision.Text(10))
	fmt.Println("                      ln(x) NumStr: ", bINumlnOfX.GetNumStr())
	fmt.Println("              ln(x) Expected Value: ", expectedValue)

	status := "FAILURE!!! Expected Value DOES NOT MATCH Actual Value"
	if expectedValue == bINumlnOfX.GetNumStr() {
		status = "SUCCESS**** Expected Value MATCHES Actual Value"
	}

	fmt.Println("                    Status: ", status)
	fmt.Println("======================================================================")
	fmt.Println("      Execution Time: ", duration)

}


func TestBigIntNumLogBaseOfX(
	base,
	xNum mathops.BigIntNum,
	maxPrecision uint,
	expectedLogValue string) {

	// Logarithm test code
	timeStart := time.Now()
	timeEnd := time.Now()
	ePrefix := "TestBigIntNumLogBaseOfX() "

	timeStart = time.Now()
	logValue, err :=
		mathops.BigIntMathLogarithms{}.BigIntNumLogBaseOfX(
			base,
			xNum,
			maxPrecision)
	timeEnd = time.Now()


	if err != nil {
		fmt.Printf(ePrefix +
			"Error: error returned from " +
			"BigIntMathLogarithms{}.BigIntNumLogBaseOfX(" +
			"base, xNum, maxPrecision) "+
			"base='%v' xNum='%v' Error='%v'",
			base.GetNumStr(), xNum.GetNumStr(), err.Error())

		return
	}


	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	status := "Success!!!!! Expected Value Matches Actual Value"

	if expectedLogValue != logValue.GetNumStr() {
		status = "FAILURE***** Expected Value DOES NOT MATCH Actual Value!"
	}

	fmt.Println()
	fmt.Println("   BigIntMathLogarithms{}.BigIntNumLogBaseOfX() ")
	fmt.Println("==================================================")
	fmt.Println("               base: ", base.GetNumStr())
	fmt.Println("               xNum: ", xNum.GetNumStr())
	fmt.Println("       maxPrecision: ", maxPrecision)
	fmt.Println("--------------------------------------------------")
	fmt.Println("          log Value: ", logValue.GetNumStr())
	fmt.Println(" Expected log Value: ", expectedLogValue)
	fmt.Println("--------------------------------------------------")
	fmt.Println(status)
	fmt.Println("--------------------------------------------------")
	fmt.Println("      Execution Time: ", duration)
	fmt.Println("==================================================")
	fmt.Println()
}


func TestBinaryLogBaseOfX(
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxInternalPrecision,
	maxPrecision,
	cycles  *big.Int,
	expectedValue string) {

	ePrefix := "TestBinaryLogBaseOfX() "
	timeStart := time.Now()
	timeEnd := time.Now()

	timeStart = time.Now()

	logResult, logResultPrecision, errX :=
		mathops.BigIntMathLogarithms{}.LogBaseOfXByDivide(
			base,
			basePrecision,
			xNum,
			xNumPrecision,
			maxInternalPrecision,
			maxPrecision,
			cycles)

		timeEnd = time.Now()

	if errX != nil {
		fmt.Printf(ePrefix + "%v", errX)
	}

	biNumBase, errX := mathops.BigIntNum{}.NewBigIntPrecision(
		base,
		basePrecision)

	if errX != nil {
		fmt.Printf(ePrefix + "Error: biNumBase - %v", errX)
	}

	biXNum, errX := mathops.BigIntNum{}.NewBigIntPrecision(
		xNum,
		xNumPrecision)

	if errX != nil {
		fmt.Printf(ePrefix + "Error: biNumBase - %v", errX)
	}

	biNumLogResult, errX := mathops.BigIntNum{}.NewBigIntPrecision(
			logResult,
			logResultPrecision)

	if errX != nil {
		fmt.Printf(ePrefix + "Error: biNumLogResult - %v", errX)
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	status := "Success!!!!! Expected Matches Actual Value"

	fmt.Println()
	fmt.Println("================================================")
	fmt.Println("   BigIntMathLogarithms{}.LogBaseOfXByDivide() ")
	fmt.Println("================================================")
	fmt.Println("                base: ", biNumBase.GetNumStr())
	fmt.Println("                xNum: ", biXNum.GetNumStr())
	fmt.Println("maxInternalPrecision: ", maxInternalPrecision.Text(10))
	fmt.Println("        maxPrecision: ", maxPrecision.Text(10))
	fmt.Println("              cycles: ", cycles)
	fmt.Println("          log Result: ", biNumLogResult.GetNumStr())
	fmt.Println("     expected Result: ", expectedValue)
	fmt.Println("      Execution Time: ", duration)

	if expectedValue != biNumLogResult.GetNumStr() {
		status = "FAILURE****** Expected DOES NOT MATCH Actual Value"
	}

	fmt.Println(status)
	fmt.Println("================================================")
	fmt.Println()

}


func TestBigIntLogBaseOfX(
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxPrecision *big.Int,
	expectedValue string) {

	ePrefix := "TestBigIntLogBaseOfX() "
	timeStart := time.Now()
	timeEnd := time.Now()

	timeStart = time.Now()
	logResult, logResultPrecision, err :=
		mathops.BigIntMathLogarithms{}.BigIntLogBaseOfX(
			base,
			basePrecision,
			xNum,
			xNumPrecision,
			maxPrecision)

	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v ", err.Error())
	}

	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	status := "Success!!!!! Expected Matches Actual Value"

	if err == nil {
		binLogResult, _ := mathops.BigIntNum{}.NewBigIntPrecision(logResult, logResultPrecision)

		fmt.Println()
		fmt.Println("================================================")
		fmt.Println("    BigIntMathLogarithms{}.BigIntLogBaseOfX() ")
		fmt.Println("================================================")
		fmt.Println("           base: ", base.Text(10))
		fmt.Println("           xNum: ", xNum.Text(10))
		fmt.Println("  xNumPrecision: ", xNumPrecision.Text(10))
		fmt.Println("      logResult: ", logResult.Text(10))
		fmt.Println("   logPrecision: ", logResultPrecision.Text(10))
		fmt.Println("      logNumStr: ", binLogResult.GetNumStr())
		fmt.Println("Expected Result: ", expectedValue)
		fmt.Println(" Execution Time: ", duration)
		fmt.Println("================================================")
		if expectedValue != binLogResult.GetNumStr() {
			status = "FAILURE****** Expected DOES NOT MATCH Actual Value"
		}

		fmt.Println(status)
		fmt.Println("================================================")
		fmt.Println()
	}

	biBase := mathops.BigIntNum{}.NewBigInt(base, 0)
	biXNum, _ := mathops.BigIntNum{}.NewBigIntPrecision(xNum, xNumPrecision)

	uiMaxPrecision := uint(maxPrecision.Uint64())

	timeStart = time.Now()

	biNumResult, err := mathops.BigIntMathLogarithms{}.BigIntNumLogBaseOfX(
		biBase,
		biXNum,
		uiMaxPrecision)
	timeEnd = time.Now()

	if err != nil {
		fmt.Printf(ePrefix + "%v ", err.Error())
		return
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)
	status = "Success!!!!! Expected Matches Actual Value"

	fmt.Println()
	fmt.Println("------------------------------------------------")
	fmt.Println("      BigIntMathLogarithms.BigIntNumLogBaseOfX() ")
	fmt.Println("------------------------------------------------")
	fmt.Println("BigIntNum Log Result: ", biNumResult.GetNumStr())
	fmt.Println("     Expected Result: ", expectedValue)
	fmt.Println("      Execution Time: ", duration)
	fmt.Println("------------------------------------------------")
	if expectedValue != biNumResult.GetNumStr() {
		status = "FAILURE****** Expected DOES NOT MATCH Actual Value"
	}
	fmt.Println(status)
	fmt.Println("------------------------------------------------")
	fmt.Println()
}



/*


func main () {

	fde := mathops.GetEulersNum20k()

	base := big.NewInt(10)
	basePrecision := big.NewInt(0)
	xNum := fde.GetInteger()
	xNumPrecision := fde.GetPrecisionBigInt()
	maxPrecision := big.NewInt(1000)
	factor := big.NewInt(4)
	cycles := big.NewInt(0).Mul(maxPrecision, factor)
	maxInternalPrecision := big.NewInt(0).Mul(cycles, factor)

	expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418688"



	TestBinaryLogBaseOfX(
		base,
		basePrecision,
		xNum,
		xNumPrecision,
		maxInternalPrecision,
		maxPrecision,
		cycles,
		expectedValue)


		m := big.NewInt(18)
	s4DivPrecisionMaxInternalPrecision := big.NewInt(6000)
	agMeanMaxInternalPrecision := big.NewInt(2000)
	agMeanMaxOutputPrecision := big.NewInt(300)
	piDivideMaxPrecision := big.NewInt(2000)
	maxFinalResultPrecision := big.NewInt(99)

	TestNatLogOfXArithmeticGeometricMean(
		xNum,
		xNumPrecision,
		m,
		s4DivPrecisionMaxInternalPrecision,
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		piDivideMaxPrecision,
		maxFinalResultPrecision,
		expectedValue)

}


func main() {

	fdEuler := mathops.GetEulersNum1k()

	xNum := big.NewInt(1500000)
	xNumPrecision := big.NewInt(0)
	maxInternalPrecision := big.NewInt(500)
	maxPrecision := big.NewInt(30)
	//                            1         2         3
	//                   1234567890123456789012345678901
	expectedValue := "14.220975666072438486085961843571"

	TestBigIntLogBaseOfX(
		fdEuler.GetInteger(),
		fdEuler.GetPrecisionBigInt(),
		xNum,
		xNumPrecision,
		maxInternalPrecision,
		maxPrecision,
		expectedValue)

}


func main() {
	ePrefix := "main() "

	num1 := big.NewInt(762939453125)
	num1Precision := big.NewInt(17)

	num2 := big.NewInt(1)
	num2Precision := big.NewInt(0)

	product, productPrecision, err :=
		mathops.BigIntMathMultiply{}.BigIntMultiply(num1, num1Precision, num2, num2Precision)

	if err !=nil {
		fmt.Printf(ePrefix + "%v", err.Error())
		return
	}

	binProduct, err := mathops.BigIntNum{}.NewBigIntPrecision(product, productPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "binProduct - %v", err.Error())
		return
	}

	fmt.Println("Product: ", binProduct.GetNumStr())
	maxInternalPrecision := big.NewInt(6000)


	fdNRoot := mathops.FixedDecimalNthRoot{}
	g, gPrecision, err :=
		fdNRoot.CalculatePositiveIntegerNthRoot(
			product,
			productPrecision,
			big.NewInt(2),
			big.NewInt(0),
			maxInternalPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "nthRoot - %v", err.Error())
		return
	}

	binG, err := mathops.BigIntNum{}.NewBigIntPrecision(g, gPrecision)

	if err !=nil {
		fmt.Printf(ePrefix + "binG - %v", err.Error())
		return
	}

	fmt.Println("binG: ", binG.GetNumStr())

}



func main() {
	xNum := big.NewInt(2)
	xNumPrecision := big.NewInt(0)
	m := big.NewInt(193)
	s4DivPrecisionMaxInternalPrecision := big.NewInt(8000)
	agMeanMaxInternalPrecision := big.NewInt(6000)
	agMeanMaxOutputPrecision := big.NewInt(800)
	maxPiDividePrecision := big.NewInt(15000)
	maxFinalOutputPrecision := big.NewInt(99)
	expectedValue := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418688"

	TestNatLogOfXArithmeticGeometricMean(
		xNum,
		xNumPrecision,
		m,
		s4DivPrecisionMaxInternalPrecision,
		agMeanMaxInternalPrecision,
		agMeanMaxOutputPrecision,
		maxPiDividePrecision,
		maxFinalOutputPrecision,
		expectedValue)

}

*/



func TestBigIntToNegativeFractionalPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	expectedResult string) {


	timeStart := time.Now()
	result,
	resultPrecision,
	err := mathops.BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	timeEnd := time.Now()

	if err != nil {
		fmt.Printf("%v ", err.Error())
	}

	binResult := mathops.BigIntNum{}.NewBigInt(result, uint(resultPrecision.Uint64()))
	timeDuration := timeEnd.Sub(timeStart)

	duration := examples.CodeDurationToStr(timeDuration)

	fmt.Println()
	fmt.Println()
	fmt.Println("BigIntMathPower{}.BigIntToNegativeFractionalPower() ")
	fmt.Println("============================================================")
	fmt.Println("                  base: ", base.Text(10))
	fmt.Println("         basePrecision: ", basePrecision.Text(10))
	fmt.Println("              exponent: ", exponent.Text(10))
	fmt.Println("     exponentPrecision: ", exponentPrecision.Text(10))
	fmt.Println("     Maximum Precision: ", maxPrecision.Text(10))
	fmt.Println("------------------------------------------------------------")
	fmt.Println("                result: ", result.Text(10))
	fmt.Println("       resultPrecision: ", resultPrecision)
	fmt.Println("         result NumStr: ", binResult.GetNumStr())
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()

	binBase, _ := mathops.BigIntNum{}.NewBigIntPrecision(base, basePrecision)
	binExponent, _ := mathops.BigIntNum{}.NewBigIntPrecision(exponent, exponentPrecision)

	timeStart = time.Now()
	binPwr, err := mathops.BigIntMathPower{}.Pwr(binBase, binExponent, uint(maxPrecision.Uint64()))
	timeEnd = time.Now()
	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(...) " +
			"Error='%v' ", err.Error())
	}

	timeDuration = timeEnd.Sub(timeStart)

	duration = examples.CodeDurationToStr(timeDuration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println( "               BigIntMathPower{}.Pwr() ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("     BigIntNum  result: ", binPwr.GetNumStr())
	fmt.Println("   BigIntNum precision: ", binPwr.GetPrecision())
	fmt.Println("        expectedResult: ", expectedResult)
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Time: ", duration)
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
}
