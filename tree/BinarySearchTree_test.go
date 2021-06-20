package tree

import "testing"

//             15
//       10           20
//    5     13      18   25
//  2   6        16
func generateBST() *BST {
	bst := NewBST(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(5)
	bst.Insert(13)
	bst.Insert(25)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(16)
	return bst
}

func TestBST_Insert(t *testing.T) {
	bst := generateBST()
	InOrderRecursion(bst.root)
}

func TestBST_Find(t *testing.T) {
	bst := generateBST()
	node := bst.Find(5)
	t.Log(node)
}

func TestBST_Minimum(t *testing.T) {
	bst := generateBST()
	bst.Insert(1)
	InOrderRecursion(bst.root)
	t.Log(bst.Minimum())
}

func TestBST_Remove(t *testing.T) {
	bst := generateBST()
	InOrderRecursion(bst.root)
	println()
	bst.Remove(15)
	InOrderRecursion(bst.root)
}

func TestBST_Delete(t *testing.T) {
	bst := generateBST()
	InOrderRecursion(bst.root)
	println()
	bst.Delete(15)
	InOrderRecursion(bst.root)
}
