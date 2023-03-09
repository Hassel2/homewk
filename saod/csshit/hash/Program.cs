namespace HHHash {
	public class Program {
		public static void Main() {
			string[] lines = System.IO.File.ReadAllLines("Astra.txt");
			/* HashTable hash = new HashTable(lines.Length); */
			MegaHashTable mhash = new MegaHashTable(lines.Length);
			foreach (string line in lines)
			{
				/* Console.WriteLine(hash.Add(line)); //+ " " + mhash.Add(line)); */

				Console.WriteLine(mhash.Add(line));
			}
			Random rnd = new Random();
			string search = lines[rnd.Next(0, lines.Length - 1)];
			Console.WriteLine("Number to search: " + search.Substring(0, 5));
			Console.WriteLine(mhash.Search(Convert.ToInt32(search.Substring(0,5))));
		}
	}
}
