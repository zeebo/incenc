package incenc

import "testing"

func TestCommonPrefix(t *testing.T) {
	cases := []struct {
		left, right string
		same        int
	}{
		{"foobar", "foo", 3},
		{"", "lol", 0},
		{"some letters", "some other", 5},
	}

	for i, test := range cases {
		got := commonPrefix(test.left, test.right)
		if got != test.same {
			t.Errorf("%d: commonPrefix(%q, %q) = %d instead of %d",
				i, test.left, test.right, got, test.same)
		}

		got = commonPrefix(test.right, test.left)
		if got != test.same {
			t.Errorf("%d: commonPrefix(%q, %q) = %d instead of %d",
				i, test.right, test.left, got, test.same)
		}
	}
}

func TestReaderWriter(t *testing.T) {
	var (
		buf []byte
		r   Reader
		w   Writer
	)

	buf = w.Append(buf, "hello")
	buf = w.Append(buf, "hello.world")
	buf = w.Append(buf, "hello.woopy")

	buf, value := r.Next(buf)
	if string(value) != "hello" {
		t.FailNow()
	}

	buf, value = r.Next(buf)
	if string(value) != "hello.world" {
		t.FailNow()
	}

	buf, value = r.Next(buf)
	if string(value) != "hello.woopy" {
		t.FailNow()
	}

	if len(buf) != 0 {
		t.FailNow()
	}
}

var corpusLength int

func init() {
	for _, v := range corpus {
		corpusLength += len(v)
	}
}

var corpus = []string{
	"abase",
	"abased",
	"abasedly",
	"abasedness",
	"abasement",
	"abasements",
	"abaser",
	"abasers",
	"abases",
	"abash",
	"abashed",
	"abashedly",
	"abashedness",
	"abashes",
	"abashing",
	"abashless",
	"abashlessly",
	"abashment",
	"abashments",
	"abasia",
	"abasias",
	"abasic",
	"abasing",
	"abasio",
	"abask",
	"abassi",
	"abastard",
	"abastardize",
	"abastral",
	"abatable",
	"abatage",
	"abated",
	"abatement",
	"abatements",
	"abater",
	"abaters",
	"abates",
	"abatic",
	"abating",
	"abatis",
	"abatised",
	"abatises",
	"abatjour",
	"abatjours",
	"abaton",
	"abator",
	"abators",
	"abattage",
	"abattis",
	"abattised",
	"abattises",
	"abattoir",
	"abattoirs",
	"abattu",
	"abattue",
	"abature",
	"abaue",
	"abave",
	"abaxial",
	"abaxile",
	"abaze",
	"abb",
	"abbacy",
	"abbacies",
	"abbacomes",
	"abbaye",
	"abbandono",
	"abbas",
	"abbasi",
	"abbassi",
	"abbatial",
	"abbatical",
	"abbatie",
	"abbeys",
	"abbeystead",
	"abbeystede",
	"abbes",
	"abbess",
	"abbesses",
	"abbest",
	"abboccato",
	"abbogada",
	"abbotcy",
	"abbotcies",
	"abbotnullius",
	"abbotric",
	"abbots",
	"abbotship",
	"abbotships",
	"abbozzo",
	"abbrev",
	"abbreviatable",
	"abbreviate",
	"abbreviated",
	"abbreviately",
	"abbreviates",
	"abbreviating",
	"abbreviation",
	"abbreviations",
	"abbreviator",
	"abbreviatory",
	"abbreviators",
	"abbreviature",
	"abbroachment",
	"abcess",
	"abcissa",
	"abcoulomb",
	"abd",
	"abdal",
	"abdali",
	"abdaria",
	"abdat",
	"abdest",
	"abdicable",
	"abdicant",
	"abdicate",
	"abdicated",
	"abdicates",
	"abdicating",
	"abdication",
	"abdications",
	"abdicative",
	"abdicator",
	"abditive",
	"abditory",
	"abdom",
	"abdomen",
	"abdomens",
	"abdomina",
	"abdominal",
	"abdominalia",
	"abdominalian",
	"abdominally",
	"abdominals",
	"abdominoanterior",
	"abdominocardiac",
	"abdominocentesis",
	"abdominocystic",
	"abdominogenital",
	"abdominohysterectomy",
	"abdominohysterotomy",
	"abdominoposterior",
	"abdominoscope",
	"abdominoscopy",
	"abdominothoracic",
	"abdominous",
	"abdominovaginal",
	"abdominovesical",
	"abduce",
	"abduced",
	"abducens",
	"abducent",
	"abducentes",
	"abduces",
	"abducing",
	"abduct",
	"abducted",
	"abducting",
	"abduction",
	"abductions",
	"abductor",
	"abductores",
	"abductors",
	"abducts",
	"abeam",
	"abear",
	"abearance",
	"abecedaire",
	"abecedary",
	"abecedaria",
	"abecedarian",
	"abecedarians",
	"abecedaries",
	"abecedarium",
	"abecedarius",
	"abed",
	"abede",
	"abedge",
	"abegge",
	"abeyance",
	"abeyances",
	"abeyancy",
	"abeyancies",
	"abeyant",
	"abeigh",
	"abele",
	"abeles",
	"abelmosk",
	"abelmosks",
	"abelmusk",
	"abeltree",
	"abend",
	"abends",
	"abenteric",
	"abepithymia",
	"aberdavine",
	"aberdevine",
	"aberduvine",
	"abernethy",
	"aberr",
	"aberrance",
	"aberrancy",
	"aberrancies",
	"aberrant",
	"aberrantly",
	"aberrants",
	"aberrate",
	"aberrated",
	"aberrating",
	"aberration",
	"aberrational",
	"aberrations",
	"aberrative",
	"aberrator",
	"aberrometer",
	"aberroscope",
	"aberuncate",
	"aberuncator",
	"abesse",
	"abessive",
	"abet",
	"abetment",
	"abetments",
	"abets",
	"abettal",
	"abettals",
	"abetted",
	"abetter",
	"abetters",
	"abetting",
	"abettor",
	"abettors",
	"abevacuation",
	"abfarad",
	"abfarads",
	"abhenry",
	"abhenries",
	"abhenrys",
	"abhinaya",
	"abhiseka",
	"abhominable",
	"abhor",
	"abhorred",
	"abhorrence",
	"abhorrences",
	"abhorrency",
	"abhorrent",
	"abhorrently",
	"abhorrer",
	"abhorrers",
	"abhorrible",
	"abhorring",
	"abhors",
	"aby",
}
