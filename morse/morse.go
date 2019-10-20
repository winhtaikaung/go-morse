package morse

import (
	"encoding/json"
	"strings"
)

type RuleStruct struct {
	From string
	To   string
}

func Encode(str string) string {
	rule := `[{"from":"a","to":".- "},{"from":"b","to":"-... "},{"from":"c","to":"-.-. "},{"from":"d","to":"-.. "},{"from":"e","to":". "},{"from":"f","to":"..-. "},{"from":"g","to":"--. "},{"from":"h","to":".... "},{"from":"i","to":".. "},{"from":"j","to":".--- "},{"from":"k","to":"-.- "},{"from":"l","to":".-.. "},{"from":"m","to":"-- "},{"from":"n","to":"-. "},{"from":"o","to":"--- "},{"from":"p","to":".--. "},{"from":"q","to":"--.- "},{"from":"r","to":".-. "},{"from":"s","to":"... "},{"from":"t","to":"- "},{"from":"u","to":"..- "},{"from":"v","to":"...- "},{"from":"w","to":".-- "},{"from":"x","to":"-..- "},{"from":"y","to":"-.-- "},{"from":"z","to":"--.. "},{"from":"1","to":".---- "},{"from":"2","to":"..--- "},{"from":"3","to":"...-- "},{"from":"4","to":"....- "},{"from":"5","to":"..... "},{"from":"6","to":"-.... "},{"from":"7","to":"--... "},{"from":"8","to":"---.. "},{"from":"9","to":"----. "},{"from":"0","to":"----- "},{"from":" ","to":"/ "}]`
	return replaceWithStringReplace(rule, strings.ToLower(str), false)
}
func Decode(str string) string {
	rule := `[{"from":".-","to":"a"},{"from":"-...","to":"b"},{"from":"-.-.","to":"c"},{"from":"-..","to":"d"},{"from":".","to":"e"},{"from":"..-.","to":"f"},{"from":"--.","to":"g"},{"from":"....","to":"h"},{"from":"..","to":"i"},{"from":".---","to":"j"},{"from":"-.-","to":"k"},{"from":".-..","to":"l"},{"from":"--","to":"m"},{"from":"-.","to":"n"},{"from":"---","to":"o"},{"from":".--.","to":"p"},{"from":"--.-","to":"q"},{"from":".-.","to":"r"},{"from":"...","to":"s"},{"from":"-","to":"t"},{"from":"..-","to":"u"},{"from":"...-","to":"v"},{"from":".--","to":"w"},{"from":"-..-","to":"x"},{"from":"-.--","to":"y"},{"from":"--..","to":"z"},{"from":".----","to":"1"},{"from":"..---","to":"2"},{"from":"...--","to":"3"},{"from":"....-","to":"4"},{"from":".....","to":"5"},{"from":"-....","to":"6"},{"from":"--...","to":"7"},{"from":"---..","to":"8"},{"from":"----.","to":"9"},{"from":"-----","to":"0"},{"from":"/","to":" "}]`
	return replaceWithStringReplace(rule, strings.ToLower(str), true)
}

func replaceWithStringReplace(rule string, output string, isDecode bool) string {

	var ruleArr []RuleStruct
	err := json.Unmarshal([]byte(rule), &ruleArr)
	if err != nil {
		panic(err.Error())
	}
	text := output
	strArray := []string{""}
	seperator := ""
	if isDecode == true {
		seperator = ""
		strArray = strings.Split(text, " ")
	} else {
		seperator = " "
		strArray = strings.Split(text, "")
	}
	for i := range strArray {
		for c := range ruleArr {
			if strArray[i] == ruleArr[c].From {
				strArray[i] = ruleArr[c].To
			}
		}
	}
	return strings.Join(strArray, seperator)
}
