/*
 * This is description.
 * 需求:
	1. 创建一个User(Name, Age, Gender).
	2. 对User进行序列化Marshal, 写入保存到user.txt文件中, 并进行测试.
	3. 对user.txt文件进行读取, 进行反序列化UnMarshal, 并进行测试.
 * @author Chris0710
 * @date 2021/2/18 11:27
*/
package user_test

import (
	"encoding/json"
	user "go_code/chapter_15/user_test"
	"io/ioutil"
	"log"
	"testing"
)

func MarshalUser(user user.User) bool {
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("user marshal err: %v", err)
		return false
	}
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_15\\user_test\\user.txt"
	// 1. 有缓冲: 通过缓冲io的方式写入.
	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// writer := bufio.NewWriter(file)
	// _, err = writer.Write(data)
	// if err != nil {
	// 	log.Fatalf("user write data err: %v", err)
	// 	return false
	// }
	// writer.Flush()

	// 2. 无缓冲: 通过ioutil的方式写入.
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		log.Fatalf("user write data err: %v", err)
		return false
	}
	return true
}

func UnMarshalUser(user *user.User) *user.User {
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_15\\user_test\\user.txt"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("open file err: %v", err)
	}

	err = json.Unmarshal(fileData, user)
	if err != nil {
		log.Fatalf("Unmarshal file err: %v", err)
	}
	return user

}

func TestMarshalUser(t *testing.T) {
	user := user.User{Name: "chris", Age: 30, Gender: true}
	flag := MarshalUser(user)
	if flag {
		t.Logf("TestMarshalUser() 成功...")
	}
}

func TestUnMarshalUser(t *testing.T) {
	u := user.User{}
	user := UnMarshalUser(&u)
	t.Logf("user.Name= %v, user.Age= %v, user.Gender= %v\n", user.Name, user.Age, user.Gender)
	if user.Name == "chris" && user.Age == 30 && user.Gender == true {
		t.Log("UnMarshalUser() 测试成功...")
	}
}

/*
E:\Code\Go\go_learn\src\go_code\chapter_15\user_test>go test
PASS
ok      go_code/chapter_15/user_test    0.043s

E:\Code\Go\go_learn\src\go_code\chapter_15\user_test>go test -v
=== RUN   TestMarshalUser
    TestMarshalUser: user_test.go:65: TestMarshalUser() 成功...
--- PASS: TestMarshalUser (0.00s)
=== RUN   TestUnMarshalUser
    TestUnMarshalUser: user_test.go:72: user.Name= chris, user.Age= 30, user.Gender= true
    TestUnMarshalUser: user_test.go:74: UnMarshalUser() 测试成功...
--- PASS: TestUnMarshalUser (0.00s)
PASS
ok      go_code/chapter_15/user_test    0.044s
*/
