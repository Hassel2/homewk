using System;

namespace MyQueue {
	public class QueueElem<T> {
		private T info;
		private QueueElem<T> next = null;

		public T Info => info;
		public QueueElem<T> Next { get { return next; } set { next = value; } }

		public QueueElem(T Inf) {
			info = Inf;
		}
	}
}
