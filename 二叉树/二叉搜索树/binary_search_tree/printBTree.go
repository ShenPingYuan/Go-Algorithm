package main

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type btNode struct {
	Data   interface{}
	Lchild *btNode
	Rchild *btNode
}

type biTree struct {
	Root *btNode
}

func Build(data interface{}) *biTree {
	var list []interface{}
	if data == nil {
		return &biTree{}
	}
	switch data.(type) {
	case []interface{}:
		list = append(list, data.([]interface{})...)
	default:
		list = append(list, data)
	}
	if len(list) == 0 {
		return &biTree{}
	}
	node := &btNode{Data: list[0]}
	list = list[1:]
	Queue := []*btNode{node}
	for len(list) > 0 {
		if len(Queue) == 0 {
			//panic("Given array can not build binary tree.")
			return &biTree{Root: node}
		}
		cur := Queue[0]
		val := list[0]
		Queue = Queue[1:]
		if val != nil {
			cur.Lchild = &btNode{Data: val}
			if cur.Lchild != nil {
				Queue = append(Queue, cur.Lchild)
			}
		}
		list = list[1:]
		if len(list) > 0 {
			val := list[0]
			if val != nil {
				cur.Rchild = &btNode{Data: val}
				if cur.Rchild != nil {
					Queue = append(Queue, cur.Rchild)
				}
			}
			list = list[1:]
		}
	}
	return &biTree{Root: node}
}

func (bt *biTree) AppendNode(data interface{}) {
	root := bt.Root
	if root == nil {
		bt.Root = &btNode{Data: data}
		return
	}
	Queue := []*btNode{root}
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		if cur.Lchild != nil {
			Queue = append(Queue, cur.Lchild)
		} else {
			cur.Lchild = &btNode{Data: data}
			return
		}
		if cur.Rchild != nil {
			Queue = append(Queue, cur.Rchild)
		} else {
			cur.Rchild = &btNode{Data: data}
			break
		}
	}
}

func Copy(bt *biTree) *biTree {
	root := bt.Root
	if root == nil {
		return &biTree{}
	}
	node := &btNode{Data: root.Data}
	Queue1, Queue2 := []*btNode{root}, []*btNode{node}
	for len(Queue1) > 0 {
		p1, p2 := Queue1[0], Queue2[0]
		Queue1, Queue2 = Queue1[1:], Queue2[1:]
		if p1.Lchild != nil {
			Node := &btNode{Data: p1.Lchild.Data}
			p2.Lchild = Node
			Queue1 = append(Queue1, p1.Lchild)
			Queue2 = append(Queue2, Node)
		}
		if p1.Rchild != nil {
			Node := &btNode{Data: p1.Rchild.Data}
			p2.Rchild = Node
			Queue1 = append(Queue1, p1.Rchild)
			Queue2 = append(Queue2, Node)
		}
	}
	return &biTree{Root: node}
}

func Mirror(bt *biTree) *biTree {
	root := bt.Root
	if root == nil {
		return &biTree{}
	}
	node := &btNode{Data: root.Data}
	Queue1, Queue2 := []*btNode{root}, []*btNode{node}
	for len(Queue1) > 0 {
		p1, p2 := Queue1[0], Queue2[0]
		Queue1, Queue2 = Queue1[1:], Queue2[1:]
		if p1.Lchild != nil {
			Node := &btNode{Data: p1.Lchild.Data}
			p2.Rchild = Node
			Queue1 = append(Queue1, p1.Lchild)
			Queue2 = append(Queue2, Node)
		}
		if p1.Rchild != nil {
			Node := &btNode{Data: p1.Rchild.Data}
			p2.Lchild = Node
			Queue1 = append(Queue1, p1.Rchild)
			Queue2 = append(Queue2, Node)
		}
	}
	return &biTree{Root: node}
}

func (bt *biTree) BForder2D() [][]interface{} {
	var res [][]interface{}
	root := bt.Root
	if root == nil {
		return res
	}
	Queue := []*btNode{root}
	for len(Queue) > 0 {
		Nodes := []interface{}{}
		Levels := len(Queue)
		for Levels > 0 {
			cur := Queue[0]
			Queue = Queue[1:]
			Nodes = append(Nodes, cur.Data)
			Levels--
			if cur.Lchild != nil {
				Queue = append(Queue, cur.Lchild)
			}
			if cur.Rchild != nil {
				Queue = append(Queue, cur.Rchild)
			}
		}
		res = append(res, Nodes)
	}
	return res
}

