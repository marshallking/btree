package main 

import "fmt"

type Node struct{
	left * Node
	right * Node
	data int
}
var gnumNodes    = 0
var gFullNodes   = 0
var gSingleNodes = 0

    //------------------------------------------------------------------------------------
	//
    //     1                                  N  O  D  E   1
    //                                /                          \
	//                               /                            \
	//     2                      node2                          node3
	//                            /   \                          /  \
	//                           /     \                        /    \
	//     3                  node 4     node 5                Node8  NULL
	//                       /   \       /   \                /    \
	//                      /     \     /     \              /      \    
	//     4               NULL   NULL node 6  NULL        NULL     Node9 
	//                                /    \                         /  \
	//                               /      \                       /    \
	//     5                        NULL    node 7               NULL   Node10
	//                                     /    \                       /   \
	//                                    /      \                     /     \
	//     6                             NULL      NULL             NULL    NULL
	//
	//
	//-----------------------------------------------------------------------------------------


func main() {
   /* local variable definition */
   Node1 := &Node{nil,nil,1}
   Node2 := &Node{nil,nil,2}   
   Node3 := &Node{nil,nil,3}
   Node4 := &Node{nil,nil,4}
   Node5 := &Node{nil,nil,5}
   Node6 := &Node{nil,nil,6}
   Node7 := &Node{nil,nil,7}
   Node8 := &Node{nil,nil,8}  
   Node9 := &Node{nil,nil,9}
   Node10 := &Node{nil,nil,10} 
   Node11 := &Node{nil,nil,11}    
   
   Node1.left     = Node2
   Node1.right    = Node3
   Node2.left     = Node4
   Node2.right    = Node5
   Node3.left     = Node8
   Node8.left     = nil
   Node8.right    = Node9
   Node9.left     = nil
   Node9.right    = Node10
   Node4.left     = nil
   Node4.right    = nil
   Node5.left     = Node6
   Node5.right    = nil
   Node6.left     = nil
   Node6.right    = Node7
   Node7.left     = nil
   Node7.right    = nil
   Node10.left    = Node11
   
   numNodes := CountNodes(Node1) + 1  //Plus one for the root node
                   
   fmt.Printf( "BTREE - Number of nodes in tree =  : %d\n"+
   	           "gnumNodes = %d\n" +
               "FullNodes = %d \n"+
               "gSingleNodes = %d", numNodes,gnumNodes,gFullNodes,gSingleNodes )
}

/* function returning the max between two numbers */
func CountNodes(node * Node ) int {
	 
	numNodes:= 0
	 
	fmt.Printf( "BTREE -  : data = %d\n", node.data )	
 
	if(node.left != nil){
		//fmt.Printf( "BTREE - LEFT : data = %d\n", node.data )		
		if(node.right!=nil){	
			gFullNodes++
		}else{
			gSingleNodes++
		}
		gnumNodes++
		numNodes += CountNodes(node.left) + 1	
		 		
	}
	if(node.right!= nil){	
		if(node.left == nil){
			gSingleNodes++			
		}
		gnumNodes++		
		numNodes += CountNodes(node.right) + 1	
		 							
	}
	 	 
   return numNodes
}