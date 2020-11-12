package monster

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("err, expent=%v, real=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功！")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("err, expent=%v, real=%v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("expent: %v, real: %v", "红孩儿", monster.Name)
	}

	t.Logf("success")
}
