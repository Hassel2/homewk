namespace List
{
    public class ListElem<T>
    {
        T info;
        private ListElem<T> next = null;
        private ListElem<T> prev = null;

		public T Info { get { return info; } set{ info = value; }}
		public ListElem<T> Next { get { return next; } set { next = value; } }
		public ListElem<T> Prev { get { return prev; } set { next = value; } }

        public ListElem(T inf) 
		{ 
			info = inf;
		}
    }
}
