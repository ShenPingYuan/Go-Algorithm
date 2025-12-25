package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

type BinarySearchTree struct {
	Root *TreeNode
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func updateHeight(node *TreeNode) {
	if node == nil {
		return
	}
	var leftHeight int = getHeight(node.Left)
	var rightHeight int = getHeight(node.Right)

	node.Height = 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

func (bst *BinarySearchTree) Search(val int) *TreeNode {
	return searchNode(bst.Root, val)
}

func searchNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchNode(root.Left, val)
	}

	if root.Val < val {
		return searchNode(root.Right, val)
	}

	return nil
}

func (bst *BinarySearchTree) InOrderTraverse() {
	inOrder(bst.Root)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}

func (bt *BinarySearchTree) Xml4Tree() string {
	var Xml, Node string
	Head := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/" +
		"1999/xlink\" version=\"1.1\" width=\"Width\" height=\"Height\">\nCONTENT</svg>"
	Line := `<line x1="X1" y1="Y1" x2="X2" y2="Y2" style="stroke:black;stroke-width:2"/>`
	List := bt.LevelNullwith()
	Levels := len(List)
	for i := Levels - 1; i >= 0; i-- {
		negative := -1
		TmpXml := ""
		for j := 0; j < Pow2(i); j++ {
			t := Pow2(Levels - i - 1)
			x, y := 50*(2*t*j+t), 120*i+50
			if List[i][j] != math.MaxInt {
				fillColor := "orange"
				if i == Levels-1 || i > 0 && i < Levels-1 &&
					List[i+1][j*2] == nil && List[i+1][j*2+1] == nil {
					fillColor = "lightgreen"
				}
				TmpStr := ""
				switch List[i][j].(type) {
				case int:
					TmpStr = strconv.Itoa(List[i][j].(int))
				case float64:
					TmpStr = strconv.FormatFloat(List[i][j].(float64), 'g', -1, 64)
				case string:
					TmpStr = List[i][j].(string)
				default:
					TmpStr = "Error Type"
				}
				val, _ := strconv.Atoi(TmpStr)
				node := bt.Search(val)
				Xml = XmlNode(i, j, x, y, node.Height, TmpStr, fillColor)
			}
			if i > 0 {
				line1 := strings.Replace(Line, "X1", strconv.Itoa(x), 1)
				line1 = strings.Replace(line1, "Y1", strconv.Itoa(y-30), 1)
				negative *= -1
				x0, y0 := 21, 21
				x += 50*negative*(2*t*j%2+t) - negative*x0
				line1 = strings.Replace(line1, "X2", strconv.Itoa(x), 1)
				line1 = strings.Replace(line1, "Y2", strconv.Itoa(y-120+y0), 1)
				Xml = strings.Replace(Xml, "<ROOT/>", line1, 1)
			}
			if List[i][j] != nil {
				TmpXml += Xml
			}
		}
		Node = TmpXml + Node
	}
	Xml = strings.Replace(Head, "CONTENT", Node, 1)
	Xml = strings.Replace(Xml, "Width", strconv.Itoa(Pow2(Levels)*50), 1)
	Xml = strings.Replace(Xml, "Height", strconv.Itoa(Levels*120), 1)
	return Xml
}

func (bt *BinarySearchTree) LevelNullwith(Fills ...interface{}) [][]interface{} {
	var Nodes [][]interface{}
	var Fill0 interface{}
	if len(Fills) == 0 {
		Fill0 = math.MaxInt
	} else if len(Fills) == 1 {
		Fill0 = Fills[0]
	} else {
		panic("Error: number of parameters is greater than 1")
	}
	root := bt.Root
	if root == nil {
		return Nodes
	}
	Count := 0
	Queue := []*TreeNode{root}
	for len(Queue) > 0 {
		nodes := []interface{}{}
		Level := len(Queue)
		for Level > 0 {
			cur := Queue[0]
			Queue = Queue[1:]
			nodes = append(nodes, cur.Val)
			Count++
			Level--
			if cur.Left != nil {
				Queue = append(Queue, cur.Left)
			}
			if cur.Right != nil {
				Queue = append(Queue, cur.Right)
			}
		}
		Nodes = append(Nodes, nodes)
	}
	newbiTree := Copy(bt)
	for i := 1; i < Pow2(len(Nodes))-Count; i++ {
		newbiTree.AppendNode(Fill0)
	}
	return newbiTree.BForder2D()
}

