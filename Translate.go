package Translate3

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/wznpp1/Translate2"
	"gopkg.in/yaml.v2"
)

type MapA map[string]string
type StrSliceA []string

var MapA1 = MapA{}
var Gtranslate = Translate2.New()

var NotTranslated = make([]string, 0)
var Regexp5 = regexp.MustCompile(`(\[.*?\])+|(\{.*?\})+|\\[\\A-Za-z]+|%[%A-Za-z]+|[\p{P}]{2,}`)

func AddOrGetStrings(Strings StrSliceA) (map[string]string, error) {
	Map1 := map[string]string{}

	for _, k := range Strings {
		if v, ok := MapA1[k]; !ok || v == "" {
			NotTranslated = append(NotTranslated, Regexp5.ReplaceAllString(k, "<a>$0</a>"))
		} else {
			Map1[k] = v
		}
	}

	return Map1, nil
}

func AddOrGetString(Key string) string {
	if v, ok := MapA1[Key]; ok && v != "" {
		return v
	} else {
		NotTranslated = append(NotTranslated, Regexp5.ReplaceAllString(Key, "<a>$0</a>"))
	}
	return ""
}

func GetMap() *MapA {
	return &MapA1
}

func Gtranslate1() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	RWMutex.Lock()
	defer RWMutex.Unlock()

	if len(NotTranslated) == 0 {
		err1 := fmt.Errorf("len(NotTranslated) == 0")
		fmt.Println(err1)
		return err1
	}

	got, err := Gtranslate.Translate(NotTranslated, "zh-CN")
	if err != nil {
		return err
	}

	if len(got) != len(NotTranslated) {
		err1 := fmt.Errorf("len(got) != len(NotTranslated)")
		fmt.Println(err1)
		return err1
	}

	for i, v := range got {
		if v1, ok := MapA1[NotTranslated[i]]; ok && v1 != "" {
			continue
		} else {
			MapA1[NotTranslated[i]] = Replacer1.Replace(v)
		}
	}
	NotTranslated = make([]string, 0)

	File1, err := os.OpenFile(
		"translate.yaml",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0664)

	if err == nil {
		Encoder := yaml.NewEncoder(File1)
		err = Encoder.Encode(&MapA1)
		if err != nil {
			fmt.Println("yaml Encode error")
		}
		err = File1.Close()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Open translate.yaml error")
	}

	return nil
}

func InitMapA1() {
	File1, err := os.Open("translate.yaml")
	if err == nil {
		Decoder1 := yaml.NewDecoder(File1)
		err = Decoder1.Decode(&MapA1)
		if err != nil {
			fmt.Println("yaml Decode error")
		}
		err = File1.Close()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Open translate.yaml error")
	}
}

var RWMutex = sync.RWMutex{}
var Replacer1 = strings.NewReplacer(
	`<code>`, ``,
	`</code>`, ``,
	`<a>`, ``,
	`</a>`, ``,
	`&#39;`, `英尺`,
	`&quot;`, `英寸`,
	` ，`, `,`,
	`，`, `,`,
	` 。`, `.`,
	`。`, `.`,
	` ！`, `!`,
	`！`, `!`,
	` ？`, `?`,
	`？`, `?`,
	` 【`, `[`,
	`【`, `[`,
	` 】`, `]`,
	`】`, `]`,
	` （`, `(`,
	`（`, `(`,
	` ）`, `)`,
	`）`, `)`,
	` ％`, `%`,
	`％`, `%`,
	` ＃`, `#`,
	`＃`, `#`,
	` ＠`, `@`,
	`＠`, `@`,
	` ＆`, `&`,
	`＆`, `&`,
	` １`, `1`,
	`１`, `1`,
	` ２`, `2`,
	`２`, `2`,
	` ３`, `3`,
	`３`, `3`,
	` ４`, `4`,
	`４`, `4`,
	` ５`, `5`,
	`５`, `5`,
	` ６`, `6`,
	`６`, `6`,
	` ７`, `7`,
	`７`, `7`,
	` ８`, `8`,
	`８`, `8`,
	` ９`, `9`,
	`９`, `9`,
	` ０`, `0`,
	`０`, `0`,
	`“ `, `\"`,
	`“`, `\"`,
	` ”`, `\"`,
	`”`, `\"`,
)
