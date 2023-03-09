namespace MyList
{
    public class List<T>
		where T : IComparable<T>
    {
        private T[] body = null;
        private int current = 0;

		public int Length => current + 1; 

        public T this[int key]
        {
            get 
			{ 
				if (key > current) throw new FormatException("Unknown key");
				return body[key]; 
			}
            set 
			{
				if (key > current) throw new FormatException("Unknown key");
				body[key] = value; 
			}
        }

        public List()
                /* where T is IComparable */
		{  
			body = new T[10];
		}

		public List(int size) 
		{
			body = new T[size];
		}

		public void Add(T item)
		{
			if (current >= body.Length) Array.Resize(ref body, body.Length * 2);
			this[current++] = item;
		}

		public bool Contains(T item)
		{
			for (int i = 0; i < Length; i++) 
			{
				if (item.CompareTo(body[i]) == 0) return true;
				if (item.CompareTo(body[Length - i - 1]) == 0) return true;
			}
			return false;
		}

		public void Clear() { current = 0; }

		public bool Remove(T item)
		{
			if (!Contains(item)) return false;
			
			for (int i = 0; i < Length; i++)
			{
				if (item.CompareTo(body[i]) == 0) 
				{
					RemoveAt(i);
					return true;
				}
			}
			return true;
		}

		public void RemoveAt(int index) 
		{
			if (index > current) throw new FormatException("Unknown key");
			for (int i = index; i < current; i++) 
			{
				body[i] = body[i + 1];
			}
			current = current - 1;
		}

	}
}
