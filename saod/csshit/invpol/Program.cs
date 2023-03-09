using System;
using System.Collections;

namespace InvPol {
	public class Program {
		public static void Main() {
			Console.WriteLine("Hello, World!");
			string simple = Console.ReadLine();
			RPN rpn = new RPN();
			Console.WriteLine(rpn.SimpToRPN(simple));
			Console.WriteLine(rpn.Calculate());
		}
	}
}
