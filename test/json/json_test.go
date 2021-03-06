package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_Json(t *testing.T) {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			t.Error(err)
			return
		}
		for k := range v {
			if k != "title" {
				v[k] = nil
			}
		}
		if err := enc.Encode(&v); err != nil {
			t.Error(err)
		}
	}
}

func TestJson1(t *testing.T) {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`
	type Mesasge struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Mesasge
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			t.Log(err)
		}
		t.Logf("%s: %s\n", m.Name, m.Text)
	}
	fmt.Println(uicodeToString("\\u7384\\u5e7b"))
}
func uicodeToString(text string) string {
	unicodeArr := strings.Split(text, "\\u")
	var context string
	for _, v := range unicodeArr {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err == nil {
			context += fmt.Sprintf("%c", temp)
		}
	}
	return context
}

type ICoupon interface {
	parse(val []interface{})
	dealDetail()
	Id() int
}

type Coupon1 struct {
}

func (this *Coupon1) Id() int {
	return 1
}
func (this *Coupon1) parse(val []interface{}) {
	for idx, item := range val {
		fmt.Println(idx, item.([]interface{})[0], item.([]interface{})[1], item.([]interface{})[2], item.([]interface{})[3])
	}
}
func (this *Coupon1) dealDetail() {

}

type Coupon2 struct {
}

func (this *Coupon2) Id() int {
	return 8
}
func (this *Coupon2) parse(val []interface{}) {
	for idx, item := range val {
		fmt.Println(idx, item.([]interface{})[0], item.([]interface{})[1], item.([]interface{})[2], item.
		([]interface{})[3])
	}
}
func (this *Coupon2) dealDetail() {

}

func getCouponHandler(couponId int) ICoupon {
	var handlerMap = make(map[int]ICoupon)
	handlerMap[1] = &Coupon1{}
	handlerMap[8] = &Coupon2{}
	return handlerMap[couponId]

}
func TestBean(t *testing.T) {
	var detail string = "{\"8\":[[0,1533683940,1533687540,37]],\"1\":[[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"]]}"
	var detailJson map[string]interface{}
	json.Unmarshal([]byte(detail), &detailJson)
	t.Log(detailJson)
	for key, val := range detailJson {
		t.Log(key, val, reflect.TypeOf(val))
		couponId, _ := strconv.Atoi(key)
		handler := getCouponHandler(couponId)
		t.Log(handler)
		handler.parse(val.([]interface{}))
	}
}

func TestSome(t *testing.T) {

	type people struct {
		Age    int      `json:"age"`
		Names  []string `json:"names"`
		IsMale bool     `json:"is_male"`
	}

	m := map[string]interface{}{
	// "age":     0,
	// "names":   []string{"bill", "zhang"},
	// "is_male": true,
	}
	by, err := json.Marshal(m)
	var pe people
	err = json.Unmarshal(by, &pe)
	fmt.Println(pe, err)

	var arr []string
	fmt.Println(len(arr))
	var a uint = 2
	var b uint = 3
	fmt.Println(reflect.TypeOf(a*b), reflect.TypeOf(b-1))
	t.Log("***",int(math.Ceil(8 / float64(10))))

}
