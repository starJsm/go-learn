1. PASS 表示测试用例运行成功,FAIL 表示测试用例运行失败
2. 测试单个文件,一定要带上被测试的原文件
    go test -v cal_test.go cal.go
3. 测试单个方法
    go test -v -test.run TestAddUpper
