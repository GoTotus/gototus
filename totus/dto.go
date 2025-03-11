package totus

import (
	"encoding/json"
	"strconv"
)

type POI map[string]any

func (p POI) String() string {
	data, _ := json.MarshalIndent(p, "", "    ")
	return string(data)
}

type IPData map[string]any

func (i IPData) IP4() string {
	return MapStrWithDef(i, "ip4", "")
}

func (i IPData) IP6() string {
	return MapStrWithDef(i, "ip6", "")
}

func (i IPData) GH() string {
	return MapStrWithDef(i, "gh", "")
}

func (i IPData) String() string {
	data, _ := json.MarshalIndent(i, "", "    ")
	return string(data)
}

type ValidatedEmail map[string]any

func (v ValidatedEmail) String() string {
	data, _ := json.MarshalIndent(v, "", "    ")
	return string(data)
}

func (v ValidatedEmail) IsValid() bool {
	return MapStrWithDef(v, "result", "") == "PASSED"
}

func (v ValidatedEmail) Score() int {
	if score, err := strconv.Atoi(MapStrWithDef(v, "score", "0")); err == nil {
		return score
	}
	return 0
}

type CheckLevel string

// Constants for CheckLevel
const (
	CheckLevelL1Syntax CheckLevel = "l1_syntax"
	CheckLevelL2DNS    CheckLevel = "l2_dns"
	CheckLevelL3Server CheckLevel = "l3_server"
	CheckLevelL4Dbs    CheckLevel = "l4_dbs"
)
