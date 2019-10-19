package morse

import (
	"encoding/json"
	"regexp"
)

type RuleStruct struct {
	From string
	To   string
}

func morseEncode(str string) string {
	rule := `[{"from":"a","to":".-"},{"from":"b","to":"-..."},{"from":"c","to":"-.-."},{"from":"d","to":"-.."},{"from":"e","to":"."},{"from":"f","to":"..-."},{"from":"g","to":"--."},{"from":"h","to":"...."},{"from":"i","to":".."},{"from":"j","to":".---"},{"from":"k","to":"-.-"},{"from":"l","to":".-.."},{"from":"m","to":"--"},{"from":"n","to":"-."},{"from":"o","to":"---"},{"from":"p","to":".--."},{"from":"q","to":"--.-"},{"from":"r","to":".-."},{"from":"s","to":"..."},{"from":"t","to":"-"},{"from":"u","to":"..-"},{"from":"v","to":"...-"},{"from":"w","to":".--"},{"from":"x","to":"-..-"},{"from":"y","to":"-.--"},{"from":"z","to":"--.."},{"from":"1","to":".----"},{"from":"2","to":"..---"},{"from":"3","to":"...--"},{"from":"4","to":"....-"},{"from":"5","to":"....."},{"from":"6","to":"-...."},{"from":"7","to":"--..."},{"from":"8","to":"---.."},{"from":"9","to":"----."},{"from":"0","to":"-----"}]`

}

func morseDecode(str string) string {
	rule := `[{"from":".-","to":"a"},{"from":"-...","to":"b"},{"from":"-.-.","to":"c"},{"from":"-..","to":"d"},{"from":".","to":"e"},{"from":"..-.","to":"f"},{"from":"--.","to":"g"},{"from":"....","to":"h"},{"from":"..","to":"i"},{"from":".---","to":"j"},{"from":"-.-","to":"k"},{"from":".-..","to":"l"},{"from":"--","to":"m"},{"from":"-.","to":"n"},{"from":"---","to":"o"},{"from":".--.","to":"p"},{"from":"--.-","to":"q"},{"from":".-.","to":"r"},{"from":"...","to":"s"},{"from":"-","to":"t"},{"from":"..-","to":"u"},{"from":"...-","to":"v"},{"from":".--","to":"w"},{"from":"-..-","to":"x"},{"from":"-.--","to":"y"},{"from":"--..","to":"z"},{"from":".----","to":"1"},{"from":"..---","to":"2"},{"from":"...--","to":"3"},{"from":"....-","to":"4"},{"from":".....","to":"5"},{"from":"-....","to":"6"},{"from":"--...","to":"7"},{"from":"---..","to":"8"},{"from":"----.","to":"9"},{"from":"-----","to":"0"}]`
}

//Replace the string by matching arrays of regex
func replaceWithRule(rule string, output string) string {

	var ruleArr []RuleStruct
	err := json.Unmarshal([]byte(rule), &ruleArr)

	if err != nil {
		panic(err.Error())
	}
	re2 := regexp.MustCompile("") //declare and assign the Regexpbefore as it is expensive to declare the regexp inside the loop

	for i := range ruleArr {

		re2 = regexp.MustCompile(ruleArr[i].From)
		output = string(re2.ReplaceAll([]byte(output), []byte(ruleArr[i].To)))

	}
	return output
}
