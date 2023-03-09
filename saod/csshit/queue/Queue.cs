using System;

namespace MyQueue {
	public class Queue<T> {
		private QueueElem<T> first = null;
		private QueueElem<T> last = null;
		
		public Queue() {  }

		public void In(T info) {
			if (this.IsEmpty()) {
				last = new QueueElem<T>(info);
				first = last;
				return;
			}
			last.Next = new QueueElem<T>(info);
			last = last.Next;
		}
		
		public T Out() {
			if (this.IsEmpty()) throw new FormatException("Queue is empty");
			T temp = first.Info;
			first = first.Next;
			return temp;
		}

		public bool IsEmpty() {
			return (last != null);
		}

		public T Peek() {
			return first.Info;
		}
	}
}