func (bt *BinarySearchTree) BForder2D() [][]interface{} {
	var res [][]interface{}
	root := bt.Root
	if root == nil {
		return res
	}
	Queue := []*TreeNode{root}
	for len(Queue) > 0 {
		Nodes := []interface{}{}
		Levels := len(Queue)
		for Levels > 0 {
			cur := Queue[0]
			Queue = Queue[1:]
			Nodes = append(Nodes, cur.Val)
			Levels--
			if cur.Left != nil {
				Queue = append(Queue, cur.Left)
			}
			if cur.Right != nil {
				Queue = append(Queue, cur.Right)
			}
		}
		res = append(res, Nodes)
	}
	return res
}

func (bt *BinarySearchTree) AppendNode(data interface{}) {
	root := bt.Root
	if root == nil {
		bt.Root = &TreeNode{Val: data.(int)}
		return
	}
	Queue := []*TreeNode{root}
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		if cur.Left != nil {
			Queue = append(Queue, cur.Left)
		} else {
			cur.Left = &TreeNode{Val: data.(int)}
			return
		}
		if cur.Right != nil {
			Queue = append(Queue, cur.Right)
		} else {
			cur.Right = &TreeNode{Val: data.(int)}
			break
		}
	}
}

func Copy(bt *BinarySearchTree) *BinarySearchTree {
	root := bt.Root
	if root == nil {
		return &BinarySearchTree{}
	}
	node := &TreeNode{Val: root.Val}
	Queue1, Queue2 := []*TreeNode{root}, []*TreeNode{node}
	for len(Queue1) > 0 {
		p1, p2 := Queue1[0], Queue2[0]
		Queue1, Queue2 = Queue1[1:], Queue2[1:]
		if p1.Left != nil {
			Node := &TreeNode{Val: p1.Left.Val}
			p2.Left = Node
			Queue1 = append(Queue1, p1.Left)
			Queue2 = append(Queue2, Node)
		}
		if p1.Right != nil {
			Node := &TreeNode{Val: p1.Right.Val}
			p2.Right = Node
			Queue1 = append(Queue1, p1.Right)
			Queue2 = append(Queue2, Node)
		}
	}
	return &BinarySearchTree{Root: node}
}

func XmlNode(M, N, X, Y, H int, Data string, Color ...string) string {
	var cColor, tColor string
	R := 30
	Node := `<Tab><g id="M,N">
	<circle cx="X" cy="Y" r="RC" stroke="black" stroke-width="2" fill="COLOR" />
	<text x="X" y="Y" fill="TextColor" font-size="20" text-anchor="middle"
		dominant-baseline="middle">DATA,HEIGHT</text>
	<ROOT/>
	</g><CrLf>`
	if len(Color) == 0 {
		cColor, tColor = "orange", "red"
	} else if len(Color) == 1 {
		cColor, tColor = Color[0], "red"
	} else {
		cColor, tColor = Color[0], Color[1]
	}
	Node = strings.Replace(Node, "M", strconv.Itoa(M), 1)
	Node = strings.Replace(Node, "N", strconv.Itoa(N), 1)
	Node = strings.Replace(Node, "X", strconv.Itoa(X), 2)
	Node = strings.Replace(Node, "Y", strconv.Itoa(Y), 2)
	Node = strings.Replace(Node, "RC", strconv.Itoa(R), 1)
	Node = strings.Replace(Node, "DATA", Data, 1)
	Node = strings.Replace(Node, "HEIGHT", strconv.Itoa(H), 1)
	Node = strings.Replace(Node, "COLOR", cColor, 1)
	Node = strings.Replace(Node, "TextColor", tColor, 1)
	Node = strings.Replace(Node, "<CrLf>", "\n", -1)
	Node = strings.Replace(Node, "<Tab>", "\t", -1)
	Node = strings.Replace(Node, "\n\t\t", " ", -1)
	return Node
}

