package main

var testinput = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

var input = []string{
	"8897989158773769986683317126949895844789898919988982846976854887698859721879997687955719111598998327",
	"9649369117599598197939482994919879998758884941925999937589283638799993484872917997897895739989579191",
	"7997329984997993983998692558683997874541922699389423763943169877296129741998999881939998256858971899",
	"3961983578889685987698899271964428529363998189639145348977949944524459593642285991798688299996794386",
	"9637694796792848698585436766693878116811574545996917988849678559889871345964796779657568676781986679",
	"7169788879791193287869596821795996868629892967969999951719973757978791815589717829379465298175531995",
	"8911787949488631799192297917351799963981912889735791891541869999472998464999691489698987859927889898",
	"4128948594143972375996798887399894752996996987988978918577298969744469496461537678996579579882595217",
	"5965817119735799927884896787876486393697742894668989929487117768359565388189518976777917559992172847",
	"2989885771199621179958199115612918624488196939349672464576989555825999788599837879438936611447999839",
	"6279897824925936956497342939378927597776898676843695669658989678881993725352579797548221858547715949",
	"5989296899714187663899166299643999588599189329596992791764919989779791999799245592599795577991769987",
	"9962655586397457644489849698996938758335479848939789666999187919956988594762939617883669879142615876",
	"1776535958888489999868889935876891995917858573888141932892871929759685633999998783565999664898445899",
	"1389588539879994918966639979248925164766861119932397884836728574914889698678748159786666526977937873",
	"2965987192899975781574967992744999989784981349689999231886864868548827859776478123744977911399789781",
	"2864838571839387537669651179791649917979994298891689815735289458799788138892176795999379986975965639",
	"9849964747772999955155291796985999656537989995993851869288898689974989868949679538961555997819939889",
	"4994773458186796316251998998878929171685897818685474989269988899995133196617738959599596865237439195",
	"5898994686666887878697969493772479368545911286595687956651945494868674919679885477778964944949648993",
	"2968958619986559978978778973867646949389849239865925991515997677989945694184537469978349868199966258",
	"9398876849795467968545494546578819796379869917199941697762253999842889893186493838755831695194769779",
	"6838985769869696983898917645938998184196834598388897588942923857959281928978982195499982739696266889",
	"3987693959997592799889599829997589889422991979916974746968791979464568312437885972816259879416599813",
	"5388591473477795949179899798594642376499979743895961139675599728572974356879536747977281975699395939",
	"8822995196987896478591777986718978966932949989753586986221987843693568677989666993578883788557594696",
	"9597298489969897175889696699818898778965999957452698762998437498959899776248997979989994696988795396",
	"7999187158967688944752967388399971975857958999992686479489997893989995929595499199477891972268484529",
	"1553498985489865967924958119458296883239463279624936498871699968997973857177784971987939915999414678",
	"8393269446357283299258661887899933929276719689459479677989982869991793997993338969994168329799419189",
	"7889959627679892765496786476969146999899666958896887715816797919391953358899877773882742829989797649",
	"9467868827849152723897169551896421198986889997397899563583997995971918839996526139342742627969178835",
	"4378992965797947991933988922911889449791679784979823698195996781269647815169671914787891967987597797",
	"7764427587729777924984118513385772497994598398399999535689997596487989488878379776879898241198126984",
	"9728895571759195284389899797669672699198898829989895795713678198716319818897885979938962939749761891",
	"6189727167897919896849798129965997294767953695198231887398386818118983699866599939872984585595639955",
	"9599695951755166926729299771994919224396769599687118496983387284956419896854998157992918463486919787",
	"9915698298772984877947527983257263849947198691736737884812819569949941589379999664999827369966666728",
	"9939557889768719894729449792298998994887857711866593991986349935299166999598991777958497899399899574",
	"9815699727649526949958988599988878388691778769979121665184967735579899858921958956764798348179979298",
	"8894997945935953298758899899873931992361439657589742797697598834633717592557186892648766514518392379",
	"7888984956718674639979383999991449648389659477389594987991273963988697983868456147688581392798839945",
	"6916778599189689557986592598821957866869819996893429464977987888439116289459699177748916971279368654",
	"9974199168449615693776359925869365841974294989613521799647438893719945923974211774958644996587997717",
	"7397929761618411679417329855688914986958978959669299815977982394779469595892197571288973599379772984",
	"3876999727584882198199195988893888949899718178219896199918999991999294969991759488199979148989479386",
	"7488988474969889495618878696725459797846572196836478574379713971799597999597845872189788978992149567",
	"9557999975997699497927658847991197156959987699588266886921938442847257159799999197949899898158947714",
	"5596898878863832871413581988799689419771982789789991898979898992599959936866999596299746661193998979",
	"9926819979315698241398994996999889821982759461723282187931993979936974796199382615585996969641947382",
	"1546868918694474796999129949751985819818999888794993967128215787839729839847799539855981856713297799",
	"4495427987936717516566722146639739667263979989458986994996879924473168758879727899879776718544593969",
	"2899699234844986314975254882399582996938997388684714295999971871999883397629897616838998341158266919",
	"2329818977586886752649934729588635519447787886889292338999736579988399166579682519798985619214387781",
	"9491299987978895559956919987319979968916299471347988568678571789469181845641879621574746799859789629",
	"3999736632569947176815919497546898949979986954841599541895899751793698835731877743186983996984224985",
	"3743313688597263279688866989269199886321995866995995589858759188688999176589967184989688957679539497",
	"9972998984997364796617437949126399619996392978478898398468986927478754958698978986689898937192399858",
	"8998793781668965959939753392828897918936542894371939677883692328754998293626192325898996526899687994",
	"3157988886115989871158918789725684983968899598788931467659698949817279699792979497899987714999788764",
	"8993892698399791169437138893433545889799119249758749791986962929896918696891787159789977785416999856",
	"5521459816853999796881784898696959751339996612946789911979836669986737269681151778269977221739479817",
	"4987619666714799998371871979299997899799928999512436682298428979919579779862929676917298989129995797",
	"4924761924765989731976697798359994919447958743739577579897629989996797377885973858391669877913792899",
	"9983572598659769918988896955387484757588285588976968191976379418848242639396189669961989528969998488",
	"4112154298959695758886398979194987911731968999598348995347779992737716991999494988834679861595175994",
	"8353769619991495481963988969875712971582797279459879731845482818898412897196947959586999397575856839",
	"9387859349249999195997399821836399568218982571719439981555958988568876779642124181999199619926698953",
	"9395997865298198395949888788497538785893399566994324899989999787498192917963957929999997897678378248",
	"6636973198999893993977169998598699465424956325448997619693889468864878285992995997681998718661589294",
	"1532766997968738593697189847955956945795463175174999991474789896393888693718994292693896919751869977",
	"4987984421658796999755537263899537979776269959259993877689261339958592248277469297915283487939899897",
	"1776398739791998796976387581969977839897549999149396876195398921574819864159768125999176966424694675",
	"6389671598498696988664848846688898948556921591799959676789299484196858565978985617699718988964874431",
	"7719567561991896964999247368191744796379489983591872829711927959932769688198998998974918591391971959",
	"3973969556559682547737994768999196789396998718971353814399534374989297988781243896289295698928828289",
	"7991938827798258998889698669249141296999727767895942736589794368619948647977721566856259891837992999",
	"9347419999879498893988812718491986859569797827971899988519988799499899959199874922469949197399887867",
	"5285685399589818577328728994816859944779479548486618885639248585143899218998991745966967556743799663",
	"9395854896927798161978679894172991341633558868997681854999178398996742798559818618813166597774989989",
	"6638744792557979498599818846168699921259919875769961687442979511277755592869199448671599859934555816",
	"7859899999581981938396949188778913618938999895939949914358964839219398789234488821981969965879571199",
	"5768298685591993956399669199469713988389469658916859481718665158897887593518986874897665763895496978",
	"7389967287945171649228411182199879977995457494927297278551687468479555273595191958798197888978552819",
	"1986968998719395299973659111849196996147927191817876618544749943958598569889381987897128898941867512",
	"9928499197444497999978914519999659127459992979275198892159113931949699979514897589277869582174849793",
	"4348762191999362996487969856399964994199788889975657791432777377119886925995669989499979489629782381",
	"9986694999824187679929418558578891199561999784419463587819572299698998159559997791815899574318695888",
	"1693391916957213992197915814267941939267699787998996349859887428396996952958966899389269769999718589",
	"8799285291917998168938694785258971249352846666723399564967831982746448387677974939989399366896894981",
	"9194433917477997149968388197818149699471379996229999149448769995397971416962895419984998165959969589",
	"9997925584499397957667157718178985599959615969516697389777739967285999485299899788359789895771799922",
	"3169887497267866997513737793727435869821989699965967678257473826919878279459796139883856587941689717",
	"2262596439848199898732957698587579919159947927839682979998979498918999985468869519799888974886894699",
	"9299813478871628678829999886993469939654868121673699899798737319697878895997488917962139199759184642",
	"6892987492138253399698167927777782896789918998912979563663381926899981998898126979894767195714983179",
	"9849163895499946394788815977819834799369864978668581593347367894199786913278698391984887358789398959",
	"8172791198939899169998751197617887598959971683688955949699214349897894879896894979132272842976334299",
	"7355751274786528969795893559877799899919916375199842966163718893993893897739947698981113986695196918",
	"6237299196797895199514949678982959691497579891735645697992197498877449274999683887676581966579317679",
}
