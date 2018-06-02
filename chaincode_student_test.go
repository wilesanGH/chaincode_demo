package main

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func initSchool(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("initSchool"), []byte(args[0]), []byte(args[1])})
	
	if res.Status != shim.OK {
		fmt.Println("InitSchool failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func deleteSchool(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("deleteSchool"), []byte(args[0])})
	
	if res.Status != shim.OK {
		fmt.Println("DeleteSchool failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func addStudent(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addStudent"), []byte(args[0]), []byte(args[1]), []byte(args[2]), []byte(args[3])})
	
	if res.Status != shim.OK {
		fmt.Println("AddStudent failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func queryStudentByID(t *testing.T, stub *shim.MockStub, userid string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryStudentByID"), []byte(userid)})
	if res.Status != shim.OK {
		fmt.Println("queryStudentByID", userid, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("queryStudentByID", userid, "failed to get value")
		t.FailNow()
	}
}

//初始化学校
func TestInitSchool(t *testing.T) {
	scc := new(StudentChaincode)
	stub := shim.NewMockStub("StudentChaincode", scc)
	mockInit(t, stub, nil)
	initSchool(t, stub, []string{"schoolId_A", "学校1"})
	initSchool(t, stub, []string{"schoolId_B", "学校2"})
}

//添加学生，学校没有创建时失败
func TestAddStudent(t *testing.T) {
	scc := new(StudentChaincode)
	stub := shim.NewMockStub("StudentChaincode", scc)
	mockInit(t, stub, nil)
	initSchool(t, stub, []string{"schoolId_A", "学校1"})
	addStudent(t, stub, []string{"张三", "1", "schoolId_A", "classA"})
	//学校schoolB没有创建，返回错误
	addStudent(t, stub, []string{"李四", "2", "schoolB", "classB"})
}

//根据ID查询学生信息
func TestQueryStudentByID(t *testing.T) {
	scc := new(StudentChaincode)
	stub := shim.NewMockStub("StudentChaincode", scc)
	mockInit(t, stub, nil)
	initSchool(t, stub, []string{"schoolId_A", "学校1"})
	addStudent(t, stub, []string{"张三", "1", "schoolId_A", "classA"})
	queryStudentByID(t, stub, "1")
}

//删除学校，一并删除学校下的所有学生
func TestDeleteSchool(t *testing.T) {
	scc := new(StudentChaincode)
	stub := shim.NewMockStub("StudentChaincode", scc)
	initSchool(t, stub, []string{"schoolId_A", "学校1"})
	addStudent(t, stub, []string{"张三", "1", "schoolId_A", "classA"})
	deleteSchool(t, stub, []string{"schoolId_A"})
	queryStudentByID(t, stub, "1")
}