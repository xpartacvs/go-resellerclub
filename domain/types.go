package domain

import "github.com/xpartacvs/go-resellerclub/core"

type SortBy string
type DomainKey string

type PrivacyState string
type DomainRegistrationStatus string
type SortOrder map[SortBy]bool

type SuggestNames map[string]struct {
	Status string         `json:"status"`
	InGa   core.JSONBool  `json:"in_ga"`
	Score  core.JSONFloat `json:"score"`
	Spin   string         `json:"spin"`
}

type Contact struct {
	Company       string   `json:"company"`
	Address1      string   `json:"address1"`
	Telno         string   `json:"telno"`
	TelnoCC       string   `json:"telnocc"`
	ContactID     string   `json:"contactid"`
	Type          string   `json:"type"`
	ContactType   []string `json:"contacttype"`
	CustomerID    string   `json:"customerid"`
	Country       string   `json:"country"`
	ParentKey     string   `json:"parentkey"`
	ContactStatus string   `json:"contactstatus"`
	State         string   `json:"state"`
	EmailAddr     string   `json:"emailaddr"`
	City          string   `json:"city"`
	Name          string   `json:"name"`
	ZIP           string   `json:"zip"`
}

type OrderDetail struct {
	Classkey          string          `json:"classkey"`
	AllowDeletion     core.JSONBool   `json:"allowdeletion"`
	OrderID           string          `json:"orderid"`
	NoOfNameServers   core.JSONUint16 `json:"noOfNameServers"`
	ParentKey         string          `json:"parentkey"`
	ProductCategory   string          `json:"productcategory"`
	CurrentStatus     string          `json:"currentstatus"`
	DomainName        string          `json:"domainname"`
	Description       string          `json:"description"`
	MultilingualFlag  string          `json:"multilingualflag"`
	TNCRequired       core.JSONBool   `json:"tnc_required"`
	PremiumDNSEnabled core.JSONBool   `json:"premiumdnsenabled"`
	GDPR              struct {
		Enabled  core.JSONBool `json:"enabled"`
		Eligible core.JSONBool `json:"eligible"`
	} `json:"gdpr"`
	CustomerID                 string          `json:"customerid"`
	Addons                     []string        `json:"addons"`
	BulkWhoIsOptOut            string          `json:"bulkwhoisoptout"`
	TechContactID              string          `json:"techcontactid"`
	IsImmediateReseller        core.JSONBool   `json:"isImmediateReseller"`
	CreationTime               core.JSONTime   `json:"creationtime"`
	DNSSec                     []string        `json:"dnssec"`
	JumpConditions             []string        `json:"jumpConditions"`
	RaaVerificationStartTime   core.JSONTime   `json:"raaVerificationStartTime"`
	CNS                        struct{}        `json:"cns"`
	Paused                     core.JSONBool   `json:"paused"`
	Admincontact               Contact         `json:"admincontact"`
	BillingContactID           string          `json:"billingcontactid"`
	PrivacyProtectedAllowed    core.JSONBool   `json:"privacyprotectedallowed"`
	DomSecret                  string          `json:"domsecret"`
	PremiumDNSAllowed          core.JSONBool   `json:"premiumdnsallowed"`
	ServiceProviderID          string          `json:"serviceproviderid"`
	Classname                  string          `json:"classname"`
	ResellerCost               core.JSONUint16 `json:"resellercost"`
	OrderStatus                []string        `json:"orderstatus"`
	EaqID                      string          `json:"eaqid"`
	EndTime                    core.JSONTime   `json:"endtime"`
	BillingContact             Contact         `json:"billingcontact"`
	AutoRenewTermType          string          `json:"autoRenewTermType"`
	RaaVerificationStatus      string          `json:"raaVerificationStatus"`
	EntityID                   string          `json:"entityid"`
	Recurring                  core.JSONBool   `json:"recurring"`
	ProductKey                 string          `json:"productkey"`
	NS1                        string          `json:"ns1"`
	NS2                        string          `json:"ns2"`
	ActionCompleted            core.JSONUint16 `json:"actioncompleted"`
	RegistrantContact          Contact         `json:"registrantcontact"`
	EntityTypeID               string          `json:"entitytypeid"`
	AutoRenewAttemptDuration   core.JSONUint16 `json:"autoRenewAttemptDuration"`
	CustomerCost               core.JSONFloat  `json:"customercost"`
	DomainStatus               []string        `json:"domainstatus"`
	OrderSuspendedByParent     core.JSONBool   `json:"orderSuspendedByParent"`
	MoneyBackPeriod            core.JSONUint16 `json:"moneybackperiod"`
	TechContact                Contact         `json:"techcontact"`
	RegistrantContactID        string          `json:"registrantcontactid"`
	AdminContactID             string          `json:"admincontactid"`
	IsOrderSuspendedUponExpiry core.JSONBool   `json:"isOrderSuspendedUponExpiry"`
	IsPrivacyProtected         core.JSONBool   `json:"isprivacyprotected"`
}

