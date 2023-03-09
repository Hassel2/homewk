namespace List
{
    public class DLLlist<T>
        where T : IComparable
    {
        private ListElem<T> first = null;
        private ListElem<T> last = null;
        private int length = 0;

        public int Length => length;

        public T this[int key]
        {
            get { return find(key).Info; }
            set { find(key).Info = value; }
        }

        public void AddFirst(T inf)
        {
            length++;
            if (first != null)
            {
                first.Prev = new ListElem<T>(inf);
                first.Prev.Next = first;
                if (last.Equals(first))
                {
                    last.Prev = first.Prev;
                }
                first = first.Prev;
            }
            else
            {
                first = new ListElem<T>(inf);
                last = first;
            }
        }

        public void AddLast(T inf)
        {
            length++;
            if (last != null)
            {
                last.Next = new ListElem<T>(inf);
                last.Next.Prev = last;
                if (last == first)
                {
                    first.Next = last.Next;
                }
                last = last.Next;
            }
            else
            {
                last = new ListElem<T>(inf);
                first = last;
            }
        }

        public void AddAtIndex(int index, T inf)
        {
            length++;
            ListElem<T> current = find(index);
            ListElem<T> temp = current.Prev;
			var tmp = new ListElem<T>(inf);
            current.Prev = tmp;
			Console.WriteLine("curr: " + current.Info);
			Console.WriteLine("prev: " + current.Prev.Info);
			Console.WriteLine("tmp: " + tmp);
            current.Prev.Next = current;
            current.Prev.Prev = temp;
            current = null;
        }

        public void RemoveFirst()
        {
            length--;
            first = first.Next;
            first.Prev.Next = null;
            first.Prev = null;
        }

        public void RemoveLast()
        {
            length--;
            last = last.Prev;
            last.Next.Prev = null;
            last.Next = null;
        }

        public void RemoveAt(int index)
        {
            length--;
            ListElem<T> current = find(index);
            current.Prev.Next = current.Next;
            current.Next.Prev = current.Prev;
            current.Prev = null;
            current.Next = null;
            current = null;
        }

        public bool Remove(T val)
        {
            ListElem<T> current = first;
            while (current.Info.CompareTo(val) != 0)
            {
                current = current.Next;
            }
            if (current.Next == null) { return false; }
            length--;
            current.Prev.Next = current.Next;
            current.Next.Prev = current.Prev;
            current.Prev = null;
            current.Next = null;
            current = null;
            return true;
        }

        public void Clear()
        {
            length = 0;
            ListElem<T> current = first;
            while (current.Next != null)
            {
                current.Prev = null;
                current = current.Next;
                current.Prev.Next = null;
            }
            first = null;
            last = null;
        }

        public void Sort()
        {
            ListElem<T> current;
            for (int i = 0; i < length - 1; ++i)
            {
                current = first;
                for (int j = 0; j < length - 1; ++j)
                {
                    if (current.Info.CompareTo(current.Next.Info) > 0)
                    {
                        T tmp = current.Info;
                        current.Info = current.Next.Info;
                        current.Next.Info = tmp;
                    }
                    current = current.Next;
                }
            }
        }

        public bool ContainsSorted(T inf)
        {
            if (inf.CompareTo(first.Info) < 0 || inf.CompareTo(last.Info) > 0)
            {
                return false;
            }

            ListElem<T> current = first;
            for (int i = 0; i < length - 1; i++)
            {
                if (current.Info.CompareTo(inf) == 0) return true;
                if (current.Info.CompareTo(inf) > 0) return false;
                current = current.Next;
            }
            return true;
        }

        public void AddSorted(T inf)
        {
			length++;
            if (inf.CompareTo(first.Info) <= 0)
            {
                AddFirst(inf);
                return;
            }
            if (inf.CompareTo(last.Info) >= 0)
            {
                AddLast(inf);
                return;
            }
			
            ListElem<T> current = first.Next;
            for (int i = 1; i < length - 2; i++)
            {
                if (current.Info.CompareTo(inf) >= 0)
                {
                    /* ListElem<T> tmp = current.Prev; */
                    /* current.Prev = new ListElem<T>(inf); */
                    /* current.Prev.Prev = tmp; */
                    /* current.Prev.Next = current; */
                    /* tmp.Next = current.Prev; */
					AddAtIndex(i, inf);
					return;
				}
				current = current.Next;
            }
        }

        public override string ToString()
        {
            ListElem<T> current = last;
            string result = "[";
            while (current != null)
            {
                result += current.Info + ", ";
                current = current.Next;
            }
            return result + "]";
        }

        private ListElem<T> find(int key)
        {
            if (key >= length || key < 0) throw new FormatException("Unknown key " + key);
            int counter = 0;
            ListElem<T> current;
            /* if (key > Length / 2) */
            /* { */
            /*     counter = Length - 1; */
            /*     current = last; */
            /*     while (key != counter) */
            /*     { */
            /*         current = current.Prev; */
            /*         counter--; */
            /*     } */
            /*     return current; */
            /* } */
            current = first;
            while (key != counter)
            {
                current = current.Next;
                counter++;
            }
            return current;
        }
    }
}
