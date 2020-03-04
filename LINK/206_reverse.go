//METHOD 1:
func reverseList(head *ListNode) *ListNode {
    var pre  *ListNode
    cur := head 
    for cur != nil {
        pre, cur, cur.Next = cur, cur.Next, pre 
    }
    return pre 
}


//METHOD 2: TODO 看一下递归的思路