type ModifyNameServersResponse struct {
	ActionTypeDesc   string `json:"actiontypedesc"`
	EntityID         string `json:"entityid"`
	ActionStatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	EaqID            string `json:"eaqid"`
	CurrentAction    string `json:"currentaction"`
	Description      string `json:"description"`
	ActionType       string `json:"actiontype"`
	ActionStatusDesc string `json:"actionstatusdesc"`
}

const (
	SortByOrderID          SortBy = "orderid"
	SortByCustomerID       SortBy = "customerid"
	SortByEndTime          SortBy = "endtime"
	SortByTimestamp        SortBy = "timestamp"
	SortByEntityTypeID     SortBy = "entitytypeid"
	SortByCreationTime     SortBy = "creationtime"
	SortByCreationDateTime SortBy = "creationdt"

	DotASIA               DomainKey = "dotasia"
	DotAU3rd              DomainKey = "thirdleveldotau"
	DotBERLIN             DomainKey = "dotberlin"
	DotBID                DomainKey = "dotbid"
	DotBIZ                DomainKey = "dombiz"
	DotBEST               DomainKey = "dotbest"
	DotBUZZ               DomainKey = "dotbuzz"
	DotBZ                 DomainKey = "dotbz"
	DotBZ3rd              DomainKey = "thirdleveldotbz"
	DotCA                 DomainKey = "dotca"
	DotCC                 DomainKey = "dotcc"
	DotCLUB               DomainKey = "dotclub"
	DotCN                 DomainKey = "dotcn"
	DotCN3rd              DomainKey = "thirdleveldotcn"
	DotCNDotCOM           DomainKey = "centralniccncom"
	DotCO                 DomainKey = "dotco"
	DotCO3rd              DomainKey = "thirdleveldotco"
	DotCODotCOM           DomainKey = "centralniccocom"
	DotCOM                DomainKey = "domcno"
	DotCOMDotDE           DomainKey = "centralniccomde"
	DotCOOP               DomainKey = "dotcoop"
	DotDE                 DomainKey = "dotde"
	DotDESI               DomainKey = "dotdesi"
	DotES                 DomainKey = "dotes"
	DotEU                 DomainKey = "doteu"
	DotHN                 DomainKey = "dothn"
	DotHN3rd              DomainKey = "thirdleveldothn"
	DotIN                 DomainKey = "dotin"
	DotIN3rd              DomainKey = "thirdleveldotin"
	DotINDotNET           DomainKey = "indotnet"
	DotINFO               DomainKey = "dominfo"
	DotME                 DomainKey = "dotme"
	DotME3rd              DomainKey = "thirdleveldotme"
	DotMN                 DomainKey = "dotmn"
	DotMOBI               DomainKey = "dotmobi"
	DotNAME               DomainKey = "dotname"
	DotNET                DomainKey = "dotnet"
	DotNL                 DomainKey = "dotnl"
	DotNYC                DomainKey = "dotnyc"
	DotNZ3rd              DomainKey = "thirdleveldotnz"
	DotOOO                DomainKey = "dotooo"
	DotORG                DomainKey = "domorg"
	DotPRO                DomainKey = "dotpro"
	DotPW                 DomainKey = "dotpw"
	DotQUEBEC             DomainKey = "dotquebec"
	DotRU                 DomainKey = "dotru"
	DotRU3rd              DomainKey = "thirdleveldotru"
	DotSC                 DomainKey = "dotsc"
	DotSX                 DomainKey = "dotsx"
	DotTEL                DomainKey = "dottel"
	DotTRADE              DomainKey = "dottrade"
	DotTV                 DomainKey = "dottv"
	DotUK                 DomainKey = "dotuk"
	DotUK3rd              DomainKey = "thirdleveldotuk"
	DotUNO                DomainKey = "dotuno"
	DotUS                 DomainKey = "domus"
	DotVC                 DomainKey = "dotvc"
	DotWEBCAM             DomainKey = "dotwebcam"
	DotWEBDotIN           DomainKey = "premiumdotin"
	DotWS                 DomainKey = "dotws"
	DotXXX                DomainKey = "dotxxx"
	DotCentralNicPremium  DomainKey = "centralnicpremium"
	DotCentralNicStandard DomainKey = "centralnicstandard"
	DotUSDotCOM           DomainKey = "centralnicuscom"
	DotZADotCOM           DomainKey = "centralniczacom"
	DotUKDotCOM           DomainKey = "centralnicukcom"
	DotDonutsGroup1       DomainKey = "donutsgroup1"
	DotDonutsGroup2       DomainKey = "donutsgroup2"
	DotACCOUNTANTS        DomainKey = "dotaccountants"
	DotCASINO             DomainKey = "dotcasino"
	DotCREDIT             DomainKey = "dotcredit"
	DotCREDITCARD         DomainKey = "dotcreditcard"
	DotENERGY             DomainKey = "dotenergy"
	DotGOLD               DomainKey = "dotgold"
	DotINVESTMENTS        DomainKey = "dotinvestments"
	DotLOANS              DomainKey = "dotloans"
	DotPICTURES           DomainKey = "dotpictures"
	DotTIRES              DomainKey = "dottires"
	DotACTOR              DomainKey = "dotactor"
	DotAIRFORCE           DomainKey = "dotairforce"
	DotARMY               DomainKey = "dotarmy"
	DotATTORNEY           DomainKey = "dotattorney"
	DotAUCTION            DomainKey = "dotauction"
	DotBAND               DomainKey = "dotband"
	DotCONSULTING         DomainKey = "dotconsulting"
	DotDANCE              DomainKey = "dotdance"
	DotDEGREE             DomainKey = "dotdegree"
	DotDEMOCRAT           DomainKey = "dotdemocrat"
	DotDENTIST            DomainKey = "dotdentist"
	DotENGINEER           DomainKey = "dotengineer"
	DotFORSALE            DomainKey = "dotforsale"
	DotFUTBOL             DomainKey = "dotfutbol"
	DotGIVES              DomainKey = "dotgives"
	DotHAUS               DomainKey = "dothaus"
	DotIMMOBILIEN         DomainKey = "dotimmobilien"
	DotKAUFEN             DomainKey = "dotkaufen"
	DotLAWYER             DomainKey = "dotlawyer"
	DotLIVE               DomainKey = "dotlive"
	DotMARKET             DomainKey = "dotmarket"
	DotMODA               DomainKey = "dotmoda"
	DotMORTGAGE           DomainKey = "dotmortgage"
	DotNAVY               DomainKey = "dotnavy"
	DotNEWS               DomainKey = "dotnews"
	DotNINJA              DomainKey = "dotninja"
	DotPUB                DomainKey = "dotpub"
	DotREHAB              DomainKey = "dotrehab"
	DotREPUBLICAN         DomainKey = "dotrepublican"
	DotREVIEWS            DomainKey = "dotreviews"
	DotRIP                DomainKey = "dotrip"
	DotROCKS              DomainKey = "dotrocks"
	DotSALE               DomainKey = "dotsale"
	DotSOCIAL             DomainKey = "dotsocial"
	DotVET                DomainKey = "dotvet"
	DotVIDEO              DomainKey = "dotvideo"
	DotBLACK              DomainKey = "dotblack"
	DotBLUE               DomainKey = "dotblue"
	DotGLOBAL             DomainKey = "dotglobal"
	DotGREEN              DomainKey = "dotgreen"
	DotKIM                DomainKey = "dotkim"
	DotLIGHT              DomainKey = "dotlgbt"
	DotPINK               DomainKey = "dotpink"
	DotPOKER              DomainKey = "dotpoker"
	DotRED                DomainKey = "dotred"
	DotSHIKSHA            DomainKey = "dotshiksha"
	DotVEGAS              DomainKey = "dotvegas"
	DotVOTE               DomainKey = "dotvote"
	DotVOTO               DomainKey = "dotvoto"
	DotChineseONLINE      DomainKey = "dotchineseonline"
	DotChineseMOBILE      DomainKey = "dotmobile"
	DotChineseWEBSITE     DomainKey = "dotchinesewebsite"
	DotBEER               DomainKey = "dotbeer"
	DotCASA               DomainKey = "dotcasa"
	DotCOOKING            DomainKey = "dotcooking"
	DotCOUNTRY            DomainKey = "dotcountry"
	DotFASHION            DomainKey = "dotfashion"
	DotFISHING            DomainKey = "dotfishing"
	DotFIT                DomainKey = "dotfit"
	DotGARDEN             DomainKey = "dotgarden"
	DotHORSE              DomainKey = "dothorse"
	DotKIWI               DomainKey = "dotkiwi"
	DotLONDON             DomainKey = "dotlondon"
	DotRODEO              DomainKey = "dotrodeo"
	DotSURF               DomainKey = "dotsurf"
	DotWEDDING            DomainKey = "dotwedding"
	DotWORK               DomainKey = "dotwork"
	DotYOGA               DomainKey = "dotyoga"
	DotAUDIO              DomainKey = "dotaudio"
	DotBLACKFRIDAY        DomainKey = "dotblackfriday"
	DotCLICK              DomainKey = "dotclick"
	DotDIET               DomainKey = "dotdiet"
	DotFLOWERS            DomainKey = "dotflowers"
	DotGAME               DomainKey = "dotgame"
	DotGIFT               DomainKey = "dotgift"
	DotGUITARS            DomainKey = "dotguitars"
	DotHELP               DomainKey = "dothelp"
	DotHIPHOP             DomainKey = "dothiphop"
	DotHOSTING            DomainKey = "dothosting"
	DotJUEGOS             DomainKey = "dotjuegos"
	DotLINK               DomainKey = "dotlink"
	DotLOL                DomainKey = "dotlol"
	DotPHOTO              DomainKey = "dotphoto"
	DotPICS               DomainKey = "dotpics"
	DotPROPERTY           DomainKey = "dotproperty"
	DotSEXY               DomainKey = "dotsexy"
	DotTATOO              DomainKey = "dottattoo"
	DotBUILD              DomainKey = "dotbuild"
	DotLUXURY             DomainKey = "dotluxury"
	DotMEN                DomainKey = "dotmen"
	DotMENU               DomainKey = "dotmenu"
	DotONE                DomainKey = "dotone"
	DotSHABAKA            DomainKey = "dotshabaka"
	DotBAR                DomainKey = "dotbar"
	DotCOLLEGE            DomainKey = "dotcollege"
	DotDESIGN             DomainKey = "dotdesign"
	DotONLINE             DomainKey = "dotonline"
	DotPRESS              DomainKey = "dotpress"
	DotRENT               DomainKey = "dotrent"
	DotREST               DomainKey = "dotrest"
	DotSITE               DomainKey = "dotsite"
	DotSPACE              DomainKey = "dotspace"
	DotWEBSITE            DomainKey = "dotwebsite"
	DotCAREER             DomainKey = "dotcareer"
	DotJOBS               DomainKey = "dotjobs"
	DotMARKETS            DomainKey = "dotmarkets"
	DotHindiBHARAT        DomainKey = "dotbharat"
	DotRussianORG         DomainKey = "dotcyrillicorg"
	DotChineseORG         DomainKey = "dotchineseorg"
	DotHindiORG           DomainKey = "dotsangathan"
	DotADULT              DomainKey = "dotadult"
	DotAMSTERDAM          DomainKey = "dotamsterdam"
	DotARCHI              DomainKey = "dotarchi"
	DotBIO                DomainKey = "dotbio"
	DotCAPETOWN           DomainKey = "dotcapetown"
	DotCYMRU              DomainKey = "dotcymru"
	DotDURBAN             DomainKey = "dotdurban"
	DotIRISH              DomainKey = "dotirish"
	DotJOBURG             DomainKey = "dotjoburg"
	DotLA                 DomainKey = "centralnicdotla"
	DotNAGOYA             DomainKey = "dotnagoya"
	DotPORN               DomainKey = "dotporn"
	DotSKI                DomainKey = "dotski"
	DotSOFTWARE           DomainKey = "dotsoftware"
	DotSOY                DomainKey = "dotsoy"
	DotSTUDIO             DomainKey = "dotstudio"
	DotTIROL              DomainKey = "dottirol"
	DotTOKYO              DomainKey = "dottokyo"
	DotTOP                DomainKey = "dottop"
	DotWALES              DomainKey = "dotwales"
	DotWANG               DomainKey = "dotwang"
	DotWIKI               DomainKey = "dotwiki"
	DotXYZ                DomainKey = "dotxyz"
	DotPHOTOGRAPHY        DomainKey = "dotphotography"
	DotSYSTEMS            DomainKey = "dotsystems"
	DotCENTER             DomainKey = "dotcenter"
	DotEMAIL              DomainKey = "dotemail"
	DotCOMPANY            DomainKey = "dotcompany"
	DotSOLUTIONS          DomainKey = "dotsolutions"
	DotTIP                DomainKey = "dottips"
	DotTODAY              DomainKey = "dottoday"
	DotCITY               DomainKey = "dotcity"
	DotBUSINESS           DomainKey = "dotbusiness"
	DotASSOCIATES         DomainKey = "dotassociates"
	DotBIKE               DomainKey = "dotbike"
	DotPLUMBING           DomainKey = "dotplumbing"
	DotGURU               DomainKey = "dotguru"
	DotCAMERA             DomainKey = "dotcamera"
	DotACADEMY            DomainKey = "dotacademy"
	DotWORLD              DomainKey = "dotworld"
	DotTOYS               DomainKey = "dottoys"
	DotCAMP               DomainKey = "dotcamp"
	DotKITCHEN            DomainKey = "dotkitchen"
	DotSHOES              DomainKey = "dotshoes"
	DotBARGAINS           DomainKey = "dotbargains"
	DotCOFFEE             DomainKey = "dotcoffee"
	DotGLASS              DomainKey = "dotglass"
	DotSOLAR              DomainKey = "dotsolar"
	DotZONE               DomainKey = "dotzone"
	DotCOOL               DomainKey = "dotcool"
	DotWORKS              DomainKey = "dotworks"
	DotCLEANING           DomainKey = "dotcleaning"
	DotTOOLS              DomainKey = "dottools"
	DotCASH               DomainKey = "dotcash"
	DotLIFE               DomainKey = "dotlife"
	DotDOG                DomainKey = "dotdog"
	DotCOUPONS            DomainKey = "dotcoupons"

	PrivacyEnabled     PrivacyState = "true"
	PrivacyDisabled    PrivacyState = "false"
	PrivacyUnsupported PrivacyState = "na"

	DomRegUnknown       DomainRegistrationStatus = "unknown"
	DomRegUnregistered  DomainRegistrationStatus = "available"
	DomRegThroughUs     DomainRegistrationStatus = "regthroughus"
	DomRegThroughOthers DomainRegistrationStatus = "regthroughothers"
)
