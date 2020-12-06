package main

import (


	binary_tree "BinaryTreeAssignment/src/Assignment"
	"fmt"
	
)

func main() {
	 

	for {
	  var tree binary_tree.Tree
      var user_input int


      fmt.Println("Binary Tree Assignment")
      fmt.Println("You want to run default inputs or manual inputs")
      fmt.Println("Press 1 for running default inputs")
      fmt.Println("Press 2 for getting inputs from users")
      fmt.Print("please enter ->")
      fmt.Scanf("%d",&user_input)

    
      if user_input == 1{

      	tree.Insert(100)
	    tree.Insert(-20)
	    tree.Insert(30)
	    tree.Insert(4)
	    tree.Insert(5)
	    tree.Insert(78)

	    fmt.Printf("Pre Order  :")
	    binary_tree.PreorderTraversal(tree.Root)
	    fmt.Println()
	    fmt.Printf("Post Order :")
	    binary_tree.PostorderTraversal(tree.Root)
	    fmt.Println()
	    fmt.Printf("In Order   :")
	    binary_tree.InorderTraversal(tree.Root)
	    fmt.Println()

      }else  if user_input == 2{

      	var total_number int
      	var manual_number int
      	fmt.Print("Enter Number of elements to be Inserted --")
      	fmt.Scanf("%d",&total_number)

      	for i:=0;i<total_number;i++{

      		fmt.Scanf("%d",&manual_number)
      		tree.Insert(manual_number)
      	}

      	fmt.Println("Thanks for entering your manual inputs")

      	fmt.Printf("Pre Order  :")
	    binary_tree.PreorderTraversal(tree.Root)
	    fmt.Println()
	    fmt.Printf("Post Order :")
	    binary_tree.PostorderTraversal(tree.Root)
	    fmt.Println()
	    fmt.Printf("In Order   :")
	    binary_tree.InorderTraversal(tree.Root)
	    fmt.Println()


      }else{

      	fmt.Println("please give valid option")

      }

   
	}

}
