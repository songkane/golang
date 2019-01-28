// Package base62 单元测试
// Created by chenguolin 2018-08-22
package base62

import (
	"fmt"
	"testing"
)

func TestUid2Base62(t *testing.T) {
	// case 1
	uid := int64(0)
	base62 := UID2Base62(uid)
	if base62 != "" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 2
	uid = int64(37)
	base62 = UID2Base62(uid)
	if base62 != "65vZcb" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 3
	uid = int64(115)
	base62 = UID2Base62(uid)
	if base62 != "65vZdL" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 4
	uid = int64(1537618052)
	base62 = UID2Base62(uid)
	if base62 != "4H3jhs" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 5
	uid = int64(38068692544)
	base62 = UID2Base62(uid)
	if base62 != "dTrvku" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 6
	uid = int64(38068692543)
	base62 = UID2Base62(uid)
	if base62 != "dTrvl9" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d", uid))
	}

	// case 7
	uid = int64(42276989)
	base62 = UID2Base62(uid)
	if base62 != "67dRzh" {
		t.Fatal(fmt.Sprintf("uid 2 base62 error, uid: %d, memo: %s", uid, base62))
	}
}

func TestBase622Uid(t *testing.T) {
	// case 1
	base62 := ""
	uid := Base622Uid(base62)
	if uid != -1 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}

	// case 2
	base62 = "65vZcb"
	uid = Base622Uid(base62)
	if uid != 37 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}

	// case 3
	base62 = "65vZdL"
	uid = Base622Uid(base62)
	if uid != 115 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}

	// case 4
	base62 = "4H3jhs"
	uid = Base622Uid(base62)
	if uid != 1537618052 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}

	// case 5
	base62 = "dTrvku"
	uid = Base622Uid(base62)
	if uid != 38068692544 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}

	// case 6
	base62 = "dTrvl9"
	uid = Base622Uid(base62)
	if uid != 38068692543 {
		t.Fatal(fmt.Sprintf("base62 2 uid error, base62: %s", base62))
	}
}
