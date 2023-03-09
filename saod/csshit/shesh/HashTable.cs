namespace HHHash {
	public class HashTable {
		private List<string>[] body = null;
		private int length = 0;
		
		public int Length => length;
		

		private int Hash(int key) 
		{
			return key % length;
		}


		public HashTable(int len) 
		{
			length = len;
			body = new List<string>[len];
		}

		public void Add(string data) 
		{
			int index = Hash(Convert.ToInt32(data.Substring(0,5)));
			if (body[index] == null)
				body[index] = new List<string>();
			body[index].Add(data);
		}

		public string Search(int key) 
		{
			int possibleIndex = Hash(key);
			for (int i = 0; i < body[possibleIndex].Count(); i++) {
				if (Convert.ToInt32(body[possibleIndex][i].Substring(0, 5)) == key) {
					return body[possibleIndex][i];
				}
			}

			return "";
		}
	}
}

