package l33tc0d3

// First attempt, runs horribly slow
/*
func hasCycle(head *ListNode) bool {
	curr := head
	visited := make(map[*ListNode]bool)
	for curr != nil {
		if visited[curr] {
			return true
		} else {
			visited[curr] = true
			curr = (*curr).Next
		}
	}
	return false
}
*/

// Second attempt
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
