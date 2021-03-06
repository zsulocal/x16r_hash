package x16r_hash

import (
	"encoding/hex"
	"testing"
)

func TestX16r(t *testing.T) {
	hash := "700000005d385ba114d079970b29a9418fd0549e7d68a95c7f168621a314201000000000578586d149fd07b22f3a8a347c516de7052f034d2b76ff68e0d6ecff9b77a45489e3fd511732011df0731000"
	hash_bin, _ := hex.DecodeString(hash)
	powhash := GetPowHash(hash_bin)
	powhash_hex := hex.EncodeToString(powhash)
	if powhash_hex != "77a19463753c27887c5697b47118719f4af6fba0647eddde71a938e7b3dd0d48" {
		t.Error("Test x16r powhash failed")
		return
	}
}
