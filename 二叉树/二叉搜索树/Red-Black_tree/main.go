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

const (
	RED   = true
	BLACK = false
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int
	Color  bool // true for red, false for black
	Parent *TreeNode
}

type Color int

const (
	Red = iota
	Black
)

// **红黑树结构**
type RedBlackTree struct {
	Root *TreeNode
}

// **左旋**
func (rbt *RedBlackTree) leftRotate(node *TreeNode) {
	if node == nil || node.Right == nil {
		return
	}

	rightNode := node.Right
	node.Right = rightNode.Left

	if rightNode.Left != nil {
		rightNode.Left.Parent = node
	}

	rightNode.Parent = node.Parent

	if node.Parent == nil {
		rbt.Root = rightNode
	} else if node == node.Parent.Left {
		node.Parent.Left = rightNode
	} else {
		node.Parent.Right = rightNode
	}

	rightNode.Left = node
	node.Parent = rightNode
}

// **右旋**
func (rbt *RedBlackTree) rightRotate(y *TreeNode) {
	if y == nil || y.Left == nil {
		return
	}

	x := y.Left
	y.Left = x.Right

	if x.Right != nil {
		x.Right.Parent = y
	}

	x.Parent = y.Parent

	if y.Parent == nil {
		rbt.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	x.Right = y
	y.Parent = x
}

// **插入修复**
func (rbt *RedBlackTree) fixInsert(node *TreeNode) {
	for node.Parent != nil && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left { // 父节点是祖父的左子节点
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RED { // 叔叔是红色
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else { // 叔叔是黑色
				if node == node.Parent.Right { // 右孩子情况
					node = node.Parent
					rbt.leftRotate(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbt.rightRotate(node.Parent.Parent)
			}
		} else { // 父节点是祖父的右子节点（对称情况）
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RED { // 叔叔是红色
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else { // 叔叔是黑色
				if node == node.Parent.Left { // 左孩子情况
					node = node.Parent
					rbt.rightRotate(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbt.leftRotate(node.Parent.Parent)
			}
		}
	}
	rbt.Root.Color = BLACK
}

// **插入节点**
func (rbt *RedBlackTree) Insert(val int) {
	newNode := &TreeNode{Val: val, Color: RED} // 新节点默认是红色
	if rbt.Root == nil {
		rbt.Root = newNode
		rbt.Root.Color = BLACK
		return
	}

	var parent *TreeNode
	current := rbt.Root

	for current != nil {
		parent = current
		if val < current.Val {
			current = current.Left
		} else if val > current.Val {
			current = current.Right
		} else {
			return // 节点已存在，直接返回
		}
	}

	newNode.Parent = parent
	if val < parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	rbt.fixInsert(newNode)
}

// **查找节点**
func (rbt *RedBlackTree) Search(val int) *TreeNode {
	current := rbt.Root
	for current != nil {
		if val == current.Val {
			return current
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// **中序遍历**
func (rbt *RedBlackTree) InOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	rbt.InOrderTraverse(node.Left)
	fmt.Printf("%d(%s) ", node.Val, colorToString(node.Color))
	rbt.InOrderTraverse(node.Right)
}

// **颜色转换**
func colorToString(color bool) string {
	if color == RED {
		return "R"
	}
	return "B"
}

// **删除修复**
func (rbt *RedBlackTree) fixDelete(node *TreeNode) {
	for node != rbt.Root && node.Color == BLACK {
		if node == node.Parent.Left {
			sibling := node.Parent.Right
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				rbt.leftRotate(node.Parent)
				sibling = node.Parent.Right
			}
			if sibling.Left.Color == BLACK && sibling.Right.Color == BLACK {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Right.Color == BLACK {
					sibling.Left.Color = BLACK
					sibling.Color = RED
					rbt.rightRotate(sibling)
					sibling = node.Parent.Right
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				sibling.Right.Color = BLACK
				rbt.leftRotate(node.Parent)
				node = rbt.Root
			}
		} else { // 对称情况
			sibling := node.Parent.Left
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				rbt.rightRotate(node.Parent)
				sibling = node.Parent.Left
			}
			if sibling.Right.Color == BLACK && sibling.Left.Color == BLACK {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Left.Color == BLACK {
					sibling.Right.Color = BLACK
					sibling.Color = RED
					rbt.leftRotate(sibling)
					sibling = node.Parent.Left
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				sibling.Left.Color = BLACK
				rbt.rightRotate(node.Parent)
				node = rbt.Root
			}
		}
	}
	node.Color = BLACK
}

func (bt *RedBlackTree) Xml4Tree() string {
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
				fillColor = "black"
				textColor := "white"

				if node.Color == RED {
					fillColor = "red"
					textColor = "black"
				}
				Xml = XmlNode(i, j, x, y, node.Height, TmpStr, fillColor, textColor)
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

func (bt *RedBlackTree) LevelNullwith(Fills ...interface{}) [][]interface{} {
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

func (bt *RedBlackTree) BForder2D() [][]interface{} {
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

func (bt *RedBlackTree) AppendNode(data interface{}) {
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

func Copy(bt *RedBlackTree) *RedBlackTree {
	root := bt.Root
	if root == nil {
		return &RedBlackTree{}
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
	return &RedBlackTree{Root: node}
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

	// 确保传入的颜色正确
	if len(Color) == 0 {
		cColor, tColor = "black", "white" // 默认黑色节点，白色文本
	} else if len(Color) == 1 {
		cColor, tColor = Color[0], "white"
		if cColor == "red" {
			tColor = "black" // 红色节点，黑色文本
		}
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

func (bst *RedBlackTree) Show() {
	var file *os.File
	var err1 error

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

func main() {
	var bst *RedBlackTree = &RedBlackTree{}

	// for i := 1; i < 100; i++ {
	// 	fmt.Print("请输入一个节点：\n")
	// 	var val int = 0
	// 	fmt.Scan(&val)
	// 	bst.Insert(val)
	// 	bst.Show()
	// }

	for i := 1; i < 30; i++ {
		num := rand.Intn(100)
		fmt.Println(num)
		bst.Insert(num)
	}
	bst.Show()

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(2)
	bst.Insert(1)
	// bst.Insert(11)
	// bst.InOrderTraverse()
	// bst.Delete(10)
	// bst.Delete(2)
	// bst.Delete(6)
	// bst.Delete(3)
	// bst.InOrderTraverse()

}
