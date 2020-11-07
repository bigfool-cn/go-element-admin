package utils

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "strconv"
  "time"
)

// 去重
func RemoveDuplicateElement(originals interface{}) (interface{}, error) {
	temp := map[string]struct{}{}
	switch slice := originals.(type) {
	case []string:
		result := make([]string, 0, len(originals.([]string)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	case []int64:
		result := make([]int64, 0, len(originals.([]int64)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	default:
		err := fmt.Errorf("Unknown type: %T", slice)
		return nil, err
	}
}

// md5
func MD5(str string) string  {
  h := md5.New()
  h.Write([]byte(str))
  return hex.EncodeToString(h.Sum(nil))
}

// 获取当前时间 yyy-mm-dd HH:ii:ss
func GetCurrntTime() string {
  return time.Now().Format("2006-01-02 15:04:05")
}

// 求交集
func Intersect(slice1 interface{}, slice2 interface{}) (interface{}, error){
  switch slice := slice1.(type) {
  case []string:
    base := make(map[string]int)
    result := make([]string, 0)
    for _, v := range slice1.([]string) {
      base[v]++
    }

    for _, v := range slice2.([]string) {
      times, _ := base[v]
      if times == 1 {
        result = append(result, v)
      }
    }
    return result, nil
  case []int64:
    base := make(map[int64]int)
    result := make([]int64, 0)
    for _, v := range slice1.([]int64) {
      base[v]++
    }
    for _, v := range slice2.([]int64) {
      times, _ := base[v]
      if times == 1 {
        result = append(result, v)
      }
    }
    return result, nil
  default:
    err := fmt.Errorf("Unknown type: %T", slice)
    return nil, err
  }
}

// 求差集 slice1-并集
func Difference(slice1 interface{}, slice2 interface{}) (interface{}, error) {
  switch slice := slice1.(type) {
  case []string:
    base := make(map[string]int)
    result := make([]string, 0)
    inter, err := Intersect(slice1, slice2)
    if err != nil {
      return nil, err
    }
    for _, v := range inter.([]string) {
      base[v]++
    }

    for _, value := range slice1.([]string) {
      times, _ := base[value]
      if times == 0 {
        result = append(result, value)
      }
    }
    return result, nil
  case []int64:
    base := make(map[int64]int)
    result := make([]int64, 0)
    inter, err := Intersect(slice1, slice2)
    if err != nil {
      return nil, err
    }
    for _, v := range inter.([]int64) {
      base[v]++
    }

    for _, value := range slice1.([]int64) {
      times, _ := base[value]
      if times == 0 {
        result = append(result, value)
      }
    }
    return result, nil
  default:
    err := fmt.Errorf("Unknown type: %T", slice)
    return nil, err
  }
}

func SliceStringToInt64(strs []string) []int64  {
  var result []int64
  for _, str:= range strs  {
    it,_ := strconv.ParseInt(str,10,64)
    result = append(result,it)
  }
  return result
}

func SliceInt64ToString(ints []int64) []string {
  var result []string
  for _, it:= range ints  {
    str := strconv.FormatInt(it,10)
    result = append(result,str)
  }
  return result
}

