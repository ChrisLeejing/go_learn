/*
 * This is description.
 * 测试注意事项:
	1. 测试文件名必须以_test.go结尾, 如: addUpper_test.go.
	2. 测试用例函数必须以Test开头, 如: TestAddUpper.
	3. 测试函数的形参必须以(t *testing.T), 如: func TestAddUpper(t *testing.T) {}.
	4. go test: 运行正确不输出日志, 运行错误则输出日志; go test -v: 不管运行正确还是错误, 都输出日志.
		E:\Code\Go\go_learn\src\go_code\chapter_15\unit_test>go test
		--- FAIL: TestAddUpper (0.00s)
			addUpper_test.go:20: AddUpper()测试失败, 期望值: 54, 实际值: 55.
		FAIL
		exit status 1
		FAIL    go_code/chapter_15/unit_test    0.042s

		E:\Code\Go\go_learn\src\go_code\chapter_15\unit_test>go test -v
		=== RUN   TestAddUpper
			TestAddUpper: addUpper_test.go:20: AddUpper()测试失败, 期望值: 54, 实际值: 55.
		--- FAIL: TestAddUpper (0.00s)
		FAIL
		exit status 1
		FAIL    go_code/chapter_15/unit_test    0.045s
	5. t.Fatalf(): 格式化输出错误日志, 并结束程序运行; t.Logf(): 格式化输出日志.
	6. 测试单个文件, 一定要带上被测试的源文件; 测试单个方法格式: go test -v test.run <function_name> 如下:
		E:\Code\Go\go_learn\src\go_code\chapter_15\unit_test>go test addUpper_test.go addUpper.go
		--- FAIL: TestAddUpper (0.00s)
			addUpper_test.go:14: AddUpper()测试失败, 期望值: 54, 实际值: 55.
		FAIL
		FAIL    command-line-arguments  0.040s
		FAIL

		E:\Code\Go\go_learn\src\go_code\chapter_15\unit_test>go test -v -test.run TestAddUpper
		=== RUN   TestAddUpper
			TestAddUpper: addUpper_test.go:14: AddUpper()测试失败, 期望值: 54, 实际值: 55.
		--- FAIL: TestAddUpper (0.00s)
		FAIL
		exit status 1
		FAIL    go_code/chapter_15/unit_test    0.040s

 * @author Chris0710
 * @date 2021/2/18 10:38
*/
package main

import "testing"

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 54 {
		t.Fatalf("AddUpper()测试失败, 期望值: %v, 实际值: %v.", 54, res)
	}
	// if res != 55 {
	// 	t.Fatalf("AddUpper()测试失败, 期望值: %v, 实际值: %v.", 55, res)
	// }

	t.Logf("AddUpper()测试成功, 期望值: %v, 实际值: %v.", 55, res)
}
