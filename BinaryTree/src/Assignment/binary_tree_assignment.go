package Assignment
import (
    "fmt"
)

type Tree struct {
    Root *Node
}

type Node struct {
    data  int
    left  *Node
    right *Node
}

func (t *Tree) Insert(data int) {
    if t.Root == nil {
        t.Root = &Node{data: data}
    } else {
        t.Root.Insert(data)
    }
}

func (n *Node) Insert(data int) {
    if data <= n.data {
        if n.left == nil {
            n.left = &Node{data: data}
        } else {
            n.left.Insert(data)
        }
    } else {
        if n.right == nil {
            n.right = &Node{data: data}
        } else {
            n.right.Insert(data)
        }
    }
}

func PreorderTraversal(n *Node) {
    if n == nil {
        return
    } else {
        fmt.Printf("%d ", n.data)
        PreorderTraversal(n.left)
        PreorderTraversal(n.right)
    }
}

func PostorderTraversal(n *Node) {
    if n == nil {
        return
    } else {
        PostorderTraversal(n.left)
        PostorderTraversal(n.right)
        fmt.Printf("%d ", n.data)
    }
}

func InorderTraversal(n *Node) {
    if n == nil {
        return
    } else {
        InorderTraversal(n.left)
        fmt.Printf("%d ", n.data)
        InorderTraversal(n.right)
    }
}

//func main() {
  
//}