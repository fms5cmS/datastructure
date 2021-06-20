## 回文链表的判断

- 思路一

正着念和反着念都一样。

遍历链表将元素放入栈中，此时栈弹出的顺序就是链表逆序了，然后每次出栈时和链表正序的元素比对，如果都相等，则是回文结构。额外空间复杂度$O(N)$。



```java
public boolean isPalindrome1(Node head) {
  Stack<Node> stack = new Stack<Node>();
  Node cur = head;
  while (cur != null) {
    stack.push(cur);
    cur = cur.next;
  }
  while (head != null) {
    if (head.value != stack.pop().value) {
      return false;
    }
    head = head.next;
  }
  return true;
}
```



## 思路二

对称

使用两个指针，一快一慢，快指针每次走两步，慢指针每次走一步，当快指针走完时，慢指针就走到了链表的中间，然后将慢指针之后遍历的元素放入栈中，然后比对出栈元素和链表的元素(从头开始)，当栈中元素全部弹出都没有元素不等，则该链表是回文结构。额外空间复杂度$O(N/2)$。

```java
public boolean isPalindrome2(Node head) {
  if (head == null || head.next == null) {
    return true;
  }
  Node right = head.next;
  Node cur = head;
  while (cur.next != null && cur.next.next != null) {
    right = right.next;
    cur = cur.next.next;
  }
  Stack<Node> stack = new Stack<Node>();
  while (right != null) {
    stack.push(right);
    right = right.next;
  }
  while (!stack.isEmpty()) {
    if (head.value != stack.pop().value) {
      return false;
    }
    head = head.next;
  }
  return true;
}
```



## 思路三

使用两个指针，一快一慢，快指针每次走两步，慢指针每次走一步，当快指针走完时，慢指针就走到了链表的中间，然后将慢指针之后的元素逆序，如：1—>2—>3—>2—>1转变为1—>2—>3<—2<—1，3指向null，然后使用两个指针从两端开始遍历并比较，如果有不等的两个节点则返回false，如果某个节点走到null时所有节点都相等则返回true。注意：无论最后的结果是true还是false，都需要将链表还原(所以需要记录中点的位置)！额外空间复杂度$O(1)$。时间复杂度$O(n)$。

注意：**要保证当链表元素为奇数时，慢指针走到中间的节点，当链表元素为偶数时，慢指针走到中间相等的两个节点的前一个节点！！！！！！**

```java
public class HuiWen {
  private static class Node {
    public int value;
    public Node next;

    public Node(int data) {
      this.value = data;
    }
  }

  public static boolean isPalindrome(Node head) {
    Node left = head;  //记录链表左侧的 head ，便于后面的比较
    Node fast = head;
    Node slow = head;
    boolean result = true;
    //确保奇数个元素时n1走到中点，偶数个时n1走到中点的前一个节点
    while (fast.next != null && fast.next.next != null) {
      fast = fast.next.next;
      slow = slow.next;
    }
    //此时slow走到了中点
    fast = slow.next;
    slow.next = null;
    Node right = reverse(fast, slow);
    //这里记录最后一个节点，便于之后还原链表
    Node last = right;
    //比较
    while (left != null && right != null) {
      if (left.value != right.value) {
        result = false;
        break;
      }
      left = left.next;
      right = right.next;
    }
    //还原链表
    reverse(last, null);
    return result;
  }

  private static Node reverse(Node head, Node pre) {
    Node next = null;
    while (head != null) {
      next = head.next;
      head.next = pre;
      pre = head;
      head = next;
    }
    return pre;
  }

  public static void printLinkedList(Node node) {
    System.out.print("Linked List: ");
    while (node != null) {
      System.out.print(node.value + " ");
      node = node.next;
    }
    System.out.println();
  }

  public static void main(String[] args) {
    Node head = new Node(1);
    head.next = new Node(2);
    head.next.next = new Node(3);
    head.next.next.next = new Node(2);
    head.next.next.next.next = new Node(1);
    //head.next.next.next.next.next = new Node(2);
    printLinkedList(head);
    System.out.println(isPalindrome(head));
    printLinkedList(head); //看是否还原了链表
  }
}
```