func Pow2(x int) int { //x>=0
	res := 1
	for i := 0; i < x; i++ {
		res *= 2
	}
	return res
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.Root = insertNode(bst.Root, val)
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val, Height: 1}
	}

	if val == node.Val {
		return node
	}

	if val < node.Val {
		node.Left = insertNode(node.Left, val)
	}

	if val > node.Val {
		node.Right = insertNode(node.Right, val)
	}

	updateHeight(node) // 插入新节点后，对应的搜索路径上的节点高度都需要更新

	return rotateIfNeeded(node)
	// return node
}

func rotateIfNeeded(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	balance := getBalanceFactor(node)

	if balance > 1 {
		if getBalanceFactor(node.Left) < 0 {
			node.Left = leftRotate(node.Left) // LR
		}
		return rightRotate(node) // RR
	}

	if balance < -1 {
		if getBalanceFactor(node.Right) > 0 {
			node.Right = rightRotate(node.Right) // RL
		}
		return leftRotate(node) // LL
	}

	return node
}

func leftRotate(node *TreeNode) *TreeNode {
	if node == nil || node.Right == nil {
		return nil
	}
	var rightNode = node.Right         // node的右子节点
	var rightLeftNode = rightNode.Left // node的右子节点的左子节点

	// 旋转
	node.Right = rightLeftNode
	rightNode.Left = node

	// 旋转后node和rightNode的高度发生变化，需要更新
	updateHeight(node)
	updateHeight(rightNode)

	// 新的rightNode成为新的node
	return rightNode
}

func rightRotate(node *TreeNode) *TreeNode {
	if node == nil || node.Left == nil {
		return nil
	}
	var leftNode = node.Left           // node的左子节点
	var leftRightNode = leftNode.Right // node的左子节点的右子节点

	// 旋转
	node.Left = leftRightNode
	leftNode.Right = node

	// 旋转后node和leftNode的高度发生变化，需要更新
	updateHeight(node)
	updateHeight(leftNode)

	// 新的leftNode成为新的node
	return leftNode
}

func (bst *BinarySearchTree) Delete(val int) {
	bst.Root = deleteNode(bst.Root, val)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = deleteNode(node.Left, val)
	} else if val > node.Val {
		node.Right = deleteNode(node.Right, val)
	} else {
		// 找到要删除的节点
		if isLeaf(node) {
			// 情况 1：叶子节点
			return nil
		} else if node.Left == nil {
			// 情况 2：只有右子节点
			return node.Right
		} else if node.Right == nil {
			// 情况 2：只有左子节点
			return node.Left
		} else {
			// 情况 3：有两个子节点
			// 找到右子树的最小值节点（后继节点）
			minNode := findMin(node.Right)
			node.Val = minNode.Val // 用后继节点的值替换当前节点
			// 删除后继节点
			node.Right = deleteNode(node.Right, minNode.Val)
		}
	}

	updateHeight(node)

	return rotateIfNeeded(node)
}

// 找到以当前节点为根的子树中的最小值节点
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	var file *os.File
	var err1 error
	var bst *BinarySearchTree = &BinarySearchTree{}

	for i := 1; i < 10; i++ {
		bst.Insert(rand.Intn(100))
	}

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(2)
	// bst.Insert(11)
	// bst.InOrderTraverse()
	bst.Delete(10)
	bst.Delete(2)
	// bst.Delete(6)
	bst.Delete(3)
	// bst.InOrderTraverse()

	texts := []string{bst.Xml4Tree()}

	for index, text := range texts {
		svgFile := "./biTree0" + strconv.Itoa(index+1) + ".svg"
		file, err1 = os.Create(svgFile)
		if err1 != nil {
			panic(err1)
		}
		_, err1 = io.WriteString(file, text)
		if err1 != nil {
			panic(err1)
		}
		file.Close()
		exec.Command("cmd", "/c", "start", svgFile).Start()
		//Linux 代码：
		//exec.Command("xdg-open", svgFile).Start()
		//Mac 代码：
		//exec.Command("open", svgFile).Start()
	}
}