func (bt *biTree) LevelNullwith(Fills ...interface{}) [][]interface{} {
	var Nodes [][]interface{}
	var Fill0 interface{}
	if len(Fills) == 0 {
		Fill0 = nil
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
	Queue := []*btNode{root}
	for len(Queue) > 0 {
		nodes := []interface{}{}
		Level := len(Queue)
		for Level > 0 {
			cur := Queue[0]
			Queue = Queue[1:]
			nodes = append(nodes, cur.Data)
			Count++
			Level--
			if cur.Lchild != nil {
				Queue = append(Queue, cur.Lchild)
			}
			if cur.Rchild != nil {
				Queue = append(Queue, cur.Rchild)
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

func XmlNode(M, N, X, Y int, Data string, Color ...string) string {
	var cColor, tColor string
	R := 30
	Node := `<Tab><g id="M,N">
	<circle cx="X" cy="Y" r="RC" stroke="black" stroke-width="2" fill="COLOR" />
	<text x="X" y="Y" fill="TextColor" font-size="20" text-anchor="middle"
		dominant-baseline="middle">DATA</text>
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

func Xml4Full(Levels int) string {
	var Xml, Node string
	Head := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/" +
		"1999/xlink\" version=\"1.1\" width=\"Width\" height=\"Height\">\nCONTENT</svg>"
	Line := `<line x1="X1" y1="Y1" x2="X2" y2="Y2" style="stroke:black;stroke-width:2"/>`
	Link := `<a xlink:href="https://blog.csdn.net/boysoft2002" target="_blank">
	<text x="5" y="20" fill="blue">Hann's CSDN Homepage</text></a>`
	for i := 0; i < Levels; i++ {
		negative := -1
		for j := 0; j < Pow2(i); j++ {
			t := Pow2(Levels - i - 1)
			x, y := 50*(2*t*j+t), 120*i+50
			if Levels != 1 && i == Levels-1 {
				Xml = XmlNode(i, j, x, y, strconv.Itoa(Pow2(i)+j), "lightgreen")
			} else {
				Xml = XmlNode(i, j, x, y, strconv.Itoa(Pow2(i)+j))
			}
			if i > 0 {
				line1 := strings.Replace(Line, "X1", strconv.Itoa(x), 1)
				line1 = strings.Replace(line1, "Y1", strconv.Itoa(y-30), 1)
				negative *= -1
				x0, y0 := 21, 21
				//过连线起始端的半径与纵轴线夹角取45度时x,y坐标修正值21,21[30/1.414]
				//取30度时 x0,y0:= 15,30-26；取60度时 x0,y0:= 26[15*1.732],15
				x += 50*negative*(2*t*j%2+t) - negative*x0
				line1 = strings.Replace(line1, "X2", strconv.Itoa(x), 1)
				line1 = strings.Replace(line1, "Y2", strconv.Itoa(y-120+y0), 1)
				Xml = strings.Replace(Xml, "<ROOT/>", line1, 1)
			}
			Node += Xml
		}
	}
	Xml = strings.Replace(Head, "CONTENT", Node, 1)
	Xml = strings.Replace(Xml, "Width", strconv.Itoa(Pow2(Levels)*50), 1)
	Xml = strings.Replace(Xml, "Height", strconv.Itoa(Levels*120), 1)
	Xml = strings.Replace(Xml, "<ROOT/>", Link, 1)
	return Xml
}

func (bt *biTree) Xml4Tree() string {
	var Xml, Node string
	Head := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/" +
		"1999/xlink\" version=\"1.1\" width=\"Width\" height=\"Height\">\nCONTENT</svg>"
	Line := `<line x1="X1" y1="Y1" x2="X2" y2="Y2" style="stroke:black;stroke-width:2"/>`
	Link := `<a xlink:href="https://blog.csdn.net/boysoft2002" target="_blank">
	<text x="5" y="20" fill="blue">Hann's CSDN Homepage</text></a>`
	List := bt.LevelNullwith()
	Levels := len(List)
	for i := Levels - 1; i >= 0; i-- {
		negative := -1
		TmpXml := ""
		for j := 0; j < Pow2(i); j++ {
			t := Pow2(Levels - i - 1)
			x, y := 50*(2*t*j+t), 120*i+50
			if List[i][j] != nil {
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
				Xml = XmlNode(i, j, x, y, TmpStr, fillColor)
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
	Xml = strings.Replace(Xml, "<ROOT/>", Link, 1)
	return Xml
}

func main() {

	var file *os.File
	var err1 error

	list1 := []interface{}{"-", "*", 6, "+", 3, nil, nil, 2, 8}
	list2 := []interface{}{1, 2, 3, 4, 5, nil, 6, 7, 8}
	tree1 := Build(list1)
	tree2 := Build(list2)
	tree3 := Mirror(tree2)
	texts := []string{tree1.Xml4Tree(), tree2.Xml4Tree(), tree3.Xml4Tree(), Xml4Full(4)}

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

	// fmt.Println("Welcome to my homepage: https://blog.csdn.net/boysoft2002")
	// exec.Command("cmd", "/c", "start", "https://blog.csdn.net/boysoft2002").Start()

}
