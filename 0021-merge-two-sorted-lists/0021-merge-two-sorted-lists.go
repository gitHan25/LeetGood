/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
	dummy := &ListNode{Val:0}
	curr := dummy

	

	for list1 != nil && list2 != nil{
		if list1.Val > list2.Val {
			curr.Next = list2
			list2 = list2.Next
		}else{
			curr.Next = list1
			list1 = list1.Next
		}
        curr = curr.Next
	}
     if list1 == nil{
         curr.Next = list2
     }else{
         curr.Next = list1
     }

    return dummy.Next
}

