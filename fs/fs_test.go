package fs

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	t.Log("Exists Test Start")
	// 目录测试
	testDirName := "./utils-fs-exists-test"
	if err := os.Mkdir(testDirName, 0777); err != nil {
		t.Error(err.Error())
	}

	if !Exists(testDirName) {
		t.Error("The test directory does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testDirName); err != nil {
		t.Error(err.Error())
	}
	if Exists(testDirName) {
		t.Error("The test directory still exists after being deleted.")
	}

	// 文件测试
	testFileName := "./utils-fs-exists-test.txt"
	if file, err := os.Create(testFileName); err != nil {
		t.Log(err.Error())
	} else {
		file.Close()
	}
	if !Exists(testFileName) {
		t.Error("The test file does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testFileName); err != nil {
		t.Error(err.Error())
	}
	if Exists(testFileName) {
		t.Error("The test file still exists after being deleted.")
	}
	t.Log("Exists Test Done")
}

func TestIsDir(t *testing.T) {
	t.Log("IsDir Test Start")
	// 目录测试
	testDirName := "./utils-fs-exists-test"
	if err := os.Mkdir(testDirName, 0777); err != nil {
		t.Error(err.Error())
	}
	if !IsDir(testDirName) {
		t.Error("The test directory does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testDirName); err != nil {
		t.Error(err.Error())
	}
	if IsDir(testDirName) {
		t.Error("The test directory still exists after being deleted.")
	}

	// 文件测试
	testFileName := "./utils-fs-exists-test.txt"
	if file, err := os.Create(testFileName); err != nil {
		t.Log(err.Error())
	} else {
		file.Close()
	}
	if IsDir(testFileName) {
		t.Error("The test file does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testFileName); err != nil {
		t.Error(err.Error())
	}
	if IsDir(testFileName) {
		t.Error("The test file still exists after being deleted.")
	}
	t.Log("IsDir Test Done")
}

func TestIsFile(t *testing.T) {
	t.Log("IsFile Test Start")
	// 目录测试
	testDirName := "./utils-fs-exists-test"
	if err := os.Mkdir(testDirName, 0777); err != nil {
		t.Error(err.Error())
	}
	if IsFile(testDirName) {
		t.Error("The test directory does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testDirName); err != nil {
		t.Error(err.Error())
	}
	if IsFile(testDirName) {
		t.Error("The test directory still exists after being deleted.")
	}

	// 文件测试
	testFileName := "./utils-fs-exists-test.txt"
	if file, err := os.Create(testFileName); err != nil {
		t.Log(err.Error())
	} else {
		file.Close()
	}
	if !IsFile(testFileName) {
		t.Error("The test file does not exist after being created.")
	}
	//删除目录
	if err := os.Remove(testFileName); err != nil {
		t.Error(err.Error())
	}
	if IsFile(testFileName) {
		t.Error("The test file still exists after being deleted.")
	}
	t.Log("IsFile Test Done")
}
