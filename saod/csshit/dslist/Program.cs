namespace List
{
    public class Program
    {
        public static void Main()
        {
			DLLlist<int> DLL = new DLLlist<int>();
			Random rnd = new Random();
			
			for(int i = 0; i < 10; ++i) {
				DLL.AddLast(rnd.Next(1, 10));
				Console.Write(" " + DLL[i] + " ");
			}
			Console.WriteLine();
			DLL.Sort();
			for(int i = 0; i < 10; ++i) {
				Console.Write(DLL[i] + " ");
			}
			Console.WriteLine(DLL.ContainsSorted(4));
			DLL.AddSorted(4);	
			for(int i = 0; i < 11; ++i) {
				Console.Write(DLL[i] + " ");
			}
        }
    }
}
