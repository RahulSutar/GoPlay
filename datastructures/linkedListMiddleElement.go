//Program to find the middle element from a given linked list
//Author 	: Rahul Sutar

package datastructures

import (
	"fmt"
	"math"
)

//Node - Linked list node
type Node struct {
	integerValue int
	nextNode     *Node
}

//FindMiddleElementFromLinkedList - Finds the middle element from a linked list
func FindMiddleElementFromLinkedList() {

	node5 := Node{5, nil}
	node4 := Node{4, &node5}
	node3 := Node{3, &node4}
	node2 := Node{2, &node3}
	node1 := Node{1, &node2}

	startNode := &node1

	elementsToTraverse := GetMiddleElementIndex(startNode)

	// fmt.Println("Elements to traverse : ", elementsToTraverse)

	if elementsToTraverse == -1 {
		fmt.Println("Probably the list does not have middle element")
		return
	}

	middleElement := GetElementValAtIndex(startNode, elementsToTraverse)

	fmt.Println("MiddleElement Value : ", middleElement.integerValue)

}

//GetMiddleElementIndex - to traverse the link
func GetMiddleElementIndex(headLink *Node) float64 {

	middleIndex := 0

	for headLink != nil {
		fmt.Println("Element : ", headLink.integerValue)
		headLink = headLink.nextNode
		middleIndex++
	}

	if middleIndex%2 == 0 {
		return -1
	}

	return math.Ceil(float64(middleIndex) / 2.0)
}

//GetElementValAtIndex - Get element at index
func GetElementValAtIndex(headLink *Node, elements float64) *Node {

	for (elements - 1) > 0 {
		headLink = headLink.nextNode
		if headLink == nil {
			return nil
		}
		elements--

	}
	return headLink
}
