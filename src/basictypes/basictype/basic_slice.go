package basictype

import (
	"database/sql/driver"
	"encoding/json"
	"sort"
)

//
// 这个文件定义在 MySQL 的一个 Field 存储 JSON 文本的 array 的场景。with gorm
// 比如：一个org有多张图片 这里存储每一张图片的URL
//

// ImageSlice 定义一个 string 数组，用于存储多张图片的 URL
type ImageSlice []string

func (is *ImageSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(is)
	return string(b), err
}

func (is *ImageSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), is)
}

// 遍历所有string，只要有一个string的长度大于0，就是 valid
func (is ImageSlice) IsValid() bool {
	if len(is) == 0 {
		return false
	}
	for _, value := range is {
		if len(value) > 0 {
			return true
		}
	}

	return false
}

func (is ImageSlice) Sort() {
	sort.Slice(is, func(i, j int) bool {
		return is[i] < is[j]
	})
}

// UintSlice 定义一个 uint 数组，用于存储多个 uint
// type UintSlice []uint

// func (us *UintSlice) Value() (driver.Value, error) {
// 	b, err := json.Marshal(*us)
// 	return string(b), err
// }

// func (us *UintSlice) Scan(input interface{}) error {
// 	return json.Unmarshal(input.([]byte), &us)
// }

// func (us UintSlice) Sort() {
// 	sort.Slice(us, func(i, j int) bool {
// 		return us[i] < us[j]
// 	})
// }

// func (us UintSlice) Len() int {
// 	return len(us)
// }

// func (us UintSlice) Less(i, j int) bool {
// 	return us[i] < us[j]
// }

// func (us UintSlice) Swap(i, j int) {
// 	us[i], us[j] = us[j], us[i]
// }

// func (us UintSlice) MarshalJSON() (output []byte, err error) {
// 	// 将 Value 编码为 JSON 字符串
// 	var buf []byte
// 	value, _ := us.Value()

// 	// fmt.Println(string(string(value)))

// 	if buf, err = json.Marshal(value); err != nil {
// 		return buf, nil
// 	}

// 	// // 将 JSON 字符串二次编码并返回
// 	// if output, err = json.Marshal(string(buf)); err != nil {
// 	// 	return
// 	// }
// 	return buf, nil
// }

// func (s *Stringify[T]) UnmarshalJSON(input []byte) (err error) {
// 	// 解码出原始 JSON 字符串
// 	var buf string
// 	if err = json.Unmarshal(input, &buf); err != nil {
// 		return
// 	}
// 	// 将原始字符串解码到 Value 上
// 	if err = json.Unmarshal([]byte(buf), &s.Value); err != nil {
// 		return
// 	}
// 	return
// }

//	type StringArray struct {
//		Values []string `json:"values"`
//	}

// type StringArray []string

// func (sary *StringArray) Value() (driver.Value, error) {
// 	b, err := json.Marshal(sary)
// 	return string(b), err
// }

// func (sary *StringArray) Scan(input interface{}) error {
// 	return json.Unmarshal(input.([]byte), sary)
// }

// type IntArray []int

// func (sary *IntArray) Value() (driver.Value, error) {
// 	b, err := json.Marshal(sary)
// 	return string(b), err
// }

// func (sary *IntArray) Scan(input interface{}) error {
// 	return json.Unmarshal(input.([]byte), sary)
// }

type BasicSlice[TYPE uint | string] []TYPE

func (bs *BasicSlice[TYPE]) Value() (driver.Value, error) {
	b, err := json.Marshal(bs)
	return string(b), err
}

func (bs *BasicSlice[TYPE]) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), bs)
}

func (bs BasicSlice[TYPE]) Sort() {
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
}

func (bs BasicSlice[TYPE]) Len() int {
	return len(bs)
}

// func (bs *BasicSlice[TYPE]) HadDuplicate() bool {
// 	if len(*bs) < 2 {
// 		return false
// 	}

// 	existing := map[TYPE]struct{}{}
// 	for _, value := range *bs {
// 		_, OK := existing[value]
// 		if OK {
// 			return true
// 		}
// 		existing[value] = struct{}{}
// 	}

// 	return false
// }

// func (s BasicSlice[TYPE]) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s BasicSlice[TYPE]) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func (bs BasicSlice[TYPE]) MarshalJSON() (output []byte, err error) {
// 	// 将 Value 编码为 JSON 字符串
// 	var buf []byte
// 	value, _ := bs.Value()

// 	// fmt.Println(string(string(value)))

// 	if buf, err = json.Marshal(value); err != nil {
// 		return buf, nil
// 	}

// 	// // 将 JSON 字符串二次编码并返回
// 	// if output, err = json.Marshal(string(buf)); err != nil {
// 	// 	return
// 	// }
// 	return buf, nil
// }

// func (s *Stringify[T]) UnmarshalJSON(input []byte) (err error) {
// 	// 解码出原始 JSON 字符串
// 	var buf string
// 	if err = json.Unmarshal(input, &buf); err != nil {
// 		return
// 	}
// 	// 将原始字符串解码到 Value 上
// 	if err = json.Unmarshal([]byte(buf), &s.Value); err != nil {
// 		return
// 	}
// 	return
// }
