using System;
using System.Collections;

namespace InvPol {
	public class RPN {
		private Stack<string> intStack = new Stack<string>();
		private Stack<char> operStack = new Stack<char>();
		private string output = "";
		
		public string SimpToRPN(string input) {
			string temp = "";
			input = input.Replace(" ", "") + " ";

			foreach(char i in input) {
				if (Char.IsDigit(i)) {
					temp += i.ToString();
				} else {
					if (temp != "") intStack.Push(temp);
					output += temp;
					temp = "";

					if (i == '(') {
						operStack.Push(i);
					} else if (i == ')') {
						char s = operStack.Pop();

						while (s != '(') {
							output += s.ToString();
							s = operStack.Pop();
						}
					} else {
						if (GetPriority(i) == -1) throw new FormatException("Неизвестный оператор " + i.ToString());
						if (GetPriority(i) == 7) {
							while (operStack.Count > 0)
								output += operStack.Pop().ToString();
							return output;
						}
						if (operStack.Count > 0) 
							if (GetPriority(i) <= GetPriority(operStack.Peek()))
								output += operStack.Pop().ToString();
						operStack.Push(i);
					}
				}
			}
			return output;
		}

		public int GetPriority(char op) {
			//return "()+-*/^ ".IndexOf(op);
			switch (op)
			{
				case '(': return 0;
				case ')': return 1;
				case '+': return 2;
				case '-': return 3;
				case '*': return 4;
				case '/': return 4;
				case '^': return 5;
				case ' ': return 7;
				default: return -1;
			}
		}

		public string Calculate() {
			Stack<string> revIntStack = new Stack<string>(intStack);
			foreach (char i in output) {
				if (!Char.IsDigit(i)) {
					/* foreach(string j in revIntStack.ToArray()) Console.Write(j + " "); */
					/* Console.WriteLine(); */
					int temp1 = Convert.ToInt32(revIntStack.Pop());
					int temp2 = Convert.ToInt32(revIntStack.Pop());
					switch(i) {
						case '+':
							revIntStack.Push((temp1 + temp2).ToString());
							break;
						case '-':
							revIntStack.Push((temp1 - temp2).ToString());
							break;
						case '*':
							revIntStack.Push((temp1 * temp2).ToString());
							break;
						case '/':
							revIntStack.Push((temp1 / temp2).ToString());
							break;
						case '^':
							revIntStack.Push((Math.Pow(temp1, temp2)).ToString());
							break;
					}
				}
			}
			return revIntStack.Pop();
		}
	}
}
