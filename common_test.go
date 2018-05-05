package incenc

import "testing"

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestFindSuffix(t *testing.T) {
	cases := []struct {
		last, value string
		expect      string
	}{
		{"foo", "foobar", "bar"},
		{"", "lol", "lol"},
		{"some letters", "some other", "other"},
	}

	for i, test := range cases {
		_, got := findSuffix(test.last, test.value)
		if got != test.expect {
			t.Errorf("%d: findSuffix(%q, %q) = %q instead of %q",
				i, test.last, test.value, got, test.expect)
		}
	}
}

func TestVarint(t *testing.T) {
	for i := 0; i < 32768; i++ {
		var buf [2]byte
		data := writeVarint(buf[:0], uint16(i))
		if i < 128 && len(data) != 1 {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "wrote too many bytes")
		}
		data, x, err := readVarint(buf[:])
		assertNoError(t, err)
		if i < 128 && len(data) != 1 {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "read too many bytes")
		}
		if int(x) != i {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "failed to round trip", x)
		}
	}
}

func TestReaderWriter(t *testing.T) {
	var (
		buf []byte
		r   Reader
		w   Writer
	)

	buf, _ = w.Append(buf, "hello")
	buf, _ = w.Append(buf, "hello.world")
	buf, _ = w.Append(buf, "hello.woopy")

	buf, value, err := r.Next(buf)
	assertNoError(t, err)
	if string(value) != "hello" {
		t.FailNow()
	}

	buf, value, err = r.Next(buf)
	assertNoError(t, err)
	if string(value) != "hello.world" {
		t.FailNow()
	}

	buf, value, err = r.Next(buf)
	assertNoError(t, err)
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

func encodeCorpus(buf []byte) []byte {
	var w Writer
	for _, v := range corpus {
		buf, _ = w.Append(buf, v)
	}
	return buf
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
