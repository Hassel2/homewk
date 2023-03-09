namespace MyList
{
    public class Program
    {
        public static void Main()
        {
            List<int> lst = new List<int>();
            Random rnd = new Random();
            for (int i = 0; i < 16; i++)
            {
                lst.Add(rnd.Next(0, 10));
				Console.Write(lst[i] + " ");
            }
			Console.WriteLine();
			Console.WriteLine(lst.Length);
			lst.Remove(lst[4]);
            for (int i = 0; i <= lst.Length ; i++)
            {
				Console.Write(lst[i] + " ");
            }
			Console.WriteLine();
			Console.WriteLine(lst.Length);
        }
    }
}

