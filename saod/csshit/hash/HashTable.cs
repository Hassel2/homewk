namespace HHHash {
	public class HashTable {
		private string[] body = null;
		private int length = 0;
		
		public int Length => length;
		

		public virtual int Hash(int key) 
		{
			return key % length;
		}


		public HashTable(int len) 
		{	
			length = len;
			body = new string[len];
		}

		public ulong Add(string data) 
		{
			int index = Hash(Convert.ToInt32(data.Substring(0,5)));
			Console.WriteLine(Convert.ToInt32(data.Substring(0,5)));
			if (String.IsNullOrEmpty(body[index])) {
				body[index] = data;
				return 0;
			}
			else 
			{
				uint collision = 1;
				int newIndex = (index + 1) % length;
				while (newIndex != index) 
				{
					if (String.IsNullOrEmpty(body[newIndex])) 
					{
						body[newIndex] = data;
						return collision;
					}
					newIndex = (newIndex + 1) % length;
					collision++;
				}
			}
			throw new Exception("Set is full");
		}

		public string Search(int key) 
		{
			int possibleIndex = Hash(key);
			if (Convert.ToInt32(body[possibleIndex].Substring(0,5)) == key) 
			{
				return body[possibleIndex];
			}
			else
			{
				int index = possibleIndex + 1;
				while (index != possibleIndex)
				{
					if (Convert.ToInt32(body[index].Substring(0,5)) == key) 
					{
						return body[index];
					}
					index = (index + 1) % length;
				}
			}

			return "";
		}
	}
}

