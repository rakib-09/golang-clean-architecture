package methodutil

import (
	"encoding/json"
	"fmt"
	"github.com/rakib-09/golang-clean-architecture/config"
	"reflect"
	"strconv"
	"time"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		PrettyPrint("Recovered from panic", r)
	}
}

func GeneratePartnerTokenKey() string {
	conf := config.Redis()

	//ex: `github.com/rakib-09/golang-clean-architecture_partner_`
	return conf.MandatoryPrefix + "token_prefix"
}

func DollarToCents(dollar float64) int {
	t := fmt.Sprintf("%.0f", dollar*100.00)
	v, _ := strconv.Atoi(t)
	return v
}

func CentsToDollar(cents int) float64 {
	return float64(cents) / 100.00
}

func DoAsync(task func()) {
	go func() {
		defer RecoverPanic()
		task()
	}()
}

func PrettyPrint(msg string, data interface{}) {
	if r, err := json.MarshalIndent(&data, "", "  "); err == nil {
		fmt.Printf("[INFO] %v %v: \n %v\n", time.Now(), msg, string(r))
	}
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func IsEmpty(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
