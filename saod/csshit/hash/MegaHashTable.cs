namespace HHHash {
	public class MegaHashTable : HashTable 
	{
		public MegaHashTable(int len)
			: base(len)
		{  }
		public override int Hash(int key) 
		{
			return Convert.ToInt32((((key * (Math.Sqrt(5) + 1) / 2) - Math.Truncate(key * (Math.Sqrt(5) + 1) / 2)) * Length));
		}
	}
}
