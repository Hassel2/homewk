namespace HHHash {
	public class Program {
		public static void Main() {
			string[] lines = System.IO.File.ReadAllLines("Astra.txt");
			HashTable hash = new HashTable(lines.Length);
			foreach (string line in lines)
			{
				hash.Add(line);
			}
			Random rnd = new Random();
			string search = lines[rnd.Next(0, lines.Length - 1)];
			Console.WriteLine("Number to search: " + search.Substring(0, 5));
			Console.WriteLine(hash.Search(Convert.ToInt32(search.Substring(0,5))));
		}
	}
}
